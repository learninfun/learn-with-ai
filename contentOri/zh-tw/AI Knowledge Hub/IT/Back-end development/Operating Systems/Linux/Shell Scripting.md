+++
title = "Shell腳本"
weight = "2"
+++
## 習題預習
{{< ask_chatgpt >}}
給我5題Shell Scripting的中文問題
{{< /ask_chatgpt >}}



1. 在Linux中如何創建一個新的空白文件？
2. 如何在Shell Scripting中使用條件語句來判斷一個數字是奇數還是偶數？
3. 如何使用Shell腳本查找並替換文件中的特定字符串？
4. 如何使用Shell腳本獲取當前日期和時間？
5. 如何使用Shell腳本計算一個文件的行數並將其輸出到屏幕上？   

## 說明知識
{{< ask_chatgpt >}}
以中文說明Shell Scripting並舉例
{{< /ask_chatgpt >}}



Shell Scripting是一種編程語言，擁有類似於命令行和終端的GUI介面。它可以用來自動化繁瑣、重複的任務，一次性處理多項計算，處理文本、系統管理等任務。

以下是一個簡單的Shell Scripting的例子：

```
#!/bin/bash
echo "What is your name?"
read name
echo "Hello, $name!"
```

這個腳本使用`echo`命令打印問候語，並用`read`命令從用戶處讀取一個變量`name`。然後使用`echo`命令打印個性化問候語，其中`$name`變量插入到字符串中。

這裡有一些更複雜的例子：

1. 查找目錄中的所有文件，並打印文件列表：

```
#!/bin/bash
for file in /path/to/directory/*; do
    echo $file
done
```

此腳本使用`for`循環遞歸地查找目錄中所有文件。使用通配符`*`可以匹配所有文件。

2. 自動刪除7天以上沒被修改的文件：

```
#!/bin/bash
find /path/to/directory -type f -mtime +7 -delete
```

此腳本使用`find`命令搜索目錄中的所有文件，並使用`-mtime`標誌過濾7天以上沒有被修改的文件。然後，使用`-delete`標誌刪除這些文件。

總的來說，Shell Scripting是一個非常有用的工具，可以幫助完成許多進程和操作。但是，需要謹慎使用，因為它可能會對系統造成意外的影響。   

## 彙總重點
{{< ask_chatgpt >}}
以中文條列Shell Scripting的重點
{{< /ask_chatgpt >}}



1. Shell環境變量：Shell環境變量是一組在當前Shell進程中存在的變量，用來保存Shell的配置信息以及其它一些有用的信息，如當前用戶名、主機名、系統路徑等。

2. Shell腳本的格式：Shell腳本文件通常使用.sh或.bash擴展名，腳本文件通常應該以#!/bin/sh或#!/bin/bash開頭，以指示腳本文件在哪個Shell下運行。

3. Shell腳本中的基本語法：Shell腳本中的基本通用語法包括：變量賦值、條件語句、循環語句以及命令執行等。

4. Shell腳本中的算術運算：Shell腳本可以使用算術運算符來進行算術運算，包括：+、-、*、/、%等。

5. Shell腳本中的邏輯運算：Shell腳本可以使用邏輯運算符來進行邏輯運算，包括：&&、||、!等。

6. Shell腳本中的數組：Shell腳本支持數組，可以通過數組實現對一組數據的處理。

7. Shell腳本中的函數：Shell腳本支持函數，可以將一些重複的代碼塊封裝成函數，以提高代碼的可讀性和可維護性。

8. Shell腳本中的命令執行和管道：Shell腳本支持命令的執行和管道操作，可以將多個命令連接起來形成一個流程。

9. Shell腳本中的錯誤處理：Shell腳本中可以通過使用set -e命令來開啟錯誤監測，並使用trap命令來處理錯誤。可以讓腳本在發生錯誤時及時處理，防止程序出錯後癱瘓。   

## 知識測驗
{{< ask_chatgpt >}}
以中文給我5題Shell Scripting的中等難度問題，並在後面列出答案
{{< /ask_chatgpt >}}



1. 請寫一支 Shell Script，能夠計算目錄下所有檔案的行數總和。要求排除所有的空行以及以 "#" 符號開頭的行。

答案：

```
#!/bin/bash

total=0

for file in *
do
 if [ -f "$file" ]
 then
  count=$(grep -v "^#" "$file" | grep -v "^$" | wc -l)
  total=$((total+count))
 fi
done

echo "Total number of lines: $total"
```

2. 請寫一支 Shell Script，能夠找出當前目錄下所有大於 1MB 的檔案，並將它們的檔名列出來。

答案：

```
#!/bin/bash

for file in *
do
 if [ -f "$file" ]
 then
  size=$(du -m "$file" | cut -f1)
  if [ "$size" -gt 1 ]
  then
   echo "$file"
  fi
 fi
done
```

3. 請寫一支 Shell Script，能夠列出目錄下所有的子目錄及其下的檔案總數。要求檔案總數不包含子目錄。

答案：

```
#!/bin/bash

for dir in */
do
  count=0
  for file in "$dir"*
  do
    if [ -f "$file" ]
    then
      count=$((count+1))
    fi
  done
  echo "$dir: $count files"
done
```

4. 請寫一支 Shell Script，能夠找出當前目錄下最新修改的 5 個檔案，並依時間排序列出它們的檔名。

答案：

```
#!/bin/bash

ls -t | head -n 5
```

5. 請寫一支 Shell Script，能夠將目錄下所有以 ".txt" 結尾的檔案，複製到另一個目錄中。要求該目錄不存在時請自動建立。

答案：

```
#!/bin/bash

dest_dir="/home/user/text_files"

if [ ! -d "$dest_dir" ]
then
  mkdir "$dest_dir"
fi

for file in *.txt
do
  cp "$file" "$dest_dir"
done
```   

