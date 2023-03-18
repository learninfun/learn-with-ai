

1. Input/Output Stream: Java提供了兩種Stream來處理IO，一種是Input Stream用於從外部資源讀取數據，另一種是Output Stream用於將數據寫入外部資源。

2. Reader/Writer: 除了Stream外，Java還提供了Reader和Writer來處理文字資料的IO操作。Reader用於從外部資源讀取文字數據，而Writer用於將文字數據寫入外部資源。

3. File IO: Java提供了File類來處理檔案IO，可以使用File類來創建、寫入、讀取和刪除檔案。

4. 字符集(Encoding): 在處理文字資料IO時，需要注意字符集的問題，Java提供了多種字符集，可以使用相應的字符集來處理不同語系的文字資料。

5. 序列化(Serialization): Java提供了序列化技術，可以將物件轉換成byte數組，以便在不同的環境中進行傳輸和存儲。

6. 緩衝區(Buffering): 為了提高IO操作的效率，Java提供了緩衝區(Buffer)，通過將IO操作的數據先緩存在緩衝區中，在一定條件下再將數據批量寫入外部資源，可以提高IO操作的效率。

7. NIO(Non-blocking IO): Java提供了NIO技術，可以實現非阻塞的IO操作，增強系統的擴展性和吞吐量。NIO技術包括通道、緩衝區和選擇器等組件。