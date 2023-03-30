1. What is Java reflection and how does it work?

Answer: Java reflection is an API in Java that allows code to access and manipulate the metadata of classes, fields, methods, and constructors at runtime. It works by providing a way for code to introspectively examine and modify the internal structure of a program during execution.

2. What is the difference between static and dynamic class loading using reflection?

Answer: Static class loading refers to the process of loading classes at compile time, whereas dynamic class loading occurs at runtime. With reflection, both static and dynamic class loading can be used interchangeably to access class metadata and instantiate objects.

3. How can you invoke a private method using reflection?

Answer: To invoke a private method using reflection, you first need to call the setAccessible() method on the method object, which will allow the method to be invoked even if it is inaccessible based on its visibility modifier.

4. What is the difference between getClass() and getDeclaredClasses() in reflection?

Answer: getClass() returns the runtime class of an object, while getDeclaredClasses() returns an array of all the inner classes declared within a class. getClass() is a method defined in the Object class, while getDeclaredClasses() is a method defined in the Class class.

5. What is the purpose of the java.lang.reflect.Method class?

Answer: The java.lang.reflect.Method class is used in reflection to represent a single method of a class. It provides methods for accessing the method's name, return type, parameter types, annotations, and modifiers. Using the Method class, you can invoke methods dynamically at runtime, regardless of their access level.