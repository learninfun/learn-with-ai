

1. 如何设定并验证网路接口的MTU？
答案：可以使用 ifconfig 命令来设定网路接口的MTU，并使用 ping 或其他工具来验证MTU是否正确运作。

2. 如何查看系统的随机数产生器的种子？
答案：可以查看 /proc/sys/kernel/random/entropy_avail 档案的内容，该档案显示了系统目前的随机数产生器种子数。

3. 如何设定系统的DNS伺服器？
答案：可以编辑 /etc/resolv.conf 档案，加入一行 nameserver IP_ADDRESS，其中 IP_ADDRESS 是您要使用的 DNS 伺服器的 IP 位址。

4. 如何复制整个档案系统？
答案：可以使用 dd 命令来复制整个档案系统，如 dd if=/dev/source_disk of=/dev/destination_disk bs=1M。

5. 如何监测系统的硬件资源利用率？
答案：可以使用 top 或 htop 命令来监测系统的 CPU，记忆体等硬件资源利用率。也可以使用 sar 或 vmstat 命令记录资源利用率的历史纪录。