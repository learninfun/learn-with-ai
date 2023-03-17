

1. 定義:
Iterator Pattern是一種行為型設計模式，它提供了一種通用的方式，可以將集合中的元素逐一經過存取或者處理。

2. 目的:
Iterator Pattern旨在解決遍歷集合時，如何儲存集合元素的問題。它將集合和遍歷分開來，使得同一種集合可以有不同的遍歷方式，而不必將處理遍歷方法的程式碼耦合在集合本身的實現中。

3. 結構:
Iterator Pattern由五個構件組成，分別是: Aggregate、ConcreteAggregate、Iterator、ConcreteIterator、Client。

4. 應用:
Iterator Pattern適用於需要迭代訪問集合中元素的場景。如:Java中的Iterator、C#中的IEnumerator、Python中的迭代器等。它可以方便地遍歷不同資料結構，如:陣列、LinkedList、二叉樹等。

5. 優點:
Iterator Pattern的主要優點是:取消了對集合的實作方式的束縛，使得集合和它的遍歷可以獨立演化和自由地改變。同時，Iterator Pattern為不同的集合提供了統一的遍歷介面，使得客戶端的程式碼簡潔而易懂。

6. 缺點:
Iterator Pattern的主要缺點是:集合元素的類型只能是固定的，無法動態地加入新的元素；另外，開發者需要實作Iterator介面，增加了開發成本。


總之，Iterator Pattern的重點就是解決了集合元素的遍歷問題，把集合的遍歷和集合本身分離開來，使得集合可以獨立演化和自由地改變。這一機制可以透過Iterator的介面，統一地遍歷不同資料結構的集合。