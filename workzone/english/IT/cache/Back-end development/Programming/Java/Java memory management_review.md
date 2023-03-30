1. What is the purpose of the Java Garbage Collector and how does it work?
Answer: The Java Garbage Collector is responsible for automatically freeing up memory by identifying unused objects and reclaiming their memory. It works by periodically scanning the heap for objects that are no longer being referenced and freeing up their memory.

2. How can you control the amount of memory used by a Java application?
Answer: You can control the amount of memory used by a Java application by setting the JVM heap size using the -Xmx and -Xms flags. The -Xmx flag sets the maximum heap size, while the -Xms flag sets the initial heap size.

3. What is a memory leak in Java and how can it be prevented?
Answer: A memory leak in Java occurs when objects are continuously created but not properly destroyed, causing the heap to fill up and eventually causing the application to crash. Memory leaks can be prevented by ensuring that all objects are properly destroyed when they are no longer needed and by monitoring the memory usage of the application.

4. How does the Java Heap differ from the Stack in terms of memory management?
Answer: The Java Heap is used to store objects that are created dynamically during runtime, while the Stack is used to store method calls and local variables. The Heap is managed automatically by the Garbage Collector, while the Stack is managed automatically by the JVM.

5. What is the difference between a strong and weak reference in Java?
Answer: A strong reference in Java holds a direct reference to an object and prevents it from being garbage collected. A weak reference, on the other hand, allows the object to be garbage collected if there are no other strong references to it. Weak references can be useful when caching objects that are expensive to create but may not be needed for long periods of time.