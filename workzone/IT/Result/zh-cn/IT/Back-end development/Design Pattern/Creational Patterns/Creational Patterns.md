+++
title = "创建型模式"
+++
## 习题预习
{{< ask_chatgpt >}}
给我5题Creational Patterns的问题
{{< /ask_chatgpt >}}



1. 什么是Creational Patterns？它们主要用于什么目的？
2. 哪些是Creational Patterns的主要类别？请述其基本原理及使用情境。
3. 为什么我们需要使用Creational Patterns？可以举出一个实际的案例来说明吗？
4. 如何在Creational Patterns中选择适合的模式？有哪些因素需要考虑？
5. 何为Singleton模式？它和其他Creational Patterns的区别是什么？该如何适当地应用Singleton模式？   

## 说明知识
{{< ask_chatgpt >}}
说明Creational Patterns并举例
{{< /ask_chatgpt >}}



Creational Patterns是指软件设计中一种与对象创建有关的模式，这些模式可以帮助我们创建不同类型的对象，并且更灵活地应对不同的需求。Creational Patterns主要涉及到以下三种模式： 

1. Singleton Pattern（单例模式）：用来保证一个类只会有一个实例，并且提供一个全局唯一的访问点来访问该实例。 

例如，一个应用程序可能需要一个全局设置或是资料库连接，Singleton Pattern可以用来保证这些对象只会有一个实例，这样可以减少管理和资源浪费。 

2. Factory Method Pattern（工厂方法模式）：定义一个用于创建对象的介面，让子类来决定实例化哪个类。 

例如，当创建一个对象时，可能需要遵循一定的流程（例如设置参数、初始化等），使用工厂方法可以将这些流程封装起来，并且交由子类实现，从而实现更加灵活的对象创建。 

3. Builder Pattern（建造者模式）：用来组合一个复杂的对象，同时隐藏其创建过程。 

例如，当创建一个复杂的对象时，可能涉及到许多子部分，并且需要按照一定的步骤来组装，使用Builder Pattern可以将这些步骤封装起来，让用户只需要指定需要的部分即可构建所需对象。 

以上就是三种Creational Patterns的基本介绍及其使用场景。   

## 汇总重点
{{< ask_chatgpt >}}
条列Creational Patterns的重点
{{< /ask_chatgpt >}}



Creational Patterns 是建立物件的软体设计模式，其重点包括：

1. 抽象化建立物件过程：Creational Patterns 通常使用工厂方法、抽象工厂、建造者、原型和单例等机制，以抽象化建立物件过程，让程式设计可以更灵活地处理物件的建立。

2. 简化物件建立：Creational Patterns 可以降低程式设计师建立物件的复杂度，让开发者专注于业务逻辑的实现。

3. 易于维护和扩展：Creational Patterns 可以帮助程式设计师维护和扩展系统，使程式更易于维护和修改。

4. 提高程式的可测性：Creational Patterns 可以改善程式的可测性，帮助开发者进行单元测试，以确保程式逻辑的正确性。

5. 适用范围广：Creational Patterns 适用于各种应用场景，包括桌面应用程式、Web 应用程式、移动应用程式等等，对于复杂的系统而言尤其有用。   

## 知识测验
{{< ask_chatgpt >}}
给我5题Creational Patterns的中等难度问题，并在后面列出答案
{{< /ask_chatgpt >}}



问题1：在使用工厂模式时，当需要创建复杂对象时，应该使用哪种变体？

问题2：以下哪个是一个建造者模式的关键元素？

A.抽象生成器

B.具体生成器

C.产品

D.主管

问题3：以下哪个创建了单例模式的图形？

A.三角形

B.正方形

C.菱形

D.圆形

问题4：在使用抽象工厂模式时，如何确定哪个具体工厂应该被使用？

问题5：下面哪个Creational Pattern可以确保在一个应用程序中只有一个实例被创建？

A.工厂模式

B.建造者模式

C.原型模式

D.单例模式

答案：

1.抽象工厂模式

2.D.主管

3.D.圆形

4.根据应用程序的需要

5.D.单例模式   

## 网络数据
{{< ask_chatgpt >}}
给我5篇Creational Patterns的网络数据
{{< /ask_chatgpt >}}



1. Factory Method Pattern
Factory Method Pattern is one of the most common creational design patterns used in object-oriented programming. This pattern defines an interface for creating objects, but allows subclasses to decide which class to instantiate. The Factory Method Pattern allows flexibility in creating objects while also providing a central point for creating objects in a system.

Source: https://www.geeksforgeeks.org/factory-method-design-pattern/

2. Abstract Factory Pattern
Abstract Factory Pattern is a creational design pattern that provides an interface for creating families of related or dependent objects without specifying their concrete classes. The Abstract Factory Pattern is used when a group of related objects must be created together across multiple classes, and it can help to eliminate coupling between the client code and the concrete classes.

Source: https://refactoring.guru/design-patterns/abstract-factory

3. Singleton Pattern
Singleton Pattern is a creational design pattern that restricts the instantiation of a class to one object. This pattern is used when there should be only one instance of a class in a system, and it provides a global point of access to that instance. Singleton Pattern is commonly used for database connections, shared resources, and logging.

Source: https://www.baeldung.com/java-singleton

4. Builder Pattern
Builder Pattern is a creational design pattern that separates the construction of a complex object from its representation, allowing the same construction process to create different representations. The Builder Pattern allows flexible creation of complex objects using a step-by-step approach, and it can help to simplify object construction and reduce the number of constructor parameters needed.

Source: https://www.javatpoint.com/builder-design-pattern-in-java

5. Prototype Pattern
Prototype Pattern is a creational design pattern that allows objects to be created by cloning existing objects. The Prototype Pattern is used when creating new objects is expensive or complex, and it can help to reduce the time and resources needed for object creation. This pattern also allows objects to be easily customized without affecting the original object.

Source: https://refactoring.guru/design-patterns/prototype   

