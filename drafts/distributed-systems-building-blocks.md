# Building Blocks of Distributed Systems

This blog post is based a chapter from [_The Architecture of Open Source Applications_](http://www.aosabook.org/en/index.html) titled "[Scalable Web Architecture and Distributed Systems](http://www.aosabook.org/en/distsys.html)."

## _The Architecture of Open Source Applications_

[The Architecture of Open Source Applications cover][]

Before getting into the details of the chapter, the book itself deserves some introduction. Its opening pages make the compelling point that architects in the traditional sense are exposed to and study thousands of real buildings, but software architects rarely make a similar investment, leading to repeated mistakes:

> Architects look at thousands of buildings during their training, and study critiques of those buildings written by masters. In contrast, most software developers only ever get to know a handful of large programs well—usually programs they wrote themselves—and never study the great programs of history. As a result, they repeat one another's mistakes rather than building on one another's successes.[^1]

Each chapter is written by experienced software developers who impart knowledge of a particular system or design. Some other chapters include [nginx](http://www.aosabook.org/en/nginx.html), [Firefox Release Engineering](http://www.aosabook.org/en/ffreleng.html), and [Git](http://www.aosabook.org/en/git.html). The chapter covered in this post primarily uses an image hosting application to explain the principles and building blocks of scalable distributed systems.

## Building blocks

[Building blocks][]

After explaining the general principles, the author asserts that the most challenging aspect of building web distributed systems is scaling access to the data. While application servers are inherently stateless and embody a shared-nothing architecture, "the heavy lifting is pushed down the stack to the database server and supporting services." The data access layer is "where the real scaling and performance challenges come into play."[^2]

Caches, proxies, indexes, load balancers, and queues are the building blocks of a scalable data access layer. Rather than covering the entire chapter, I will focus the remainder of this post on these building blocks.

### Caches

Caches are ubiquitous in computing. Their ability to scale read access in a system is clear. They "take advantage of the locality of reference principle: recently requested data is likely to be requested again."[^2]

In a [previous blog post](https://backendology.com/2018/08/27/multiple-layers-caching/), I wrote at length about the importance of having multiple layers of caching, including client-side caching. The author of this chapter reached a similar conclusion:

> Caches can exist at all levels in architecture, but are often found at the level nearest to the front end, where they are implemented to return data quickly without taxing downstream levels.[^2]

When clients avoid "taxing downstream levels", they enable more growth in the system without the need to scale out. For example, assuming linear scaling and equally taxing requests, if clients implement caching that reduces their usage of the backend by 50%, then the backend can handle twice as many clients without purchasing more resources.

The chapter's coverage of caching augments my previous post with a fascinating discussion of cache placement.

#### Cache placement

[Cache placement strategies][]

* _Request Node_: collocate the cache with the node that requests the data
    - Pros
        + Each time a request is made, the node can quickly return cached data if it exists, avoiding any network hops
        + Often in-memory and very fast
    - Cons
        + When you have multiple request nodes that are load balanced, you may have to cache the same item on all the nodes
* _Global Cache_: central cache used by all request nodes
    - Pros
        + A given item will only be cached only once
        + Multiple requests for an item can be _collapsed_ into one request to the backend
    - Cons
        + Easy to overwhelm a single cache as the number of clients and requests increase
    - Types
        + Reverse proxy cache: cache is responsible for retrieval on cache miss (more common, handles its own eviction)
        + Cache as a service: request nodes are responsible for retrieval on cache miss (typically used when the request nodes understand the eviction strategy or hot spots better than the cache)
* _Distributed Cache_: each of the nodes that make up the cache own part of the cached data; divided using a [consistent hashing function](https://en.wikipedia.org/wiki/Consistent_hashing)
    - Pros
        + Cache space and load capacity can be increased by scaling out (increasing the number of nodes)
    - Cons
        + Node failure must be handled or intentionally ignored

### Proxies

> At a basic level, a proxy server is an intermediate piece of hardware/software that receives requests from clients and relays them to the backend origin servers. Typically, proxies are used to filter requests, log requests, or sometimes transform requests (by adding/removing headers, encrypting/decrypting, or compression).[^2]

Proxies are a deceptively simple building block in an architecture: their very nature is to be lightweight, nearly invisible components yet they can provide incredible value to a system by reducing load on the backend servers, providing a convenient location for caching layers, and funneling traffic appropriately.

#### Collapsed forwarding

Collapsed forwarding is an example of a technique that proxies can employ to decrease load on a downstream server. In this technique, similar requests are _collapsed_ into a single request that is made to the downstream server; the result of this request is then written to all similar requests, thus reducing the number of requests made.

A simple example of collapsed forwarding is **deduplication**. If a resource X is 100 times, the proxy can make a single request to retrieve X from the downstream server and then write the same response body to the 99 other requests for X.

This is particularly helpful for the downstream server when the resource X is large in size. Let's assume a 5 MB payload that must be read into memory (rather than streamed). Without deduplication, the hundred requests would require the server to wastefully read 5 * 100 = 500 MB into memory. The deduplication step in the proxy can smooth over spikes and reduce the memory usage dramatically.

Let's implement a simple proxy and server in Go with collapsed forwarding!

Our implementation of collapsed forwarding will batch together requests for the same URL. Every five seconds, our `requestBatcher` will flush its batches of requests by handling one request from each batch and giving the same result to the rest of the requests in that batch.

```go
type requestBatcher struct {
    batch   map[string][]*request
    handler http.HandlerFunc

    mu    *sync.Mutex
    close chan struct{}
}

type request struct {
    w    http.ResponseWriter
    r    *http.Request
    done chan struct{}
}
```

`requestBatcher` will store the batches of requests as a map from string (URL) to a slice of requests. The `handler` will indicate how the requests should be handled. We'll protect our map with a mutex `mu` and have a `close` channel for a clean shutdown.

A `request` is everything we need to process the request, the original request `r` and the writer `w` for writing the response. We'll also have a `done` channel that will allow the flush step to tell the goroutine handling the request that the request has been handled.

```go
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
```

The constructor for the `requestBatcher` will initialize all the variables and kick off a goroutine that calls Flush every five seconds. For a real production proxy, five seconds is likely far too long. It will give us enough time to see that our batching logic is working, however.

```go
func (rc *requestBatcher) Append(key string, request *request) {
    rc.mu.Lock()
    defer rc.mu.Unlock()

    rc.batch[key] = append(rc.batch[key], request)
}
```

`Append` appends the given request to the slice of requests under the given key.

```go
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
```

`Flush` handles one request from each batch of requests and writes the same result to all requests in the batch. Finally, it deletes the batch of requests so we can start fresh after each flush.

```go
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
```

Our main function does the work of creating a new `requestBatcher`, standing up our simple proxy and server, and waiting for an interrupt.

* The proxy runs on `localhost:8080` and will handle requests by Appending them to the batch of requests and waiting on the `done` channel for completion. If we didn't wait for this channel, the request's goroutine would exit and leave the response unwritten.
* The server runs on `localhost:8081`. It's pretty straightforward: it handles requests by writing the response to be the URL of the request, like an echo server.
* `proxyRequest` is the handler that we pass to `newRequestBatcher`. It tells the batcher what it should do with requests when flushing them. In this case, we're indicating that requests should be made to the server at `localhost:8081`.

The entirety of this code is [available on GitHub](https://github.com/jaredririe/backendology/tree/master/code/collapsed-forwarding).

If we make several requests to the proxy in the background, like so:

```bash
$ curl localhost:8080/test &
$ curl localhost:8080/test &
$ curl localhost:8080/test2 &
$ curl localhost:8080/test &
```

The application logs each request, finds two sets of batched requests, and makes a single request to the server for each batch:

```bash
[proxy] /test
[proxy] /test
[proxy] /test2
[proxy] /test
3 batched requests under key "/test"
[server] /test
1 batched requests under key "/test2"
[server] /test2
```

#### Reverse proxy cache

A reverse proxy cache is as the name implies the combination of a proxy and cache. Requests are made to a proxy in front of an **origin server** which performs best-effort caching. It always reserves the right to fall back on the origin for a definitive response, which is a convenient property that makes failure scenarios relatively straightforward.

A less straightforward problem is handling cache eviction. Let's consider a few options:

##### Automatic expiration after a TTL

This option works well with in-memory caches that are intended to protect against spikes of requests. The data isn't cached very long, so its usefulness can be limited, however.

##### Intercept modifications and handle evictions

If all modifications to the underlying data go through the proxy layer, the cached data can be evicted as needed.

##### Only cache immutable data

In cases where modifications cannot be intercepted or the cached data is a computed result, more advanced techniques must be used. One such technique is to only cache unchanging, immutable data that never becomes stale or needs eviction. While this may seem impractical, it's usually not.

Let's say that you want to cache the result of running a query against some data. If you run the query today, you get five rows of data back. If you run it tomorrow, you get seven because new data arrived. The query result is therefore **mutable**. How can we make it immutable? What we'll do is store the data under a cache key computed like so:`hash(resource identifier, hash(query string), timestamp)`

* **resource identifier**: the ID that references the data, like a customerId or datasetId within a customer's account
* **query string**: the string that identifies the query, perhaps provided as a query parameter in the request's URL or a JSON representation in the request's body
* **timestamp**: the last updated time of the data stored under the resource identifier

Assuming the table like the following, if we make a query for all data stored under A, we'll cache the response (`[1, 2, 3]`) under the cache key `hash(A, hash(query string), 1539322037479)`. Then subsequent requests will only be cache hits if the data has not changed.

| ID  | LastUpdated   | Data          |
|-----|---------------|---------------|
| A   | 1539322037479 | [1, 2, 3]     |
| B   | 1538431688314 | [8, 2, 3, 1]  |
| C   | 1537899135166 | [1, 10, 1]    |
| D   | 1538116563215 | [10, 9, 8, 7] |

Using this technique works best when the consumer provides a `LastUpdated` value as part of their request. Preferably they retrieved this value once and will use it across multiple queries (to populated a dashboard, for example). If `LastUpdated` is not passed in on the request, the proxy can quickly retrieve it in the consumer's behalf and use it to check the cache. Usually it's much easier to get the `LastUpdated` value than compute a (potentially complex) query, so the caching layer still provides a lot of value.

### Indexes

When most developers hear the word "indexes", they immediately follow an index jump to database indexes. At least this is the case for me. While I find databases indexes to be an interesting topic (to the point that I wrote a [blog post which describes how database indexes work at a low level](https://backendology.com/2018/07/23/database-indexes/)), this chapter's explanation helped broaden my understanding of indexes beyond databases.

> Using an index to access your data quickly is a well-known strategy for optimizing data access performance; probably the most well known when it comes to databases. An index makes the trade-offs of increased storage overhead and slower writes (since you must both write the data and update the index) for the benefit of faster reads. ... Just as to a traditional relational data store, you can also apply this concept to larger data sets.[^2]

(Multiple layers of indexes)

### Load balancers

> Load balancers are a principal part of any architecture, as their role is to distribute load across a set of nodes responsible for servicing requests. This allows multiple nodes to transparently service the same function in a system.[^2]

Like caches, load balancers are placed in many strategic places throughout an architecture. They are also implemented in a variety of ways. There are several comparisons to be aware of in this space:

#### Software and hardware

Load balancers can be implemented either in software or hardware. A common commercial hardware offering is [F5](https://www.f5.com/) while [HAProxy](http://www.haproxy.org/) is best known on the software side.

#### Layer 4 and Layer 7

> Load balancers are generally grouped into two categories: Layer 4 and Layer 7. Layer 4 load balancers act upon data found in network and transport layer protocols (IP, TCP, FTP, UDP). Layer 7 load balancers distribute requests based upon data found in application layer protocols such as HTTP.[^3]

#### North-south and east-west

**North-south traffic** is client to server traffic that originates outside of the datacenter (e.g. edge firewalls, routers). **East-west traffic** is server to server traffic originates is internal traffic within a datacenter (e.g. LAN connection between microservices in a Microservices Architecture.

Many businesses stand up a hardware load balancer at the edge of their datacenters and then use software load balancing for communication within each datacenter. These additional layers of software load balancing avoid the need to return back to the edge of the network to distribute load to a downstream service.

#### Client-side, server-side, and service mesh

Traditional load balancing strategies encourage either the client or the server to take responsibility for load balancing. The client might ensure that it properly sends traffic to a server in a distributed manner. A server, on the other hand, could protect itself with a reverse proxy layer that offers load balancing.

When both clients and servers are part of the same service mesh, clients and servers do not directly involve themselves in load balancing. Instead, calls from the client to the server are transparently load balanced at the cost of some additional latency for the service mesh to distributed the load. Service meshes like [Istio](https://istio.io/) are gaining traction as they can provide load balancing, automatic retries, and other helpful features without direct participation from the involved services.

### Queues

Unlike proxies and load balancers which augment an existing architecture and _scale reads_, queues have a more dramatic impact on the data flow of the architecture and _scale writes_. Queues have this impact by forcing the introduction of **asynchronous processing**.

While a synchronous system tightly couples a request to its immediate response, an asynchronous system separates the two. This is achieved by having clients provide a work request to the queue which is not immediately processed while the clients waits. "While the client is waiting for an asynchronous request to be completed it is free to perform other work, even making asynchronous requests of other services."

In a synchronous system where clients are actively waiting for responses, service outages and intermittent failures are exposed directly to clients. High availability is difficult to provide, especially when the underlying database(s) are under high load and requests time out. Due to the asynchronous nature of queues, they can provide protection from failed requests as they can easily retry requests. This takes away the stress of ensuring that every single request succeeds at the cost of great engineering effort. Retry logic is also much easier to implement in asynchronous processing, avoiding the need for "complicated and often inconsistent client-side error handling."

This added protection from a lack of availability in a downstream service and improved retry logic makes a strong argument for the introduction of more queues into an architecture. The client of a queue can often be unaware that a downstream service was temporarily unavailable.

[^1]: http://www.aosabook.org/en/index.html
[^2]: http://www.aosabook.org/en/distsys.html
[^3]: https://www.f5.com/services/resources/glossary/load-balancer

---

# Notes (supplementary to blog post)

## 1.1. Principles of Web Distributed Systems Design

> This chapter is largely focused on web systems, although some of the material is applicable to other distributed systems as well.

> When designing any sort of web application it is important to consider these key principles, even if it is to acknowledge that a design may sacrifice one or more of them.

* Availability
    - Uptime
    - > requires the careful consideration of redundancy for key components, rapid recovery in the event of partial system failures, and graceful degradation when problems occur
* Performance
    - User satisfaction
    - Usability
    - Search engine rankings
* Reliability
    - Request for data will consistently return the same data
* Scalability
    - Ability to increase capacity to handle greater amounts of load
* Manageability
    - Scalability of operations: maintenance and updates
    - Ease of diagnosing and understanding problems when they occur
* Cost
    - Hardware and software cost
    - Amount of developer time the system takes to build
    - Amount of operational effort required to run the system
    - Amount of training required

## 1.2. The Basics

> what are the right pieces, how these pieces fit together, and what are the right tradeoffs

> This section is focused on some of the core factors that are central to almost all large web applications: services, redundancy, partitions, and handling failure. Each of these factors involves choices and compromises, particularly in the context of the principles described in the previous section. In order to explain these in detail it is best to start with an example.

### Services

> When considering scalable system design, it helps to decouple functionality and think about each part of the system as its own service with a clearly defined interface. In practice, systems designed in this way are said to have a Service-Oriented Architecture (SOA).

> Another potential problem with this design is that a web server like Apache or lighttpd typically has an upper limit on the number of simultaneous connections it can maintain (defaults are around 500, but can go much higher) and in high traffic, writes can quickly consume all of those. Since reads can be asynchronous, or take advantage of other performance optimizations like gzip compression or chunked transfer encoding, the web server can switch serve reads faster and switch between clients quickly serving many more requests per second than the max number of connections (with Apache and max connections set to 500, it is not uncommon to serve several thousand read requests per second). Writes, on the other hand, tend to maintain an open connection for the duration for the upload, so uploading a 1MB file could take more than 1 second on most home networks, so that web server could only handle 500 such simultaneous writes.

How can reads be asynchronous? Is low latency not a goal?

What is chunked transfer encoding?

The conclusion from this paragraph is that splitting out reads and writes is beneficial. I have never considered this for the CRUD storage services that I have written. Should I have? Or is this scenario with large objects different than what I encountered before? As Go does not have the problem of a being restricted to a fixed number of open connections, is splitting still valuable?

> The advantage of this approach is that we are able to solve problems independently of one another—we don't have to worry about writing and retrieving new images in the same context. Both of these services still leverage the global corpus of images, but they are free to optimize their own performance with service-appropriate methods (for example, queuing up requests, or caching popular images—more on this below). And from a maintenance and cost perspective each service can scale independently as needed, which is great because if they were combined and intermingled, one could inadvertently impact the performance of the other as in the scenario discussed above.

Flickr solves the connection problem in a different way:

> Flickr solves this read/write issue by distributing users across different shards such that each shard can only handle a set number of users, and as users increase more shards are added to the cluster

There are tradeoffs to consider. See below (where former is splitting out reads and writes):

> In the former an outage or issue with one of the services brings down functionality across the whole system (no-one can write files, for example), whereas an outage with one of Flickr's shards will only affect those users. In the first example it is easier to perform operations across the whole dataset—for example, updating the write service to include new metadata or searching across all image metadata—whereas with the Flickr architecture each shard would need to be updated or searched (or a search service would need to be created to collate that metadata—which is in fact what they do).

### Redundancy

Storing multiple copies of the data and versions of a given service improves redundancy and eliminates single points of failures.

> This same principle also applies to services. If there is a core piece of functionality for an application, ensuring that multiple copies or versions are running simultaneously can secure against the failure of a single node.

Shared-nothing architecture

> Another key part of service redundancy is creating a shared-nothing architecture. With this architecture, each node is able to operate independently of one another and there is no central "brain" managing state or coordinating activities for the other nodes. This helps a lot with scalability since new nodes can be added without special conditions or knowledge. However, and most importantly, there is no single point of failure in these systems, so they are much more resilient to failure.

### Partitions

Scaling vertically vs. scaling horizontally

> There may be very large data sets that are unable to fit on a single server. It may also be the case that an operation requires too many computing resources, diminishing performance and making it necessary to add capacity. In either case you have two choices: scale vertically or horizontally.
> ...
> Scaling vertically means adding more resources to an individual server
> ...
> To scale horizontally, on the other hand, is to add more nodes. In the case of the large data set, this might be a second server to store parts of the data set, and for the computing resource it would mean splitting the operation or load across some additional nodes.

Challenge: Data locality

> Of course there are challenges distributing data or functionality across multiple servers. One of the key issues is data locality; in distributed systems the closer the data to the operation or point of computation, the better the performance of the system.

Challenge: Inconsistency

> When there are different services reading and writing from a shared resource, potentially another service or data store, there is the chance for race conditions—where some data is supposed to be updated, but the read happens prior to the update—and in those cases the data is inconsistent.

## 1.3. The Building Blocks of Fast and Scalable Data Access

Most simple web applications look like this:

Internet -> App server -> Database server

App servers are written to be shared-nothing and stateless, making them horizontally scalable. Scaling the database server, on the other hand, is a real challenge:

> the heavy lifting is pushed down the stack to the database server and supporting services; it's at this layer where the real scaling and performance challenges come into play

> memory access is as little as 6 times faster for sequential reads, or 100,000 times faster for random reads, than reading from disk

### Caches

> Caches take advantage of the locality of reference principle: recently requested data is likely to be requested again. They are used in almost every layer of computing: hardware, operating systems, web browsers, web applications and more.

The author came to a similar conclusion that I did in my blog post on having multiple layers of caching:

> Caches can exist at all levels in architecture, but are often found at the level nearest to the front end, where they are implemented to return data quickly without taxing downstream levels.

Cache placement

* Request node
    - Enables local storage of response data
    - Each time a request is made, the node can quickly return cached data if it exists
    - Often in-memory and very fast
    - Drawback: when you have multiple request nodes that are load balanced, you may have to cache the same item on all the nodes
* Global Cache
    - All nodes use the same single cache space
    - Drawback: easy to overwhelm a single cache as the number of clients and requests increase
    - Very effective in some architectures (like those with specialized hardware to make the global cache very fast)
    - Types
        + Global cache where cache is responsible for retrieval (reverse proxy cache)
            * More common
            * Cache itself manages eviction and fetching data to prevent flood of requests for the same data from the clients
        + Global cache where request nodes are responsible for retrieval (cache as a service)
            * Makes more sense for large files where a low cache hit percentage would cause the cache buffer to become overwhelmed with cache misses
            * Also used when the application logic understands the eviction strategy or hot spots better than the cache
* Distributed Cache
    - Each of the nodes that make up the cache own part of the cached data
    - Typically divided up using a consistent hashing function (if a request node is looking for a certain piece of data, it can quickly know where to look within the distributed cache to determine if it's cached)
    - Increase cache space simply by adding new nodes
    - Drawback: remedying a missing node
        + Could store multiple copies
        + Could accept the cache misses (the requests will just pull from the origin)

> Caches are wonderful for making things generally faster, and moreover provide system functionality under high load conditions when otherwise there would be complete service degradation.

### Proxies

Proxy definition

> a proxy server is an intermediate piece of hardware/software that receives requests from clients and relays them to the backend origin servers. Typically, proxies are used to filter requests, log requests, or sometimes transform requests (by adding/removing headers, encrypting/decrypting, or compression).

Collapsed forwarding

> One way to use a proxy to speed up data access is to collapse the same (or similar) requests together into one request, and then return the single result to the requesting clients. This is known as collapsed forwarding.

Deduplicating requests for the same value is an example of collapsed forwarding:

> Imagine there is a request for the same data (let's call it littleB) across several nodes, and that piece of data is not in the cache. If that request is routed through the proxy, then all of those requests can be collapsed into one, which means we only have to read littleB off disk once

Individual requests may experience more latency in order to accomplish collapsed forwarding because they are slightly delayed to be grouped with similar requests.

> But it will improve performance in high load situations, particularly when that same data is requested over and over. This is similar to a cache, but instead of storing the data/document like a cache, it is optimizing the requests or calls for those documents and acting as a proxy for those clients.

Proxies can also collapse requests for data that is spatially close together in the origin store. Perhaps you can read an entire block of data just as easily as a row within that block. The proxy could read the block of data and return it to several concurrent requests asking for parts of that block.

Proxies can also batch several requests into one, thus limiting the number of calls made to the origin.

Proxies work well with caches.

> It is worth noting that you can use proxies and caches together, but generally it is best to put the cache in front of the proxy, for the same reason that it is best to let the faster runners start first in a crowded marathon race. This is because the cache is serving data from memory, it is very fast, and it doesn't mind multiple requests for the same result. But if the cache was located on the other side of the proxy server, then there would be additional latency with every request before the cache, and this could hinder performance.

### Indexes

> Using an index to access your data quickly is a well-known strategy for optimizing data access performance; probably the most well known when it comes to databases. An index makes the trade-offs of increased storage overhead and slower writes (since you must both write the data and update the index) for the benefit of faster reads.

> Just as to a traditional relational data store, you can also apply this concept to larger data sets. The trick with indexes is you must carefully consider how users will access your data

> An index can be used like a table of contents that directs you to the location where your data lives.

Multiple layers of indexes

> Often there are many layers of indexes that serve as a map, moving you from one location to the next, and so forth, until you get the specific piece of data you want. (See Figure 1.17.)

Views, avoiding copies of the data

> Indexes can also be used to create several different views of the same data. For large data sets, this is a great way to define different filters and sorts without resorting to creating many additional copies of the data.

Inverse indexes

> First, inverse indexes to query for arbitrary words and word tuples need to be easily accessible; then there is the challenge of navigating to the exact page and location within that book, and retrieving the right image for the results. So in this case the inverted index would map to a location (such as book B), and then B may contain an index with all the words, locations and number of occurrences in each part.

Inverted index example:

> each word or tuple of words provide an index of what books contain them

Importance of indexes for big data problems

> Creating these intermediate indexes and representing the data in smaller sections makes big data problems tractable. Data can be spread across many servers and still accessed quickly. Indexes are a cornerstone of information retrieval, and the basis for today's modern search engines.

### Load Balancers

> Load balancers are a principal part of any architecture, as their role is to distribute load across a set of nodes responsible for servicing requests. This allows multiple nodes to transparently service the same function in a system. (See Figure 1.18.) Their main purpose is to handle a lot of simultaneous connections and route those connections to one of the request nodes, allowing the system to scale to service more requests by just adding nodes.

Load balancing algorithms

> There are many different algorithms that can be used to service requests, including picking a random node, round robin, or even selecting the node based on certain criteria, such as memory or CPU utilization.

Multiple layers of load balancing

> In a distributed system, load balancers are often found at the very front of the system, such that all incoming requests are routed accordingly. In a complex distributed system, it is not uncommon for a request to be routed to multiple load balancers as shown in Figure 1.19.

Reverse proxies

> Like proxies, some load balancers can also route a request differently depending on the type of request it is. (Technically these are also known as reverse proxies.)

Sometimes round robin DNS is sufficient

> If a system only has a couple of a nodes, systems like round robin DNS may make more sense since load balancers can be expensive and add an unneeded layer of complexity.

### Queues

> So far we have covered a lot of ways to read data quickly, but another important part of scaling the data layer is effective management of writes.

Accepting the tradeoffs of asynchronous processing allows improvements to performance and availability

> In the cases where writes, or any task for that matter, may take a long time, achieving performance and availability requires building asynchrony into the system; a common way to do that is with queues.

Synchronous systems behave poorly under high load. Failed requests cause cascading failures.

> However, when the server receives more requests than it can handle, then each client is forced to wait for the other clients' requests to complete before a response can be generated.
>
> This kind of synchronous behavior can severely degrade client performance; the client is forced to wait, effectively performing zero work, until its request can be answered. Adding additional servers to address system load does not solve the problem either; even with effective load balancing in place it is extremely difficult to ensure the even and fair distribution of work required to maximize client performance. Further, if the server handling requests is unavailable, or fails, then the clients upstream will also fail.

API of queues

> When a client submits task requests to a queue they are no longer forced to wait for the results; instead they need only acknowledgment that the request was properly received. This acknowledgment can later serve as a reference for the results of the work when the client requires it.

Queues provide a separation between the request and response, rather than tightly integrating the two

> In an asynchronous system the client requests a task, the service responds with a message acknowledging the task was received, and then the client can periodically check the status of the task, only requesting the result once it has completed. While the client is waiting for an asynchronous request to be completed it is free to perform other work, even making asynchronous requests of other services.

Protection from service outages and failures, providing a much improved client experience--they're not directly exposed to a struggling synchronous system.

> Queues also provide some protection from service outages and failures. For instance, it is quite easy to create a highly robust queue that can retry service requests that have failed due to transient server failures. It is more preferable to use a queue to enforce quality-of-service guarantees than to expose clients directly to intermittent service outages, requiring complicated and often-inconsistent client-side error handling.

