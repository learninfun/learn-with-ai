1. What is the purpose of Java generics?
Answer: The purpose of Java generics is to provide type safety at compile-time by enabling the programmer to specify the type of objects that a collection, class or method can hold or operate on.

2. What is a type erasure in Java generics?
Answer: Type erasure is the process where the generic type information is removed from the bytecode and replaced with the appropriate type casts. This is done to preserve backward compatibility with pre-existing Java code that does not use generics.

3. What is a wildcard in Java generics and when would you use it?
Answer: A wildcard is a symbol used to specify an unknown type in Java generics. It can be used in situations where the exact type is unknown or when type variance is required. There are three types of wildcards - '?' for an unknown type, '? extends T' for a type that extends T, and '? super T' for a type that is a superclass of T.

4. What is the difference between bounded and unbounded wildcards in Java generics?
Answer: Bounded wildcards use the syntax '? extends T' or '? super T' to specify a range of allowed types. Unbounded wildcards use the '?' symbol to indicate that any type is accepted. Bounded wildcards allow the programmer to restrict the range of allowed types, while unbounded wildcards provide greater flexibility.

5. What is type variance in Java generics and how is it expressed?
Answer: Type variance refers to the relationship between a generic type and its subtypes or supertypes. Java supports covariance, contravariance, and invariance in generic types, which can be expressed in Java code using the keywords extends, super, and neither, respectively. Covariant types allow the subtype to be more specific than the supertype, whereas contravariant types allow the subtype to be less specific. Invariant types require the type to match exactly.