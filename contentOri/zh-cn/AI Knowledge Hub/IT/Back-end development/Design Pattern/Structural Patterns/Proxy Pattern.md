+++
title = "代理模式"
weight = "7"
+++
## 習題預習
{{< ask_chatgpt >}}
給我5題Proxy Pattern的中文問題
{{< /ask_chatgpt >}}



1. 何謂Proxy Pattern？該模式在軟體設計中有何用處？
2. 請舉一個實際案例，說明如何使用Proxy Pattern來實現軟體功能需求？
3. Proxy Pattern和Decorator Pattern有什麼區別？請列出至少兩點差異。
4. 何時應該使用Proxy Pattern？請描述一個案例，說明使用Proxy Pattern的好處。
5. Proxy Pattern在Java中的實現方式有哪些？何時使用哪種方式比較適合？   

## 說明知識
{{< ask_chatgpt >}}
以中文說明Proxy Pattern並舉例
{{< /ask_chatgpt >}}



Proxy Pattern 是一種靜態設計模式，它可以將物件的存取權限控制在一個代理物件中，代理物件可以在真實的物件之前或之後執行一些額外的功能，例如遠端代理、虛擬代理、安全代理等。通過使用代理模式，客戶端可以透明地訪問一個物件，而不需要知道它的實際實現細節。

舉例來說，假設有一個 Image 接口，它有一個 display() 方法用來顯示圖片，一個 RealImage 類用來實現 Image 接口，另外還有一個 ProxyImage 類也實現了 Image 接口，它在顯示圖片之前先檢查是否有權限。

```
interface Image {
    void display();
}

class RealImage implements Image {
    private String imagePath;
    
    RealImage(String imagePath) {
        this.imagePath = imagePath;
        loadImageFromDisk();
    }
    
    private void loadImageFromDisk() {
        System.out.println("Loading image from disk: " + imagePath);
    }
    
    @Override
    public void display() {
        System.out.println("Displaying image: " + imagePath);
    }
}

class ProxyImage implements Image {
    private RealImage realImage;
    private String imagePath;
    
    ProxyImage(String imagePath) {
        this.imagePath = imagePath;
    }
    
    @Override
    public void display() {
        if (realImage == null) {
            realImage = new RealImage(imagePath);
        }
        realImage.display();
    }
}

public class Main {
    public static void main(String[] args) {
        Image image = new ProxyImage("test.jpg");
        image.display();
    }
}
```

在上面的例子中，當使用 ProxyImage 類顯示圖片時，它首先會檢查是否有權限，如果沒有權限，就不會顯示圖片；如果有權限，就會調用 RealImage 的 display() 方法來顯示圖片。這讓客戶端可以透明地訪問圖片，而不需要知道代理物件是否存在。   

## 條列重點
{{< ask_chatgpt >}}
以中文條列Proxy Pattern的重點
{{< /ask_chatgpt >}}



1. 為了控制物件的訪問，使用一個代理物件來代替原始物件。
2. 代理物件與原始物件實現相同的介面，因此可以在不影響使用者的情況下進行替換。
3. 代理模式可以實現許多不同的功能，例如安全檢查、遠端訪問、記錄日誌等。
4. 代理模式可以分為靜態代理和動態代理，靜態代理需要在代理類中手動編寫額外方法，動態代理則是實現了java.lang.reflect.InvocationHandler接口，運行時動態生成代理類。
5. 代理模式可以製造出一些看似替代品的代理物件，實際上他們是指向同一個目標物件，可以在不影響使用者情況下替換。
6. 代理模式常常使用在業務邏輯與外部數據或服務的交互中，將一些複雜或不安全的操作委託給代理處理。   

## 知識測驗
{{< ask_chatgpt >}}
以中文給我5題Proxy Pattern的中等難度問題，並在後面列出答案
{{< /ask_chatgpt >}}



1. 試著用 Proxy Pattern 實作一個簡單的網路瀏覽器，讓使用者可以輸入一個網址並瀏覽該網站。但是，當瀏覽一些特定網站時，我們想要在載入頁面之前先將網站標題改為「Dangerous Site」。該如何實現這個功能？

答案：

```java
public interface WebPage {
    public void display();
}

public class RealWebPage implements WebPage {
    private String url;

    public RealWebPage(String url) {
        this.url = url;
        load();
    }

    private void load() {
        System.out.println("Loading " + url);
    }

    public void display() {
        System.out.println("Displaying " + url);
    }
}

public class ProxyWebPage implements WebPage {
    private RealWebPage realWebPage;

    public ProxyWebPage(String url) {
        // 如果網址是 "dangerous.com"，則在實現 WebPage 的顯示方法顯示 "Dangerous Site"，否則顯示網址本身
        if (url.equals("dangerous.com")) {
            System.out.println("Access Denied");
        } else {
            realWebPage = new RealWebPage(url);
        }
    }

    public void display() {
        if (realWebPage != null) {
            realWebPage.display();
        }
    }
}
```

