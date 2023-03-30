1. What is the Factory Pattern used for?
Answer: The Factory Pattern is used to create objects without specifying the exact class of object that will be created.

2. What are the components of the Factory Pattern?
Answer: The components of the Factory Pattern include: a Factory interface or class, a Concrete Factory that implements the Factory, and various Product classes that the Concrete Factory creates.

3. What is the difference between a Simple Factory and a Factory Method?
Answer: A Simple Factory creates objects using a single method, while a Factory Method delegates object creation to a subclass.

4. How does the Factory Pattern help to decouple dependent code?
Answer: The Factory Pattern allows dependent code to create objects without knowing anything about the concrete classes being used to create them, which means that changes to the concrete classes will not affect the dependent code.

5. What are some disadvantages of using the Factory Pattern?
Answer: Some disadvantages of using the Factory Pattern include: increased complexity of code, increased number of classes, and decreased performance due to the additional method calls and object creation.