

1. 如何在Hyper-V中建立一个虚拟交换机(Virtual Switch)？ 
2. 如何将现有的物理硬碟加入到Hyper-V的虚拟环境中？ 
3. 如何配置虚拟机器(Virtual Machine)的CPU和内存资源？ 
4. 如何将可移动的虚拟硬碟(Virtual Hard Disk)移动到不同的运行Hyper-V的伺服器？ 
5. 如何在Hyper-V中建立一个新的虚拟网路选项卡(New Virtual Network Adapter)？ 

答案：
1. 右键单击Hyper-V管理员，从下拉选单中选择"虚拟交换机管理器"，并按下"新增虚拟交换机"按钮。
2. 打开Hyper-V管理员，找到目标虚拟机器，点击设置，点击"硬碟"，然后点击"加入"，找到要使用的现有物理硬碟。
3. 打开Hyper-V管理员，点击目标虚拟机器，点击"设置"，选择"处理器"或"记忆体"。在那里，您可以配置分配给虚拟机器的CPU和RAM资源。
4. 转换虚拟硬碟时，您必须停止使用虚拟硬碟的虚拟机器。然后，将虚拟硬碟的位置修改为你想移到的那台Hyper-V伺服器上的硬碟路径。接着，您可以启动虚拟机器，并执行虚拟硬碟上的作业系统(OS)。
5. 打开Hyper-V管理员，从下拉选单中选择"虚拟交换机管理器"，然后在右侧中按下新的虚拟交换机，并选择"内部"或"外部"（桥接到物理网路），最后将其命名并按下"OK"。