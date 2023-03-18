

1. What is the difference between process and thread in Java? 
Answer: A process is an independent unit that executes a program, while a thread is a lightweight unit of execution within a process.

2. What is synchronization in Java and why is it important in multithreading? 
Answer: Synchronization is the mechanism of controlling the access of multiple threads to the shared resources. Without proper synchronization, concurrent access to shared resources can lead to errors and inconsistency in the program's output.

3. What is a deadlock in Java multithreading and how can it be avoided? 
Answer: A deadlock is a situation where two or more threads are blocked, waiting for each other to release the shared resources. Deadlocks can be avoided by using a proper locking order, limiting the use of shared resources, and avoiding nested locks.

4. What is the purpose of the volatile keyword in Java concurrency? 
Answer: The volatile keyword is used to indicate that a variable may be modified by multiple threads concurrently. It ensures that the updates made by one thread are immediately visible to all other threads, preventing synchronization errors.

5. What is a thread pool in Java and how can it improve performance in multithreaded applications? 
Answer: A thread pool is a collection of pre-created threads that are available for executing tasks. Thread pools improve performance by avoiding the overhead of creating and destroying threads frequently, and by limiting the number of threads that can be created, preventing the system from being overwhelmed with too many threads.