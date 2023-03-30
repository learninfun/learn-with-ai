1) What is the purpose of the Decorator Pattern? 
Answer: The Decorator Pattern allows for adding behavior or features to individual objects at runtime, without affecting other instances of the same class.

2) How do the Decorator Pattern and Inheritance differ in their approach to adding behavior to objects? 
Answer: Inheritance adds behavior to a class at compile time, while the Decorator Pattern adds behavior to an object at runtime.

3) Why is the Decorator Pattern useful for complex systems with many feature combinations? 
Answer: The Decorator Pattern allows for modular and incremental feature additions, rather than creating a separate class for every possible feature combination.

4) How does the Component role in the Decorator Pattern differ from the ConcreteComponent role? 
Answer: The Component defines the abstract interface for objects that can have behavior dynamically added to them, while the ConcreteComponent implements the Component interface and defines the default behavior.

5) What is the disadvantage of using the Decorator Pattern on objects with many features or methods?
Answer: A large number of decorator classes may be required, leading to added complexity and decreased performance.