

Chain of Responsibility pattern is one of the behavioral design patterns, which is used to create a chain of objects that can handle different types of requests. 

In this pattern, a series of handlers are created and linked together. Each handler has the ability to handle the request it receives, or it can pass it to the next handler in the chain. The request is passed through the chain until it reaches a handler that can handle it. 

Example: Consider a support system of a software company where customers raise tickets to complain about software errors. When a customer raises a ticket, the first level support team handles the issue. If the issue cannot be resolved by the first level support team, it is escalated to the second level support team. If the second level support team cannot handle the issue, it goes to the third level support team. The chain of responsibility pattern can be used to implement this support system. 

In this scenario, the first level support team is the first handler in the chain. If the first level support team cannot handle the issue, they pass it to the second level support team, which is the second handler in the chain. If the second level support team cannot handle the issue, they pass it to the third level support team, which is the last handler in the chain.

This way, the chain of responsibility pattern provides flexibility and scalability, as different types of handlers can be added or removed from the chain without affecting the entire system.