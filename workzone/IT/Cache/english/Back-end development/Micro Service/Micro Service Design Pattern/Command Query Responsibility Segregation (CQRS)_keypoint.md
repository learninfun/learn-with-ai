

1. Separation of concerns between reading and writing operations
2. Using separate models for querying and updating data 
3. Queries go to the read model which is optimized for performance and scalability 
4. Commands go to the write model which is responsible for updating the system state 
5. Use of event sourcing to maintain a history of state changes 
6. Decoupling of the read and write models to enable independent scaling and evolution 
7. Allows for greater flexibility and better performance in complex applications 
8. Can simplify testing and maintenance 
9. Requires careful design and planning to implement effectively.