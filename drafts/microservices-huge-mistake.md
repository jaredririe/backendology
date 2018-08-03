# Microservices: Are We Making a Huge Mistake?

There is a clear trend in the software industry moving away from large, monolithic systems to fine-grained services in the form of "microservices." While compelling, microservices introduce their own set of challenges and fallacies. This post considers the benefits and drawbacks of a microservices architecture (MSA).

## Relationship to "distributed systems"

First, let's clear up some terminology that you may find confusing. A distributed system is a general term which refers to a network of independent computers that form a coherent system.  A microservices architecture, on the other hand, is a software development technique for building software systems that run on a distributed system. This kind of architecture encourages fine-grained services with lightweight, easy-to-use communication protocols like HTTP/S and RPC.

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

### Change of mindset

Working with distributed systems requires a change of mindset.

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
