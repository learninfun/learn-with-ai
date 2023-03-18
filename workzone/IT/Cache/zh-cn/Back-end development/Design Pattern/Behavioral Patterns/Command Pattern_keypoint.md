

1. Command Pattern是一种行为型设计模式。
2. Command Pattern让你能够将特定操作的信息从其执行中分离出来，并封装成一个独立的物件中。
3. Command Pattern让你可以将特定的操作序列化、日志记录、取消或延迟其执行。
4. Command Pattern中的关键角色有Command、Invoker、Receiver和Client。
5. Command是行为请求的选择接口，Invoker引用并调用命令，Receiver实现命令和最终操作，Client则创建一个具体的Command对象并将其传递给Invoker。
6. Command Pattern的优点包括解耦程式码、易于修改、简化操作层级和支援撤销和恢复功能等。
7. Command Pattern的缺点包括生成大量命令物件可能会影响效能、需要额外实现的模式如果实现不好可能会产生更多的问题。