

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

5. 問題：如何在PowerShell中使用Get-ADUser命令查找一個特定群組中的所有用戶，并將它們的電子郵件地址輸出到一個CSV文件中？
答案：

$users = Get-ADGroupMember "MyGroup" | Where-Object {$_.objectClass -eq "user"}
$userEmails = @()
foreach ($user in $users) {
    $userEmails += Get-ADUser $user.SamAccountName -Properties EmailAddress | Select-Object EmailAddress
}
$userEmails | Export-Csv "C:\userEmails.csv" -NoTypeInformation