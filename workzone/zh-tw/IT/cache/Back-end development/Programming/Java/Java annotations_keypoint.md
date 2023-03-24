

1. Annotations是Java 5中引入的新特性，可以為Java程序提供元數據信息，使得程序的開發、部署、測試等工作更加靈活。
2. Annotations可以在編譯時、運行時或甚至在部署時通過反射機制來讀取信息，對程序的調試和優化工作非常有幫助。
3. 常用的Java Annotations包括Override、Deprecated、SuppressWarnings、Inherited等。
4. Override用於標注方法覆蓋了父類的方法，編譯時可以檢查是否正確覆蓋。
5. Deprecated用於標注已經過期的方法或類，建議不再使用。
6. SuppressWarnings用於關閉Java編譯器的警告信息。
7. Inherited用於標注子類是否繼承父類的Annotation。
8. 自定義註解可以通過@Target和@Retention等註解來定義作用域和保留期。
9. 註解處理器可以通過apt工具來自動化生成代碼，簡化開發工作。