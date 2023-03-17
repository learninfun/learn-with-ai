## 習題預習
{{< ask_chatgpt >}}
給我5題PowerShell的問題
{{< /ask_chatgpt >}}



1. 如何列出指定目錄下的所有檔案和子目錄？
2. 如何列出所有執行中的進程和其相關資訊，如PID和記憶體使用量？
3. 如何訪問和修改Windows註冊表中的數值？
4. 如何創建一個新的資料夾，並在其中編寫文字檔？
5. 如何通過PowerShell啟動另一個應用程序並向其傳遞參數？   

## 說明知識
{{< ask_chatgpt >}}
說明PowerShell並舉例
{{< /ask_chatgpt >}}



PowerShell是一種動態的命令列介面和腳本語言，用於自動化許多Windows操作系統的任務和管理。

PowerShell的優點之一是其處理.NET框架的能力，使得PowerShell能夠與多種平台和應用程序協同工作，並提供龐大的擴展性和功能。

以下是一些示例：

1. 列出磁碟上的所有文件：Get-ChildItem C:\

2. 查詢現有服務的信息：Get-Service

3. 創建一個新文件夾：New-Item -ItemType Directory -Path C:\NewFolder

4. 執行檔案的程序：Invoke-Item -Path "C:\Example.exe"

5. 設置系統環境變量：[Environment]::SetEnvironmentVariable("Path",$env:Path + ";C:\NewPath", "User")

6. 壓縮和解壓縮檔案：Compress-Archive -Path C:\Files\* -DestinationPath C:\Archive.zip， Expand-Archive -Path C:\Archive.zip -DestinationPath C:\UnzippedFiles

7. 自動安裝應用軟件：Install-Package -Name ExampleSoftware

PowerShell是一個非常強大的工具，可以使任務自動化，提高效率並減少錯誤。 它的學習曲線可能有些陡峭，但一旦熟練，它可以大大簡化日常工作。   

## 彙總重點
{{< ask_chatgpt >}}
條列PowerShell的重點
{{< /ask_chatgpt >}}



下面列出了 PowerShell 的重點：

1. 管理和自動化任務： PowerShell是用於管理和自動化Windows操作系統的工具。它允許管理員執行一些複雜的任務，如批量管理Windows Server，編寫腳本來自動化常用的管理任務。

2. 命令行工具： PowerShell提供了一個命令行界面，管理員可以使用該界面在Windows上執行各種任務。

3. 網絡支持： PowerShell通過PowerShell Remoting提供網絡支持，以允許管理員在網絡上執行遠程命令。

4. 管道： PowerShell中的管道功能允許管理員將命令鏈在一起以完成特定任務，同時將輸出從一個命令傳遞到另一個命令。

5. 腳本語言： PowerShell是一種編程語言，允許管理員編寫自己的腳本來執行各種管理任務。

6. 可擴展性： PowerShell可以通過添加模塊來擴展其功能。模塊是一組命令、函數和其他資源，可以在PowerShell中加載和使用。

7. 追蹤和分析日誌： PowerShell提供了用於追蹤和分析日誌的工具，允許管理員監控各種系統事件和活動。   

## 知識測驗
{{< ask_chatgpt >}}
給我5題PowerShell的中等難度問題，並在後面列出答案
{{< /ask_chatgpt >}}



1. 問題：如何使用PowerShell將一個文本文件的內容分成多個文件？
答案：

$content = Get-Content "C:\input.txt"
$newFiles = [System.Collections.ArrayList]@()
$i = 1
$size = 3
while ($i -lt $content.Length) {
    $newFile = $content[$i..($i+$size-1)] | Out-File "C:\file$i.txt" -Encoding utf8 -Width ($content[$i..($i+$size-1)] | Measure-Object -Maximum Length | Select-Object -ExpandProperty Maximum)
    $newFiles.Add($newFile.Name) | Out-Null
    $i += $size
}
$newFiles

2. 問題：如何創建一個包含隨機數字的日期文件夾，並將文件夾中的文件名稱隨機改變？
答案：

$dateFolder = Get-Date -Format yyyy-MM-dd-HH-mm-ss-ffff
New-Item -ItemType Directory -Path "C:\$dateFolder"
Get-ChildItem | Where-Object {$_.PsIsContainer -eq $False} | ForEach-Object {
    $newName = Get-Random -Maximum 1000 -Minimum 100
    Rename-Item $_.FullName -NewName "$newName.txt"
}

3. 問題：如何使用PowerShell查找一個文件夾中的所有文件，並將它們存儲到一個數組中？
答案：

$files = Get-ChildItem "C:\MyFolder" -File | Select-Object FullName
$files

4. 問題：如何在PowerShell中使用Start-Process命令啟動一個應用程序，並將它的輸出保存到一個變量中？
答案：

$output = Start-Process -FilePath "C:\Program Files (x86)\MyApp\MyApp.exe" -ArgumentList "-arg1", "-arg2" -NoNewWindow -PassThru -Wait -RedirectStandardOutput "C:\output.txt" -WindowStyle Hidden
Get-Content "C:\output.txt"

5. 問題：如何在PowerShell中使用Get-ADUser命令查找一個特定群組中的所有用戶，並將它們的電子郵件地址輸出到一個CSV文件中？
答案：

$users = Get-ADGroupMember "MyGroup" | Where-Object {$_.objectClass -eq "user"}
$userEmails = @()
foreach ($user in $users) {
    $userEmails += Get-ADUser $user.SamAccountName -Properties EmailAddress | Select-Object EmailAddress
}
$userEmails | Export-Csv "C:\userEmails.csv" -NoTypeInformation   

## 網路資料
{{< ask_chatgpt >}}
給我5篇PowerShell的網路資料
{{< /ask_chatgpt >}}



1. "Getting Started with PowerShell" - https://docs.microsoft.com/en-us/powershell/scripting/setup/getting-started-with-powershell

This article provides a great introduction to PowerShell, covering everything from installation and setup to basic scripting concepts. It includes a list of resources for further learning, making it a great starting point for beginners.

2. "PowerShell Scripting Best Practices" - https://devblogs.microsoft.com/powershell/powershell-scripting-best-practices/

This article offers some best practices for writing efficient and effective PowerShell scripts. It covers topics such as error handling, commenting, and parameter validation, and includes examples of good and bad code.

3. "PowerShell Basics: Variables, Arrays, Hash Tables, and Conditional Statements" - https://www.altaro.com/msp-dojo/powershell-basics-variables-arrays-hashtables-conditional-statements/

This article provides a useful overview of some of the key features of PowerShell, including variables, arrays, and conditional statements. It includes examples of how to use these features in scripts, making it a great resource for those who are new to PowerShell scripting.

4. "PowerShell for Active Directory" - https://activedirectorypro.com/powershell-for-active-directory/

This article focuses specifically on using PowerShell to manage Active Directory. It covers a range of tasks, from creating and managing user accounts to querying AD data and generating reports. This is a great resource for those who need to manage AD using PowerShell.

5. "PowerShell 101: How to use PowerShell scripts to automate tasks on Windows 10" - https://www.windowscentral.com/how-use-powershell-scripts-automate-tasks-windows-10

This article offers a practical guide to using PowerShell to automate tasks on Windows 10. It includes step-by-step instructions for creating and running scripts, as well as some examples of common automation tasks. This is a great resource for those who want to automate their workflow using PowerShell.   

