層次聚類是一種聚類分析方法，它將數據集中的對象進行分組，以便與其他組內的對象相比具有更相似的特徵。該方法的結果可以表示為樹狀圖，稱為譜系圖。

在層次聚類中，可以使用兩種方法來組織數據：自上而下（稱為分裂）或自下而上（稱為合併）。如果採用自下而上的方法，開始時把每個對象看作一個單獨的簇，然後將它們逐步合併為更大的簇，直到形成完整數據集的最終簇。如果採用自上而下的方法，開始時把整個數據集看作一個簇，然後逐步分裂它成更小的簇，直到每個簇只包含一個對像為止。

以下是一個簡單的示例：

假設有一組數據：A(1,1)，B(2,1)，C(4,3)，D(5,4)和E(6,5)。

自下而上的層次聚類方法會從單獨的對象開始，將A和B合併成一個簇，然後再將C、D和E合併成另一個簇。這樣就得到了一個包含兩個簇的樹狀圖。可以根據需要通過樹狀圖中的劃分來確定聚類的數量。

自上而下的層次聚類方法則從整個數據集開始。將整個數據集視為一個簇，然後將它分成兩個簇。每個簇會再次分裂成更小的簇，如此循環，直到每個簇只包含一個對象。這樣也會得到一個包含兩個簇的樹狀圖。

需要注意的是，在層次聚類中，選擇哪種方法取決於數據的結構和目標。通常，自下而上的方法更適合於密集的簇，而自上而下的方法更適合於稀疏的簇。此外，還可以根據需要使用不同的距離度量方法來計算簇之間的距離。