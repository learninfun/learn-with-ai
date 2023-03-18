

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