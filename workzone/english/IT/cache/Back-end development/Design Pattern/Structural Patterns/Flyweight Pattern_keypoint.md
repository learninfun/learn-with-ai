

1. The Flyweight pattern is a structural design pattern that aims to minimize the memory usage and object initialization time.

2. The pattern achieves this by sharing objects that have similar or identical state and only creating new objects for variations in state.

3. A shared flyweight object stores intrinsic state, which is common across objects, and an extrinsic state, which can vary across objects.

4. Flyweight objects can be stored in a Flyweight Factory that manages the creation, storage and retrieval of flyweight objects.

5. The pattern improves performance by reducing the number of objects created.

6. Flyweight objects are immutable, meaning their state cannot be modified after creation.

7. The pattern is used extensively in graphical user interfaces, where numerous objects such as images, icons, and fonts need to be displayed on the screen.

8. Flyweight objects can be shared between multiple threads and processes, further reducing the memory usage and initialization time.