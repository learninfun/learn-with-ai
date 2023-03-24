

1. 请写一支 Shell Script，能够计算目录下所有档案的行数总和。要求排除所有的空行以及以 "#" 符号开头的行。

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

2. 请写一支 Shell Script，能够找出当前目录下所有大于 1MB 的档案，并将它们的档名列出来。

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

3. 请写一支 Shell Script，能够列出目录下所有的子目录及其下的档案总数。要求档案总数不包含子目录。

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

4. 请写一支 Shell Script，能够找出当前目录下最新修改的 5 个档案，并依时间排序列出它们的档名。

答案：

```
#!/bin/bash

ls -t | head -n 5
```

5. 请写一支 Shell Script，能够将目录下所有以 ".txt" 结尾的档案，复制到另一个目录中。要求该目录不存在时请自动建立。

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