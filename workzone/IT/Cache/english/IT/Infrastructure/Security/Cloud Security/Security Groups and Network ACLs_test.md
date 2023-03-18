

1) Q: What is the main difference between a Security Group and a Network ACL?
   A: Security Groups control inbound and outbound traffic for an EC2 instance based on allow or deny rules, whereas Network ACLs control traffic at the subnet level based on sequential numbered rules. 

2) Q: How many Security Groups can be assigned to an EC2 instance, and how many Network ACLs can be assigned to a subnet?
   A: An EC2 instance can have multiple Security Groups assigned to it, while a subnet can only have one Network ACL assigned to it. 

3) Q: Can a Security Group reference another Security Group within the same VPC?
    A: Yes, security groups can reference other security groups in the same VPC to allow traffic to be allowed or denied between instances that belong to different security groups. 

4) Q: Can Network ACLs be used to restrict traffic between subnets in the same VPC?
   A: Yes, Network ACLs can be used to restrict traffic between subnets in the same VPC by creating rules that block traffic coming in or going out of a particular subnet. 

5) Q: What happens when a Security Group is deleted?
   A: When a Security Group is deleted, all rules associated with the Security Group are also deleted. And if any instance had that Security Group assigned to it, it will lose its configuration for that Security Group, losing connectivity outside the VPC.