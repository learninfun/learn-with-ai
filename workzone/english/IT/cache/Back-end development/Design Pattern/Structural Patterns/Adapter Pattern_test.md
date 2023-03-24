

1. What is Adapter Pattern?
Answer: Adapter Pattern is a structural design pattern that allows two incompatible interfaces to work together. It converts the interface of a class into another interface that the client expects.

2. What are the two main classes involved in Adapter Pattern?
Answer: The two main classes involved in Adapter Pattern are the Adapter class and the Adaptee class. The Adapter class implements the target interface and uses an instance of the Adaptee class to provide the required functionality.

3. When should Adapter Pattern be used?
Answer: Adapter Pattern should be used when you need to make two incompatible interfaces work together without modifying the existing code. It is also useful when you want to reuse an existing class, but its interface is not compatible with the rest of the system.

4. What are the advantages of using Adapter Pattern?
Answer: The advantages of using Adapter Pattern are: It allows two incompatible interfaces to work together, it enables reuse of existing classes, it helps in providing a stable interface to clients, it allows for flexibility in design and supports the open/closed principle of design.

5. Can Adapter Pattern be used in both class and object level?
Answer: Yes, Adapter Pattern can be used in both class and object level. In class level, a class adapts its interface to the target interface, while in object level, the adapter is created as an object that wraps around the Adaptee object and implements the target interface.