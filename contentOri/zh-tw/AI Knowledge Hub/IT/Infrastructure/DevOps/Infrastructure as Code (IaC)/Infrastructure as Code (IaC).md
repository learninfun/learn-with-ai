+++
title = "基礎架構即代碼 (IaC)"
weight = "3"
+++
## 習題預習
{{< ask_chatgpt >}}
給我5題Infrastructure as Code (IaC)的中文問題
{{< /ask_chatgpt >}}



1. 什麼是 Infrastructure as Code (IaC)？
2. IaC的優勢是什麼？
3. IaC使用哪些技術工具？
4. 如何實現IaC的流程與方法？
5. 如何進行IaC的測試與驗證？   

## 說明知識
{{< ask_chatgpt >}}
以中文說明Infrastructure as Code (IaC)並舉例
{{< /ask_chatgpt >}}



Infrastructure as Code (IaC)是指利用程式碼去管理與自動化整個資訊基礎架構的建置、設定和管理。這種方法是為了強調應用程式和基礎架構的平等性，而非像以往一樣只關注應用程式。

IaC的優點在於：

1.自動化程式碼：IaC可以使整個環境編碼化，讓管理者只需編寫程式碼，就可以簡化複雜的設置。

2.易於管理：程式碼可以被存儲在版本控制系統中，用於檢查更改、回顧過去記錄以及測試。

3.可靠性：所使用的檔案是相同的，每次重建環境時都會生成確定的執行檔，且不會遺漏任何環境設置。

4.可擴展性：使用IaC可以使新的資源快速部署，而替換現有的系統也會更容易。

以下是幾個IaC的應用範例：

1. CloudFormation: CloudFormation是Amazon Web Services (AWS)的基於模板的服務，可用於定義基礎設施作為項目。

2. Ansible: Ansible是一種開放原始碼自動化工具，可用於配置、部署和管理基礎設施。它支援文本編輯器、Git、Vagrant等工具。

3. Terraform: Terraform 是 HashiCorp 公司開發的可讓管理員定義基礎設施的工具。它提供了豐富的功能，支援多種基礎架構提供商，包括 Amazon Web Services、Google Cloud 和 Azure。

4. Chef: Chef是一種開源的系統管理自動化工具，它使用基於 Ruby 的語言來描述系統配置。它提供了一個命令式管理架構和配置的抽像層。

5. Puppet: Puppet是一種自動化IT土建設的解決方案，它可以編寫程式碼來部署、配置和管理伺服器，並自動檢測和修補偏差。   

## 彙總重點
{{< ask_chatgpt >}}
以中文條列Infrastructure as Code (IaC)的重點
{{< /ask_chatgpt >}}



1. 自動化：

利用IaC，可以自動建立、配置和管理IT設施，從而減少了人工干預，消除了人為錯誤，進一步提高了生產力和效率。

2. 代碼化：

IaC使得建立和管理IT設施成為了代碼化，這意味著可以更好地管理，維護和優化IT基礎架構，同時也增加了可讀性和可重用性。

3. 適應性：

IaC使得IT設施在不同的環境中產生變化時更容易進行管理和更改，這種適應性可大大提高系統的可攜性和可擴展性。

4. 清晰的文檔：

IaC產生了清晰的文檔，使得IT人員可以更快地查找和理解系統的功能，從而更快地發現問題和解決問題。

5. 版本控制：

IaC使得IT設施的版本控制更易於管理和控制，從而簡化了維護。

6. 增強安全性：

IaC可以幫助IT人員更好地實現安全性，從而保護企業的數據和應用不受威脅。

7. 優化IT管理：

IaC可以幫助IT管理更好地管理和自動化IT設施，從而降低了成本和風險。

8. 流程優化：

IaC可以幫助IT人員建立和優化流程，從而提高生產力和效率。   

## 知識測驗
{{< ask_chatgpt >}}
以中文給我5題Infrastructure as Code (IaC)的中等難度問題，並在後面列出答案
{{< /ask_chatgpt >}}



1. 在Terraform中如何使用變數來依據多個條件來決定是否新增resource？
答: 可以使用conditional expression，例如:
```
resource "aws_instance" "example" {
  count = var.create_instance ? 1 : 0
  instance_type = "t2.micro"
  // 其他設定
}
```

2. 如何在Ansible的playbook中使用模組來安裝Apache服務及相關模組？
答: 可以使用apt模組(只有在Debian系統上有用)，例如:
```
- name: 安裝Apache
  apt:
    name: apache2
    state: present

- name: 安裝PHP支援
  apt:
    name: libapache2-mod-php
    state: present
```

3. 如何在Chef中使用resource來設定Linux系統上的防火牆規則？
答: 可以使用iptables resource，例如:
```
installation = search(:node, "name:#{node['firewall']['installation']['node_name']}")
iptables_rule 'port_http' do
  source installation
  action :enable
end
```

4. 如何在Puppet的manifest中使用ERB模板來建立Nginx虛擬主機？
答: 可以在manifest中使用file resource和ERB模板，例如:
```
file { "/etc/nginx/sites-available/${fqdn}":
  content => template('nginx/site.conf.erb'),
}
```

5. 在CloudFormation中如何定義一個S3 Bucket的Life Cycle策略？
答: 可以使用AWS::S3::Bucket resource和aws_s3_bucket_lifecycle_configuration data source，例如:
```
Resources:
  myBucket:
    Type: AWS::S3::Bucket
    Properties:
      BucketName: my-bucket-name
      LifecycleConfiguration:
        Rules:
          - Status: Enabled
            NoncurrentVersionExpirationInDays: 365
            NoncurrentVersionTransition:
              StorageClass: GLACIER
              TransitionInDays: 30
```
   

