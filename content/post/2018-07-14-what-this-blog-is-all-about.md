---
title: What This Blog is All About
featured_image: /public/images/topics-outline.png
author: Jared Ririe
categories: General
tags:
- microservices
- golang
- architecture
- databases
- nosql
date: "2018-07-14"
slug: what-this-blog-is-all-about
---

## "Backendology"?

> There are only two hard things in Computer Science: cache invalidation and naming things (Phil Karlton).

Or, my favorite variant:

> There are only two hard things in Computer Science: cache invalidation, naming things, and off by one errors (Unknown)

Naming is hard. Backendology isn't a real word. Like the programmer who decides to write a lengthy comment rather than coming up with a better name, I'll try to explain what I intended when I named this blog Backendology. The name draws obvious reference to the term "backend” as in backend or server-side web development. The suffix "-logy" is a branch of learning, or study of a particular subject. Put together, this blog is a detailed study of the concepts and technologies related to backend web development.

## Topics I plan to cover

While this list does not aim to be exhaustive, it should give a good sense for topics I will cover in this blog. I’ll also briefly elaborate on why each topic deserves attention.

### Distributed systems

Many backend systems are an entanglement of services which together can be considered a distributed system. Rather than being monolithic in nature, these systems benefit from things like **independent deployability**, **focused development**, and **loose coupling** yet suffer from the complexities of **partial failure**, **lack of concurrency guarantees**, and **network boundaries**.

This list of tradeoffs is clearly incomplete. Each complexity deserves its own blog post! Once network boundaries exist between components of a system, for example, problems such as discoverability must be dealt with. No longer can I make an easy function call; rather, I must build in some way for one service to find another service and then make a network call. This call could fail (should I retry?) or time out (how long should I wait?). Once I've figured out how to handle these failure scenarios, does *every* service I write need to duplicate this logic or can I abstract it away through a form of middleware?

Fascinating!

Topics like caching, consistency, reverse proxying, API gateways, and service meshes are fair game.

### Go

Go is a phenomenal language for backend development. I started writing Go in October 2015 and haven't looked back. I'll be writing several blog posts where I solve problems using Go or otherwise talk about it.

Go is easy to learn, proven in production, and designed for the cloud. It's deployed in production by high-traffic companies like Google, Dropbox, Uber, and Facebook in cases where stability and high performance are critical. Many open-source cloud computing tools are written in Go like Docker and Consul.

I enjoy its strong focus on simplicity which translates to less ramp-up time for new developers. Another benefit is in maintainability of Go codebases, as it’s very readable and easy to understand existing code.

### NoSQL and general database concepts

NoSQL is a movement that started in response to a need for increased scalability in large cloud companies like Google and Amazon. In my mind, NoSQL is less of a rebuttal of relational databases (i.e. *No, SQL!*) and more of an alternative to SQL when it makes sense for the problem being solved (i.e. *Not Only SQL*).

I have worked with a large variety of NoSQL databases in my time at Qualtrics. Some have turned out extraordinarily well while others turned out quite the opposite! These experiences left me with this conclusion: **database choice is often more important than programming language choice**. In order to make an informed decision, you need to be well-educated in general database concepts such as consistency and data modeling.

### Software and non-software books

