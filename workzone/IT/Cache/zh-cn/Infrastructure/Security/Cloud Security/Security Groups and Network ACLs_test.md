

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