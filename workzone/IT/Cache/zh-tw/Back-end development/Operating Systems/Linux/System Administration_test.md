

1. 如何設定並驗證網路接口的MTU？
答案：可以使用 ifconfig 命令來設定網路接口的MTU，並使用 ping 或其他工具來驗證MTU是否正確運作。

2. 如何查看系統的隨機數產生器的種子？
答案：可以查看 /proc/sys/kernel/random/entropy_avail 檔案的內容，該檔案顯示了系統目前的隨機數產生器種子數。

3. 如何設定系統的DNS伺服器？
答案：可以編輯 /etc/resolv.conf 檔案，加入一行 nameserver IP_ADDRESS，其中 IP_ADDRESS 是您要使用的 DNS 伺服器的 IP 位址。

4. 如何複製整個檔案系統？
答案：可以使用 dd 命令來複製整個檔案系統，如 dd if=/dev/source_disk of=/dev/destination_disk bs=1M。

5. 如何監測系統的硬件資源利用率？
答案：可以使用 top 或 htop 命令來監測系統的 CPU，記憶體等硬件資源利用率。也可以使用 sar 或 vmstat 命令記錄資源利用率的歷史紀錄。