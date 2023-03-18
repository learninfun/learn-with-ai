+++
title = "原型继承"
+++
## 习题预习
{{< ask_chatgpt >}}
给我5题Prototypes的问题
{{< /ask_chatgpt >}}



1. 什麼是Prototype模式，它有甚麼作用和優點？ 
2. Prototype模式和工廠模式的區別是甚麼？
3. Prototype模式和Singleton模式有甚麼區別和相似點？
4. Prototype模式的實現方式有哪些？哪種方式更好？
5. Prototype模式在使用時應該注意哪些問題？   

## 说明知识
{{< ask_chatgpt >}}
说明Prototypes并举例
{{< /ask_chatgpt >}}



Prototype（原型）是指在開發完成之前，開發人員在設計階段創建的模型或樣本。Prototype通常用來評估用戶體驗、設計風格、交互設計、內容等方面的問題。通過Prototype，開發人員可以在開發之前找到並解決問題，從而節省成本和時間。

舉例：

1.網頁開發中的Prototype：在網頁開發中，Prototype通常是指靜態網頁設計圖或效果圖。開發人員通過Prototype可以確保將來的網頁實際內容和實際需求匹配。

2.產品開發中的Prototype：在產品開發中，Prototype通常是指物理或數字樣本。例如，在製造新產品之前，製造商可能會製作一個3D打印樣本或一個手工原型來驗證產品的功能和設計。

3.軟件設計中的Prototype：在軟件開發中，Prototype通常是指初步的軟件版本或測試版本。開發人員可以使用Prototype進行用戶測試和找到問題，從而進一步優化軟件。   

## 汇总重点
{{< ask_chatgpt >}}
条列Prototypes的重点
{{< /ask_chatgpt >}}



1. Prototypes 是 JavaScript 中的一個重要的概念。
2. 原型鏈（Prototype Chain）是從一個對象找到它的屬性和方法的過程。
3. 繼承是一種面向對象程式設計中的重要概念，而在 JavaScript 中，原型繼承是現代 JavaScript 程式設計中的一個關鍵概念。
4. JavaScript 中的所有對象都有一個原型（prototype）屬性。
5. 使用 Object.create() 方法創建對象。
6. 調用對象的方法時，首先在本地尋找對應的方法，如果找不到就到原型鏈上繼續尋找。
7. 原型可以繼承屬性和方法，所以可以通過對原型進行修改，來影響所有繼承它的對象。
8. 原型可以用來擴展現有的對象，並且在創建新對象時，可以從現有的對象繼承原型。
9. Prototype 鏈的終點為 null，即 Object.prototype.__proto__ === null，表示沒有更高的原型繼承，Object.prototype 是所有內建對象的原型對象。
10. 改變原型對象可能會導致代碼中的不良後果，因此應該謹慎使用。   

## 知识测验
{{< ask_chatgpt >}}
给我5题Prototypes的中等难度问题，并在后面列出答案
{{< /ask_chatgpt >}}



1. 試著給一個汽車原型增加一個自動停止引擎的偵測系統，當車輛在停車模式下一定時間內沒有活動，引擎就會自動關閉。如何實現這個系統？（提示：可以利用車載電腦和傳感器來偵測車輛的狀態。）

答案：可以利用車載電腦和傳感器來偵測車輛的狀態，並設計一個程式來控制引擎。當車輛在停車模式下一定時間內沒有活動，程式會自動關閉引擎。

2. 你想要設計一個能夠根據不同手勢移動的機器人原型，這樣人們可以使用手勢控制它的移動。如何實現這個原型？

答案：可以使用陀螺仪和加速度计这类传感器来控制机器人的移动。手势可以通过卷曲手指、拍击、挥手等方式来识别，随后可以通过编程来实现机器人的移动。

3. 你要设计一个无人机原型，可以上下左右自由移动，并能够通过语音控制。如何实现？

答案：可以使用GPS 和 互联网的通讯来控制无人机的移动，说话时，麦克风可以捕捉语音信号，随后需要使用自然语言处理算法将语音信号转换为机器可理解的指令，然后机器人将自动执行相应的任务。

4. 你想开发一款可翻转拍糖豆的钥匙扣原型，如何实现？

答案：可以设计一个小型机械臂来进行糖豆的翻转，并将它集成在钥匙扣上。需要进行一定的编程来控制机械臂的移动和抓取糖豆的方式。

5. 设计一个可以以各种方式自动移动的玩具车原型。 如何实现？

答案：可以使用DC电机来驱动车轮的运动，并通过陀螺仪和加速度计来控制其移动方向。还可以使用遥控器或者简单的编程语言来实现不同的运动操作，例如倒车，直线行驶以及绕场地运动等操作。   

## 网络数据
{{< ask_chatgpt >}}
给我5篇Prototypes的网络数据
{{< /ask_chatgpt >}}



1. "The Importance of Prototyping in Product Development." Medium, 16 Nov. 2020, medium.com/@SohaibKhan/the-importance-of-prototyping-in-product-development-9ebcc2cb494e.

This article discusses the important role that prototyping plays in product development, highlighting how it helps to identify potential design flaws and issues before products are released to the market. It also outlines some of the key benefits of prototyping, including the ability to test and refine product functionality, and the ability to reduce development costs.

2. "Understanding Prototypes in Software Development." Scopic, 11 May 2020, scopicsoftware.com/blog/understanding-prototypes-in-software-development.

This article provides an overview of software prototypes, including how they differ from other types of prototypes, and why they are important in software development. It also delves into the different types of software prototypes, such as proof of concept prototypes, visual prototypes, and functional prototypes, and describes how they can be used in different stages of the development process.

3. "10 Different Types of Prototyping Techniques." Justinmind, Justinmind, 8 Mar. 2019, justinmind.com/blog/10-different-types-of-prototyping-techniques/.

This article outlines the different types of prototyping techniques that are commonly used in product design and development. These include low-fidelity prototypes, paper prototypes, digital prototypes, and interactive prototypes, among others. The article highlights the benefits and drawbacks of each type of prototyping technique, and offers advice on when to use them.

4. "Prototyping: From Concept to Finished Product." Autodesk, 18 Nov. 2020, autodesk.com/redshift/prototyping-concept-to-finished-product/.

This article explores the process of prototyping, from the initial concept phase to the creation of a finished product. It describes the different stages of prototyping, such as defining project goals and requirements, developing and testing prototypes, and refining the final product. It also discusses the role of different tools and technologies in the prototyping process, such as 3D printing and virtual reality.

5. "The Benefits of Prototyping & Why Every Startup Should Prototype Their Product." TEC, Technology Evaluation Centers Inc., 16 Oct. 2020, tec.com.au/blog/benefits-prototyping-startup/.

This blog post explains why prototyping is essential for startups that want to design and launch successful products. The post outlines the key benefits of prototyping, which include enabling rapid iteration and feedback, reducing development costs, and improving the user experience. It also provides practical tips and advice for startups that want to get started with prototyping.   

