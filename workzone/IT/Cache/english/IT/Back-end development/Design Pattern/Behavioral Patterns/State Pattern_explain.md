

State pattern is a behavioral design pattern that allows an object to alter its behavior when its state changes. It provides a way for an object to change its internal state and subsequently change its behavior without changing its classes. 

The State pattern consists of the following entities: 

1. Context: It is the object whose behavior depends on its state. 
2. State: Represents the different states that the context can transition through. 
3. Concrete State: The specific implementation of the State interface. 

Example: 

A vending machine can be used to demonstrate the State pattern. The vending machine can be in different states, such as 'No Coin', 'Has Coin', 'Sold', and 'Sold Out'. Each of these states will have different actions that can be performed based on the user's input. 

In the 'No Coin' state, the vending machine will not dispense any item until the user inserts a coin. When the user inserts a coin, the vending machine will change to the 'Has Coin' state. In this state, the user can select the item they want. If the item is available, the vending machine will dispense it and change to the 'Sold' state. If the item is not available, the vending machine will remain in the 'Has Coin' state. 

Once the item is dispensed or if there is no item available, the vending machine will transition to the 'No Coin' state. If the vending machine runs out of stock, it will change to the 'Sold Out' state, and no further purchases can be made. 

In this example, the vending machine's behavior is dependent on its state, and each state has different actions that can be performed. The State pattern allows the vending machine to change its behavior based on its internal state without changing its code.