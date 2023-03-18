+++
title = "安全性群组和网路存取控制列表（ACL）"
+++
## 习题预习
{{< ask_chatgpt >}}
给我5题Security Groups and Network ACLs的问题
{{< /ask_chatgpt >}}



1. What is the primary difference between Security Groups and Network ACLs?
2. How do Security Groups and Network ACLs work together to protect a network?
3. What are the limitations of Security Groups and Network ACLs?
4. Can Security Groups or Network ACLs be used to protect against insider threats? If yes, how? If no, why not?
5. How do you ensure that Security Groups and Network ACLs are configured correctly to protect your network?   

## 说明知识
{{< ask_chatgpt >}}
说明Security Groups and Network ACLs并举例
{{< /ask_chatgpt >}}



Security Groups和Network ACLs是AWS中用于网路安全的两种机制。

Security Groups是一种虚拟防火墙，将一个或多个EC2实例分配到一个安全组中。每个安全组中定义了一些入站和出站的规则，只允许特定的流量进出EC2实例。Security Groups是基于实例级别的，因此它们对根据需要的EC2实例进行流量管理非常有用，并且非常简单易用。

例如，一个Web应用程序需要在80端口上接受HTTP流量，但需要在其他端口上阻止所有流量。在这种情况下，可以在Web服务器的安全组中添加一条入站规则，允许80端口的HTTP流量，然后阻止所有其他入站流量。

Network ACLs是另一种虚拟防火墙，通过定义一组规则，控制VPC子网路（Subnets）中所有流入和流出EC2实例的流量。Network ACLs是基于子网路配置的，因此在特定的子网路上定义一些规则对子网路内的所有实例都有效。

例如，一个企业需要用一个子网路作为Web服务器群集的位置，另一个子网路会被用作数据库服务器群集的位置，需要隔离二者。在这种情况下，可以设置Network ACLs，允许Web子网路上的HTTP流量，但不允许访问数据库子网路。同时，在数据库子网路上，列出浏览器通过Web服务器访问的流量是允许的，但是其他流量都被阻止。这将确保Web和数据库子网路保持隔离，不会受到对方群集上的攻击。   

## 汇总重点
{{< ask_chatgpt >}}
条列Security Groups and Network ACLs的重点
{{< /ask_chatgpt >}}



Security Groups:
1. 为EC2 Instance提供安全性和隔离性。
2. 基于protocol、port、以及IP address进行安全限制。
3. 预设拒绝所有流量，只开放明确授权的流量。
4. 可控制进入或离开Instance的流量。
5. 可将Security Groups与不同subnet或VPC关联。

Network ACLs:
1. 为VPC或subnet提供安全性。
2. 基于protocol、port以及IP address进行安全限制。
3. 以规则集为基础建立，每个规则皆有一定的优先顺序，从高到低依序检查与启用。
4. 预设开启所有流量，需要明确设定拒绝的规则。
5. 控制进入或离开VPC或subnet的流量。
6. 可与subnet或VPC关联。

重点：
1. Security Groups和Network ACLs都是AWS提供的安全性控制机制，分别用于保护EC2 Instance和VPC或subnet。
2. Security Groups控制进出EC2 Instance的流量，而Network ACLs控制进出VPC或subnet的流量。
3. Security Groups与Instance一一对应，而Network ACLs与VPC或subnet一一对应。
4. Security Groups以授权为基础，预设拒绝所有流量；而Network ACLs以拒绝为基础，预设开启所有流量。
5. Security Groups的限制精细，只能设定允许的流量；而Network ACLs的限制较宽松，可以设定拒绝的流量。
6. Security Groups实现时，可以参考需求进行改变，当需要修改规则时只需修改Security Groups就好了；而Network ACLs则是基于subnet和VPC建立，且规则要依一定的先后次序配置。   

## 知识测验
{{< ask_chatgpt >}}
给我5题Security Groups and Network ACLs的中等难度问题，并在后面列出答案
{{< /ask_chatgpt >}}



1. 哪些协议可以被 Security Group 过滤掉？

答案：Security Group 可以过滤掉所有来源或目标 IP 地址、TCP/UDP 端口和 ICMP 协议，但是不能过滤内容。

2. 您如何设置安全性组允许两个 VPC 之间的电脑通讯？

答案：要允许 VPC 之间的电脑通讯，请在每个 VPC 中设置安全性组，以允许来自对方 VPC 的流量。然后，将安全性组附加到 VPC 的子网路中。

3. 您如何设置安全组允许 EC2 实例之间通讯？

答案：要允许 EC2 实例之间通讯，请在每个安全性组中创建规则，允许来自对方 IP 地址范围的流量。然后，将这些安全性组附加到您的 EC2 实例。

4. 您如何跟踪网络中的流量，以便确定哪些流量被阻挡或允许？

答案：您可以在 VPC 流量镜像中创建一个流量镜像会话，用于复制来自任何网路接口的流量。然后，您可以定向该流量镜像会话到一个或多个 EC2 实例或网路资料库。使用设定管理器软件追踪流量。

5. 您如何确保网路 ACL 不会阻止升级或配置更改的流量？

答案：请确保路由表中有一个组织升级的IP地址，并且有适当的入站和出站规则。监控网路ACL维持最新资料，快速回应网路问题，确保网路资源的安全性和可靠性。   

## 网络数据
{{< ask_chatgpt >}}
给我5篇Security Groups and Network ACLs的网络数据
{{< /ask_chatgpt >}}



1. What are Security Groups and Network ACLs?
https://aws.amazon.com/security/identity-amp-compliance/amazon-vpc/

This article from Amazon's official website describes what Security Groups and Network ACLs are, how they differ, and how they work together. It also provides information on how they are used in Amazon Web Services (AWS) Virtual Private Cloud (VPC).

2. AWS Security Groups vs Network ACLs: What’s the Difference?
https://www.twistlock.com/2017/10/04/aws-security-groups-vs-network-acls-whats-difference/

Twistlock offers a comprehensive explanation of the differences between Security Groups and Network ACLs, including their primary functions, strengths, and limitations. The article also shares internal links to resources for learning more.

3. AWS Security Groups and Network ACLs: An Overview
https://blog.gigamon.com/2018/10/23/aws-security-groups-and-network-acls-an-overview/

Gigamon's blog provides an overview of both Security Groups and Network ACLs in AWS. It explains how they work and how they are used to secure applications hosted in VPC. The author also discusses best practices and tips for optimizing security.

4. What are Amazon VPC security groups and network ACLs?
https://searchcloudsecurity.techtarget.com/definition/Amazon-VPC-Security-Groups-and-Network-ACLs

This techtarget article discusses the role of Security Groups and Network ACLs in AWS, providing definitions and examples of each. It also provides an analysis of how these tools can be used to protect cloud-based infrastructure.

5. Understanding AWS Security Groups and Network ACLs
https://securitytrails.com/blog/aws-security-groups-and-network-acls

Security Trails provides an analysis of Security Groups and Network ACLs in AWS. The article provides a brief overview, comparing and contrasting the two, and then delves into the specifics of each, including how they function and how they are implemented.   

