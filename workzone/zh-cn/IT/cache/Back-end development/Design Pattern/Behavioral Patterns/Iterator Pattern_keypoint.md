

1. 定义:
Iterator Pattern是一种行为型设计模式，它提供了一种通用的方式，可以将集合中的元素逐一经过存取或者处理。

2. 目的:
Iterator Pattern旨在解决遍历集合时，如何储存集合元素的问题。它将集合和遍历分开来，使得同一种集合可以有不同的遍历方式，而不必将处理遍历方法的程式码耦合在集合本身的实现中。

3. 结构:
Iterator Pattern由五个构件组成，分别是: Aggregate、ConcreteAggregate、Iterator、ConcreteIterator、Client。

4. 应用:
Iterator Pattern适用于需要迭代访问集合中元素的场景。如:Java中的Iterator、C#中的IEnumerator、Python中的迭代器等。它可以方便地遍历不同资料结构，如:阵列、LinkedList、二叉树等。

5. 优点:
Iterator Pattern的主要优点是:取消了对集合的实作方式的束缚，使得集合和它的遍历可以独立演化和自由地改变。同时，Iterator Pattern为不同的集合提供了统一的遍历介面，使得客户端的程式码简洁而易懂。

6. 缺点:
Iterator Pattern的主要缺点是:集合元素的类型只能是固定的，无法动态地加入新的元素；另外，开发者需要实作Iterator介面，增加了开发成本。


总之，Iterator Pattern的重点就是解决了集合元素的遍历问题，把集合的遍历和集合本身分离开来，使得集合可以独立演化和自由地改变。这一机制可以透过Iterator的介面，统一地遍历不同资料结构的集合。