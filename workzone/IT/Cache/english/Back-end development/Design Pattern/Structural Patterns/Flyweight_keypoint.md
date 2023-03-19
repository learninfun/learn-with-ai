

1. Flyweight pattern is a structural design pattern.
2. It is used to optimize code that involves a large number of small and lightweight objects.
3. The pattern achieves optimization by sharing common data between multiple objects instead of creating separate copies of it.
4. The common data is maintained in a centralized data structure called the flyweight factory.
5. Flyweight pattern is particularly useful when the application has a large number of identical objects that consume a lot of memory.
6. It separates the intrinsic state (attributes that are common to all objects) from the extrinsic state (attributes that differ between objects).
7. The pattern allows for flexible representation of objects, as extrinsic state can be easily assigned or modified when needed.
8. In Java, the Flyweight pattern is implemented using interfaces, abstract classes or inheritance.
9. The pattern can be combined with other design patterns such as singleton and factory method.
10. Flyweight pattern can contribute to improved performance and reduced memory consumption, but may also introduce complexity to the code.