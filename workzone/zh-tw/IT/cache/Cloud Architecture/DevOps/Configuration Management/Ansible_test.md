1. 如何從Ansible playbook中選擇特定的Play來執行？
2. 如何從Ansible playbook中選擇特定的Tasks來執行？
3. 如何在Ansible playbook中傳遞變量？如何在該變量中包含特殊字符？
4. 如何將Ansible playbook中使用的SSH金鑰指定為變量？
5. 如何讓Ansible playbook在特定的日期和時間執行？

答案：
1. 使用 --limit 選項指定Playbook要運行的主機或主機組。
2. 使用 --tags 選項指定要運行的Task標籤。
3. 可以使用vars_prompt在Playbook中提示變量值，也可以在命令行中使用 --extra-vars 參數指定變量。使用{{ }}表示法編寫變量，如需使用特殊符號，可以使用反斜線進行轉義。
4. 可以在 Playbook 中使用 vars 或 vars_file 定義 SSH 金鑰的變量，也可以在命令行中使用 --extra-vars 參數指定 SSH 金鑰。
5. 使用 Ansible 的 cron 模塊，創建一個 cron 任務，按照特定的日期和時間執行 Playbook。