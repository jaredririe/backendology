package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"os/signal"
)

func main() {
	rc := newRequestBatcher(proxyRequest)
	proxy := newProxy(rc)
	server := newServer()

	defer func() {
		proxy.Close()  // stop accepting new requests at the proxy layer
		rc.Close()     // close and flush out any pending requests
		server.Close() // stop the server
	}()

	// run until interrupted
	stop := make(chan os.Signal)
	signal.Notify(stop, os.Interrupt)
	<-stop
}

func newProxy(rc *requestBatcher) *http.Server {
	proxy := &http.Server{
		Addr: ":8080",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("[proxy]", r.URL.String())

			ch := make(chan struct{})
			rc.Append(r.URL.String(), &request{
				w:    w,
				r:    r,
				done: ch,
			})

			<-ch
		}),
	}

	// start the proxy
	go func() {
		fmt.Println("[proxy] running on port", proxy.Addr)
		fmt.Println(proxy.ListenAndServe())
	}()

	return proxy
}

func newServer() *http.Server {
	server := &http.Server{
		Addr: ":8081",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("[server]", r.URL.String())
			w.Write([]byte(r.URL.String()))
		}),
	}

	// start the server
	go func() {
		fmt.Println("[server] running on port", server.Addr)
		fmt.Println(server.ListenAndServe())
	}()

	return server
}

func proxyRequest(w http.ResponseWriter, r *http.Request) {
	u := url.URL{
		Scheme: "http",
		Host:   "localhost:8081",
		Path:   r.URL.Path,
	}

	request, err := http.NewRequest(r.Method, u.String(), r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Println(err)
		return
	}

	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	response.Body.Close()

	w.WriteHeader(response.StatusCode)
	w.Write(bytes)
}
