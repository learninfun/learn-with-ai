1) What is the Memento Pattern?
Answer: The Memento Pattern is a software design pattern that allows the state of an object to be saved and restored at a later time without violating encapsulation.

2) What are the three main components of the Memento Pattern?
Answer: The three main components of the Memento Pattern are the originator, the caretaker, and the memento. The originator creates the memento to store its state, the caretaker stores the memento, and the memento stores the saved state.

3) Can multiple mementos be saved for a single originator object?
Answer: Yes, multiple mementos can be saved for a single originator object, allowing for the ability to restore to previous saved states.

4) How does the Memento Pattern differ from the Command Pattern?
Answer: The Memento Pattern focuses on saving and restoring the state of an object, while the Command Pattern focuses on encapsulating a request to perform an action.

5) In what situations is the Memento Pattern useful in software development?
Answer: The Memento Pattern is useful in situations where there is a need to save and restore the state of an object, such as in undo/redo functionality or in creating checkpoints during a long process.