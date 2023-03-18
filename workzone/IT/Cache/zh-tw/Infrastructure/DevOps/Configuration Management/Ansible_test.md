1. 如何在playbook中設置變量的默認值？

答案：使用default關鍵字。例如：{{ my_var | default('my_default_value') }}

2. 如何在Ansible中使用條件語句（If-else statements）？

答案：使用when關鍵字進行條件判斷。例如：

```
- name: Check if my_var is true
  debug:
    msg: My var is true
  when: my_var == true
```

3. 如何定義和使用Ansible角色（Role）？

答案：使用ansible-galaxy命令行工具創建和安裝角色。然後在playbook中使用角色。例如：

```
- hosts: my_host
  roles:
    - my_role
```

4. 如何在Ansible中管理SSH密鑰（SSH keys）？

答案：使用ssh_keygen模組創建SSH密鑰，並使用ssh_authorized_key模組將公鑰添加到受管計算機上的authorized_keys文件中。

5. 如何使用Ansible Vault來保護敏感信息？

答案：使用ansible-vault命令創建加密文件（使用密碼或密鑰），然後在playbook中使用vars_files將其引用。例如：

```
- hosts: my_host
  vars_files:
    - /path/to/my_vault_file.yml
  tasks:
    - name: My task
      debug:
        msg: "My encrypted variable is {{ my_encrypted_var }}"
```