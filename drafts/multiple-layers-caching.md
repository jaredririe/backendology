# Multiple Layers of Caching

I'm a huge fan of caching! Done right, it has the potential to make a impressive gains in the performance of a system.

This blog post isn't just about caching as a general concept, however. Specifically, I would like to cover the importance of having **multiple layers** of caching in a system.

## A caching anecdote

One of the systems I built at Qualtrics could be described as the "back of the backend" as it was a critical storage system that many services relied on, yet had no service dependencies itself. In this position, it was subject to any and all misuse by its consumers. It suffered from what I grew to call the **multiplier effect**, a condition where a single request (perhaps from a user clicking a button in a UI) caused a few hundred requests and for each one of those requests, issued several more requests, and so on, leading to potentially thousands of requests making it to the back of the backend. I recall determining that the cause of 17,000 requests to my service within a few minutes was caused by the click of an Export button.

Let's talk about those 17,000 requests some more. If each request was unique and necessary for the result to be produced, then so be it. We should scale our services to handle the load. Maybe throw in some rate limiting so we can throttle the traffic at the cost of increased latency. As you can probably guess in the context of a blog post on caching, however, these 17,000 requests were for a mere handful of objects _that rarely change_!

While our service handled the request load adequately in this case, my team became motivated to start tracking down obviously bad patterns like this to see what we could do to help our consumers utilize our service more efficiently. It was in this process that I came to understand a core principle: **performant systems require multiple layers of caching with participation from both clients and servers**.

## Reasons consumers chose not to cache

In talking to consumers about why they were not caching, we heard several reasons:

* Our service is just a stateless worker processing a job queue
* There isn't any way for us to know when we should invalidate our cache
* The backend already has a few layers of caching, so it's unnecessary

Let's talk through each of these cases and see if we can come up with a way to still introduce caching.

### Stateless worker

Example scenario: the client is pulling jobs off of a queue, doing some work, making requests to external services, and moving onto the next job. Each job starts with an allotted amount of memory and can’t reasonably cache things in that space.

Potential solution: make requests through a reverse proxy cache. This is a stateful cache that represents an external API. When a request is received, it determines whether a cached response is ready; if not, it makes the request to the external API and caches it for later. Not all requests will be cacheable, so this service acts as a best-effort caching layer that reserves the right to fall back on the "origin" external service.

### Invalidation concerns

This is a fair point. You are just a consumer so how will you know when the data changes and therefore should be invalidated.

Solution: configure your cache to expire after a given amount of time, known as TTL (time to live). If the cached data rarely changes, risk a long TTL. If consistency is of utmost importance and then data changes regularly, have a short TTL.

Consider an example API that returns a time value. If the value is years, it’s obviously highly cacheable. If it’s milliseconds, it’s not very cacheable. But even in the millisecond case, a short TTL cache could still make sense to give several requests a "snapshot" of the data. Requests made to populate a dashboard have this quality: would you rather make a slew of requests against a consistent view of the data, or have half of the requests run against a different version of the underlying data because it changed mid-load? The consistent view makes a lot of sense for most use cases, making a short (maybe 5-20 seconds) TTL cache very attractive.

### Redundant caching

The argument that "the server is caching, so why should the client?" is very common in my experience. Perhaps it's because we're taught as developers to avoid code duplication (_Don't Repeat Yourself_[^1]). One might argue that multiple teams implementing caching is a form of wasteful duplication of effort. My response is that it's certainly not wasteful and it's not even true duplication because client-side and server-side caching are potentially very different.

#### Memory hierarchy

My first in-depth introduction to the concept of caching was in the context of hardware in a Computer Architecture class. In this class, I learned about the fascinating topic of the memory hierarchy[^2] which I think has some striking parallels to caching in a distributed system.

The memory hierarchy organizes the various forms of computer storage (CPU caches, RAM, disks) according to response time. The registers in a processor are the fastest possible, usually requiring only a single CPU cycle to retrieve their contents. Next is the processor cache which itself has multiple levels (L0, L1, and so on). Access speed for L1 data cache is around 700 GB/s. Next is the hierarchy is main memory (RAM) with speeds of 10 GB/s. Then comes disk storage at 2000 MB/s. The bottom of the hierarchy varies depending on use case, but could be an cloud-based storage system or nearline storage which allows exabytes of data at an access speed of 160 MB/s.

Each storage system in the hierarchy can be thought of as a caching layer. Just like a computer system would be terribly slow and unusable without RAM, a distributed system without some layers of caching would likewise not perform optimally. What we're seeing here is the power of data locality. A RAM cache is much faster than a disk-based cache, but cache memory is much faster than a RAM cache because it's so close to the CPU! When we can get data close to the system that processing it, our throughput increases dramatically.

A key property of the memory hierarchy is that one of the main ways to increase system performance is to "minimize how far down the memory hierarchy one has to go to manipulate data."[^2] The same is true of a request as it flows through a system: if the data is cached in a client-side application, it doesn't have to make a few network hops to a backend system, wait for a database query to run, and propagate the response back up.

## The power of client-side caching

Have a better sense than the backend about workflows and where a short TTL cache could make a world of a difference. The backend can try to identify highly-requested objects and cache them proactively.

## Protect the backend through caching

> Google quote about making their APIs incredibly resilient rather than pushing clients

### Gotcha: the system died before we were able to start caching

Deduplicate request

> Google deduplication logic can be impressively paired with caching. Send one request to get the data and give it to everyone who is waiting, then cache it for later requests.

[^1]: https://en.wikipedia.org/wiki/Don%27t_repeat_yourself
[^2]: https://en.wikipedia.org/wiki/Memory_hierarchy

---

# Notes (supplementary to blog post)

## Memory hierarchy

> Cache memory: Random access memory, or RAM, that a microprocessor can access faster than it can access regular RAM. Cache memory is often tied directly to the CPU and is used to cache instructions that are frequently accessed. A RAM cache is much faster than a disk-based cache, but cache memory is much faster than a RAM cache because it's so close to the CPU.

> Disk cache: Holds recently read data and perhaps adjacent data areas that are likely to be accessed soon. Some disk caches cache data based on how frequently it's read. Frequently read storage blocks are referred to as hot blocks and are automatically sent to the cache.

https://en.wikipedia.org/wiki/Memory_hierarchy

> One of the main ways to increase system performance is minimizing how far down the memory hierarchy one has to go to manipulate data.[4]

- Processor registers – the fastest possible access (usually 1 CPU cycle). A few thousand bytes in size
- Processor cache
    - Level 0 (L0) Micro operations cache – 6 KiB [8] in size
    - Level 1 (L1) Instruction cache – 128 KiB in size
    - Level 1 (L1) Data cache – 128 KiB in size. Best access speed is around 700 GiB/second[9]
    - Level 2 (L2) Instruction and data (shared) – 1 MiB in size. Best access speed is around 200 GiB/second[9]
    - Level 3 (L3) Shared cache – 6 MiB in size. Best access speed is around 100 GB/second[9]
    - Level 4 (L4) Shared cache – 128 MiB in size. Best access speed is around 40 GB/second[9]
- Main memory (Primary storage) – Gigabytes in size. Best access speed is around 10 GB/second.[9] In the case of a NUMA machine, access times may not be uniform
- Disk storage (Secondary storage) – Terabytes in size. As of 2017, best access speed is from a consumer solid state drive is about 2000 MB/second [10]
- Nearline storage (Tertiary storage) – Up to exabytes in size. As of 2013, best access speed is about 160 MB/second[11]
- Offline storage

