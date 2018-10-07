# Building Blocks of Web Distributed Systems

[Scalable Web Architecture and Distributed Systems](http://www.aosabook.org/en/distsys.html)

[_The Architecture of Open Source Applications_](http://www.aosabook.org/en/index.html)

> Architects look at thousands of buildings during their training, and study critiques of those buildings written by masters. In contrast, most software developers only ever get to know a handful of large programs well—usually programs they wrote themselves—and never study the great programs of history. As a result, they repeat one another's mistakes rather than building on one another's successes.[^1]

## A

[^1]: http://www.aosabook.org/en/index.html

---

# Notes (supplementary to blog post)

## My High-level Summary

This article asserts that the most challenging aspect of building web distributed systems is scaling access to the data. While app servers are inherently stateless and embody a shared-nothing architecture, "the heavy lifting is pushed down the stack to the database server and supporting services." The data access Layer is "where the real scaling and performance challenges come into play."

Caches, queues, indexes, and load balancers are the building blocks of a scalable web system.

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

> Imagine there is a request for the same data (let's call it littleB) across several nodes, and that piece of data is not in the cache. If that request is routed thought the proxy, then all of those requests can be collapsed into one, which means we only have to read littleB off disk once

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


