## 習題預習
{{< ask_chatgpt >}}
給我5題Caching的問題
{{< /ask_chatgpt >}}



1. 什麼是快取(Caching)？在什麼情況下使用快取可以提高系統的效能？

2. 快取的失效率是什麼？它對系統性能有什麼影響？

3. 常見的快取策略有哪些？它們之間有什麼區別？

4. 如何測試和評估系統的快取策略是否有效？

5. 快取穿透和快取雪崩是什麼？它們有什麼解決方案？   

## 說明知識
{{< ask_chatgpt >}}
說明Caching並舉例
{{< /ask_chatgpt >}}



Caching是將常用的數據存儲在快速、易於存取的高速記憶體中，以加快數據訪問速度的行為。這通常是因為存取主記憶體（RAM）或者硬碟等設備的速度較慢，載入數據需要大量時間和資源，如果非常依賴主存和硬碟來加載資料，將導致速度緩慢。

例如，在Web應用程序中，經常會將比較大的數據（例如圖像、CSS和JS文件）存儲在瀏覽器的快取中，以加快使用者訪問網站的速度。另一個常見的例子是瀏覽器的歷史紀錄。你最近訪問的網站將保存在你的瀏覽器的快取中，因此下一次訪問時將不需要重新載入所有數據。

在一些資料密集型的應用程序中，也可以使用緩存網格，將資料存儲在高速讀取和儲存的內存裡，而不是在慢速硬碟上。這樣可以大幅提高數據訪問速度，從而減少服務器負載。

總之，緩存是提高訪問速度並減少服務器負載的一種重要技術，可以應用到各種應用場景中。   

## 彙總重點
{{< ask_chatgpt >}}
條列Caching的重點
{{< /ask_chatgpt >}}



1. 提高响应速度和性能：缓存可以提高应用程序的响应速度和性能，因为它可以减少对数据库或其他资源的频繁访问，从而减少了处理时间。

2. 降低服务器负载：缓存可以减少服务器负载，因为它可以通过缓存响应减少服务器处理的请求。

3. 提高用户体验：快速的加载时间可以提高用户体验，因为用户可以在较短的时间内访问到所需的内容。

4. 减少网络流量：应用程序使用缓存可以减少网络流量，因为它可以减少对远程资源的请求。

5. 提高可伸缩性：缓存可以提高应用程序的可伸缩性，并可以更好地处理增加的请求。

6. 缓存数据一致性：缓存可以确保数据一致性，并确保用户访问到的数据是最新的。

7. 避免重复计算和访问：缓存可以避免重复计算和访问，以减少处理时间和网络流量。

8. 提高系统可用性：缓存可以提高系统可用性，在后端系统中出现故障时，缓存可以继续提供服务。

9. 反向代理功能：缓存还可以作为反向代理服务器，为用户提供更快的访问速度。

10. 优化内存使用：缓存使用内存存储数据，可以优化内存使用，并改善应用程序的吞吐量。   

## 知識測驗
{{< ask_chatgpt >}}
給我5題Caching的中等難度問題，並在後面列出答案
{{< /ask_chatgpt >}}



1. 給定一個有 $n$ 個元素的整數陣列和一個整數 $k$，請設計一個支持查询區間和的數據結構，要求時間複雜度為 $O(n\sqrt{n} + q\sqrt{n}\log n)$，其中 $q$ 為查詢數量。

2. 設計一個線性時間複雜度的算法，查詢一個區間中出現次數最多的元素及其出現次數。

3. 給定一個長度為 $n$ 的字符串 $S$，設計一個數據結構支持以下操作：

   - 將某個位置的字符修改為另一個字符。
   - 查詢某個子串在原串中出現的次數。

   要求修改和查詢操作的時間複雜度均為 $O(\sqrt{n})$。

4. 設計一個數據結構，支持以下操作：

   - 將一個元素加入集合中。
   - 從集合中刪除一個元素。
   - 查詢某個區間內元素的個數。

   其中，添加和刪除操作的時間複雜度均為 $O(\sqrt{n})$，查詢操作的時間複雜度為 $O(q\sqrt{n}\log n)$，其中 $q$ 為查詢次數。

5. 設計一個數據結構，支持以下操作：

   - 在序列中插入一個元素。
   - 刪除序列中的某個元素。
   - 查詢某個位置後面第 $k$ 小的元素。

   其中，插入和刪除操作的時間複雜度均為 $O(\sqrt{n})$，查詢操作的時間複雜度為 $O(q\sqrt{n}\log n)$，其中 $q$ 為查詢次數。

答案：

