+++
title = "入侵检测和预防系统（IDPS）"
weight = "2"
+++
## 习题预习
{{< ask_chatgpt >}}
给我5题Intrusion Detection and Prevention Systems (IDPS)的中文问题
{{< /ask_chatgpt >}}



1. 什么是入侵侦测和预防系统（IDPS）？它可以保护哪些资源？
2. IDPS有哪些不同的类型？它们如何运作？
3. IDPS如何检测和阻止来自外部或内部的入侵？
4. IDPS如何与其他安全措施（例如防火墙）协同工作，以提高整体安全性?
5. 如何评估一个IDPS的效能和成本效益？它应该具备哪些核心功能和特性？   

## 说明知识
{{< ask_chatgpt >}}
以中文说明Intrusion Detection and Prevention Systems (IDPS)并举例
{{< /ask_chatgpt >}}



Intrusion Detection and Prevention Systems (IDPS)是一种IT安全解决方案，用于侦测网络或系统中的可疑活动，并根据设定的规则自动阻止该活动，以保护网络安全。

举例来说，一些IDPS可以监控网络流量，并通过比对特定网络行为来检测恶意活动，例如攻击者搜集使用者登陆信息、超时和暴力破解等行为。评估系统提供的系统日志以查找针对主机或应用程序的攻击也是一项常见的IDPS功能。当IDPS检测到可疑活动时，它会触发警报或阻挡该活动。

其他示例包括：

1. 外部入侵检测：监测外部网络传输，以检测入侵或未经授权的存取。

2. 内部入侵检测：监测内部网络流量，以检测企图骇入内部系统的人员。

3. 应用程式安全：监控应用程式以检测异常行为，并防止应用程式被攻击时进行注入式攻击。

4. 主机入侵检测：监测主机以检测异常并防止未经授权的存取。

总的来说，IDPS是保护信息系统安全不可或缺的组件，可以在安全事件发生时提供早期警告，并进行应急反应。   

## 条列重点
{{< ask_chatgpt >}}
以中文条列Intrusion Detection and Prevention Systems (IDPS)的重点
{{< /ask_chatgpt >}}



1. IDPS是安全防御系统的关键组成部分，旨在监测和预防网络攻击和威胁。

2. IDPS系统通常采用多种技术，包括主机和网络基础设施监测、威胁情报分析、流量捕获和分析等。

3. IDPS的主要功能包括实时监测网络流量，检测入侵行为，识别威胁并采取及时的措施进行防御。

4. IDPS发现威胁后，可以自动执行防御措施，如封锁威胁来源、改变访问控制策略、通知安全管理员等。

5. IDPS不仅能够帮助企业实时监测网络安全状态，还可以对安全事件进行记录和报告，为安全管理和合规性提供支持。

6. IDPS系统需要定期更新和维护，以保持对最新威胁的防御能力。

7. 在选择IDPS系统时，需要考虑其适应性、可扩展性、可管理性和可靠性等因素，并根据企业的特定需求进行定制化部署。   

## 知识测验
{{< ask_chatgpt >}}
以中文给我5题Intrusion Detection and Prevention Systems (IDPS)的中等难度问题，并在后面列出答案
{{< /ask_chatgpt >}}



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

