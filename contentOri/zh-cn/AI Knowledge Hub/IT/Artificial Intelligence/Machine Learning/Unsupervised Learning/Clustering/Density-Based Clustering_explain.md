Density-Based Clustering（基于密度的聚类）是一种聚类方法，它根据样本在特征空间中的分布情况寻找样本空间中密度较高的区域，将高密度区域视作聚类中心，将低密度区域视作噪音或离群值。

最常用的Density-Based Clustering算法是DBSCAN（Density-Based Spatial Clustering of Applications with Noise），它透过指定一定的半径Eps和密度阈值MinPts，来将样本点分为三种不同的类型：
1. 核心点（Core Point）：在以样本点为中心且半径为Eps的区域内样本数大于等于MinPts的点。
2. 边界点（Border Point）：在以样本点为中心且半径为Eps的区域内样本数小于MinPts的点，但它是一个核心点的邻居，即可以连接到一个核心点。
3. 噪音点（Noise Point）：既不是核心点也不是边界点的点。

以二维空间为例，下图展示了从5个核心点开始的DBSCAN聚类过程：
![DBSCAN clustering process](https://miro.medium.com/max/600/1*WH1cji_S35tbybvloRQL4A.gif)
 
图中颜色不同的点表示它们属于不同的聚类，灰色的点表示噪音点。整个过程可以看作是将样本空间依照密度进行分割的过程，同时保留了一定程度的空间信息，避免了像k-means聚类那样仅根据样本在特征空间中的距离对样本进行划分，缺乏对样本分布情况的考虑。