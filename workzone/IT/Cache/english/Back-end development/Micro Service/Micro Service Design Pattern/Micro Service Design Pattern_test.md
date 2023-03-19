

1. What is the purpose of the "service registry" pattern in microservices? 

Answer: The service registry pattern is used to provide a centralized location where microservices can register themselves and discover other services they depend on. This allows for flexibility in deployment and enables dynamic scaling within a microservices architecture.

2. How does the "circuit breaker" pattern help avoid cascading failures in microservices? 

Answer: The circuit breaker pattern is used to prevent failing microservices from causing cascading failures throughout the system. It involves introducing a layer of logic that monitors requests to a service and trips a circuit breaker if the service is experiencing too many failures. This reduces the load on the service and prevents other microservices from being affected by its failure.

3. What is the difference between the "gateway" pattern and the " API gateway" pattern in microservices? 

Answer: The basic gateway pattern is used to provide a single entry point for incoming requests from clients. It can be used to manage security and routing of requests to different microservices. An API gateway builds on the basic gateway pattern and is specifically designed to manage requests to APIs. It provides features such as caching, rate limiting, and versioning of API endpoints.

4. How does the "event sourcing" pattern store and retrieve data in microservices? 

Answer: The event sourcing pattern involves storing data in the form of a sequence of events that describe changes to an entity over time. This allows for easy reconstruction of the entity's state at any point in time. When a microservice needs to retrieve data, it can query the event log and reconstruct the entity's state by applying the events in order.

5. What is the role of the "polyglot persistence" pattern in microservices? 

Answer: The polyglot persistence pattern is used to allow microservices to use multiple types of databases to store data. This allows each microservice to choose the most appropriate database for their specific needs, such as scalability or consistency. It also allows for easier data migration and reduces the risk of data loss due to a single point of failure.