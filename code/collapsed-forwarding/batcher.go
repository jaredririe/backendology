package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"sync"
	"time"
)

type requestBatcher struct {
	batch   map[string][]*request
	handler http.HandlerFunc // handler describes what Flush should do to handle the request

	mu    *sync.Mutex
	close chan struct{}
}

type request struct {
	w    http.ResponseWriter
	r    *http.Request
	done chan struct{} // done provides a way to indicate when the request has been handled
}

func newRequestBatcher(handler http.HandlerFunc) *requestBatcher {
	rc := &requestBatcher{
		batch:   make(map[string][]*request),
		handler: handler,
		mu:      &sync.Mutex{},
		close:   make(chan struct{}),
	}

	go func() {
		for {
			ticker := time.NewTicker(5 * time.Second)
			select {
			case <-ticker.C:
				rc.Flush()
			case <-rc.close:
				return
			}
		}
	}()

	return rc
}

// Append appends the given request to the slice of requests under the
// given key.
func (rc *requestBatcher) Append(key string, request *request) {
	rc.mu.Lock()
	defer rc.mu.Unlock()

	rc.batch[key] = append(rc.batch[key], request)
}

// Close notifies the flushing goroutine created in the
// constructor that it should finish execution. It then ensures
// that a final Flush is performed to clear out any pending requests.
func (rc *requestBatcher) Close() {
	rc.close <- struct{}{}
	rc.Flush()
}

// Flush handles one request from each batch of requests and
// writes the same result to all requests in the batch. Finally, it
// deletes the batch of requests.
func (rc *requestBatcher) Flush() {
	rc.mu.Lock()
	defer rc.mu.Unlock()

	for key, batch := range rc.batch {
		fmt.Printf("%d batched requests under key %q\n",
			len(batch), key)

		// handle one candidate request
		candidateRequest := batch[0]

		w := httptest.NewRecorder()
		rc.handler.ServeHTTP(w, candidateRequest.r)
		responseBody := w.Body.Bytes()

		// write the same result to all requests in this batch
		for _, request := range batch {
			request.r.Body.Close()
			request.w.WriteHeader(w.Result().StatusCode)
			request.w.Write(responseBody)

			// let the goroutine for this request know that
			// it has been handled
			request.done <- struct{}{}
		}

		// delete the batch
		delete(rc.batch, key)
	}
}
