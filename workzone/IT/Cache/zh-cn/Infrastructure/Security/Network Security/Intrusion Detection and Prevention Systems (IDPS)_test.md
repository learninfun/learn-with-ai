

1. 什麼是IDPS典型的部署位置，以保護企業免受內部和外部攻擊？

答：IDPS通常放置在內部網路或DMZ上，以保護企業的內部和外部網絡免受攻擊。

2. 什麼是「false positive」，在IDPS中的意義是什麼？

答：「False positive」指的是IDPS認為發現了攻擊，但實際上並沒有攻擊發生。在IDPS中，「false positive」通常會被視為一個問題，因為它會導致資源浪費和不必要的警報。

3. IDPS中的「signature-based」和「anomaly-based」檢測系統的差異在哪裡？

答：在IDPS中，「signature-based」檢測系統是基於已知攻擊程式的固定特徵來檢測攻擊的。而「anomaly-based」檢測系統基於系統正常行為建立基線，並檢測異常行為。它們的差異在於，前者只能偵測已知攻擊，後者則可以偵測未知攻擊。

4. IDPS中的「inline」和「passive」模式有什麼區別？

答：在IDPS中，「inline」模式會主動阻攔攻擊，而「passive」模式只會監視流量並生成報告。通常，"inline" 模式比 "passive" 模式具有更好的保護能力，但也可能因誤報或阻止了合法流量而導致問題。

5. 當攻擊者進行「fragmentation攻擊」時，IDPS如何偵測和防禦？

答：Fragmentation攻擊是指將大型封包分片，繞過檢測系統的攻擊方式。IDPS可以使用「Reassembly-free Packet Inspection」技術，對每個分片進行檢測並重建封包，以檢驗它是否符合攻擊規則。此外，IDPS還可以使用「Fragmentation Attack Detection/Prevention」模組來檢測和防禦分片攻擊。