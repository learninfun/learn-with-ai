+++
title = "安全性群組和網路存取控制列表（ACL）"
weight = "2"
+++
## 習題預習
{{< ask_chatgpt >}}
給我5題Security Groups and Network ACLs的中文問題
{{< /ask_chatgpt >}}



1. 什麼是安全群組（Security Groups）？可以如何使用它來保護您的網絡？
2. 什麼是網絡存取控制列表（Network ACL）？它們如何區別於安全群組？
3. 安全群組和網絡存取控制列表在提供安全性方面有何不同？哪種更有效？
4. 在AWS中如何創建和管理安全群組和網絡存取控制列表？
5. 安全群組和網絡存取控制列表的規則是如何工作的？它們可以設置那些安全性功能？   

## 說明知識
{{< ask_chatgpt >}}
以中文說明Security Groups and Network ACLs並舉例
{{< /ask_chatgpt >}}



Security Groups和Network ACLs是AWS中用於網路安全的兩種機制。

Security Groups是一種虛擬防火牆，將一個或多個EC2實例分配到一個安全組中。每個安全組中定義了一些入站和出站的規則，只允許特定的流量進出EC2實例。Security Groups是基於實例級別的，因此它們對根據需要的EC2實例進行流量管理非常有用，並且非常簡單易用。

例如，一個Web應用程序需要在80端口上接受HTTP流量，但需要在其他端口上阻止所有流量。在這種情況下，可以在Web服務器的安全組中添加一條入站規則，允許80端口的HTTP流量，然後阻止所有其他入站流量。

Network ACLs是另一種虛擬防火牆，通過定義一組規則，控制VPC子網路（Subnets）中所有流入和流出EC2實例的流量。Network ACLs是基於子網路配置的，因此在特定的子網路上定義一些規則對子網路內的所有實例都有效。

例如，一個企業需要用一個子網路作為Web服務器群集的位置，另一個子網路會被用作數據庫服務器群集的位置，需要隔離二者。在這種情況下，可以設置Network ACLs，允許Web子網路上的HTTP流量，但不允許訪問數據庫子網路。同時，在數據庫子網路上，列出瀏覽器通過Web服務器訪問的流量是允許的，但是其他流量都被阻止。這將確保Web和數據庫子網路保持隔離，不會受到對方群集上的攻擊。   

## 彙總重點
{{< ask_chatgpt >}}
以中文條列Security Groups and Network ACLs的重點
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
以中文給我5題Security Groups and Network ACLs的中等難度問題，並在後面列出答案
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

