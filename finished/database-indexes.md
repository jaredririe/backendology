# Breaking Down Abstractions: Database Indexes

In the introduction to this blog, I mentioned that I love breaking down abstractions to understand what makes them tick. This blog post is the first of many that will break down a fascinating abstraction.

## Abstraction

*Database indexes provide the abstraction of performant queries with the cost of some overhead during writes and additional storage space. As most workflows are read-heavy, the overhead of indexes is almost always worthwhile.*

Indexes are incredibly powerful. The difference between a query that can utilize one or more indexes and a query without indexes can be breathtaking on datasets larger than a few hundred objects. In fact, the difference can be that the query with indexes finishes in milliseconds while the index-less query causes the database to fall over!

I mentored an intern through a project with a database component last summer. Once the database grew to a substantial size, he remarked that his queries were taking seconds to complete, which was much longer than on his test dataset of a few items. I questioned whether indexes were in place and had him tack on `EXPLAIN` to the start of his query (most databases, SQL or NoSQL, have some notion of `EXPLAIN` which details the steps in executing the given query). While he did have some indexes, it was evident from the results of the `EXPLAIN` that none of them were being used; rather the entire database was being scanned (known as a [full table scan](https://en.wikipedia.org/wiki/Full_table_scan)). We identified and added the missing index. The query wasn't just twice as fast, it was two orders of magnitude faster, so about 10 milliseconds. I think his reaction may resonate with some of you: "Clearly I need to learn more about indexes."

Why do the right indexes lead to potentially massive performance increases? Conversely, why can the wrong index cause slow queries or even the database to fail? Why are indexes so important? Let's break down this abstraction.

## Breaking it down

To understand how database indexes work under the hood, let's define a simple table which lists several people, their ages, and associated IDs:

| id | name            | age |
|----|-----------------|-----|
| 1  | Arvilla Hawks   | 24  |
| 2  | Maryellen Gourd | 18  |
| 6  | Corliss Henline | 38  |
| 8  | Lidia Haught    | 17  |
| 9  | Leo Thurlow     | 84  |
| 14 | Raymundo Vavra  | 23  |
| 26 | Lyn Stucky      | 28  |
| 28 | Lura Apodaca    | 11  |
| 30 | Brook Milum     | 81  |

Given an ID, we'd like to find out the name and age. Here's a SQL query to discover that Leo Thurlow (age 84) is stored at ID 9:

```sql
SELECT name, age FROM people WHERE id = 9
```

How does this query find the right answer? Well, without indexes, it merely scans the rows of the table until it happens upon ID 9. This ID was just a few entries from the start, so that doesn't sound too bad. Consider, however, that this table could be thousands, millions, or billions of rows long. What if ID 9 was the *last* row? We'd have to scan through a lot of rows! This process of scanning each row has a time complexity of O(N). Not great.

It's also worth noting that if the `id` column was not unique, it would become necessary to scan all entries to find all the ones with the given ID. When N is large, scanning N items every time we need to respond to a query is going to take a long time. It's possible that a long-running query combined with lots of other concurrent queries could just cause a database to fail (it could run out of memory, go into swap, become unable to respond to healthchecks, and be removed from the cluster).

What we need is some way to store the IDs in a data structure more suited to this kind of searching. Enter: the index!

> Indexes are used to quickly locate data without having to search every row in a database table every time a database table is accessed. Indexes can be created using one or more columns of a database table, providing the basis for both rapid random lookups and efficient access of ordered records.[^1]

There are many [different types of database indexes](https://en.wikipedia.org/wiki/Comparison_of_relational_database_management_systems#Indexes), including B-tree, R-tree, Hash, Bitmap, and Spacial. Rather than trying to cover all of these index types, let's deep dive on the ubiquitous B-tree.

### Binary Search Tree vs. B-Tree vs. B+ Tree

Many software engineers are understandably confused about the differences between these three very similar sounding data structures: binary search tree, B-tree, and B+ tree. While they are all tree data structures, they are *not* synonymous. A B-tree is not a "binary" tree (no matter how convenient that would be)! In fact: "What, if anything, the B stands for has never been established."[^2]

#### Binary Search Tree

Binary search trees are binary trees that keep their keys in sorted order by enforcing the requirement that all left children of a node have values less than the node's and all right children have values greater than the node's.

![Binary Search Tree](../static/public/images/binary-search-tree.png)

#### B-tree

> B-tree is a self-balancing tree data structure that keeps data sorted and allows searches, sequential access, insertions, and deletions in logarithmic time. The B-tree is a generalization of a binary search tree in that a node can have more than two children.[^2]

![B-Tree](../static/public/images/balanced-nary-tree.png)

B-trees have these properties:

- Can have more than two children
- Self-balancing
    - Internal (non-leaf) nodes may be joined or split in order to maintain balance
    - All leaf nodes must be at the same depth
- The number of children for a particular node is equal to the number of keys in it plus one
- The minimum and maximum number of child nodes are typically fixed (a **2-3 B-tree** allows each internal node to have 2 or 3 child nodes)
- Minimizes wasted storage space by ensuring the interior nodes are at least half full

In our simple B-tree, we store only a few keys in each node and each node only has a small number of children. A practical B-tree would have far more keys and children--the exact number is related to the size of a full disk block such that each read will get as much data as possible.

To use a B-tree as a database index, we must either store the entire table rows or hold pointers to the rows. Let's take the pointer approach:

![B-Tree as Database Index](../static/public/images/b-tree.png)

Note that this B-tree has the properties outlined above:

- All leaf nodes are at the same depth
- The `[8  28]` node has two keys and therefore three children
- Values less than 8 are in the left subtree, values between 8 and 28 and in the middle subtree, and values greater than 28 are in the right subtree
- Wasted space is minimized and the tree is balanced

Let's return to the people table we started with and rerun the query now that we have an index structure in place. When evaluating the WHERE clause (`WHERE id = 9`), the query optimizer now sees there is an index on the `id` column. It follows the B-tree data structure until it finds ID 9 and follows the pointer to the associated table row.

1. Is 9 equal to 8? No
2. Is 9 less than or equal to 8? No
3. Is 9 between 8 and 28 (inclusive)? Yes!
4. Is 9 equal to 9? Yes!
5. Which table row does 9 point to? `9 | Leo Thurlow | 84`

The time complexity for this search drops from O(N) to O(log(N))!

#### B+ tree

> A B+ tree is an N-ary tree with a variable but often large number of children per node. ... A B+ tree can be viewed as a B-tree in which each node contains only keys (not key–value pairs), and to which an additional level is added at the bottom with linked leaves.[^3]

![B+ Tree as Database Index](../static/public/images/b+-tree.png)

B+ trees have these additional properties relative to a B-tree:

- Leaf nodes are linked together
- All keys (and pointers to table rows) are stored in the leaves
- Copies of the keys are stored in the internal (non-leaf) nodes
- Typically have a large number of children per node
- Each node *may* store a pointer to the next node for faster sequential access

Comparing the example B-tree and B+ tree reveals that the same data is stored in each, but the additional properties of the B+ tree force the keys down to the leaf nodes. The linked list this forms makes range queries more efficient.

#### Progression

Given those definitions, it becomes clear that each of these tree data structures builds on the previous one in a progression towards a structure fit for a database index. The added complexity with each step is a trade-off for improved performance in particular use cases. Some databases use a combination of B-trees and B+ trees depending on the type of index and nature of queries.

Here is one way of looking at the progression from binary tree to B+ tree:

1. Start with a simple binary tree. Add self-balancing and enforce the properties of a binary search tree. Result: **balanced binary search tree**
2. Allow this balanced binary search tree to have more than two children and enforce other properties of a B-tree. Result: **B-tree**
3. Add pointers from the stored keys to rows in the database table. Result: **B-tree used as a database index**
4. Push all keys into the leaf nodes and link the leaves. Result: **B+ tree used as a database index**

## Further reading

* [Wikipedia: Database Index](https://en.wikipedia.org/wiki/Database_index)
* [Wikipedia: B-Tree](https://en.wikipedia.org/wiki/B-tree)
* [Wikipedia: B+ Tree](https://en.wikipedia.org/wiki/B%2B_tree)
* [JavaScript B-Tree Demo](http://goneill.co.nz/btree-demo.php)
* [B+ Tree Visualization](https://www.cs.usfca.edu/~galles/visualization/BPlusTree.html)
* [Straightforward B+ Tree Implementation in JavaScript](http://blog.conquex.com/?p=84)
* [Comparison of B-Tree and Hash Indexes](https://dev.mysql.com/doc/refman/5.5/en/index-btree-hash.html)

[^1]: https://en.wikipedia.org/wiki/Database_index
[^2]: https://en.wikipedia.org/wiki/B-tree
[^3]: https://en.wikipedia.org/wiki/B%2B_tree

---

# Notes (supplementary to blog post)

## [B-Tree](https://en.wikipedia.org/wiki/B-tree)

*Not* a binary tree!

> B-tree is a self-balancing tree data structure that keeps data sorted and allows searches, sequential access, insertions, and deletions in logarithmic time. The B-tree is a generalization of a binary search tree in that a node can have more than two children.
>
> Unlike self-balancing binary search trees, the B-tree is well suited for storage systems that read and write relatively large blocks of data, such as discs. It is commonly used in databases and filesystems.
>
> What, if anything, the B stands for has never been established.
>
> In B-trees, internal (non-leaf) nodes can have a variable number of child nodes within some pre-defined range. When data is inserted or removed from a node, its number of child nodes changes. In order to maintain the pre-defined range, internal nodes may be joined or split. Because a range of child nodes is permitted, B-trees do not need re-balancing as frequently as other self-balancing search trees, but may waste some space, since nodes are not entirely full. The lower and upper bounds on the number of child nodes are typically fixed for a particular implementation. For example, in a 2-3 B-tree (often simply referred to as a 2-3 tree), each internal node may have only 2 or 3 child nodes.
>
> For example, if an internal node has 3 child nodes (or subtrees) then it must have 2 keys: a1 and a2. All values in the leftmost subtree will be less than a1, all values in the middle subtree will be between a1 and a2, and all values in the rightmost subtree will be greater than a2.
>
> A B-tree is kept balanced by requiring that all leaf nodes be at the same depth. This depth will increase slowly as elements are added to the tree, but an increase in the overall depth is infrequent, and results in all leaf nodes being one more node farther away from the root.
>
> By maximizing the number of keys within each internal node, the height of the tree decreases and the number of expensive node accesses is reduced. In addition, rebalancing of the tree occurs less often. The maximum number of child nodes depends on the information that must be stored for each child node and the size of a full disk block or an analogous size in secondary storage. While 2-3 B-trees are easier to explain, practical B-trees using secondary storage need a large number of child nodes to improve performance.
>
> In some designs, the leaves may hold the entire data record; in other designs, the leaves may only hold pointers to the data record. Those choices are not fundamental to the idea of a B-tree.

- keeps keys in sorted order for sequential traversing
- uses a hierarchical index to minimize the number of disk reads
- uses partially full blocks to speed insertions and deletions
- keeps the index balanced with a recursive algorithm
- In addition, a B-tree minimizes waste by making sure the interior nodes are at least half full.

https://www.geeksforgeeks.org/b-tree-set-1-introduction-2/

- All leaves are at same level.
- A B-Tree is defined by the a minimum degree 't'. The value of t depends upon disk block size.
- Every node except root must contain at least t-1 keys. Root may contain minimum 1 key.
- All nodes (including root) may contain at most 2t – 1 keys.
- Number of children of a node is equal to the number of keys in it plus 1.
- All keys of a node are sorted in increasing order. The child between two keys k1 and k2 contains all keys in the range from k1 and k2.
- B-Tree grows and shrinks from the root which is unlike Binary Search Tree. Binary Search Trees grow downward and also shrink from downward.
- Like other balanced Binary Search Trees, time complexity to search, insert, and delete is O(log(n)).

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
>
> In the B+ tree, copies of the keys are stored in the internal nodes; the keys and records are stored in leaves; in addition, a leaf node may include a pointer to the next leaf node to speed sequential access

[JavaScript B-tree Demo](http://goneill.co.nz/btree-demo.php)

[B+ Tree Visualization](https://www.cs.usfca.edu/~galles/visualization/BPlusTree.html)

[JavaScript implementation](http://blog.conquex.com/?p=84)

## [Comparison of B-Tree and Hash Indexes](https://dev.mysql.com/doc/refman/5.5/en/index-btree-hash.html)

## MySQL

> MySQL uses both BTREE (B-Tree and B+Tree) and HASH indexes
(http://www.vertabelo.com/blog/technical-articles/all-about-indexes-part-2-mysql-index-structure-and-performance)

## [Database Indexes](https://en.wikipedia.org/wiki/Database_index)

> A database index is a data structure that improves the speed of data retrieval operations on a database table at the cost of additional writes and storage space to maintain the index data structure. Indexes are used to quickly locate data without having to search every row in a database table every time a database table is accessed. Indexes can be created using one or more columns of a database table, providing the basis for both rapid random lookups and efficient access of ordered records.
>
> An index is a copy of selected columns of data from a table that can be searched very efficiently that also includes a low-level disk block address or direct link to the complete row of data it was copied from.

### Improved time complexity for lookups

> Suppose a database contains N data items and one must be retrieved based on the value of one of the fields. A simple implementation retrieves and examines each item according to the test. If there is only one matching item, this can stop when it finds that single item, but if there are multiple matches, it must test everything. This means that the number of operations in the worst case is O(N)
>
> Many index designs exhibit logarithmic (O(log(N))) lookup performance and in some applications it is possible to achieve flat (O(1)) performance
>
> With an index the database simply follows the B-tree data structure until the Smith entry has been found; this is much less computationally expensive than a full table scan.

### Overhead

On write, the time required to update the B-tree index and rebalance this data structure is overhead.

### When is a B-tree index inefficient?

> A bitmap index is a special kind of indexing that stores the bulk of its data as bit arrays (bitmaps) and answers most queries by performing bitwise logical operations on these bitmaps. The most commonly used indexes, such as B+ trees, are most efficient if the values they index do not repeat or repeat a small number of times. In contrast, the bitmap index is designed for cases where the values of a variable repeat very frequently. For example, the sex field in a customer database usually contains at most three distinct values: male, female or unknown (not recorded). For such variables, the bitmap index can have a significant performance advantage over the commonly used trees.

### Index implementations

> Indices can be implemented using a variety of data structures. Popular indices include balanced trees, B+ trees and hashes.

### Covering index

> A covering index is a special case where the index itself contains the required data field(s) and can answer the required data.
>
> A covering index can dramatically speed up data retrieval but may itself be large due to the additional keys, which slow down data insertion & update. To reduce such index size, some systems allow including non-key fields in the index. Non-key fields are not themselves part of the index ordering but only included at the leaf level, allowing for a covering index with less overall index size.
