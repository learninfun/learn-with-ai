

1. Input/Output Stream: Java提供了两种Stream来处理IO，一种是Input Stream用于从外部资源读取数据，另一种是Output Stream用于将数据写入外部资源。

2. Reader/Writer: 除了Stream外，Java还提供了Reader和Writer来处理文字资料的IO操作。Reader用于从外部资源读取文字数据，而Writer用于将文字数据写入外部资源。

3. File IO: Java提供了File类来处理档案IO，可以使用File类来创建、写入、读取和删除档案。

4. 字符集(Encoding): 在处理文字资料IO时，需要注意字符集的问题，Java提供了多种字符集，可以使用相应的字符集来处理不同语系的文字资料。

5. 序列化(Serialization): Java提供了序列化技术，可以将物件转换成byte数组，以便在不同的环境中进行传输和存储。

6. 缓冲区(Buffering): 为了提高IO操作的效率，Java提供了缓冲区(Buffer)，通过将IO操作的数据先缓存在缓冲区中，在一定条件下再将数据批量写入外部资源，可以提高IO操作的效率。

7. NIO(Non-blocking IO): Java提供了NIO技术，可以实现非阻塞的IO操作，增强系统的扩展性和吞吐量。NIO技术包括通道、缓冲区和选择器等组件。