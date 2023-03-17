## 習題預習
{{< ask_chatgpt >}}
給我5題Security Groups and Network ACLs的問題
{{< /ask_chatgpt >}}



1. What is the primary difference between Security Groups and Network ACLs?
2. How do Security Groups and Network ACLs work together to protect a network?
3. What are the limitations of Security Groups and Network ACLs?
4. Can Security Groups or Network ACLs be used to protect against insider threats? If yes, how? If no, why not?
5. How do you ensure that Security Groups and Network ACLs are configured correctly to protect your network?   

## 說明知識
{{< ask_chatgpt >}}
說明Security Groups and Network ACLs並舉例
{{< /ask_chatgpt >}}



Security Groups和Network ACLs是AWS中用於網路安全的兩種機制。

Security Groups是一種虛擬防火牆，將一個或多個EC2實例分配到一個安全組中。每個安全組中定義了一些入站和出站的規則，只允許特定的流量進出EC2實例。Security Groups是基於實例級別的，因此它們對根據需要的EC2實例進行流量管理非常有用，並且非常簡單易用。

例如，一個Web應用程序需要在80端口上接受HTTP流量，但需要在其他端口上阻止所有流量。在這種情況下，可以在Web服務器的安全組中添加一條入站規則，允許80端口的HTTP流量，然後阻止所有其他入站流量。

Network ACLs是另一種虛擬防火牆，通過定義一組規則，控制VPC子網路（Subnets）中所有流入和流出EC2實例的流量。Network ACLs是基於子網路配置的，因此在特定的子網路上定義一些規則對子網路內的所有實例都有效。

例如，一個企業需要用一個子網路作為Web服務器群集的位置，另一個子網路會被用作數據庫服務器群集的位置，需要隔離二者。在這種情況下，可以設置Network ACLs，允許Web子網路上的HTTP流量，但不允許訪問數據庫子網路。同時，在數據庫子網路上，列出瀏覽器通過Web服務器訪問的流量是允許的，但是其他流量都被阻止。這將確保Web和數據庫子網路保持隔離，不會受到對方群集上的攻擊。   

## 彙總重點
{{< ask_chatgpt >}}
條列Security Groups and Network ACLs的重點
{{< /ask_chatgpt >}}



Security Groups:
1. 為EC2 Instance提供安全性和隔離性。
2. 基於protocol、port、以及IP address進行安全限制。
3. 預設拒絕所有流量，只開放明確授權的流量。
4. 可控制進入或離開Instance的流量。
5. 可將Security Groups與不同subnet或VPC關聯。

Network ACLs:
1. 為VPC或subnet提供安全性。
2. 基於protocol、port以及IP address進行安全限制。
3. 以規則集為基礎建立，每個規則皆有一定的優先順序，從高到低依序檢查與啟用。
4. 預設開啟所有流量，需要明確設定拒絕的規則。
5. 控制進入或離開VPC或subnet的流量。
6. 可與subnet或VPC關聯。

重點：
1. Security Groups和Network ACLs都是AWS提供的安全性控制機制，分別用於保護EC2 Instance和VPC或subnet。
2. Security Groups控制進出EC2 Instance的流量，而Network ACLs控制進出VPC或subnet的流量。
3. Security Groups與Instance一一對應，而Network ACLs與VPC或subnet一一對應。
4. Security Groups以授權為基礎，預設拒絕所有流量；而Network ACLs以拒絕為基礎，預設開啟所有流量。
5. Security Groups的限制精細，只能設定允許的流量；而Network ACLs的限制較寬鬆，可以設定拒絕的流量。
6. Security Groups實現時，可以參考需求進行改變，當需要修改規則時只需修改Security Groups就好了；而Network ACLs則是基於subnet和VPC建立，且規則要依一定的先後次序配置。   

## 知識測驗
{{< ask_chatgpt >}}
給我5題Security Groups and Network ACLs的中等難度問題，並在後面列出答案
{{< /ask_chatgpt >}}



1. 哪些協議可以被 Security Group 過濾掉？

答案：Security Group 可以過濾掉所有來源或目標 IP 地址、TCP/UDP 端口和 ICMP 協議，但是不能過濾內容。

2. 您如何設置安全性組允許兩個 VPC 之間的電腦通訊？

答案：要允許 VPC 之間的電腦通訊，請在每個 VPC 中設置安全性組，以允許來自對方 VPC 的流量。然後，將安全性組附加到 VPC 的子網路中。

3. 您如何設置安全組允許 EC2 實例之間通訊？

答案：要允許 EC2 實例之間通訊，請在每個安全性組中創建規則，允許來自對方 IP 地址範圍的流量。然後，將這些安全性組附加到您的 EC2 實例。

4. 您如何跟蹤網絡中的流量，以便確定哪些流量被阻擋或允許？

答案：您可以在 VPC 流量鏡像中創建一個流量鏡像會話，用於複製來自任何網路接口的流量。然後，您可以定向該流量鏡像會話到一個或多個 EC2 實例或網路資料庫。使用設定管理器軟件追蹤流量。

5. 您如何確保網路 ACL 不會阻止升級或配置更改的流量？

答案：請確保路由表中有一個組織升級的IP地址，並且有適當的入站和出站規則。監控網路ACL維持最新資料，快速回應網路問題，確保網路資源的安全性和可靠性。   

## 網路資料
{{< ask_chatgpt >}}
給我5篇Security Groups and Network ACLs的網路資料
{{< /ask_chatgpt >}}



1. What are Security Groups and Network ACLs?
https://aws.amazon.com/security/identity-amp-compliance/amazon-vpc/

This article from Amazon's official website describes what Security Groups and Network ACLs are, how they differ, and how they work together. It also provides information on how they are used in Amazon Web Services (AWS) Virtual Private Cloud (VPC).

2. AWS Security Groups vs Network ACLs: What』s the Difference?
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

