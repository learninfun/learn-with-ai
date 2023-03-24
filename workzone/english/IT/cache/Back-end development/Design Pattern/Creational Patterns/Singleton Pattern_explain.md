

Singleton pattern is a design pattern in software engineering that is used when you want to make sure that the class has only one instance, and it provides a global point of access to that instance. In other words, it's a way to ensure that there's only one object of certain class and that object can be accessed globally.

Here's an example to understand this better:

Let's say you have a class called Logger, that is responsible for logging messages to a file. You want to make sure that no matter how many times the Logger class is instantiated, there should only be one instance of it.

class Logger:
    _instance = None
    def __new__(cls):
        if cls._instance is None:
            print("Creating new Logger object")
            cls._instance = super().__new__(cls)
        return cls._instance

In the Logger class, we define a _instance variable, which is initially set to None. We also define a __new__() method which allows us to override the process of creating a new object. In this method, we check if the _instance variable is None. If it is, we create a new Logger object, assign it to _instance and return it. If _instance is not None, we simply return it instead of creating a new one.

Now, let's test the Logger class:

logger1 = Logger()
logger2 = Logger()

print(logger1 is logger2)

Output:
Creating new Logger object
True

As you can see, when we create logger1, it creates a new Logger object, and the message "Creating new Logger object" is printed. When we create logger2, it simply returns the instance created by logger1, and no new object is created. This ensures that there's only one instance of the Logger class.