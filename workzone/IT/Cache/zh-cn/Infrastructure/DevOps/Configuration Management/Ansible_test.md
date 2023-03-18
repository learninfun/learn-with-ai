1. 如何在playbook中设置变量的默认值？

答案：使用default关键字。例如：{{ my_var | default('my_default_value') }}

2. 如何在Ansible中使用条件语句（If-else statements）？

答案：使用when关键字进行条件判断。例如：

```
- name: Check if my_var is true
  debug:
    msg: My var is true
  when: my_var == true
```

3. 如何定义和使用Ansible角色（Role）？

答案：使用ansible-galaxy命令行工具创建和安装角色。然后在playbook中使用角色。例如：

```
- hosts: my_host
  roles:
    - my_role
```

4. 如何在Ansible中管理SSH密钥（SSH keys）？

答案：使用ssh_keygen模组创建SSH密钥，并使用ssh_authorized_key模组将公钥添加到受管计算机上的authorized_keys文件中。

5. 如何使用Ansible Vault来保护敏感信息？

答案：使用ansible-vault命令创建加密文件（使用密码或密钥），然后在playbook中使用vars_files将其引用。例如：

```
- hosts: my_host
  vars_files:
    - /path/to/my_vault_file.yml
  tasks:
    - name: My task
      debug:
        msg: "My encrypted variable is {{ my_encrypted_var }}"
```