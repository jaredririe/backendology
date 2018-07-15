# Breaking Down Abstractions: Database Indexes

## Prerequisite Terminology and Concepts

## Abstraction 

Database indexes provide the abstraction of performant queries with minimal overhead during writes. As most workflows are read-heavy, the overhead of indexes is almost always worthwhile. 

Indexes are incredibly powerful. The difference between a query that can utilize one or more indexes and a query without indexes can be breathtaking on datasets larger than a few hundred objects. In fact, the difference can be that the query with indexes finishes in milliseconds while the index-less query causes the database to fall over!

I mentored an intern through a project with a database component last summer. Once the database grew to a substantial size, he remarked that his queries were taking seconds to complete, which was much longer than on his test dataset of a few items. I questioned whether indexes were in place and had him tack on `EXPLAIN` to the start of his query (most databases, SQL or NoSQL, have some notion of `EXPLAIN` which details the steps in executing the given query). While he did have some indexes, it was evident from the results of the `EXPLAIN` that none of them were being used; rather the entire database was being scanned. We identified and added the missing index. The query wasn't just twice as fast, it was two orders of magnitude faster, so about 10 milliseconds. I think his reaction may resonate with some of you: "Clearly I need to learn more about indexes." 

Why do the right indexes lead to potentially massive performance increases? Conversely, why can the wrong index cause slow queries or even the database to fail? Why are indexes so important? Let's break down this abstraction.

## Breaking it down

## Further reading
