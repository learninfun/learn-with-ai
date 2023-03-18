

1. 在Linux中，如何查看正在運行的服務（services）及其相關的端口號？
答案：可以使用命令“netstat -tulpn”查看所有正在運行的服務及其關聯的端口號。

2. 如何禁止使用者在Linux中通過SSH密碼的方式進入系統？
答案：可以在SSH設定文件（/etc/ssh/sshd_config）中將“PasswordAuthentication”項目設置為“No”。

3. 如何限制某個使用者只能訪問特定的目錄？
答案：可以使用“chroot”命令來限制使用者的訪問權限。也可以為該使用者配置特定的權限和權限設置。

4. 如何保護Linux中的目錄免受未授權訪問？
答案：可以使用“chmod”命令更改文件或目錄的權限設置，以控制對它們的訪問許可權。

5. 如何使Linux中的文件實現加密保護？
答案：可以使用加密文件系統（如LUKS或eCryptfs）或加密軟件（如GnuPG或OpenSSL）將文件加密保護。