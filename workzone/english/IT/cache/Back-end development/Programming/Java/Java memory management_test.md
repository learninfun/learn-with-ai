

1. What is the difference between stack memory and heap memory allocation in Java?
A: Stack memory is used for static allocation during compile-time while heap memory is used for dynamic allocation during runtime.

2. What is the function of the Garbage Collector in Java, and how does it manage memory?
A: The Garbage Collector cleans up unused objects from heap memory, freeing up space for other objects to be allocated. It does this by identifying objects that are no longer being used and removing them from memory.

3. What is a memory leak in Java, and how can it be prevented?
A: A memory leak occurs when unused objects are not properly disposed of, causing memory to be consumed unnecessarily over time. It can be prevented by ensuring that all objects are properly disposed of after use and by properly managing the lifecycle of objects.

4. How does the finalize() method in Java relate to memory management?
A: The finalize() method is called by the Garbage Collector before an object is removed from memory. It can be used to perform any necessary cleanup or finalization tasks before the object is deleted.

5. What is the difference between a weak reference and a strong reference in Java, and how are they used in memory management?
A: A strong reference keeps an object in memory until it is explicitly set to null or no longer referenced, while a weak reference allows the object to be garbage collected if there are no other references to it. Weak references are often used for caching data where the data is not critical to the application and can be easily recomputed if needed.