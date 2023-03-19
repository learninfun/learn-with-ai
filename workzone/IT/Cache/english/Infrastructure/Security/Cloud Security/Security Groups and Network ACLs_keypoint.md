

Security Groups:
- Firewall that controls inbound and outbound traffic for EC2 instances.
- Provide stateful protection.
- Can be associated with multiple instances.
- Can be used to control traffic within VPC or between VPCs.
- Can be used to control traffic from internet or specific IP addresses.
- Can be updated on the fly.
- Can be used to configure egress filters.

Network ACLs:
- Firewall that acts like a subnet border.
- Provide stateless protection.
- Can be associated with one subnet.
- Control inbound and outbound traffic at a subnet level.
- Can be used to block or allow specific IP addresses or protocols.
- Can't be updated on the fly.
- Can be used to configure ingress filters.