

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