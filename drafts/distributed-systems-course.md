# Designing a Comprehensive Course in Distributed Systems

In a recent conversation at work, I learned about [MIT's Distributed Systems course](https://pdos.csail.mit.edu/6.824/schedule.html). As the majority of the content is available  online through the course website, I was just about ready to dive in, follow the course, and report on what I learned. But then I had this thought: I am not formally a student anymore, so I should take more ownership of my learning. I should design my own comprehensive course in Distributed Systems!

My course doesn't have to be constrained to a semester in length or only lightly cover topics that I find important and compelling. I'll be able to take a more depth-first approach to learning. For example, after implementing the Raft consensus algorithm, if I feel like there is more to learn, I can take time to build an simple system which uses my Raft code. A university course on a broad topic such as Distributed Systems has to avoid depth like this in order to cover all the required content. As someone who has been in industry for a few years, I also have a rough idea of which concepts matter in practice because I've run up against them and struggled to find time to understand them as well as I wanted.

## Why study Distributed Systems?

First, some motivation. Of all the things to learn about in our field, why study distributed systems? In that same conversation I mentioned above, one of my coworkers said it better than I could:

> Developers _really_ need to know how to deal with concurrency, think in that space, and stop thinking about programs as serial lists of instructions. There is _no_ demand for someone who can’t natively think about concurrency all over their algorithm and their code.

He went on to argue that an understanding of distributed systems is increasingly becoming a hard prerequisite of work done by software developers. While testing and software practice can safely and easily be learned on the job, the concepts related to asynchronous concurrent systems often cannot be. Misunderstandings in this regard lead to system outages and incur serious technical debt. To avoid making common mistakes and architecting poor systems, it's critical to take time to seriously learn distributed systems.

Personally, I feel like I have a good grasp of the fundamentals, but an important next step in my career is learning advanced concepts like consensus and broadcast. _Understanding abstractions a level or two below your usage is incredibly worthwhile, not to mention satisfying._ I want to understand the content necessary to be considered a so-called "distributed systems engineer." I also think that a few years of distributed systems practice in industry have primed me to better understand these concepts. I've covered many of them at a high level, so revisiting them with a new perspective will be helpful.

## Designing my course

What follows is the initial design of my own Distributed Systems course. To create it, I scoured the internet, literally following hundreds of links. If you are also interested in this topic, I recommend you do the same: make your own course and then find some way to hold yourself accountable to it while allowing depth. These are the resources I found most useful in designing my course:

* [Awesome Distributed Systems](https://github.com/theanalyst/awesome-distributed-systems)
* [Readings in Distributed Systems](http://christophermeiklejohn.com/distributed/systems/2013/07/12/readings-in-distributed-systems.html)
* [The Paper Trail: necessary distributed systems theory and recommendations](https://www.the-paper-trail.org/post/2014-08-09-distributed-systems-theory-for-the-distributed-systems-engineer/)
* [MIT Distributed Systems Reading Group Paper List](http://dsrg.pdos.csail.mit.edu/papers/)

Distributed Systems research is known for an abundance of papers. My course will include papers as they are the primary source of information. Papers are also challenging and I want to get better at understanding them and using them as part of my learning.

## Concepts

### Basics

The basics give a sense for why distributed systems present challenging problems. With the problems framed, it will make more sense why we have to think so carefully about the advanced concepts below. Understanding that the network is unreliable and packets regularly get dropped, for example, makes it clear why it's hard to get multiple nodes of a database to agree on the state of the data.

* [Distributed Systems for Fun and Profit](http://book.mixu.net/distsys/single-page.html)
    - Free book that was recommended by several articles
* [Scalable Web Architecture and Distributed Systems](http://www.aosabook.org/en/distsys.html)
    - Chapter from a free book, _The Architecture of Open Source Applications_
    - Introduces the building blocks of distributed systems, including caches, indexes, load balancers, and queues
* [Notes on Distributed Systems for Young Bloods](https://www.somethingsimilar.com/2013/01/14/notes-on-distributed-systems-for-young-bloods/)
    - Short blog post written by an experienced distributed systems engineer with an audience of new engineers
    - Calls out that coordination is hard, failure of individual components is common, and that metrics and percentiles enable visibility
* [Failure Modes in Distributed Systems](http://alvaro-videla.com/2013/12/failure-modes-in-distributed-systems.html)
    - Short blog post which explains what could otherwise be confusing terminology of types of failure (performance, omission, fail-stop, crash, etc).
* [An introduction to distributed systems](https://github.com/aphyr/distsys-class)
    - Outline to a course taught by Kyle Kingsbury, the creator of the distributed systems tester Jepsen
    - Kyle Kingsbury has also given several talks which are [available on the Jepsen website](http://jepsen.io/talks)
* [Microservices: Are We Making a Huge Mistake?](https://backendology.com/2018/08/21/microservices-huge-mistake/)
    - My previous blog post which covered Microservices (a software development technique for building a software system that runs on a distributed system) and some of the [Fallacies of Distributed Computing](https://en.wikipedia.org/wiki/Fallacies_of_distributed_computing)
    - See this more detailed [Explanation of the Fallacies](https://www.rgoarchitects.com/Files/fallacies.pdf) as well

### MapReduce

* ["Lecture 1: Introduction"](https://pdos.csail.mit.edu/6.824/notes/l01.txt)
* [MapReduce Paper](https://pdos.csail.mit.edu/6.824/papers/mapreduce.pdf)
* [Lab 1: MapReduce](https://pdos.csail.mit.edu/6.824/labs/lab-1.html)

### RPC and Threads

* ["Lecture 2: Infrastructure: RPC and threads"](https://pdos.csail.mit.edu/6.824/notes/l-rpc.txt)
* [Crawler](https://pdos.csail.mit.edu/6.824/notes/crawler.go), [K/V](https://pdos.csail.mit.edu/6.824/notes/kv.go)
* [RPC Package](https://golang.org/pkg/net/rpc/)

### Distributed Consensus

* [The Byzantine Generals Problem](http://www.cs.cornell.edu/courses/cs614/2004sp/papers/LSP82.pdf)
    - One of the classic papers which presents a fictitious scenario in war to explain a problem faced by any distributed system
* [Time, Clocks, and the Ordering of Events in a Distributed System (1978)](https://www.microsoft.com/en-us/research/publication/time-clocks-ordering-events-distributed-system/?from=http%3A%2F%2Fresearch.microsoft.com%2Fen-us%2Fum%2Fpeople%2Flamport%2Fpubs%2Ftime-clocks.pdf)
* [Distributed Snapshots: Determining Global States of a Distributed System (1984)](http://research.cs.wisc.edu/areas/os/Qual/papers/snapshots.pdf)
* [Impossibility of Distributed Consensus with One Faulty Process](http://groups.csail.mit.edu/tds/papers/Lynch/jacm85.pdf)

#### Raft

Paxos 1989. Raft 2013 meant to be a simplified, understandable version of Paxos

* [In Search of an Understandable Consensus Algorithm](https://www.usenix.org/node/184041)
* [Consul: Raft Protocol Overview](https://www.consul.io/docs/internals/consensus.html)
* [ETCD](https://github.com/etcd-io/etcd)
* MIT Lab 2

#### Paxos

* [Paxos Made Simple](http://www.cs.utexas.edu/users/lorenzo/corsi/cs380d/past/03F/notes/paxos-simple.pdf)
* [Paxos Made Live](https://static.googleusercontent.com/media/research.google.com/en//archive/paxos_made_live.pdf)
    - > We describe our experience building a fault-tolerant database using the Paxos consensus algorithm. Despite the existing literature in the field, building such a database proved to be non-trivial.

#### Two-Phase and Three-Phase Commit

* [Two-Phase Commit](https://www.the-paper-trail.org/post/2008-11-27-consensus-protocols-two-phase-commit/)
* [Three-Phase Commit](https://www.the-paper-trail.org/post/2008-11-29-consensus-protocols-three-phase-commit/)

### Broadcast

> Atomic broadcast is exactly as hard as consensus - in a precise sense, if you solve atomic broadcast, you solve consensus, and vice versa. Chandra and Toueg prove this, but you just need to know that it’s true.

* [SWIM: Scalable Weakly-consistent Infection-style Process Group Membership Protocol](http://www.cs.cornell.edu/info/projects/spinglass/public_pdfs/swim.pdf)

Serf

* [Serf](https://www.serf.io/intro/vs-consul.html)
* [Serf: Gossip Protocol](https://www.serf.io/docs/internals/gossip.html)
* [Practical Golang: Building a simple, distributed one-value database with Hashicorp Serf](https://jacobmartins.com/2017/01/29/practical-golang-building-a-simple-distributed-one-value-database-with-hashicorp-serf/)

hashicorp/membership

* [hashicorp/memberlist](https://github.com/hashicorp/memberlist)
* [Pilosa's use of hashicorp/memberlist for Gossip](https://github.com/pilosa/pilosa/blob/10eea2db4cca35dd6b173377edf36790a5f164e6/gossip/gossip.go)

### Distributed file systems

* [The Google File System](https://pdos.csail.mit.edu/6.824/papers/gfs.pdf)

### Availability

* [Distributed Systems: Take Responsibility for Failover](http://ilyavolodarsky.com/distributed-systems-take-responsibility-for-failover/)

### Replication

### Consistency

* [Highly Available Transactions: Virtues and Limitations](http://www.bailis.org/papers/hat-vldb2014.pdf)
* [Consistency Models](http://jepsen.io/consistency):
    - Graph showing the relationships between consistency models in databases like Strict Serializable and Linearizability
* [Life beyond Distributed Transactions](https://www.ics.uci.edu/~cs223/papers/cidr07p15.pdf)
* [Consistency Tradeoffs in Modern Distributed Database System Design](http://cs-www.cs.yale.edu/homes/dna/papers/abadi-pacelc.pdf)
* [Building on Quicksand](https://arxiv.org/abs/0909.1788)
    - > Emerging patterns of eventual consistency and probabilistic execution may soon yield a way for applications to express requirements for a "looser" form of consistency while providing availability in the face of ever larger failures.
* [Eventually Consistent - Revisited](https://www.allthingsdistributed.com/2008/12/eventually_consistent.html)
* [There is No Now](https://queue.acm.org/detail.cfm?id=2745385)

(consistency models image)

### Big Data

### Tangential concepts

* [Consistent Hashing and Random Trees](https://www.akamai.com/us/en/multimedia/documents/technical-publication/consistent-hashing-and-random-trees-distributed-caching-protocols-for-relieving-hot-spots-on-the-world-wide-web-technical-publication.pdf)
* Queues [Everything Will Flow](https://www.youtube.com/watch?v=1bNOO3xxMc0)

### Distributed systems in the wild

* [Bigtable: A Distributed Storage System for Structured Data](http://static.googleusercontent.com/media/research.google.com/en//archive/bigtable-osdi06.pdf)
* [Chubby Lock Manager](http://static.googleusercontent.com/media/research.google.com/en//archive/chubby-osdi06.pdf)
* [Spanner](https://static.googleusercontent.com/media/research.google.com/en//archive/spanner-osdi2012.pdf)
* [Dynamo: Amazon’s Highly Available Key-value Store](http://www.read.seas.harvard.edu/~kohler/class/cs239-w08/decandia07dynamo.pdf)
* Kafka
* [ZooKeeper](https://www.usenix.org/legacy/event/usenix10/tech/full_papers/Hunt.pdf)
* [The Tail at Scale](https://ai.google/research/pubs/pub40801)

## Resources

### Blogs

* [High Scalability](http://highscalability.com/)
    - [Twitter](http://highscalability.com/blog/2013/7/8/the-architecture-twitter-uses-to-deal-with-150m-active-users.html)
    - [WhatsApp](http://highscalability.com/blog/2014/2/26/the-whatsapp-architecture-facebook-bought-for-19-billion.html)
* [All Things Distributed](https://www.allthingsdistributed.com/)

### Courses

* [Carnegie Mellon University: Distributed Systems](http://www.cs.cmu.edu/~dga/15-440/F12/syllabus.html)
* [University of Washington: Distributed Systems](https://courses.cs.washington.edu/courses/csep552/16wi/)

I found these courses in [this curated list of awesome Computer Science courses](https://github.com/prakhar1989/awesome-courses) avaialble online. Here is a similar list, except it focuses on [courses using Go](https://github.com/golang/go/wiki/Courses), many of which cover concurrency and distributed systems.

# Notes

> Developers _really_ need to know how to deal with concurrency, think in that space, and stop thinking about programs as serial lists of instructions. There is _no_ demand for someone that can’t natively think about concurrency all over their algorithm and their code.
> It’s super easy to learn testing and software practice out in the field. In fact, if you’re not always learning that, then you’re going to fall behind. Is it essential? Sure. But you’d better learn how to learn that, because you always will be.
On the other hand, learning about asynchronous concurrent systems in the field is….. why we have a lot of the problems that we do. You really should sit down and learn that seriously, then go to the field. Because otherwise we have all the problems that we do.

> You should know about safety and liveness properties:
> safety properties say that nothing bad will ever happen. For example, the property of never returning an inconsistent value is a safety property, as is never electing two leaders at the same time.
>liveness properties say that something good will eventually happen. For example, saying that a system will eventually return a result to every API call is a liveness property, as is guaranteeing that a write to disk always eventually completes.
