Hierarchical clustering is a clustering method in which the data objects are represented in the form of a tree-like structure (dendrogram). It is a bottom-up approach, where each object is initially considered as a separate cluster and then, based on the similarity between objects, the clusters are merged recursively until a single cluster is formed.

Example: Suppose we have a dataset of 6 points in 2-D space as follows:

(1,2), (2,4), (2,5), (3,5), (4,6), (5,7)

We can start with each point being a cluster on its own, and then start merging the closest points into a new cluster. For example, we can merge (2,4) and (1,2) into a new cluster with an average coordinate (1.5,3), then merge that with (2,5) into a new cluster with an average coordinate (1.8,3.7). 

We then proceed to form a cluster with (3,5) and the previous cluster, resulting in a new cluster with an average coordinate (2.1,4.2). Finally, we merge the last two points with the previous cluster, creating a new cluster with an average coordinate (4.5,6.5).

The dendrogram representation of the hierarchical clustering can be seen as:

                (1,2)
         /                \
   (2,4)                  (2,5)
  /        \              |
(3,5)   (4,6)          (5,7)


The resulting dendrogram shows us how the clusters were formed and at which distance level they were merged.