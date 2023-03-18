+++
title = "創建型模式"
+++
## 習題預習
{{< ask_chatgpt >}}
給我5題Creational Patterns的問題
{{< /ask_chatgpt >}}



1. 甚麼是Creational Patterns？它們主要用於什麼目的？
2. 哪些是Creational Patterns的主要類別？請述其基本原理及使用情境。
3. 為什麼我們需要使用Creational Patterns？可以舉出一個實際的案例來說明嗎？
4. 如何在Creational Patterns中選擇適合的模式？有哪些因素需要考慮？
5. 何為Singleton模式？它和其他Creational Patterns的區別是甚麼？該如何適當地應用Singleton模式？   

## 說明知識
{{< ask_chatgpt >}}
說明Creational Patterns並舉例
{{< /ask_chatgpt >}}



Creational Patterns是指軟件設計中一種與對像創建有關的模式，這些模式可以幫助我們創建不同類型的對象，並且更靈活地應對不同的需求。Creational Patterns主要涉及到以下三種模式： 

1. Singleton Pattern（單例模式）：用來保證一個類只會有一個實例，並且提供一個全局唯一的訪問點來訪問該實例。 

例如，一個應用程序可能需要一個全局設置或是資料庫連接，Singleton Pattern可以用來保證這些對像只會有一個實例，這樣可以減少管理和資源浪費。 

2. Factory Method Pattern（工廠方法模式）：定義一個用於創建對象的介面，讓子類來決定實例化哪個類。 

例如，當創建一個對像時，可能需要遵循一定的流程（例如設置參數、初始化等），使用工廠方法可以將這些流程封裝起來，並且交由子類實現，從而實現更加靈活的對象創建。 

3. Builder Pattern（建造者模式）：用來組合一個複雜的對象，同時隱藏其創建過程。 

例如，當創建一個複雜的對象時，可能涉及到許多子部分，並且需要按照一定的步驟來組裝，使用Builder Pattern可以將這些步驟封裝起來，讓用戶只需要指定需要的部分即可構建所需對象。 

以上就是三種Creational Patterns的基本介紹及其使用場景。   

## 彙總重點
{{< ask_chatgpt >}}
條列Creational Patterns的重點
{{< /ask_chatgpt >}}



Creational Patterns 是建立物件的軟體設計模式，其重點包括：

1. 抽像化建立物件過程：Creational Patterns 通常使用工廠方法、抽像工廠、建造者、原型和單例等機制，以抽像化建立物件過程，讓程式設計可以更靈活地處理物件的建立。

2. 簡化物件建立：Creational Patterns 可以降低程式設計師建立物件的複雜度，讓開發者專注於業務邏輯的實現。

3. 易於維護和擴展：Creational Patterns 可以幫助程式設計師維護和擴展系統，使程式更易於維護和修改。

4. 提高程式的可測性：Creational Patterns 可以改善程式的可測性，幫助開發者進行單元測試，以確保程式邏輯的正確性。

5. 適用範圍廣：Creational Patterns 適用於各種應用場景，包括桌面應用程式、Web 應用程式、移動應用程式等等，對於複雜的系統而言尤其有用。   

## 知識測驗
{{< ask_chatgpt >}}
給我5題Creational Patterns的中等難度問題，並在後面列出答案
{{< /ask_chatgpt >}}



問題1：在使用工廠模式時，當需要創建複雜對像時，應該使用哪種變體？

問題2：以下哪個是一個建造者模式的關鍵元素？

A.抽像生成器

B.具體生成器

C.產品

D.主管

問題3：以下哪個創建了單例模式的圖形？

A.三角形

B.正方形

C.菱形

D.圓形

問題4：在使用抽像工廠模式時，如何確定哪個具體工廠應該被使用？

問題5：下面哪個Creational Pattern可以確保在一個應用程序中只有一個實例被創建？

A.工廠模式

B.建造者模式

C.原型模式

D.單例模式

答案：

1.抽像工廠模式

2.D.主管

3.D.圓形

4.根據應用程序的需要

5.D.單例模式   

## 網路資料
{{< ask_chatgpt >}}
給我5篇Creational Patterns的網路資料
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

