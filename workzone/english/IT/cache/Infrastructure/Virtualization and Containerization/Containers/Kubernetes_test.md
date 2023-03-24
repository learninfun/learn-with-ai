

1. What is the role of Kubernetes Master Node in the Kubernetes architecture?
Answer: The Master Node is responsible for managing and scheduling the workloads, maintaining the state of the cluster, and responding to the API server.

2. What is a Pod in Kubernetes?
Answer: A Pod is the smallest and simplest unit in the Kubernetes object model, which represents a single instance of a running process in a cluster.

3. What are Kubernetes Labels and Selectors used for?
Answer: Labels are key-value pairs used to organize and manage Kubernetes objects, while Selectors are used to identify and group objects that share common labels.

4. What is a Kubernetes ReplicaSet and what is its purpose?
Answer: A ReplicaSet is a Kubernetes object that ensures a specified number of replicas of a Pod are running at all times, thus ensuring high availability and scalability.

5. What is the difference between a Kubernetes Deployment and a ReplicaSet?
Answer: A Deployment is responsible for managing and updating a ReplicaSet and provides additional features such as rollback and scaling, while a ReplicaSet ensures a specified number of replicas of a Pod are running at all times.