+++
title = "Strangler模式"
+++
## 习题预习
{{< ask_chatgpt >}}
给我5题Strangler Pattern的问题
{{< /ask_chatgpt >}}



1. 什么是Strangler Pattern，它可以解决什么问题？

2. Strangler Pattern的实现需要哪些步骤？如何使用它来重构遗留系统？

3. Strangler Pattern的优点和缺点是什么？在何种情况下使用更适合？

4. Strangler Pattern如何与微服务架构相关？如何使用它来实现微服务化？

5. 除了Strangler Pattern，还有哪些常用的重构技术可以应用于遗留系统？它们各自的优点和缺点是什么？   

## 说明知识
{{< ask_chatgpt >}}
说明Strangler Pattern并举例
{{< /ask_chatgpt >}}



Strangler Pattern是一种改善软体架构的设计模式，其目的是将旧的系统逐步取代为新的系统，而不会对现有系统造成任何损害或风险。

该模式是以植物学家David Snowden的比喻为名，他说：“当一株树从别的植物中榨取养分，它将逐渐生长，使其他植物枯萎，然后最终完全代替那些过时的植物。”

Strangler Pattern的关键在于，它在导入新系统的同时，不会中断现有系统的运行。相反地，新系统将逐渐地“侵蚀”旧系统，直到完全取代它。

例如，假设有一个旧的电子商务网站，它的外观与功能已经过时了。为了提供更好的用户体验和更高的性能，公司决定开发一个全新的网站。但是，他们还不能停止旧网站的运行，因为它会对商业运营造成严重影响。

使用 Strangler Pattern，公司可以逐步导入新的网站元素，例如新的搜寻引擎、新的结帐流程等等。这些新元素不会影响旧网站的运作，但它们为用户提供了更好的功能和体验。最终，当所有新元素都准备就绪时，旧网站可以完全被新网站所取代。   

## 汇总重点
{{< ask_chatgpt >}}
条列Strangler Pattern的重点
{{< /ask_chatgpt >}}



以下是Strangler Pattern 的重点：

1. 不必完全重写现有应用程式：Strangler Pattern 的目标是逐步将现有应用程式移向新的架构或平台，而不是一次性重写它。这有助于降低成本和风险。

2. 使用“缓慢杀手”策略：Strangler Pattern 通常使用名为“缓慢杀手”的策略来逐步取代现有应用程式的功能。这种方法可能需要几个月或几年的时间，但能确保在不中断现有服务的情况下进行过度。

3. 保持应用程式的稳定性：在应用 Strangler Pattern 时，需要确保系统保持稳定，并且不会因为过渡期间的问题而中断服务。为了达到这个目标，必须进行仔细的规划和测试，使用自动化测试等工具来减少错误。

4. 建立良好的架构：在进行过度时，需要建立一个良好的架构，以便在未来继续扩展和维护。应该考虑现有基础设施的限制，并选择适合的架构来克服这些限制。   

## 知识测验
{{< ask_chatgpt >}}
给我5题Strangler Pattern的中等难度问题，并在后面列出答案
{{< /ask_chatgpt >}}



1. 如何利用Strangler Pattern對一個電子商務網站進行升級改造？

答案：可以先在旧的网站的侧边栏或是首页上添加新的功能，引导用户尝试新的功能，并逐步将用户引到新的网站上，最后将旧的网站完全替换掉。

2. 如何在使用Strangler Pattern时避免重要的数据丢失？

答案：可以将关键数据分别保存在旧的和新的系统之中，然后通过某些方式将它们同步，以确保不会出现数据丢失的情况。

3. 如何在使用Strangler Pattern时保持前后端沟通无障碍？

答案：可以使用API接口，通过接口的方式将新后台和旧前台之间的数据传递，以保持前后台沟通的无障碍状态。

4. 如何使用Strangler Pattern将一个较大的应用分解成多个模块？

答案：可以通过将不同模块拆解成单独的应用程序，然后使用API接口进行通讯，将多个单独的应用程序整合在一起，以实现较大应用程序的分解。

5. 如何使用Strangler Pattern保持系统运行的稳定性？

答案：可以使用一些工具来监控系统的运行情况并及时发现异常，另外，需要遵守一些优秀的软件开发规范，以保障系统运行的高效性和稳定性。   

## 网络数据
{{< ask_chatgpt >}}
给我5篇Strangler Pattern的网络数据
{{< /ask_chatgpt >}}



1. Martin Fowler的博客：Strangler Fig Application Migration Pattern。

https://martinfowler.com/bliki/StranglerFigApplication.html

這是一篇由Martin Fowler撰寫的博客，透過這篇博客，讀者可以瞭解到什麼是Strangler Pattern，以及如何使用這個模式遷移應用程式。這篇博客是Strangler Pattern的經典文章之一。

2. StranglerPattern.org

https://stranglerpattern.org/

這是Strangler Pattern模式的官方網站，該網站提供有關模式的詳細信息，包括什麼是Strangler Pattern，何時應該使用它以及如何實施它的詳細步驟。

3. InfoQ的文章：Strangler Fig Approach to Legacy Modernization。

https://www.infoq.com/articles/strangler-fig-legacy-modernization/

這篇文章是一個關於如何使用Strangler Pattern來現代化應用程序的案例研究。該案例研究探討了一家企業如何使用Strangler Pattern來進行現代化，以及它們從這個遷移中學到的教訓。

4. Medium的文章：The Strangler Pattern for Application Transformation。

https://medium.com/@randyshoup/the-strangler-pattern-for-application-transformation-8f43de751d47

這篇文章將Strangler Pattern描述為一種用於應用程序轉換的有效方法。它解釋了為什麼使用這種模式是一種有利的做法，以及如何在實踐中實現這種模式。

5. DZone的文章：Strangler Pattern: A Way to Modernize Legacy Applications。

https://dzone.com/articles/strangler-pattern-a-way-to-modernize-legacy-applica

這篇文章探討了 Strangler 模式作為現代化傳統應用程序的一種方法。它提供了一些簡單的例子來解釋為什麼該模式是有用的，以及如何使用它來實現應用程序的遷移。   
