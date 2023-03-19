

1. What is the Command pattern and what problem does it solve?
Answer: The Command pattern is a design pattern that encapsulates an action as an object, allowing it to be decoupled from the object that requested it. It solves the problem of coupling the object that invokes an operation with the object that performs the operation. 

2. How does the Command pattern differ from the Strategy pattern?
Answer: The Command pattern is focused on encapsulating an operation as an object that can be passed around, whereas the Strategy pattern is focused on encapsulating interchangeable algorithms that can be used to implement a specific behavior. 

3. What are the four key components of the Command pattern?
Answer: The four key components of the Command pattern are the Command interface, the ConcreteCommand class(es), the Receiver object, and the Invoker object. 

4. How does the Command pattern support undo/redo functionality?
Answer: By storing a history of executed commands, the Command pattern allows you to undo or redo previous actions by replaying or reversing individual commands from the history. 

5. What are some common use cases for the Command pattern?
Answer: Some common use cases for the Command pattern include implementing menu items, toolbar buttons, and keyboard shortcuts in a graphical user interface, implementing multi-level undo/redo functionality, and implementing a transactional system for database updates.