

Adapter pattern is a design pattern in which an interface can be adapted to another interface. It is used when there is a need to connect two incompatible interfaces or classes that cannot be used directly. The adapter pattern converts the original interface or class into a new interface that the client can use.

For example, suppose we have a client that expects an object to have a certain interface, but we have an object that has a different interface. In this scenario, we can create an adapter class that implements the expected interface and wraps the original object, to make it compatible with the client.

One example of adapter pattern is the use of sockets. Suppose we have an application that uses a specific socket type to communicate with a remote server. However, the server uses a different socket type. In this case, we could create an adapter class that implements the expected socket type, and internally, it uses the server's socket.

Another example is the use of a card reader. A card reader may use an interface that is incompatible with the computer or device it is attached to. In this situation, an adapter can be created to convert the card reader interface to the computer's interface. 

In summary, the Adapter pattern is used to make two incompatible interfaces work together by creating an intermediary that translates one interface to another.