1. 利用分塊和前綴和，在每個塊上建立一棵平衡樹（如紅黑樹），支持區間查詢的時間複雜度為 $O(\sqrt{n}\log n)$，將每個塊區間和緩存起來，單次查詢的時間複雜度為 $O(\sqrt{n})$。
2. 將序列分為若干塊，對每個塊統計出現次數最多的元素和出現次數，總時間複雜度為 $O(n)$。對於一個查詢區間 $[l,r]$，若 $l$ 和 $r$ 落在同一個塊中，直接遍歷區間統計出現次數即可；否則分別遍歷區間左右端點所在的塊，並統計其中出現次數最多的元素及其出現次數，然後在這些塊中遍歷 $l$ 和 $r$ 所在的塊中非區間的元素，統計其出現次數，最終比較三者出現次數，返回出現次數最多的那個元素及其出現次數。
3. 將字符串划分成 $\sqrt{n}$ 個塊，對每個塊建立一棵 Trie（字典樹），在 Trie 上標記每個節點對應的字串在原串中出現的次數，緩存每個塊中所有子串的出現次數，查詢時統計相應塊中子串的出現次數，然後在相應的 Trie 上遍歷所查詢的子串，統計其出現次數，最終返回所有統計值的和。修改操作可以在 Trie 上進行，時間複雜度為 $O(|T|)$，其中 $T$ 為 Trie 的節點數，對於一個子串修改，只需找到相應的 Trie 和其對應的節點，然後修改該節點的標記，時間複雜度為 $O(\sqrt{n}+\log m)$，其中 $m$ 為字元集大小。
4. 將集合分割成 $\sqrt{n}$ 個塊，對於每個塊使用哈希表和平衡樹（如紅黑樹）進行支持添加和刪除操作，時間複雜度均為 $O(\sqrt{n})$。對於一個查詢區間 $[l,r]$，若 $l$ 和 $r$ 落在同一個塊中，直接遍歷區間統計元素個數；否則分別遍歷區間左右端點所在的塊，然後在這些塊中遍歷 $l$ 和 $r$ 所在的塊中非區間的元素，統計其出現次數，最終返回統計值的和。總時間複雜度為 $O(n\sqrt{n}+q\sqrt{n}\log n)$。
5. 將序列分為 $\sqrt{n}$ 個塊，對於每個塊使用一棵支持重複元素的排序算法（如 std::multiset）進行排序，時間複雜度為 $O(\sqrt{n}\log\sqrt{n})$。將每個塊中的第 $k$ 小的元素緩存起來，總時間複雜度為 $O(n\sqrt{n})$。對於一個插入或刪除操作，只需找到相應的塊，更新該塊中的排序算法，更新緩存中的第 $k$ 小元素，時間複雜度為 $O(\sqrt{n}\log\sqrt{n})$。對於一個查詢操作，若要查詢的位置位於某個塊中，直接在該塊中進行查詢；否則先在位置左邊的塊中查詢其後第 $k$ 小的元素，然後在位置右邊的塊中查詢其前 $k-1$ 小的元素，最終找出所有結果的前 $k$ 小的元素，總時間複雜度為 $O(q\sqrt{n}\log\sqrt{n})$。   

## 網路資料
{{< ask_chatgpt >}}
給我5篇Caching的網路資料
{{< /ask_chatgpt >}}



1. "Caching: What it is and How it Works" by Akamai Technologies Inc. (https://www.akamai.com/us/en/resources/caching.jsp)

This article provides a comprehensive overview of caching and how it works in the context of web applications and content delivery networks. It covers topics such as caching principles, benefits, types of caching, caching techniques, and best practices for cache management.

2. "Caching Strategies and Best Practices" by Cloudflare. (https://developers.cloudflare.com/cache/about/caching-strategies)

This guide outlines caching strategies and best practices for improving website performance and reducing server load. It covers a range of topics, including cache expiration, cache key design, caching for dynamic content, and dealing with cache invalidation.

3. "HTTP Caching" by MDN Web Docs. (https://developer.mozilla.org/en-US/docs/Web/HTTP/Caching)

This article provides a detailed overview of HTTP caching, including cache headers, cache revalidation, and cache control directives. It also includes information on how to configure caching for different types of responses and how to troubleshoot caching issues.

4. "Introduction to Caching in ASP.NET Core" by Microsoft. (https://docs.microsoft.com/en-us/aspnet/core/performance/caching/introduction)

This article provides a tutorial on caching in ASP.NET Core applications. It covers topics such as in-memory caching, distributed caching, cache tag helpers, and cache invalidation strategies.

5. "Server-Side Caching in Node.js Applications" by RisingStack Engineering. (https://blog.risingstack.com/server-side-caching-in-node-js/)

This blog post provides an overview of server-side caching techniques for Node.js applications. It covers topics such as in-memory caching, caching with Redis, cache expiration, and cache control. It also includes examples of how to implement caching in Node.js applications.   

