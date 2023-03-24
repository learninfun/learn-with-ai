

1. What is the default file permission for a newly created file in Linux and how can you modify it?
Answer: The default file permission for a newly created file in Linux is 666 (rw-rw-rw-). You can modify it by using the chmod command and specifying the appropriate permission settings, such as 755 (rwxr-xr-x) or 644 (rw-r--r--).

2. How can you configure a firewall in Linux to block incoming traffic from a specific IP address?
Answer: You can use the iptables command to configure a firewall in Linux. To block incoming traffic from a specific IP address, you can use the following command: iptables -A INPUT -s <IP address> -j DROP

3. What is the purpose of SELinux in Linux and how can you enable it?
Answer: SELinux is a security mechanism in Linux that provides fine-grained access control and mandatory access control (MAC) over files, processes, and other system resources. You can enable SELinux by modifying the /etc/selinux/config file and setting the SELINUX parameter to enforcing or permissive mode.

4. How can you encrypt a file or directory in Linux and what is the recommended algorithm to use?
Answer: You can encrypt a file or directory in Linux using the GnuPG (GPG) tool, which supports symmetric-key and public-key encryption. The recommended algorithm to use for symmetric-key encryption is AES-256, which provides strong encryption and is widely used in security applications.

5. How can you configure Linux to rotate log files automatically and prevent disk space from being filled?
Answer: You can configure Linux to rotate log files automatically by using the logrotate utility, which can be configured to rotate log files based on size, age, or other criteria. To prevent disk space from being filled, you can set the maxsize parameter in the logrotate configuration file and specify the maximum size that a log file can grow before it is rotated.