+++
title = "Windows服务器基础"
weight = "1"
+++
## 习题预习
{{< ask_chatgpt >}}
给我5题Windows Server Basics的问题
{{< /ask_chatgpt >}}



1. Windows Server 2019中的Hyper-V是什么? 它是如何协助虚拟化管理?
2. 在Windows Server 2019中，什么是Active Directory? 它是用来什么目的的?
3. Windows Server 2019中的Group Policy是什么? 它可以用来做什么?
4. 什么是Windows Server 2019中的Remote Desktop Services? 它如何协助使用者远程存取?
5. 在Windows Server 2019中，什么是Failover Clustering? 它可以用来维护什么？   

## 说明知识
{{< ask_chatgpt >}}
说明Windows Server Basics并举例
{{< /ask_chatgpt >}}



Windows Server Basics是指使用Microsoft Windows Server作为伺服器操作系统，提供基本的伺服器功能和服务。以下是一些Windows Server Basics的示例：

1. Active Directory：管理和存储组织中所有的用户和计算机的信息，可以统一管理帐号、密码和设定。

2. File Sharing：分享文件和资源，使多个用户可以访问同一个文件，并可按照访问权限访问数据。

3. DHCP：使用动态主机配置协议为用户分配IP地址，这样可以自动配置用户的网络设置，减少了手动配置的麻烦。

4. DNS：域名解析服务将域名转换成IP地址，使得用户可以通过域名访问网站，而不必记住IP地址。

5. Web Server：提供Web网页的托管和分发，使用户和客户端可以通过网络访问伺服器端提供的讯息和资讯。

6. Email Server：为组织提供电子邮件功能，支援多种邮件协议并可以存储邮件和邮件附件。

以上这些功能和服务都是Windows Server Basics的一些范例。Windows Server也具有更高级的功能和服务，如虚拟化、负载平衡、数据库服务等等。不过，以上这些功能基本涵盖了Windows Server提供的基本功能。   

## 汇总重点
{{< ask_chatgpt >}}
条列Windows Server Basics的重点
{{< /ask_chatgpt >}}



1. 操作系统：Windows Server是一个基于Windows操作系统的伺服器操作系统，它提供了强大的安全性和稳定性。

2. 安全性：Windows Server提供了多种安全功能，包括防火墙、病毒防护、安全更新和身份验证等。

3. 伺服器角色：Windows Server可以配置为不同的伺服器角色，例如档案伺服器、印表机伺服器、网路存储伺服器、Web伺服器、应用程式伺服器等。

4. Active Directory：Windows Server提供了Active Directory服务，它可以管理用户、电脑、群组和其他资源，并提供凭证管理和单一签入等功能。

5. PowerShell：Windows Server支持使用PowerShell进行自动化和批量管理。

6. 虚拟化：Windows Server支持虚拟化技术，例如Hyper-V，这可以让您在同一个硬体平台上运行多个虚拟机器。

7. 储存解决方案：Windows Server提供了多种储存解决方案，包括本机磁碟、网路储存、存取协定和储存空间等。

8. 管理工具：Windows Server提供了多种管理工具，包括Server Manager、Active Directory Users and Computers、PowerShell和Event Viewer等。

9. 远端管理：Windows Server支持使用远端桌面连接和PowerShell进行远端管理。

10. 高可用性：Windows Server提供了多种高可用性解决方案，例如故障转移丛集和网路负载平衡。   

## 知识测验
{{< ask_chatgpt >}}
给我5题Windows Server Basics的中等难度问题，并在后面列出答案
{{< /ask_chatgpt >}}



1. 在Windows Server中，如何设定使用者帐户的密码复杂度限制？ 
答案：请点选“管理工具”，在“本机安全原则”中设定密码原则。然后启用密码必须符合以下条件之一的选项。 

2. Windows Server中，哪个工具可以控制网路执行状态？ 
答案：“服务”工具 

3. 如何设定Windows Server上的防火墙？ 
答案：点击“开始”，输入防火墙，然后点击“Windows防火墙”。选择设置管理员权限，分别开启“入站规则”和“出站规则” 

4. 在Windows Server中，如何添加新的LDAP目录？ 
答案：点击“管理工具”，输入“Active Directory执行中”と选择该选项。然后，使用左侧导航栏，进入“AD FS管理”和“信任关系”选项，添加新的LDAP目录。 

5. 如何设置Windows Server上的DHCP伺服器？ 
答案：点击“管理工具”，并选择 DHCP 指令码。然后输入所需的网路和可用IP地址范围等配置内容。 最后，启用DHCP伺服器。   

