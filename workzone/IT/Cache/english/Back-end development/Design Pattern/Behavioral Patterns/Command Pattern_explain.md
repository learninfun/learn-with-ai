

The Command Pattern is a software design pattern that is used to encapsulate a request as an object. This pattern enables us to decouple the client that issues the request from the object that receives the request. The command pattern has four basic components: 

1. Command: This is the interface that all commands implement.

2. Concrete Command: This class implements the Command interface and provides the necessary information to execute the command.

3. Receiver: This class contains the logic to execute the command.

4. Invoker: This class is responsible for executing the command.

An example of the Command Pattern is a media player application that has a play, pause, and stop button. When the user clicks on the play button, it sends a command object to the media player. The command object is responsible for telling the media player to play. The receiver is the media player object, and the invoker is the button handling the user input. The concrete commands are play, pause and stop which implements the command interface. With this pattern, each command can be executed in multiple ways, allowing the application to easily extend the actions performed by the buttons.