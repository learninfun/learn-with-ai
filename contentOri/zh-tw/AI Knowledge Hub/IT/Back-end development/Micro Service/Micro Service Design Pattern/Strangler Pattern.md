+++
title = "Strangler模式"
+++
## 習題預習
{{< ask_chatgpt >}}
給我5題Strangler Pattern的問題
{{< /ask_chatgpt >}}



1. 什麼是Strangler Pattern，它可以解決什麼問題？

2. Strangler Pattern的實現需要哪些步驟？如何使用它來重構遺留系統？

3. Strangler Pattern的優點和缺點是什麼？在何種情況下使用更適合？

4. Strangler Pattern如何與微服務架構相關？如何使用它來實現微服務化？

5. 除了Strangler Pattern，還有哪些常用的重構技術可以應用於遺留系統？它們各自的優點和缺點是什麼？   

## 說明知識
{{< ask_chatgpt >}}
說明Strangler Pattern並舉例
{{< /ask_chatgpt >}}



Strangler Pattern是一種改善軟體架構的設計模式，其目的是將舊的系統逐步取代為新的系統，而不會對現有系統造成任何損害或風險。

該模式是以植物學家David Snowden的比喻為名，他說：「當一株樹從別的植物中搾取養分，它將逐漸生長，使其他植物枯萎，然後最終完全代替那些過時的植物。」

Strangler Pattern的關鍵在於，它在導入新系統的同時，不會中斷現有系統的運行。相反地，新系統將逐漸地「侵蝕」舊系統，直到完全取代它。

例如，假設有一個舊的電子商務網站，它的外觀與功能已經過時了。為了提供更好的用戶體驗和更高的性能，公司決定開發一個全新的網站。但是，他們還不能停止舊網站的運行，因為它會對商業運營造成嚴重影響。

使用 Strangler Pattern，公司可以逐步導入新的網站元素，例如新的搜尋引擎、新的結帳流程等等。這些新元素不會影響舊網站的運作，但它們為用戶提供了更好的功能和體驗。最終，當所有新元素都準備就緒時，舊網站可以完全被新網站所取代。   

## 彙總重點
{{< ask_chatgpt >}}
條列Strangler Pattern的重點
{{< /ask_chatgpt >}}



以下是Strangler Pattern 的重點：

1. 不必完全重寫現有應用程式：Strangler Pattern 的目標是逐步將現有應用程式移向新的架構或平臺，而不是一次性重寫它。這有助於降低成本和風險。

2. 使用「緩慢殺手」策略：Strangler Pattern 通常使用名為「緩慢殺手」的策略來逐步取代現有應用程式的功能。這種方法可能需要幾個月或幾年的時間，但能確保在不中斷現有服務的情況下進行過度。

3. 保持應用程式的穩定性：在應用 Strangler Pattern 時，需要確保系統保持穩定，並且不會因為過渡期間的問題而中斷服務。為了達到這個目標，必須進行仔細的規劃和測試，使用自動化測試等工具來減少錯誤。

4. 建立良好的架構：在進行過度時，需要建立一個良好的架構，以便在未來繼續擴展和維護。應該考慮現有基礎設施的限制，並選擇適合的架構來克服這些限制。   

## 知識測驗
{{< ask_chatgpt >}}
給我5題Strangler Pattern的中等難度問題，並在後面列出答案
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

## 網路資料
{{< ask_chatgpt >}}
給我5篇Strangler Pattern的網路資料
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

