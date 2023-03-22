+++
title = "中介者模式"
+++
## 習題預習
{{< ask_chatgpt >}}
給我5題Mediator Pattern的問題
{{< /ask_chatgpt >}}



1. 什麼是Mediator Pattern？
2. 在何種情況下會需要使用Mediator Pattern？
3. Mediator Pattern的核心實現原理是什麼？
4. Mediator Pattern和Observer Pattern有什麼不同？
5. 你可以舉出一個使用Mediator Pattern的實際應用場景嗎？   

## 說明知識
{{< ask_chatgpt >}}
說明Mediator Pattern並舉例
{{< /ask_chatgpt >}}



中介者模式是一種行為模式，它的主要作用是減少多個對象之間的直接通信，並將這些對象的互動轉化為通過中介者對像進行間接通信。中介者模式通常被用來簡化系統中的複雜關係，同時可以使得系統的設計更加具有靈活性和可維護性。

在中介者模式中，中介者對像當作多個子對像之間的調度者。每個子對像之間的通信都必須通過中介者對象，這樣可以確保彼此之間的通信維持在一個良好的狀態，同時中介者對象可以根據不同的情況和需求來控制這些子對像之間的互動。

舉例來說，假設一個社交媒體平台包含了多個用戶和多個群組，這些用戶和群組之間需要相互通信，但直接通信會使得系統變得混亂不堪且難以維護。這時候，就可以使用中介者模式來解決這個問題。這個情況下，中介者對象可以當作一個控制中心，用來調度不同用戶之間的通信，同時可以控制用戶和群組之間的互動。通過使用中介者模式，我們可以簡化系統的設計，降低系統的複雜度和耦合度，提高系統的可維護性和擴展性。   

## 彙總重點
{{< ask_chatgpt >}}
條列Mediator Pattern的重點
{{< /ask_chatgpt >}}



- Mediator Pattern是一種行為型設計模式，它可以用來減少對像之間的直接耦合，以及更好地管理它們之間的交互。
- 在Mediator Pattern中，所有的對像都通過中介者對像進行通信，而不是直接相互交互。
- Mediator Pattern可以降低系統的複雜度，增強系統的可維護性和可擴展性，並且可以有效地減少代碼的重複性。
- Mediator Pattern中的中介者對像需要維護一個對像集合，以便管理這些對像之間的交互。
- Mediator Pattern可以應用於不同類型的系統中，包括GUI設計、多人遊戲、企業級應用等。   

## 知識測驗
{{< ask_chatgpt >}}
給我5題Mediator Pattern的中等難度問題，並在後面列出答案
{{< /ask_chatgpt >}}



1. 什麼是中介者模式？
2. Mediator Pattern的優點是什麼？
3. Mediator Pattern的缺點是什麼？
4. Mediator Pattern的實現方式有哪些？
5. Mediator Pattern適用於哪些場景？

1. 什麼是中介者模式？

中介者模式是一個行為型設計模式，它允許將一個對象與其他對象進行通信，而不需要直接與它們互相作用。在中介者模式中，在不會直接對彼此談判的對象之間，對彼此進行溝通的中介對象充當著統一的調解者。

2. Mediator Pattern的優點是什麼？

(1) 減少對象之間的耦合度。

(2) 提供了一個方便的方法，使得對象能夠通信而不需要互相認識。

(3) 支持開放/關閉原則。 你可以新增中介者並且不必改變現有的中介者或其他對象的代碼。

(4) 對於多個對象之間的通信，在使用者角度上變得更加簡單和直觀。

3. Mediator Pattern的缺點是什麼？

(1) 中介物件的過於複雜，可以使代碼量增加。

(2) 通過中介物件進行交互可以導致性能瓶頸。

4.Mediator Pattern的實現方式有哪些？

(1) 事件通知器。

(2) 中介者基類。

(3) 對象鏈。

5.Mediator Pattern適用於哪些場景？

(1) 跨界處理器。

(2) 對象之間交互的場景。

(3) 多對多交互的場景。   

## 網路資料
{{< ask_chatgpt >}}
給我5篇Mediator Pattern的網路資料
{{< /ask_chatgpt >}}



1. Mediator Pattern in Java: Explained with an Example - https://www.baeldung.com/java-mediator-pattern

This article provides a detailed explanation of the Mediator pattern in Java, along with a practical example of how it can be used to simplify communication between objects in a complex system.

2. Mediator Pattern: Overview - https://refactoring.guru/design-patterns/mediator

This review of the Mediator pattern provides a high-level overview of its purpose, benefits, and implementation, along with examples of how the pattern can be applied in various scenarios.

3. The Mediator Pattern: A More Controlled Approach to Object Interaction - https://www.dummies.com/programming/software-development/the-mediator-pattern-a-more-controlled-approach-to-object-interaction/

This article explains the Mediator pattern in plain language, with a focus on how it can help simplify object interaction and improve code maintainability.

4. Practical Design Patterns: Mediator - https://www.toptal.com/design-patterns/mediator-pattern

This tutorial offers a practical guide to implementing the Mediator pattern in real-world applications, with code examples and detailed explanations of the key concepts.

5. Mediator Pattern in JavaScript - https://www.javascripttutorial.net/design-patterns/mediator-pattern/

This article provides a detailed guide to implementing the Mediator pattern in JavaScript, including a step-by-step example and tips for best practices.   
