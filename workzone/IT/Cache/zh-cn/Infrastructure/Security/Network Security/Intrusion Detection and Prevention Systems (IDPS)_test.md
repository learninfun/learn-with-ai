

1. 什么是IDPS典型的部署位置，以保护企业免受内部和外部攻击？

答：IDPS通常放置在内部网路或DMZ上，以保护企业的内部和外部网络免受攻击。

2. 什么是“false positive”，在IDPS中的意义是什么？

答：“False positive”指的是IDPS认为发现了攻击，但实际上并没有攻击发生。在IDPS中，“false positive”通常会被视为一个问题，因为它会导致资源浪费和不必要的警报。

3. IDPS中的“signature-based”和“anomaly-based”检测系统的差异在哪里？

答：在IDPS中，“signature-based”检测系统是基于已知攻击程式的固定特征来检测攻击的。而“anomaly-based”检测系统基于系统正常行为建立基线，并检测异常行为。它们的差异在于，前者只能侦测已知攻击，后者则可以侦测未知攻击。

4. IDPS中的“inline”和“passive”模式有什么区别？

答：在IDPS中，“inline”模式会主动阻拦攻击，而“passive”模式只会监视流量并生成报告。通常，"inline" 模式比 "passive" 模式具有更好的保护能力，但也可能因误报或阻止了合法流量而导致问题。

5. 当攻击者进行“fragmentation攻击”时，IDPS如何侦测和防御？

答：Fragmentation攻击是指将大型封包分片，绕过检测系统的攻击方式。IDPS可以使用“Reassembly-free Packet Inspection”技术，对每个分片进行检测并重建封包，以检验它是否符合攻击规则。此外，IDPS还可以使用“Fragmentation Attack Detection/Prevention”模组来检测和防御分片攻击。