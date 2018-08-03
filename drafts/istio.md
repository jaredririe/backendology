# Up-and-coming Technology Review: Istio

## Prerequisite Terminology and Concepts

### [Load balancing]()

### Service mesh

Should I do a blog post on service mesh and another on load-balancing/proxy layers and then do the Istio post?

## Overview: what is Istio?

### Client-side and server-side load balancing

< Lucidchart diagram >

## Strengths

## Weaknesses

## Further reading

---

# Notes (supplementary to blog post)

## TODO

* Make Lucidchart diagram to visually explain the sidecar and connected control plane

## Outstanding Questions

* If I want to use Istio, do I need to rewrite my entire backend infrastructure or can I start small?

## What is Istio?

> This document introduces Istio: an open platform to connect, manage, and secure microservices. Istio provides an easy way to create a network of deployed services with load balancing, service-to-service authentication, monitoring, and more, without requiring any changes in service code. You add Istio support to services by deploying a special sidecar proxy throughout your environment that intercepts all network communication between microservices, configured and managed using Istio’s control plane functionality.

* *without* requiring any changes in service code
* New term: sidecar
* intercepts all network communication between microservices, configured and managed using Istio’s control plane functionality.

> Istio currently supports service deployment on Kubernetes, as well as services registered with Consul or Eureka and services running on individual VMs

### Data Plane and Control Plane

> The data plane is composed of a set of intelligent proxies (Envoy) deployed as sidecars that mediate and control all network communication between microservices, along with a general-purpose policy and telemetry hub (Mixer).

> The control plane is responsible for managing and configuring proxies to route traffic, and configuring Mixers to enforce policies and collect telemetry.

* Kubernetes, Consul, Eureka

## Why use Istio?

> addresses many of the challenges faced by developers and operators as monolithic applications transition towards a distributed microservice architecture

### "Service Mesh"

> The term service mesh is often used to describe the network of microservices that make up such applications and the interactions between them. As a service mesh grows in size and complexity, it can become harder to understand and manage. Its requirements can include discovery, load balancing, failure recovery, metrics, and monitoring, and often more complex operational requirements such as A/B testing, canary releases, rate limiting, access control, and end-to-end authentication.

Istio itself is not a service mesh. A service mesh is a network of microservices. Istio solves the problems that are inherent in a service mesh and helps manage complexity. These are some of the problems faced by services meshes that can be managed through Istio:

* discovery
* load balancing
* failure recovery
* metrics
* monitoring
* A/B testing
* canary releases
* rate limiting
* access control
* end-to-end authentication

### Key capabilities of Istio

> Traffic Management. Control the flow of traffic and API calls between services, make calls more reliable, and make the network more robust in the face of adverse conditions.
>
> Service Identity and Security. Provide services in the mesh with a verifiable identity and provide the ability to protect service traffic as it flows over networks of varying degrees of trustability.
>
> Policy Enforcement. Apply organizational policy to the interaction between services, ensure access policies are enforced and resources are fairly distributed among consumers. Policy changes are made by configuring the mesh, not by changing application code.
>
> Telemetry. Gain understanding of the dependencies between services and the nature and flow of traffic between them, providing the ability to quickly identify issues.

> These capabilities greatly decrease the coupling between application code, the underlying platform, and policy. This decreased coupling not only makes services easier to implement, but also makes it simpler for operators to move application deployments between environments or to new policy schemes.
