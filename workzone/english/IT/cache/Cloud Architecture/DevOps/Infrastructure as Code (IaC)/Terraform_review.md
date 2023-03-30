1. What is Terraform and what problem does it solve?
Answer: Terraform is a tool for building, changing, and versioning infrastructure safely and efficiently. It solves the problem of infrastructure automation, specifically by providing a way to define infrastructure as code.

2. What are some advantages of using Terraform?
Answer: Some advantages of using Terraform include being able to manage infrastructure across multiple providers, having the ability to version control infrastructure, and being able to apply changes safely and efficiently.

3. How does Terraform work?
Answer: Terraform works by reading configuration files, usually in the form of .tf files, which describe the desired state of infrastructure. From there, Terraform creates an execution plan, which details the steps required to achieve the desired state. Finally, Terraform applies the plan, making changes to the infrastructure as needed.

4. How does Terraform handle dependencies between resources?
Answer: Terraform handles dependencies between resources by creating a dependency graph based on the order in which resources are declared in the configuration file. Resources listed earlier in the file will be created first, allowing later resources to reference them.

5. What are some best practices for using Terraform?
Answer: Some best practices for using Terraform include using modules to abstract infrastructure, keeping configuration files versioned in source control, using a consistent naming convention for resources, and minimizing the use of hardcoded values in configuration files.