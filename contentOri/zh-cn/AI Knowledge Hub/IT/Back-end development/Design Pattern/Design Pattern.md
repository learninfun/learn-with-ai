+++
title = "设计模式"
+++
## 习题预习
{{< ask_chatgpt >}}
给我5题Design Pattern的问题
{{< /ask_chatgpt >}}



1. 什麼是Design Pattern？它們的作用是什麼，為什麼要使用它們？
2. 請說明什麼是MVC（Model-View-Controller）Design Pattern以及它的優缺點。
3. Singleton Design Pattern是什麼？它要如何使用，什麼時候會使用到它？請解釋。
4. 規範Design Pattern與Framework Design Pattern之間的區別。
5. 簡述Factory Method Design Pattern和Abstract Factory Design Pattern之間的相異之處，以及它們的應用場景。   

## 说明知识
{{< ask_chatgpt >}}
说明Design Pattern并举例
{{< /ask_chatgpt >}}



Design Pattern（設計模式）是軟體工程中的一種解決常見設計問題的方案集。它們是已經經過漫長時間、經驗和考驗的設計架構和方法，幫助開發人員通過重複使用共通的設計模式來解決複雜的程式問題。依據它們可以幫助我們更有效率地完成特定任務，同時有助於實現可重用軟體，提高系統的可維護性和可擴展性。

以下是幾個代表性的Design Pattern：

1. Singleton（單例模式）:

這個模式保證只有一個相應類的實例。它使用一個私有的建構函式和一個私有的靜態實例變數，讓外部的其他類不能創建一個新實例。

舉例：在一個應用程式中，只希望有一個資料存儲庫物件被創建並提供給其他物件。Singleton 模式可確保只有一個資料存儲庫實例存在。

2. Factory Method（工廠方法）:

Factor Method 模式將對象創建委託給類的靜態方法，這些方法可以根據不同的參數創建不同的對象實例。工廠方法適用於一個類不能創建對象但必須給外部提供對象時，或需要基於特定條件創建不同對象的場合。

舉例：一個製造汽車的工廠根據客戶的需求製造不同型號的汽車。

3. Observer（觀察者模式）:

Observer 模式定義了對象間的一對多關係。當一個對象的狀態發生變化時，所有依賴於它的物件都會自動收到通知。Observer 模式可以使對象間解耦並在運行時動態添加或刪除依賴性。

舉例：一個氣象站搜集到溫度、濕度、氣壓等資料，並通知到不同的氣象觀測站。

4. Decorator（裝飾器模式）:

Decorator 模式允許我們動態地給一個對象增加一些功能，同時保持類的介面不變。這種模式就像是一個外在包裝，顯示類進行功能增強。

舉例：在一個圖像處理程式中，不同的圖像濾鏡可以被套用到圖像上，並且可以連續堆疊使用。

5. Command（命令模式）:

Command 模式將指令對象（Command）和接收者（Receiver）解耦。這樣，我們可以為具有不同功能的命令創建不同的命令對象。當需要執行時，可以將命令對象傳遞給調用者。

舉例：在一個影音檔案播放程式中，可以使用不同的控制命令（播放、暫停、停止、音量調整等）來控制播放進程。這些命令可以被保存、傳遞或堆疊在一起，最終實現對播放器的全面控制。   

## 汇总重点
{{< ask_chatgpt >}}
条列Design Pattern的重点
{{< /ask_chatgpt >}}



1. 模式的目的和意义：设计模式是一种有效的软件设计思想和工具，通过提供可重用的解决方案来解决常见的软件设计问题。

2. 可重用的解决方案：设计模式提供了一系列解决常见软件设计问题的经过测试的解决方案，可直接应用于开发过程中。

3. 模式的分类：设计模式可以分为三种类型：创建型模式、结构型模式和行为型模式，每种模式都解决一类特定的软件设计问题。

4. 模式的特点：设计模式具有几个特点，如可重用性、可扩展性、可管理性、可理解性等。

5. 模式的应用：设计模式在实际开发中应用广泛，如常用的单例模式、工厂模式、观察者模式等。

