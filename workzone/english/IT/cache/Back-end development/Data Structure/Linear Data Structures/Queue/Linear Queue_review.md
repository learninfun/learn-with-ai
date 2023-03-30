1. What is a Linear Queue?
Answer: A linear queue is a type of queue data structure in which the first item inserted is the first item removed (FIFO).

2. What is the difference between a Linear Queue and a Circular Queue?
Answer: A Linear Queue has a fixed size, whereas a Circular Queue can be resized dynamically. Also, in a Circular Queue, the first and last item points to each other, forming a circle, whereas a Linear Queue has a single-directional flow.

3. What are the operations that can be performed on a Linear Queue?
Answer: Enqueue (insertion) and Dequeue (removal) are the primary operations performed on a Linear Queue. Other operations include front (display the front item), rear (display the rear item), and isFull/isEmpty (check if the queue is full or empty).

4. What happens when a Linear Queue becomes full?
Answer: When a Linear Queue becomes full, a new item cannot be inserted until some items are removed or dequeued. This is because, in a Linear Queue, items are inserted at the rear end and removed from the front end.

5. Does a Linear Queue support random access?
Answer: No, a Linear Queue only supports sequential access, meaning that items can only be inserted and removed in a linear fashion, starting from the front or rear end. To access a specific item, one has to traverse through all the preceding items.