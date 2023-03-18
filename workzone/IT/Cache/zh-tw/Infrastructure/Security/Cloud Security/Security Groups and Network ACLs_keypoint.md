

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