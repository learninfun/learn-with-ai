

Containers（容器）是指在一台主機上，使用容器技術將應用程式所需的所有運行環境、庫、配置文件等打包在一起的一種環境隔離技術。每個容器都在共享的操作系統核心上運行，但是被隔離開來，獨立運行，並且具有自己的應用程序和庫。容器中的應用程序可以在不影響主機和其他容器的情況下運行。

舉例來說，一個網路應用程式可能需要運行在特定版本的作業系統和軟件庫上。使用容器化技術，可以將該應用程式及其相關文件打包在一起，確保它可以在一致的環境中運行。這樣可以消除在不同機器或不同環境中運行應用程序時可能出現的問題。此外，如果應用程序需要更新或升級，只需更新容器即可，而不需要更改操作系統或應用程序的配置文件。常見的容器平台包括Docker，Kubernetes等。