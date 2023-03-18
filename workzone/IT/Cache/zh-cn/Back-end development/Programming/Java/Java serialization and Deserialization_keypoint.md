

以下是 Java serialization 和 deserialization 的重点：

1. Serialization 是将 Java 对象转换为字节流的过程，而 deserialization 是将字节流转换回 Java 对象的过程。
2. Java 对象必须实现 Serializable 接口才能进行序列化。
3. 序列化过程中保存对象的状态，而不是对象本身。
4. 序列化的结果可以保存到文件、数据库或通过网络传输。
5. 可以使用 ObjectOutputStream 和 ObjectInputStream 类来进行序列化和反序列化。
6. 序列化过程中，要注意避免跨平台或版本不兼容的问题。
7. 可以通过自定义序列化方法和控制序列化版本来解决兼容性问题。
8. 序列化过程可以加密或压缩，以提高传输安全性和效率。
9. 序列化和反序列化是 Java RMI、JMS 和 Web Service 等技术的核心组成部分。