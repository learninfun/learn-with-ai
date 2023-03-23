+++
title = "模板方法模式"
weight = "10"
+++
## 习题预习
{{< ask_chatgpt >}}
给我5题Template Method Pattern的问题
{{< /ask_chatgpt >}}



1. 什么是Template Method Pattern？ 
2. 透过Template Method Pattern如何实现程式码重复利用及扩展性？ 
3. 与Strategy Pattern有何不同之处？ 
4. 在哪些场景应用Template Method Pattern比较适合？ 
5. Template Method Pattern有哪些优点和缺点？   

## 说明知识
{{< ask_chatgpt >}}
说明Template Method Pattern并举例
{{< /ask_chatgpt >}}



Template Method Pattern是一种行为设计模式，其目的是将一种算法的框架定义在抽像类中，而将算法的具体步骤延迟到子类中实现。这个设计模式可用于在不修改算法结构的情况下，通过子类来扩展算法的特定步骤。

例如，我们可以想像一个制作咖啡和茶的过程。它们都具有相同的基本结构，但是具体步骤略有不同。我们可以使用Template Method Pattern来减少代码复制并确保系统的分层结构。

首先，我们定义一个抽像类，称之为"饮料"。这个抽像类有一个名为"prepareBeverage"的方法，它包含类似于下面的一些步骤：

- 烧开水
- 将所需的杯子放在桌子上
- 将一些基本材料加入杯子中
- 调整饮品中的一些特定成分
- 将该饮料倒入杯子中
- 向杯子中加入一些附加物，如糖、奶油或蜂蜜

其中，步骤1、2和6是相同的，我们可以在"饮料"类中定义它们。但是，步骤3、4和5是具体步骤，因此我们需要将它们定义为抽像方法。

然后，我们可以定义两个具体类，称之为"咖啡"和"茶"，它们继承自"饮料"类。这些类可以实现在抽像类中定义的特定步骤，并覆盖抽像方法以提供其在制作咖啡或茶时的特定实现。

例如，如果我们实现"咖啡"类，我们可以覆盖"addIngredients"方法以添加咖啡豆和水。如果我们实现"茶"类，我们可以覆盖"addIngredients"方法以添加茶叶和水。但是，两种饮料的剩余步骤仍然是相同的。

这就是Template Method Pattern的作用。通过将一些共同步骤定义在抽像类中，它可以为子类提供代码复用且易于扩展的基础。   

## 汇总重点
{{< ask_chatgpt >}}
条列Template Method Pattern的重点
{{< /ask_chatgpt >}}



- Template Method Pattern是一种行为模式，它定义了一个算法的骨架，将某些步骤延迟到子类实现。
- 模板方法使用了继承机制，使得子类可以不改变算法骨架的情况下，重新定义某些步骤。
- 模板方法包含一些抽像步骤和具体步骤，并且可以定义一些钩子方法，这些方法可以由子类实现以改变算法骨架的流程。
- 模板方法的优点包括，提高代码的重用性、提高代码的可维护性、封装算法的实现细节等。
- 但是，使用模板方法也存在一些缺点，如增加了系统的类的数量、涉及到类的层次关系、可能造成性能问题等。   

## 知识测验
{{< ask_chatgpt >}}
给我5题Template Method Pattern的中等难度问题，并在后面列出答案
{{< /ask_chatgpt >}}



1. 在游戏中，有多种不同的角色可以选择，每种角色都有不同的技能和能力。使用Template Method Pattern设计一个角色选择的系统，每个角色都有以下步骤：
- 选择角色
- 显示角色资讯
- 显示角色能力详情
- 选择角色技能
- 开始游戏

答案：
抽像类别Character：
- 选择角色(selectCharacter方法)
- 显示角色资讯(displayInformation方法)
- 显示角色能力详情(displayAbility方法)
- 选择角色技能(selectSkill方法)
- 开始游戏(startGame方法)

具体类别Warrior, Mage, Ranger继承Character，并实做其方法

2. 设计一个简单的游戏关卡系统，游戏分为多个关卡，每个关卡有以下步骤：
- 进入关卡
- 显示关卡地图
- 开始战斗
- 显示战斗结果
- 过关

答案：
抽像类别Level：
- 进入关卡(enterLevel方法)
- 显示关卡地图(displayMap方法)
- 开始战斗(startFight方法)
- 显示战斗结果(displayResult方法)
- 过关(passLevel方法)

具体类别Level1, Level2, Level3继承Level，并实做其方法

3. 一个文档编辑器，可以输入不同格式的文档，例如txt和pdf。使用Template Method Pattern设计，每次输入一个新文档，编辑器会执行以下步骤：
- 开启文档
- 读取文档
- 显示文档内容

答案：
抽像类别Document：
- 开启文档(openDocument方法)
- 读取文档(readDocument方法)
- 显示文档内容(displayDocument方法)

具体类别TxtDocument, PdfDocument继承Document，并实做其方法

4. 一个网路商店的结帐系统，可以应用Template Method Pattern设计。每次结帐时，系统执行以下步骤：
- 确认订单明细
- 计算运费
- 确认客户付款
- 更新库存

答案：
抽像类别CheckOut：
- 确认订单明细(confirmOrder方法)
- 计算运费(calculateFreight方法)
- 确认客户付款(confirmPayment方法)
- 更新库存(updateInventory方法)

具体类别OnlineCheckOut, CODCheckOut继承CheckOut，并实做其方法

5. 一个聊天室，使用Template Method Pattern设计。每次聊天时，系统执行以下步骤：
- 连接聊天室
- 发送讯息
- 接收讯息
- 关闭聊天室

答案：
抽像类别ChatRoom：
- 连接聊天室(connectChatRoom方法)
- 发送讯息(sendMessage方法)
- 接收讯息(receiveMessage方法)
- 关闭聊天室(closeChatRoom方法)

具体类别PublicChatRoom, PrivateChatRoom继承ChatRoom，并实做其方法   

