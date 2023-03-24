

1. What is a linear queue and how does it differ from other types of queues?
Answer: A linear queue is a collection of elements arranged in a linear order where insertion is done at one end and deletion is done at the other end. This differs from other types of queues such as circular queues, where the last element is connected to the first element, creating a circular structure.

2. How do you check if a linear queue is empty or full?
Answer: To check if a linear queue is empty, we can simply check if the front and rear pointers are equal. To check if a linear queue is full, we can compare the rear pointer with the maximum size limit of the queue element.

3. How do you implement a linear queue using an array in C++?
Answer: We can implement a linear queue using an array in C++ by defining a fixed-size array and two pointers, front and rear. The front pointer indicates the first element in the queue, and the rear pointer indicates the last element. An element is added to the queue by incrementing the rear pointer and placing the element at this position, and an element is removed by incrementing the front pointer.

4. How do you calculate the time complexity for inserting an element in a linear queue?
Answer: The time complexity for inserting an element in a linear queue is O(1) because we only need to increment the rear pointer and add the element to the end of the queue.

5. How can you prevent overflow and underflow in a linear queue?
Answer: To prevent overflow in a linear queue, we can check if the rear pointer is equal to the maximum size limit of the queue element before inserting an element. To prevent underflow, we can check if the front pointer is equal to the rear pointer before deleting an element.