

Structural patterns are design patterns that deal with object composition to form larger structures, such as class or object hierarchies. These patterns help to make the system more organized, modular, and extensible.

Below is an example of a popular structural pattern called the Adapter pattern:

Adapter Pattern: It is used to convert the interface of a class into another interface that the client expects. This pattern involves two types of interfaces, i.e., the adapter class, which adapts the interface of the adaptee class, to that of the target interface.

Example: Consider a scenario where we want to connect a USB device to a laptop which has only a PS/2 port. In this case, we can use an adapter to convert the interface of the USB to PS/2 port. The Adapter Pattern can be used here as a mediator between the USB adapter and the Laptop port – the adapter class adapts the USB interface to the PS/2 interface, and hence, the laptop can now connect to the USB device via the adapter.