I have become an avid reader of technical books, as well as popular non-fiction books like [*Grit*](https://www.amazon.com/gp/product/1501111108/ref=as_li_tl?ie=UTF8&camp=1789&creative=9325&creativeASIN=1501111108&linkCode=as2&tag=jaredririeblo-20&linkId=c173ddc20b9a9fcd700e582440ca8479), [*Mindset*](https://www.amazon.com/gp/product/0345472322/ref=as_li_tl?ie=UTF8&camp=1789&creative=9325&creativeASIN=0345472322&linkCode=as2&tag=jaredririeblo-20&linkId=28b35ebce32bc00c963a529c58070d49), and [*Work Rules!*](https://www.amazon.com/gp/product/1455554790/ref=as_li_tl?ie=UTF8&camp=1789&creative=9325&creativeASIN=1455554790&linkCode=as2&tag=jaredririeblo-20&linkId=fff6e98d9dd5016e1aa4be73e0368874) ever since I graduated from college. I read five books each quarter, so 20 books/year. And by read, I really mean read *or* listen; I'm an advocate of [Audible](https://www.amazon.com/gp/product/B00NB86OYE/ref=as_li_tl?ie=UTF8&camp=1789&creative=9325&creativeASIN=B00NB86OYE&linkCode=as2&tag=jaredririeblo-20&linkId=627d0e41b121bbc9b5a33b365e23a2d7) and think it's a solid investment.

I’m planning on writing a post for each book I read with a summary of the content, further learning it inspired, and my overall recommendation. Here are some books on my reading list:

* [*Clean Code*](https://www.amazon.com/gp/product/B001GSTOAM/ref=as_li_tl?ie=UTF8&camp=1789&creative=9325&creativeASIN=B001GSTOAM&linkCode=as2&tag=jaredririeblo-20&linkId=2596e9caf8f63700450812054449c5d0)
* [*SQL Antipatterns*](https://www.amazon.com/gp/product/1934356557/ref=as_li_tl?ie=UTF8&camp=1789&creative=9325&creativeASIN=1934356557&linkCode=as2&tag=jaredririeblo-20&linkId=2bc3044e49259e2a806ec0d84738be0c)
* [*Cracking the Coding Interview*](https://www.amazon.com/gp/product/0984782850/ref=as_li_tl?ie=UTF8&camp=1789&creative=9325&creativeASIN=0984782850&linkCode=as2&tag=jaredririeblo-20&linkId=06a672d4319440a648fcea507d939810)
* [*Multipliers: How the Best Leaders Make Everyone Smarter*](https://www.amazon.com/gp/product/0062663070/ref=as_li_tl?ie=UTF8&camp=1789&creative=9325&creativeASIN=0062663070&linkCode=as2&tag=jaredririeblo-20&linkId=f07f15aaa881d773ccfdbe396f4c7560)
* [*The Effective Executive*](https://www.amazon.com/gp/product/0060833459/ref=as_li_tl?ie=UTF8&camp=1789&creative=9325&creativeASIN=0060833459&linkCode=as2&tag=jaredririeblo-20&linkId=67180a0eee99a76f7aca0cf432e84625)

### Research papers

I was one week away from attending graduate school. I had accepted an offer and scholarship from the University of Wisconsin: Madison, enrolled in classes, and found an apartment. Then, in a last-minute decision, I walked away from it all a week before classes started. I'll share the full story in a later blog post.

While I still believe this was the correct choice given my circumstances, I regret not being able to delve into Computer Science research. Reading research papers has been more of a hit and miss for me than reading books. I hope this blog can serve as the necessary motivation to read more research papers and review them as blog entries. The research papers I have read, such as [Amazon’s well-known Dynamo paper](https://www.allthingsdistributed.com/files/amazon-dynamo-sosp2007.pdf), have been influential in improving my design skills and identifying weak areas in my understanding.

### Up-and-coming backend technology

The backend is notably more stable than the frontend. The database terminology or [distributed consensus algorithm](http://thesecretlivesofdata.com/raft/) you learned a few years ago will still be relevant for a long time. Meanwhile, if you picked up AngularJS around the same time, you know it was soon eclipsed by Angular and then React and now maybe Vue.js.

<img src="/public/images/service-mesh-istio-google-trends.png" width="100%" alt="Google Trends for Service Mesh and Istio" />

That said, the backend is still encapsulated in the ever-changing thing which is technology. “Service mesh” is one example of a backend idea that has only recently entered my vocabulary. It is a solution to dealing with the varied interactions between services in a network of microservices. [Istio](https://istio.io/docs/concepts/what-is-istio/overview/) is an example project I’ll cover in a later blog post. I plan to regularly write about new technologies of this nature.

### Architecture and system design

One of the benefits of working at a smaller company is the opportunity to be involved in key architectural discussions even early in your career. Such has been the case for me at Qualtrics where I have been able to influence large chunks of the backend. I know, however, that I'm still in my infancy in terms of my ability to design elegant solutions to cross-cutting problems in a system. I am confident that as I improve my system design skills, I will be able to make a bigger impact on the technical direction of my software team.
