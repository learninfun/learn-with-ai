Terraform is an open-source infrastructure as code (IaC) tool that enables users to manage cloud infrastructure in a declarative way. It uses configuration files to define the desired state of the infrastructure, and then applies those configurations to create, modify or delete resources.

For example, let's say we want to create a web application infrastructure on AWS. We can define the desired infrastructure using Terraform configuration language in a file named "main.tf" which looks like:

```
provider "aws" {
  region = "us-west-2"
}

resource "aws_instance" "web" {
  ami           = "ami-0947d2ba12ee1ff75"
  instance_type = "t2.micro"
  key_name      = "mykey"
  tags = {
    Name = "web-server"
  }
}

resource "aws_security_group" "web_group" {
  name_prefix = "web"
  ingress {
    from_port   = 80
    to_port     = 80
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }
}
```

This code defines a provider (AWS) and two resources: an EC2 instance and a security group. The EC2 instance will use the "t2.micro" instance type and "mykey" key pair. The security group will allow incoming traffic on port 80 from any IP address.

After defining the infrastructure, we can use the Terraform CLI to initialize our working directory and apply our configuration. Terraform will analyze the configuration, plan and create the resources accordingly. For instance, running the following commands:

```
terraform init
terraform apply
```

In a few moments, Terraform will create the EC2 instance and the security group on AWS. This provides a repeatable and predictable way of deploying and managing infrastructure in cloud environments.