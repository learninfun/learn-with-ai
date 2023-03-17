

1. What is the difference between ClassCastException and TypeCastException in Java generics?
Answer: ClassCastException occurs when we try to cast a non-compatible object to a generic type, whereas TypeCastException occurs when we try to cast a generic type to a non-compatible type.

2. Can we use primitive data types with Java generics?
Answer: No, we cannot use primitive data types with Java generics. We need to use their corresponding object types instead.

3. What is the difference between wildcard and bounded type parameters in Java generics?
Answer: Wildcard type parameters allow us to use any type as an argument, while bounded type parameters restrict the argument to a specific type or its subtypes.

4. Can we create an instance of a generic type in Java?
Answer: No, we cannot create an instance of a generic type directly in Java. We need to create an instance of a concrete type and then cast it to the generic type.

5. What is the purpose of type erasure in Java generics?
Answer: Type erasure is a process of replacing generic type parameters with their upper bounds during compilation. This is done to ensure backward compatibility with pre-generics code and to make the code more efficient.