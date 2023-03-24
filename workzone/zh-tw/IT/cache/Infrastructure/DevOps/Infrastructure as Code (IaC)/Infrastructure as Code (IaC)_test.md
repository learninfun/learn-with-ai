

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
