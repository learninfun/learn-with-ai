## 習題預習
{{< ask_chatgpt >}}
給我5題Shell Scripting的問題
{{< /ask_chatgpt >}}



1. 在Linux中如何創建一個新的空白文件？
2. 如何在Shell Scripting中使用條件語句來判斷一個數字是奇數還是偶數？
3. 如何使用Shell腳本查找並替換文件中的特定字符串？
4. 如何使用Shell腳本獲取當前日期和時間？
5. 如何使用Shell腳本計算一個文件的行數並將其輸出到屏幕上？   

## 說明知識
{{< ask_chatgpt >}}
說明Shell Scripting並舉例
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
條列Shell Scripting的重點
{{< /ask_chatgpt >}}



1. Shell环境变量：Shell环境变量是一组在当前Shell进程中存在的变量，用来保存Shell的配置信息以及其它一些有用的信息，如当前用户名、主机名、系统路径等。

2. Shell脚本的格式：Shell脚本文件通常使用.sh或.bash扩展名，脚本文件通常应该以#!/bin/sh或#!/bin/bash开头，以指示脚本文件在哪个Shell下运行。

3. Shell脚本中的基本语法：Shell脚本中的基本通用语法包括：变量赋值、条件语句、循环语句以及命令执行等。

4. Shell脚本中的算术运算：Shell脚本可以使用算术运算符来进行算术运算，包括：+、-、*、/、%等。

5. Shell脚本中的逻辑运算：Shell脚本可以使用逻辑运算符来进行逻辑运算，包括：&&、||、!等。

6. Shell脚本中的数组：Shell脚本支持数组，可以通过数组实现对一组数据的处理。

7. Shell脚本中的函数：Shell脚本支持函数，可以将一些重复的代码块封装成函数，以提高代码的可读性和可维护性。

8. Shell脚本中的命令执行和管道：Shell脚本支持命令的执行和管道操作，可以将多个命令连接起来形成一个流程。

9. Shell脚本中的错误处理：Shell脚本中可以通过使用set -e命令来开启错误监测，并使用trap命令来处理错误。可以让脚本在发生错误时及时处理，防止程序出错后瘫痪。   

## 知識測驗
{{< ask_chatgpt >}}
給我5題Shell Scripting的中等難度問題，並在後面列出答案
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

## 網路資料
{{< ask_chatgpt >}}
給我5篇Shell Scripting的網路資料
{{< /ask_chatgpt >}}



1. Shell Scripting Tutorial – A Beginner's Guide
網址: https://www.shellscript.sh/
此教程專門針對初學者，從基礎介紹到進階技巧都有覆蓋,範圍廣泛。

2. Shell Scripting Examples
網址: https://www.cyberciti.biz/tips/category/shell-scripting
此網站會提供各種 Shell Scripting 的範例，分類很清楚並且會介紹每一段 Script 的功用。

3. Bash Shell Scripting Tutorial
網址: https://bash.cyberciti.biz/guide/Main_Page
此教程專門介紹 Bash Shell Scripting，從基礎介紹到進階技巧都會有覆蓋。

4. Linux Shell Scripting Cookbook
網址: https://www.packtpub.com/product/linux-shell-scripting-cookbook-third-edition/9781800563769
此書籍特別針對 Linux Shell Scripting 進行詳細的介紹，包含很多實用的範例和技巧。

5. Advanced Bash-Scripting Guide
網址: http://www.tldp.org/LDP/abs/html/
此指南介紹了 Linux 系統中常用的 Bash Scripting 技巧和實用範例，是一份全面而且詳盡的學習資源。   