2. 現在有一個很大的圖像檔案需要讀到記憶體中並且開始編輯，但是讀取時需要較長的時間，我們希望在讀取圖像時加入 Proxy Pattern，當使用者輸入指令編輯圖像時，再將圖像檔案從 Proxy 中取出並開始編輯，而不是在載入圖像時就開始編輯。請實現此 Proxy。

答案：

```java
public interface Image {
    void display();
}

public class RealImage implements Image {
    private String filename;

    public RealImage(String filename) {
        this.filename = filename;
        loadFromDisk();
    }

    private void loadFromDisk() {
        System.out.println("Loading " + filename);
    }

    public void display() {
        System.out.println("Displaying " + filename);
    }
}

public class ProxyImage implements Image {
    private RealImage image;
    private String filename;

    public ProxyImage(String filename) {
        this.filename = filename;
    }

    public void display() {
        if (image == null) {
            image = new RealImage(filename);
        }
        image.display();
    }
}
```

3. 在一個線上音樂串流服務中，當使用者點擊一首歌曲來收聽時，我們希望可以透過 Proxy Pattern 在呼叫播放器之前先檢查使用者是否有收費帳號。如果是付費用戶，則可以收聽音樂；否則，將回傳錯誤提示。請完成此 Proxy。

答案：

```java
public interface MusicPlayer {
    void playMusic();
}

public class RealMusicPlayer implements MusicPlayer {
    private String song;

    public RealMusicPlayer(String song) {
        this.song = song;
        loadMusic();
    }

    private void loadMusic() {
        System.out.println("Loading " + song);
    }

    public void playMusic() {
        System.out.println("Playing " + song);
    }
}

public class MusicPlayerProxy implements MusicPlayer {
    private String username;
    private String password;
    private RealMusicPlayer realMusicPlayer;

    public MusicPlayerProxy(String username, String password, String song) {
        this.username = username;
        this.password = password;
        // 檢查使用者帳號是否有付費。如果有付費，則需要建立 RealMusicPlayer 物件，否則將不會建立
        if (isPaidUser(username, password)) {
            realMusicPlayer = new RealMusicPlayer(song);
        } else {
            System.out.println("Access Denied");
        }
    }

    public void playMusic() {
        if (realMusicPlayer != null) {
            realMusicPlayer.playMusic();
        }
    }

    private boolean isPaidUser(String username, String password) {
        // 檢查使用者是否有付費帳號
        return true;
    }
}
```

4. 現在有一個較大的資料庫需要建立快照，但是一次將所有資料寫入儲存機制可能會花費太長的時間。我們希望實作一個 Proxy 管理物件的存取，當有物件被寫入時，不要馬上紀錄到資料庫中，而是暫存，當暫存滿一定數量時再將所有物件寫入資料庫，以節省時間。請完成此 Proxy。

答案：

```java
public interface Database {
    void save(String data);
}

public class RealDatabase implements Database {
    public void save(String data) {
        System.out.println("Saving " + data);
        // 將 data 寫入到資料庫中
    }
}

public class DatabaseProxy implements Database {
    private RealDatabase realDatabase;
    private List<String> buffer = new ArrayList<String>();
    private int bufferSize = 10;

    public void save(String data) {
        if (realDatabase == null) {
            realDatabase = new RealDatabase();
        }
        buffer.add(data);
        if (buffer.size() >= bufferSize) {
            for (String s : buffer) {
                realDatabase.save(s);
            }
            buffer.clear();
        }
    }
}
```

5. 試著寫出一個簡單的 Messenger，當使用者輸入想要聊天的對象時，系統會建立一個新的對話框，而不是在原本的視窗中顯示。為了避免建立太多的對話框，我們希望用 Proxy Pattern 管理已經建立的對話框。請完成此 Proxy。

答案：

```java
public interface MessengerWindow {
    void display();
}

public class RealMessengerWindow implements MessengerWindow {
    private String username;

    public RealMessengerWindow(String username) {
        this.username = username;
        openWindow();
    }

    private void openWindow() {
        System.out.println("Opening chat window for " + username);
    }

    public void display() {
        System.out.println("Chatting with " + username);
    }
}

public class MessengerWindowProxy implements MessengerWindow {
    private static Map<String, RealMessengerWindow> chatWindows = new HashMap<String, RealMessengerWindow>();
    private String username;

    public MessengerWindowProxy(String username) {
        this.username = username;
    }

    public void display() {
        if (!chatWindows.containsKey(username)) {
            chatWindows.put(username, new RealMessengerWindow(username));
        }
        chatWindows.get(username).display();
    }
}
```   

