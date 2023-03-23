

以下是 Java serialization 和 deserialization 的重點：

1. Serialization 是將 Java 對像轉換為字節流的過程，而 deserialization 是將字節流轉換回 Java 對象的過程。
2. Java 對像必須實現 Serializable 接口才能進行序列化。
3. 序列化過程中保存對象的狀態，而不是對像本身。
4. 序列化的結果可以保存到文件、數據庫或通過網絡傳輸。
5. 可以使用 ObjectOutputStream 和 ObjectInputStream 類來進行序列化和反序列化。
6. 序列化過程中，要注意避免跨平台或版本不兼容的問題。
7. 可以通過自定義序列化方法和控制序列化版本來解決兼容性問題。
8. 序列化過程可以加密或壓縮，以提高傳輸安全性和效率。
9. 序列化和反序列化是 Java RMI、JMS 和 Web Service 等技術的核心組成部分。