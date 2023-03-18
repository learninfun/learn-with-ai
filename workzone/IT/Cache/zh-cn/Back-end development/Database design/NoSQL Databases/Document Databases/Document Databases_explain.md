

Document Databases是一种NoSQL（Not Only SQL）非关联式资料库，其储存资料的方式以“文件”（Document）的形式储存。

与传统的关联式资料库不同，Document Databases没有固定的资料表和栏位，而是使用一种称为“集合”（Collection）的方式储存，每个集合可以容纳多个文件。

每个文件都是一个独立的单位，可以包含任意数量的栏位，甚至是嵌套的结构。以JavaScript Object Notation（JSON）格式表示，Document Databases能够很好地处理非结构化或半结构化的资料，如文章、新闻、照片、影片、音乐等。

常见的Document Databases包括：

1. MongoDB：目前最为流行的Document Databases，使用JSON格式与二进制格式储存资料，支援复杂的查询和索引。

2. Couchbase：基于CouchDB的Document Databases，支援多种插件和API，提供高可用性和可扩展性。

3. RavenDB：使用C＃开发的Document Databases，支援多种平台和语言，提供全文检索和实时资料同步等功能。

4. CouchDB：早期的Document Databases，使用JSON格式储存资料，支援离线同步和MapReduce。

5. Amazon DynamoDB：基于AWS的Document Databases，提供高度可扩展性和灵活的资料模型，支援多种API和SDK。

总体而言，Document Databases能够快速处理大量的非结构化或半结构化的资料，并且提供高度的可扩展性和可用性。