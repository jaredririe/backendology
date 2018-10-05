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

### Indexes

### Load Balancers
