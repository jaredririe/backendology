# What this blog is all about

## "Backended"?

> There are only two hard things in Computer Science: cache invalidation and naming things (Phil Karlton)

Or, my favorite variant:

> There are only two hard things in Computer Science: cache invalidation, naming things, and off by one errors (Unknown - probably some clever person on the internet)

Naming is hard. Backended isn't a real word. "Back ended" and its hyphened version "back-ended" do appear to be real words, though rarely used and completely unrelated to the purpose of this blog. Like the programmer who decides to write a lengthy comment rather than coming up with a better name, I'll try to explain what I intended when I named this blog Backended.

Backended draws obvious reference to the the term "backend", as in backend or server-side web development. If it was a real word, you might say, "Because of my backended nature, writing frontend code is inherently painful." I find the problems that get solved on the backend to be incredibly interesting and therefore choose it as the underlying theme of my blog. 

## Topics I plan to cover 

### Distributed systems 

Many backend systems are an entanglement of services which together can be considered a distributed system. Rather than being monolithic in nature, these systems benefit from things like **independent deployability**, **focused development**, and **loose coupling** yet suffer from the complexities of **partial failure**, **lack of concurrency guarantees**, and **network boundaries**.

This list of tradeoffs is clearly incomplete. Each complexity deserves its own blog post! Once network boundaries exist between components of a system, for example, problems such as discoverability must be dealt with. No longer can I make an easy function call; rather, I must build in some way for one service to find another service and then make a network call. This call could fail (should I retry?) or time out (how long should I wait?). Once I've figured out how to handle these failure scenarios, does *every* service I write need to duplicate this logic or can I abstract it away through a form of middleware?

Fascinating! 

Topics like caching, consistency, reverse proxying, API gateways, and service meshes are fair game. 

### NoSQL and databases 


### Go

Go is a phenomenal language for backend development. I started writing Go in October 2015 and haven't looked back. I'll be writing several blog posts where I solve problems using Go or otherwise talk about it. 

Go is easy to learn, proven in production, and designed for the cloud. It's deployed in production by high-traffic companies like Google, Dropbox, Uber, and Facebook in cases where stability and high performance are critical. Many open-source cloud computing tools are written in Go like Docker and Consul.

I enjoy its strong focus on simplicity which translates to less ramp-up time for new developers. Another benefit is in maintainability of Go codebases, as it’s very readable and easy to understand existing code.

### Software and non-software books

Ever since college, I have become an avid reader of technical/coding books, as well as popular non-fiction books like "Grit", "Mindset", and "How Google Works". I read five books each quarter, so 20 books/year. And by read, I mean read **or** listen to; I'm an advocate of Audible and think it's a solid investment. 

### Research papers

I was literally one week away from attending graduate school--I accepted an offer and scholarship from the University of Wisconsin: Madison, enrolled in classes, found an apartment, everything and then walked away from this opportunity a week before classes started. I'll share the full story in a later blog. 

While I still believe this was the correct choice in my circumstances, I regret not being able to delve into research around topics in Computer Science. Unlike reading books, reading research papers has been more of a hit and miss. I hope this blog can serve as the necessary motivation to read more research papers and review them as blog entries. 

### Up-and-coming backend technology

### Best practices 

### Architecture and system design 

--- 

# Notes (supplementary to blog post)

## Go

There are a growing number of Go codebases at Qualtrics. What makes Go well-suited for these use cases? What are its strengths and design principles, and what makes Go a good choice for my next project?

### Why use Go?

As a language, Go has a number of properties that make it attractive. At a high level, Go is battle-tested, easy to learn, and designed for cloud software.

Most importantly, Go has been a safe language to adopt at Qualtrics. The language enjoys strong adoption among high-profile cloud-computing companies and open-source infrastructure tools.

* Proven in production: Go is deployed in production by high-traffic companies like Google, Dropbox, the New York Times, BBC, SpaceX, and Facebook in cases where stability and high performance are critical.
* Powers our critical tools: Many open-source cloud computing tools are written in Go, including components that Qualtrics uses as fundamental infrastructure underpinnings like Docker and Consul.
Upward trajectory: In the last several years, Go has reliably climbed up the rankings in programming language indices like Redmonk and TIOBE.

Apart from being well-suited to Qualtrics’ use cases and a relatively future-proof choice, the language also has features that encourage developer productivity:

* Stable API: The developers of Go have promised to maintain compatibility with Go 1, meaning that programs written today will continue to compile and run correctly for every minor release of Go until Go 2 (for perspective, Go 1 was released in early 2012; Go 1.9 is the current version).
* Simple by design: The language’s syntax is extremely simple, which translates to less ramp-up time for new developers. Another benefit is in maintainability of Go codebases, as it’s very readable and easy to understand existing code.
* Well-designed standard library: The Go standard library has a surprisingly thorough collection of well-tested building blocks, reducing our dependence on third-party libraries and addressing common problems with proven implementations.

Go isn’t the ideal language for every project. Some projects still might benefit from the mature language features in Java or the expressiveness of Node.js or Python. However, Go is attractive for a class of applications that we encounter a lot at Qualtrics: high-performance, maintainable servers.

