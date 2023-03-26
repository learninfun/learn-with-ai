

1. 在Terraform中如何使用变数来依据多个条件来决定是否新增resource？
答: 可以使用conditional expression，例如:
```
resource "aws_instance" "example" {
  count = var.create_instance ? 1 : 0
  instance_type = "t2.micro"
  // 其他设定
}
```

2. 如何在Ansible的playbook中使用模组来安装Apache服务及相关模组？
答: 可以使用apt模组(只有在Debian系统上有用)，例如:
```
- name: 安装Apache
  apt:
    name: apache2
    state: present

- name: 安装PHP支援
  apt:
    name: libapache2-mod-php
    state: present
```

3. 如何在Chef中使用resource来设定Linux系统上的防火墙规则？
答: 可以使用iptables resource，例如:
```
installation = search(:node, "name:#{node['firewall']['installation']['node_name']}")
iptables_rule 'port_http' do
  source installation
  action :enable
end
```

4. 如何在Puppet的manifest中使用ERB模板来建立Nginx虚拟主机？
答: 可以在manifest中使用file resource和ERB模板，例如:
```
file { "/etc/nginx/sites-available/${fqdn}":
  content => template('nginx/site.conf.erb'),
}
```

5. 在CloudFormation中如何定义一个S3 Bucket的Life Cycle策略？
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
