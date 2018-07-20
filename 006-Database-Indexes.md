# Breaking Down Abstractions: Database Indexes

In the [introduction to this blog]() I mentioned that I love breaking down abstractions to understand what makes them tick. This blog post is the first of many that will break down a fascinating abstraction.

## Prerequisite Terminology and Concepts

These are concepts that I assume you are somewhat familiar with that may not be included in a typical Computer Science education:

* 

## Abstraction 

Database indexes provide the abstraction of performant queries with minimal overhead during writes and additional storage space. As most workflows are read-heavy, the overhead of indexes is almost always worthwhile. 

Indexes are incredibly powerful. The difference between a query that can utilize one or more indexes and a query without indexes can be breathtaking on datasets larger than a few hundred objects. In fact, the difference can be that the query with indexes finishes in milliseconds while the index-less query causes the database to fall over!

I mentored an intern through a project with a database component last summer. Once the database grew to a substantial size, he remarked that his queries were taking seconds to complete, which was much longer than on his test dataset of a few items. I questioned whether indexes were in place and had him tack on `EXPLAIN` to the start of his query (most databases, SQL or NoSQL, have some notion of `EXPLAIN` which details the steps in executing the given query). While he did have some indexes, it was evident from the results of the `EXPLAIN` that none of them were being used; rather the entire database was being scanned (known as a [full table scan](https://en.wikipedia.org/wiki/Full_table_scan)). We identified and added the missing index. The query wasn't just twice as fast, it was two orders of magnitude faster, so about 10 milliseconds. I think his reaction may resonate with some of you: "Clearly I need to learn more about indexes." 

Why do the right indexes lead to potentially massive performance increases? Conversely, why can the wrong index cause slow queries or even the database to fail? Why are indexes so important? Let's break down this abstraction.

## Breaking it down

### Binary Tree vs. B-Tree vs. B+ Tree

Many software engineers are understandably confused about the differences between these three very similar sounding data structures: binary tree, B-tree, and B+ tree. While they are all tree data structures, they are *not* synonomous. A B-tree is not a "binary" tree (no matter how convenient that would be)! In fact: "What, if anything, the B stands for has never been established."[^1] If a binary tree is a tree where each node has at most two children, how do we define B-tree and B+ tree?

#### B-tree

> B-tree is a self-balancing tree data structure that keeps data sorted and allows searches, sequential access, insertions, and deletions in logarithmic time. The B-tree is a generalization of a binary search tree in that a node can have more than two children.[^1]

- Self-balancing
- Keeps data sorted 
- Can have more than two children 

#### B+ tree

> A B+ tree is an N-ary tree with a variable but often large number of children per node. ... A B+ tree can be viewed as a B-tree in which each node contains only keys (not key–value pairs), and to which an additional level is added at the bottom with linked leaves.[^2]

- Large number of children per node
- Nodes contain only keys (not key-value pairs)
- Link the leaves together 

#### Progression

Given those definitions, it becomes clear that each of these tree data structures builds on the previous one in a progression towards a structure fit for a database index. The added complexity with each step is a tradeoff for improved performance in particular use cases.

1. Start with a simple binary tree. Add self-balancing and enforce the properties of a binary search tree. Result: **balanced binary search tree**
2. Allow this balanced binary search tree to have more than two children. Result: **B-tree**
3. Link the leaves and constrain the nodes to store only keys. Result: **B+ tree**

## Further reading


---

# Notes (supplementary to blog post)

## [B-Tree](https://en.wikipedia.org/wiki/B-tree)

*Not* a binary tree!

> B-tree is a self-balancing tree data structure that keeps data sorted and allows searches, sequential access, insertions, and deletions in logarithmic time. The B-tree is a generalization of a binary search tree in that a node can have more than two children.
> 
> Unlike self-balancing binary search trees, the B-tree is well suited for storage systems that read and write relatively large blocks of data, such as discs. It is commonly used in databases and filesystems.
> 
> What, if anything, the B stands for has never been established.

## [B+ Tree](https://en.wikipedia.org/wiki/B%2B_tree)

> A B+ tree is an N-ary tree with a variable but often large number of children per node. A B+ tree consists of a root, internal nodes and leaves. The root may be either a leaf or a node with two or more children.
> 
> A B+ tree can be viewed as a B-tree in which each node contains only keys (not key–value pairs), and to which an additional level is added at the bottom with linked leaves.
> 
> The primary value of a B+ tree is in storing data for efficient retrieval in a block-oriented storage context — in particular, filesystems. This is primarily because unlike binary search trees, B+ trees have very high fanout (number of pointers to child nodes in a node, typically on the order of 100 or more), which reduces the number of I/O operations required to find an element in the tree.
> 
> Relational database management systems such as IBM DB2, Informix, Microsoft SQL Server, Oracle 8, Sybase ASE, and SQLite support this type of tree for table indices. Key–value database management systems such as CouchDB and Tokyo Cabinet support this type of tree for data access.
> 
> The leaves (the bottom-most index blocks) of the B+ tree are often linked to one another in a linked list; this makes range queries or an (ordered) iteration through the blocks simpler and more efficient


## [Comparison of B-Tree and Hash Indexes](https://dev.mysql.com/doc/refman/5.5/en/index-btree-hash.html)

## MySQL

> MySQL uses both BTREE (B-Tree and B+Tree) and HASH indexes 
(http://www.vertabelo.com/blog/technical-articles/all-about-indexes-part-2-mysql-index-structure-and-performance)

[^1]: https://en.wikipedia.org/wiki/B-tree
[^2]: https://en.wikipedia.org/wiki/B%2B_tree

