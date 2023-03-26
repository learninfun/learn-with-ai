

Virtualization与Containerization都是用来区隔不同应用程式之间的环境，让它们可以共存及隔离，但两者概念不太一样。

Virtualization是将物理主机的硬体资源，如 CPU、记忆体、硬碟等等，透过虚拟化技术，让多个虚拟机器（Virtual Machine）在同一台主机上运作，每个虚拟机器都像是一台独立的电脑，拥有自己的作业系统、应用程式和文件。举例来说，在一台物理主机上可以运作多个不同的虚拟机器，每个虚拟机器都可以执行不同的作业系统，例如在 Windows 主机上同时运作 Linux 和 Windows Server 两个虚拟机器。

Containerization则是利用容器技术，将应用程式及相关依赖套件放在一个独立的环境中，形成一个称为容器（Container）的独立执行环境。容器间相互隔离，而且不需要额外的作业系统层，因此可以更轻量、更快速地运作，并有助于开发人员在不同环境中快速部署和运行应用程式。举例来说，同一个主机上可以运作多个不同的容器，每个容器可以拥有不同的应用程式，例如在一台主机上运行多个网站容器，每个容器中都有不同的网站应用程式及相关的资料库。