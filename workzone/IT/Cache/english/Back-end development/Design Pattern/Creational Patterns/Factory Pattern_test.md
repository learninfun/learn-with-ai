

1. What is the Factory Pattern?
Answer: Factory Pattern is a creational design pattern that provides an interface for creating objects of a class, but lets subclasses decide which class to instantiate.

2. What are the advantages of using the Factory Pattern?
Answer: Some of the advantages of using the Factory Pattern include improved code modularity, flexibility and extensibility, separation of concerns, enhanced code readability, and easier testing and debugging.

3. What is the difference between a Simple Factory and a Factory Method?
Answer: A Simple Factory creates objects of the same type, while a Factory Method creates objects of different types. Additionally, Simple Factory uses static methods, while Factory Method uses virtual methods.

4. What are some common use cases for the Factory Pattern?
Answer: Some common use cases for the Factory Pattern include object creation scenarios where object instantiation can be complex, such as when objects require a lot of configuration or initialization, or when there are different dependencies or versions of the object.

5. How does the Factory Pattern handle object creation and instantiation?
Answer: The Factory Pattern delegates object creation to its subclasses, which are responsible for instantiating and returning the appropriate object based on its specific logic and requirements. This allows for greater flexibility and abstraction in the code, as well as ease of maintenance and reuse.