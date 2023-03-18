

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