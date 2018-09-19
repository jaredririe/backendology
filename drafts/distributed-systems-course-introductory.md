# Designing a Comprehensive Course in Distributed Systems: Introductory Course

In my [previous blog post](https://backendology.com/2018/09/10/distributed-systems-course-reading-list/), I created a detailed reading list of the distributed systems content I deemed most important and interesting. This post is the next step towards designing a comprehensive course in distributed systems: creating an introductory course.

Why not simply start at the top of the reading list and work down? I think any good introduction to this topic should provide a sampling of the various concepts and encourage experimentation and hands on learning. I also organized my reading list such that the "Distributed systems in the wild" section is one of the later sections, but I believe that failing to cover some real systems in an introductory course is a lost opportunity. Students and practitioners often have experience interacting with the user-facing parts of systems like Kafka, Memcache, or Cassandra. Satisfying the curiosity to learn how these systems work under the hood can fuel the harder parts of learning distributed systems, such as grappling with Paxos.

Speaking of Paxos, it was designed first but the newcomer Raft is intended to be much easier to learn and implement. For this reason, I think an introductory course doesn't need to dive into the details of Paxos. A high-level understanding is sufficient. Implementing Raft, on the other hand, is not too much to ask and worth the effort.

Without further ado, here is my course, _Introduction to Distributed Systems_!

## Course

### Unit 1: The Problem

Questions

* _Why are distributed systems challenging to build correctly?_
* _What are some of the building blocks of distributed systems?_
* _What are the failure modes?_
* _How do web architectures and microservices relate to distributed systems?_

Real world example: MapReduce?
Hands on learning: MapReduce lab? (how hard is it?)

Concepts

* Challenging problems which must be faced
    - Byzantine Generals Problem
    - Time, Clocks, and the Ordering of Events in a Distributed System
* Fallacies of distributed systems
* Building blocks by reviewing Scalable Web Architecture
* Failure modes
* Relation to microservices

### Unit 2: Gossip Protocols

Questions

*

Real world example: hashicorp/memberlist, Pilosa's usage of it
Hands on learning: Implement a global counter using hashicorp/memberlist or Implement a distributed in-memory cache freecache?

Concepts

* Gossip protocol
* SWIM

### Unit 3: Solving Distributed Consensus with Raft

Questions

*

Real world example: etcd
Hands on learning: Implement Raft (MIT 6.824 Lab 2)

Concepts

* Impossibility of Distributed Consensus with One Faulty Process (FLP Impossibility)
* Distributed consensus
* Raft paper
* Raft visualization
* etcd
* Consul
* Managing Critical State: Distributed Consensus for Reliability
* Paxos (high-level)

### Unit 4: Consistency and Availability

Questions

*

Real world example: Dynamo, Spanner
Hands on learning: Build a fault-tolerant key/value storage service using Raft library (MIT 6.824 Lab 3)

Concepts

* A Critique of the CAP Theorem
    - Please stop calling databases CP or AP
* Highly Available Transactions
* Consistency models
* Life beyond Distributed Transactions
* Building on Quicksand

---

# Notes (supplementary to blog post)


