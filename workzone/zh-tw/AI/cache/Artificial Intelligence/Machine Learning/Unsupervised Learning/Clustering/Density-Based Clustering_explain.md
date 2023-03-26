Density-Based Clustering（基於密度的聚類）是一種聚類方法，它根據樣本在特徵空間中的分佈情況尋找樣本空間中密度較高的區域，將高密度區域視作聚類中心，將低密度區域視作噪音或離群值。

最常用的Density-Based Clustering算法是DBSCAN（Density-Based Spatial Clustering of Applications with Noise），它透過指定一定的半徑Eps和密度閾值MinPts，來將樣本點分為三種不同的類型：
1. 核心點（Core Point）：在以樣本點為中心且半徑為Eps的區域內樣本數大於等於MinPts的點。
2. 邊界點（Border Point）：在以樣本點為中心且半徑為Eps的區域內樣本數小於MinPts的點，但它是一個核心點的鄰居，即可以連接到一個核心點。
3. 噪音點（Noise Point）：既不是核心點也不是邊界點的點。

以二維空間為例，下圖展示了從5個核心點開始的DBSCAN聚類過程：
![DBSCAN clustering process](https://miro.medium.com/max/600/1*WH1cji_S35tbybvloRQL4A.gif)
 
圖中顏色不同的點表示它們屬於不同的聚類，灰色的點表示噪音點。整個過程可以看作是將樣本空間依照密度進行分割的過程，同時保留了一定程度的空間信息，避免了像k-means聚類那樣僅根據樣本在特徵空間中的距離對樣本進行劃分，缺乏對樣本分佈情況的考慮。