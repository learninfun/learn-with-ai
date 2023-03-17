## 習題預習
{{< ask_chatgpt >}}
給我5題Prototypes的問題
{{< /ask_chatgpt >}}



1. 什麼是Prototype模式，它有甚麼作用和優點？ 
2. Prototype模式和工廠模式的區別是甚麼？
3. Prototype模式和Singleton模式有甚麼區別和相似點？
4. Prototype模式的實現方式有哪些？哪種方式更好？
5. Prototype模式在使用時應該注意哪些問題？   

## 說明知識
{{< ask_chatgpt >}}
說明Prototypes並舉例
{{< /ask_chatgpt >}}



Prototype（原型）是指在開發完成之前，開發人員在設計階段創建的模型或樣本。Prototype通常用來評估用戶體驗、設計風格、交互設計、內容等方面的問題。通過Prototype，開發人員可以在開發之前找到並解決問題，從而節省成本和時間。

舉例：

1.網頁開發中的Prototype：在網頁開發中，Prototype通常是指靜態網頁設計圖或效果圖。開發人員通過Prototype可以確保將來的網頁實際內容和實際需求匹配。

2.產品開發中的Prototype：在產品開發中，Prototype通常是指物理或數字樣本。例如，在製造新產品之前，製造商可能會製作一個3D打印樣本或一個手工原型來驗證產品的功能和設計。

3.軟件設計中的Prototype：在軟件開發中，Prototype通常是指初步的軟件版本或測試版本。開發人員可以使用Prototype進行用戶測試和找到問題，從而進一步優化軟件。   

## 彙總重點
{{< ask_chatgpt >}}
條列Prototypes的重點
{{< /ask_chatgpt >}}



1. Prototypes 是 JavaScript 中的一個重要的概念。
2. 原型鏈（Prototype Chain）是從一個對像找到它的屬性和方法的過程。
3. 繼承是一種面向對像程式設計中的重要概念，而在 JavaScript 中，原型繼承是現代 JavaScript 程式設計中的一個關鍵概念。
4. JavaScript 中的所有對象都有一個原型（prototype）屬性。
5. 使用 Object.create() 方法創建對象。
6. 調用對象的方法時，首先在本地尋找對應的方法，如果找不到就到原型鏈上繼續尋找。
7. 原型可以繼承屬性和方法，所以可以通過對原型進行修改，來影響所有繼承它的對象。
8. 原型可以用來擴展現有的對象，並且在創建新對像時，可以從現有的對象繼承原型。
9. Prototype 鏈的終點為 null，即 Object.prototype.__proto__ === null，表示沒有更高的原型繼承，Object.prototype 是所有內建對象的原型對象。
10. 改變原型對象可能會導致代碼中的不良後果，因此應該謹慎使用。   

## 知識測驗
{{< ask_chatgpt >}}
給我5題Prototypes的中等難度問題，並在後面列出答案
{{< /ask_chatgpt >}}



1. 試著給一個汽車原型增加一個自動停止引擎的偵測系統，當車輛在停車模式下一定時間內沒有活動，引擎就會自動關閉。如何實現這個系統？（提示：可以利用車載電腦和傳感器來偵測車輛的狀態。）

答案：可以利用車載電腦和傳感器來偵測車輛的狀態，並設計一個程式來控制引擎。當車輛在停車模式下一定時間內沒有活動，程式會自動關閉引擎。

2. 你想要設計一個能夠根據不同手勢移動的機器人原型，這樣人們可以使用手勢控制它的移動。如何實現這個原型？

答案：可以使用陀螺儀和加速度計這類傳感器來控制機器人的移動。手勢可以通過捲曲手指、拍擊、揮手等方式來識別，隨後可以通過編程來實現機器人的移動。

3. 你要設計一個無人機原型，可以上下左右自由移動，並能夠通過語音控制。如何實現？

答案：可以使用GPS 和 互聯網的通訊來控制無人機的移動，說話時，麥克風可以捕捉語音信號，隨後需要使用自然語言處理算法將語音信號轉換為機器可理解的指令，然後機器人將自動執行相應的任務。

4. 你想開發一款可翻轉拍糖豆的鑰匙扣原型，如何實現？

答案：可以設計一個小型機械臂來進行糖豆的翻轉，並將它集成在鑰匙扣上。需要進行一定的編程來控制機械臂的移動和抓取糖豆的方式。

5. 設計一個可以以各種方式自動移動的玩具車原型。 如何實現？

答案：可以使用DC電機來驅動車輪的運動，並通過陀螺儀和加速度計來控制其移動方向。還可以使用遙控器或者簡單的編程語言來實現不同的運動操作，例如倒車，直線行駛以及繞場地運動等操作。   

## 網路資料
{{< ask_chatgpt >}}
給我5篇Prototypes的網路資料
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

