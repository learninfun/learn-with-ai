+++
title = "快取"
+++
## 习题预习
{{< ask_chatgpt >}}
给我5题Caching的问题
{{< /ask_chatgpt >}}



1. 什麼是快取(Caching)？在什麼情況下使用快取可以提高系統的效能？

2. 快取的失效率是什麼？它對系統性能有什麼影響？

3. 常見的快取策略有哪些？它們之間有什麼區別？

4. 如何測試和評估系統的快取策略是否有效？

5. 快取穿透和快取雪崩是什麼？它們有什麼解決方案？   

## 说明知识
{{< ask_chatgpt >}}
说明Caching并举例
{{< /ask_chatgpt >}}



Caching是將常用的數據存儲在快速、易於存取的高速記憶體中，以加快數據訪問速度的行為。這通常是因為存取主記憶體（RAM）或者硬碟等設備的速度較慢，載入數據需要大量時間和資源，如果非常依賴主存和硬碟來加載資料，將導致速度緩慢。

例如，在Web應用程序中，經常會將比較大的數據（例如圖像、CSS和JS文件）存儲在瀏覽器的快取中，以加快使用者訪問網站的速度。另一個常見的例子是瀏覽器的歷史紀錄。你最近訪問的網站將保存在你的瀏覽器的快取中，因此下一次訪問時將不需要重新載入所有數據。

在一些資料密集型的應用程序中，也可以使用緩存網格，將資料存儲在高速讀取和儲存的內存裡，而不是在慢速硬碟上。這樣可以大幅提高數據訪問速度，從而減少服務器負載。

總之，緩存是提高訪問速度並減少服務器負載的一種重要技術，可以應用到各種應用場景中。   

## 汇总重点
{{< ask_chatgpt >}}
条列Caching的重点
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

## 知识测验
{{< ask_chatgpt >}}
给我5题Caching的中等难度问题，并在后面列出答案
{{< /ask_chatgpt >}}



1. 给定一个有 $n$ 个元素的整数阵列和一个整数 $k$，请设计一个支持查询区间和的数据结构，要求时间复杂度为 $O(n\sqrt{n} + q\sqrt{n}\log n)$，其中 $q$ 为查询数量。

2. 设计一个线性时间复杂度的算法，查询一个区间中出现次数最多的元素及其出现次数。

3. 给定一个长度为 $n$ 的字符串 $S$，设计一个数据结构支持以下操作：

   - 将某个位置的字符修改为另一个字符。
   - 查询某个子串在原串中出现的次数。

   要求修改和查询操作的时间复杂度均为 $O(\sqrt{n})$。

4. 设计一个数据结构，支持以下操作：

   - 将一个元素加入集合中。
   - 从集合中删除一个元素。
   - 查询某个区间内元素的个数。

   其中，添加和删除操作的时间复杂度均为 $O(\sqrt{n})$，查询操作的时间复杂度为 $O(q\sqrt{n}\log n)$，其中 $q$ 为查询次数。

5. 设计一个数据结构，支持以下操作：

   - 在序列中插入一个元素。
   - 删除序列中的某个元素。
   - 查询某个位置后面第 $k$ 小的元素。

   其中，插入和删除操作的时间复杂度均为 $O(\sqrt{n})$，查询操作的时间复杂度为 $O(q\sqrt{n}\log n)$，其中 $q$ 为查询次数。

答案：

1. 利用分块和前缀和，在每个块上建立一棵平衡树（如红黑树），支持区间查询的时间复杂度为 $O(\sqrt{n}\log n)$，将每个块区间和缓存起来，单次查询的时间复杂度为 $O(\sqrt{n})$。
2. 将序列分为若干块，对每个块统计出现次数最多的元素和出现次数，总时间复杂度为 $O(n)$。对于一个查询区间 $[l,r]$，若 $l$ 和 $r$ 落在同一个块中，直接遍历区间统计出现次数即可；否则分别遍历区间左右端点所在的块，并统计其中出现次数最多的元素及其出现次数，然后在这些块中遍历 $l$ 和 $r$ 所在的块中非区间的元素，统计其出现次数，最终比较三者出现次数，返回出现次数最多的那个元素及其出现次数。
3. 将字符串划分成 $\sqrt{n}$ 个块，对每个块建立一棵 Trie（字典树），在 Trie 上标记每个节点对应的字串在原串中出现的次数，缓存每个块中所有子串的出现次数，查询时统计相应块中子串的出现次数，然后在相应的 Trie 上遍历所查询的子串，统计其出现次数，最终返回所有统计值的和。修改操作可以在 Trie 上进行，时间复杂度为 $O(|T|)$，其中 $T$ 为 Trie 的节点数，对于一个子串修改，只需找到相应的 Trie 和其对应的节点，然后修改该节点的标记，时间复杂度为 $O(\sqrt{n}+\log m)$，其中 $m$ 为字元集大小。
4. 将集合分割成 $\sqrt{n}$ 个块，对于每个块使用哈希表和平衡树（如红黑树）进行支持添加和删除操作，时间复杂度均为 $O(\sqrt{n})$。对于一个查询区间 $[l,r]$，若 $l$ 和 $r$ 落在同一个块中，直接遍历区间统计元素个数；否则分别遍历区间左右端点所在的块，然后在这些块中遍历 $l$ 和 $r$ 所在的块中非区间的元素，统计其出现次数，最终返回统计值的和。总时间复杂度为 $O(n\sqrt{n}+q\sqrt{n}\log n)$。
5. 将序列分为 $\sqrt{n}$ 个块，对于每个块使用一棵支持重复元素的排序算法（如 std::multiset）进行排序，时间复杂度为 $O(\sqrt{n}\log\sqrt{n})$。将每个块中的第 $k$ 小的元素缓存起来，总时间复杂度为 $O(n\sqrt{n})$。对于一个插入或删除操作，只需找到相应的块，更新该块中的排序算法，更新缓存中的第 $k$ 小元素，时间复杂度为 $O(\sqrt{n}\log\sqrt{n})$。对于一个查询操作，若要查询的位置位于某个块中，直接在该块中进行查询；否则先在位置左边的块中查询其后第 $k$ 小的元素，然后在位置右边的块中查询其前 $k-1$ 小的元素，最终找出所有结果的前 $k$ 小的元素，总时间复杂度为 $O(q\sqrt{n}\log\sqrt{n})$。   

## 网络数据
{{< ask_chatgpt >}}
给我5篇Caching的网络数据
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

