# Microservices: Are We Making a Huge Mistake?

There is a clear trend in the software industry moving away from large, monolithic systems to fine-grained services in the form of "microservices." While compelling, microservices introduce their own set of challenges and fallacies. This post considers the benefits and drawbacks of a microservices architecture (MSA) and contemplates the question: are we making a huge mistake in adopting this kind of architecture?

## Relationship to "distributed systems"

First, let's clear up some terminology that you may find confusing. A distributed system is a general term which refers to a network of independent computers that form a coherent system. A microservices architecture, on the other hand, is a software development technique for building software systems that _run on a distributed system_. This kind of architecture encourages fine-grained services with lightweight, easy-to-use communication protocols like HTTP/S and RPC.

Building a microservices architecture implies utilizing a distributed system[^3]. This is *not* an implication to be taken lightly. The foundation of a distributed system has several key fallacies that must be understood in order to build a reliable system!

### Fallacies

L. Peter Deutsch and others at Sun Microservices stated eight false assumptions that new programmers have regarding distributed systems:

> 1. The network is reliable.
> 2. Latency is zero.
> 3. Bandwidth is infinite.
> 4. The network is secure.
> 5. Topology doesn't change.
> 6. There is one administrator.
> 7. Transport cost is zero.
> 8. The network is homogeneous.[^4]

#### Never assume the network is reliable.

When writing code that calls an external service, remember to handle failure scenarios. What happens if your request never makes it to the other service? Will you retry, propagate the failure upwards, log an error, or do nothing? If you choose to retry, at what rate will you retry and for how long? What if your request is making it to the external service and being processed but the response is never making it back to you?

When a system is a traditional monolith, the external service is just another part of the monolith so you’re just making a function call, not a network call; if you’re up, they’re up. The failure modes are limited.

#### Network latency is not zero.

Although it takes time to make a function call, it’s so minimal that for most purposes it can be considered zero latency. A network call on the other hand will take a nontrivial amount of time (maybe 2ms). When moving to microservices, it’s important to be mindful when crossing service boundaries. Perhaps your algorithm used to make a few thousand function calls in a monolith. For MSA, the system was split into services such that your function calls become network calls. This does not impact the Big-O runtime of the algorithm but it does change the constant! Could the result of the call be cached? Could the number of calls be reduced?



Topology doesn’t change

### Change of mindset

Working with distributed systems requires a change of mindset. Most of us learned how to program by writing simple command-line applications. Perhaps your first program was the iconic “Hello, world” and then you progressed to code that could read in a file, make some changes and spit out another file. We develop a mental model of how our code executes with the simplified assumptions that can be made in a **single computer environment**. Once a system of computers must solve a problem *together*, the assumptions we could make before turn into misconceptions.

Communicating with another computer is fundamentally different. The network which connects the computers is not reliable. Requests will take some non-zero amount of time to travel to and be processed by the other computer. Technically a function call within a computer is also a non-zero amount of time, but in the context of a networked system, we can disregard it (the optimizer may choose to in line the code in the function, literally removing it)

## Service mesh pattern



Working within a microservices architecture (MSA) requires a changed mindset.

## Benefits

## Drawbacks

### Network latency

> Inter-service calls over a network have a higher cost in terms of network latency and message processing time than in-process calls within a monolithic service process.[^2]

### Complexity / Cognitive Load

> The architecture introduces additional complexity and new problems to deal with, such as network latency, message formats, load balancing and fault tolerance.[^2]

The first few microservices which are built will likely require modernizations to the infrastructure to handle problems like service discovery, independent deployability, etc. These improvements, though difficult, push the architecture in the right direction. Often, the first services to implement are carefully deliberated and their need is clear. An example of this is a reverse proxy cache that sits in front of a monolithic service.

As more microservices are added, however, the entanglement begins, raising the cognitive load required to understand the details of how the system really works. The flow of requests through the system is challenging and convoluted.

A -> B -> C -> D
A -> D
A -> D -> C -> E

The need for orchestration, service-to-service software load balancing, and improved fault tolerance becomes evident. The service mesh is one response to the growing complexity of operating microservices at scale.

### Nanoservices

> Too-fine-grained microservices have been criticized as an anti-pattern, dubbed a nanoservice by Arnon Rotem-Gal-Oz.[^2]

### Boundaries, coordinating efforts

> Dividing the backend into 8 separate services followed by a decision to assign services to people enforced ownership of specific services by developers. This led to developers complaining of their service being blocked by tasks on other services and refusing to help out by working with these blocking tasks. The separation also lead to developers losing sight of the overarching system goals, instead only focusing on the service they were working on.[^1]

## Conclusion: Are We Making a Huge Mistake?

## Further Reading

* Wikipedia:

[^1]: https://www.infoq.com/news/2014/08/failing-microservices
[^2]: https://en.wikipedia.org/wiki/Microservices
[^3]: https://www.infoq.com/news/2016/09/microservices-distributed-system
[^4]: https://en.wikipedia.org/wiki/Fallacies_of_distributed_computing

# Notes (Supplementary to blog post)

## Service Mesh

> Microservices architectures have given rise to many approaches for making them work well at varying degrees of scale. One widely discussed pattern for large-scale microservices development and delivery is the service mesh pattern.[39]
>
> In a service mesh, each service instance is paired with an instance of a reverse proxy server, called a service proxy, sidecar proxy, or sidecar. The service instance and sidecar proxy share a container, and the containers are managed by a container orchestration tool such as Kubernetes.[40]
>
> The service proxies are responsible for communication with other service instances and can support capabilities such as service (instance) discovery, load balancing, authentication and authorization, secure communications, and others.
>
> In a service mesh, the service instances and their sidecar proxy are said to make up the data plane, which includes not only data management but also request processing and response. The service mesh also includes a control plane for managing the interaction between services, mediated by their sidecar proxies. The most widely discussed service mesh architecture today is Istio, a joint project among Google, IBM, and Lyft.

## Books

*Building Microservices*: helped me make sense of the wonderful and horrific world of microservices

* Service discovery and DNS
* Monitoring
* Continuous integration
* Splitting monolithic systems
* Building resilient systems with partial degradation of functionality
* Asynchronous event-based collaboration

*Designing Data-Intensive Applications*: Rigorous explanations of concepts within distributed and database systems that will give you depth in these areas. This book could be explained a combination of Building Microservices and NoSQL Distilled, supplemented with a whole lot more goodness.

## Distributed System

Many backend systems are an entanglement of services which together can be considered a distributed system. Rather than being monolithic in nature, these systems benefit from things like **independent deployability**, **focused development**, and **loose coupling** yet suffer from the complexities of **partial failure**, **lack of concurrency guarantees**, and **network boundaries**.

This list of tradeoffs is clearly incomplete. Each complexity deserves its own blog post! Once network boundaries exist between components of a system, for example, problems such as discoverability must be dealt with. No longer can I make an easy function call; rather, I must build in some way for one service to find another service and then make a network call. This call could fail (should I retry?) or time out (how long should I wait?). Once I've figured out how to handle these failure scenarios, does *every* service I write need to duplicate this logic or can I abstract it away through a form of middleware?

Topics like caching, consistency, reverse proxying, API gateways, and service meshes are fair game.