6. 模式的实现：设计模式的实现涉及到一些关键点，如抽象、封装、多态、继承等。

7. 模式的优缺点：设计模式有其优点和缺点，应根据具体需求和情况进行选择和应用。

8. 模式的重要性：设计模式是软件工程中的重要组成部分，熟练掌握设计模式可以大大提高软件开发质量和效率。   

## 知识测验
{{< ask_chatgpt >}}
给我5题Design Pattern的中等难度问题，并在后面列出答案
{{< /ask_chatgpt >}}



1. Factory Method Pattern 製造者模式

問題：你想要一個工廠來生產不同種類的車，但你不確定你想要的哪一種車，怎麼辦？

答案：Factory Method Pattern（製造者模式）可以提供一個類別（通常稱為Creator），它的主要責任是建立其他類別的實例，這些類別通常都是一個介面或抽象類別的子類別。在本例中，你可以創建一個名為CarFactory的Creator類別，它可以創建不同種類的車，例如SedanCars和SUVs。由使用者來決定想要哪種車的實例。

2. Observer Pattern 觀察者模式

問題：你想要一個方法，當一個對象改變時，可以通知其他對象。怎麼辦？

答案：Observer Pattern（觀察者模式）可以提供一個可觀察的對象和一組觀察者。當可觀察的對象發生改變時，它通知所有的觀察者，以便它們可以更新自己。在本例中，你可以創建一個名為Subject的可觀察對象類別，在通知項目發生改變時，它會使用自己的方法通知觀察者。觀察者可以在它們自己的類別中實現被通知時所需要執行的方法。

3. Adapter Pattern 適配器模式

問題：你要將兩個不相容的類別結合在一起工作，怎麼辦？

答案：Adapter Pattern（適配器模式）可以提供一個稱為適配器的類別，它可以讓原始的類別可以和目標接口相容。在本例中，你可以創建一個名為Adapter的類別來充當兩個不相容類別之間的介面。Adapter類的方法內部實現為將目標接口的方法委託給原始的類別方法，以便在目標接口中正確地實現操作。我們可以舉一個例子，例如適配器在一個開發平台上提供對另一個開發平台的支援。

4. Strategy Pattern 策略模式

問題：你想要對一個對象進行操作，但是想要在執行時能夠決定需要使用哪一個演算法。怎麼辦？

答案：Strategy Pattern（策略模式）可以提供一個方法，以便當你需要在運行時選擇一種演算法時使用。在本例中，你可以創建一組名為Strategy的類別，每個類別都實現一種不同的演算法。應用程式將會依據需要在運行時選擇一個Strategy類別。Strategy Pattern可以在需要不同演算法的情況下幫助你維護代碼庫。

5. Decorator Pattern 裝飾器模式

問題：你想要在不修改現有對象代碼的情況下，擴展特定對象的功能。怎麼辦？

答案：Decorator Pattern（裝飾器模式）可以提供一個對象，它可以擴展特定對象的功能。在本例中，你可以將現有的對象傳遞到一個稱為Decorator的裝飾器對象中，這個裝飾器對象可以擴展原有對象的功能但不影響原有對象的代碼。這可以幫助你在不修改本質代碼的前提下，增加對象的功能。舉一個例子，你可以創建一個名為CarDecorator的類別，它可以將一個Car的類別與其他類別結合起來完成額外的操作。   

## 网络数据
{{< ask_chatgpt >}}
给我5篇Design Pattern的网络数据
{{< /ask_chatgpt >}}



1. Design Patterns by the Gang of Four (GoF) - Java Tutorials. https://www.javatpoint.com/design-patterns-tutorial

2. What is Software Design Pattern? Types of Design Patterns. https://www.guru99.com/design-patterns.html

3. Introduction to Design Patterns. https://www.tutorialspoint.com/design_pattern/index.htm

4. The 23 Gang of Four Design Patterns. https://sourcemaking.com/design_patterns

5. Design patterns in Java. https://www.javatpoint.com/design-patterns-in-java   
