# Designing a Comprehensive Course in Distributed Systems: Introductory Course

In my [previous blog post](https://backendology.com/2018/09/10/distributed-systems-course-reading-list/), I created a detailed reading list of the distributed systems content I deemed most important and interesting. This post is the next step towards designing a comprehensive course in distributed systems: creating an introductory course.

Why not simply start at the top of the reading list and work down? I think any good introduction to this topic should provide a sampling of the various concepts and encourage experimentation and hands on learning. I also organized my reading list such that the "Distributed systems in the wild" section is one of the later sections, but I believe that failing to cover some real systems in an introductory course is a lost opportunity. Students and practitioners often have experience interacting with the user-facing parts of systems like Kafka, Memcache, or Cassandra. Satisfying the curiosity to learn how these systems work under the hood can fuel the harder parts of learning distributed systems, such as grappling with Paxos.

Speaking of Paxos, it was designed first but the newcomer Raft is intended to be much easier to learn and implement. For this reason, I think an introductory course doesn't need to dive into the details of Paxos. A high-level understanding is sufficient. Implementing Raft, on the other hand, is not too much to ask and worth the effort.

Without further ado, here is my course, _Introduction to Distributed Systems_!

## Course

### Unit 1: The Problem

This unit frames the problem presented by distributed systems. It explains why they are challenging to build correctly, as well as their building blocks, failure modes, and fallacies.

* Building blocks of distributed systems
    - Read and summarize this chapter from the book _The Architecture of Open Source Applications_: [Scalable Web Architecture and Distributed Systems](http://www.aosabook.org/en/distsys.html)
* Core problems presented by distributed systems
    - Read this blog post and list the challenges: [Notes on Distributed Systems for Young Bloods](https://www.somethingsimilar.com/2013/01/14/notes-on-distributed-systems-for-young-bloods/)
    - Read and summarize this paper: [Time, Clocks, and the Ordering of Events in a Distributed System (Lamport 1978)](https://www.microsoft.com/en-us/research/publication/time-clocks-ordering-events-distributed-system/?from=http%3A%2F%2Fresearch.microsoft.com%2Fen-us%2Fum%2Fpeople%2Flamport%2Fpubs%2Ftime-clocks.pdf)
    - Read and summarize this paper: [The Byzantine Generals Problem (Lamport 1982)](http://www.cs.cornell.edu/courses/cs614/2004sp/papers/LSP82.pdf)
* Failure modes
    - Read these resources on failure modes: [Failure Modes in Distributed Systems](http://alvaro-videla.com/2013/12/failure-modes-in-distributed-systems.html), [Wikipedia: Failure Semantics](https://en.wikipedia.org/wiki/Failure_semantics)
* Real world example: MapReduce
    - Read and summarize this paper: [MapReduce (2004)](https://pdos.csail.mit.edu/6.824/papers/mapreduce.pdf)
* Hands on learning: MapReduce
    - Complete this lab: [MIT 6.824 Lab 1: MapReduce](https://pdos.csail.mit.edu/6.824/labs/lab-1.html)
        + "In this lab you'll build a MapReduce library as an introduction to programming in Go and to building fault tolerant distributed systems. In the first part you will write a simple MapReduce program. In the second part you will write a Master that hands out tasks to MapReduce workers, and handles failures of workers. The interface to the library and the approach to fault tolerance is similar to the one described in the original MapReduce paper."

### Unit 2: Filesystems?

MapReduce, HDFS, GFS

### Unit 3: Gossip Protocols

Gossip protocols have many important use cases in distributed systems, such as detecting node failure, spreading configuration data, or sharing state among multiple nodes in a cluster. They can elegantly solve problems with relaxed consistency requirements where a distributed consensus algorithm (like Paxos) or a centralized database would be impractical or unwise.

* Gossip protocol
    - Read and summarize these introductory resources: [Wikipedia: Gossip protocol](https://en.wikipedia.org/wiki/Gossip_protocol), [Using Gossip Protocols For Failure Detection, Monitoring, Messaging And Other Good Things](http://highscalability.com/blog/2011/11/14/using-gossip-protocols-for-failure-detection-monitoring-mess.html)
    - Answer this question: when is a gossip-based solution better than a centralized database or distributed consensus algorithm?
* SWIM
    - Read and summarize this paper: [SWIM: Scalable Weakly-consistent Infection-style Process Group Membership Protocol](http://www.cs.cornell.edu/info/projects/spinglass/public_pdfs/swim.pdf)
* Real world example: hashicorp/memberlist
    - Study this codebase: [hashicorp/memberlist](https://github.com/hashicorp/memberlist)
        + "The use cases for such a library are far-reaching: all distributed systems require membership, and memberlist is a re-usable solution to managing cluster membership and node failure detection."
        + Based on the SWIM protocol with some adaptions
    - Study how the distributed bitmap index utilizes the memberlist library: [Usage in Pilosa's codebase](https://github.com/pilosa/pilosa/blob/10eea2db4cca35dd6b173377edf36790a5f164e6/gossip/gossip.go)
* Hands on learning: global counter
    - Implement a global counter which uses hashicorp/memberlist; use these resources:
        + [github.com/nphase/go-clustering-example](https://github.com/nphase/go-clustering-example/blob/master/final/main.go)
        + [github.com/asim/memberlist](https://github.com/asim/memberlist)

### Unit 4: Solving Distributed Consensus with Raft

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

### Unit 5: Consistency and Availability

Questions

* _Is the CAP theorem useful in practice?_
*

Real world example: Dynamo, Spanner
Hands on learning: Build a fault-tolerant key/value storage service using Raft library (MIT 6.824 Lab 3)

Concepts

* A Critique of the CAP Theorem
    - Please stop calling databases CP or AP
* Highly Available Transactions
* Distributed Transactions
* Consistency models
* Life beyond Distributed Transactions
* Building on Quicksand
* Cache coherence/consistency
* Concurrency control (such as Optimistic Concurrency Control)

[^1]: https://en.wikipedia.org/wiki/Gossip_protocol

---

# Notes (supplementary to blog post)

Missing from my Course -- which of these topics should be covered in more advanced courses?

* Distributed state
* Big data analytics
* Security
* Consistency - Classical synchronization + Go-style synchronization
* RPC
* Fault tolerance (kind of)
* Logging and crash recovery
* Consistent hashing
* DNS and content delivery
* Peer-to-peer
* Virtual machines

Topics covered in the [University of Washington's Course](https://courses.cs.washington.edu/courses/csep552/16wi/)

* Time, Clocks, and Global States
* Distributed State
* Consensus
* Scalability
* Transactions
* Weak consistency
* Big Data Analytics
* Security

Topics covered in [CMU's Course](http://www.cs.cmu.edu/~dga/15-440/F12/syllabus.html)

* Communication over the internet
* Consistency - Classical synchronization + Go-style synchronization
* Distributed Filesystems, MapReduce, HDFS
* RPC
* Time and Synchronization
* Fault Tolerance, Byzantine Fault Tolerance
* RAID
* Concurrency Control
* Logging and Crash Recovery
* Consistent hashing and name-by-hash
* Distributed Replication
* Data-Intensive Computing and MapReduce/Hadoop
* DNS and Content Delivery Networks
* Peer-to-peer
* Virtual Machines
* Security Protocols

Topics covered in [MIT's Course](https://pdos.csail.mit.edu/6.824/schedule.html)

* MapReduce
* RPC and Threads
* Filesystems
* Primary-Backup Replication
* Fault Tolerance: Raft
* Spinnaker
* Zookeeper
* Distributed Transactions
* Optimistic Concurrency Control
* Big Data: Spark, Naiad
* Distributed Machine Learning: Parameter Server
* Cache Consistency: Frangipani, Memcached at Facebook
* Disconnected Operation, Eventual Consistency
* Peer-to-peer, DHTs
* Dynamo
* Bitcoin
