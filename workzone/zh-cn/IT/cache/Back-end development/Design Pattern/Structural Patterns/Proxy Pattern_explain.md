

Proxy Pattern 是一种静态设计模式，它可以将物件的存取权限控制在一个代理物件中，代理物件可以在真实的物件之前或之后执行一些额外的功能，例如远端代理、虚拟代理、安全代理等。通过使用代理模式，客户端可以透明地访问一个物件，而不需要知道它的实际实现细节。

举例来说，假设有一个 Image 接口，它有一个 display() 方法用来显示图片，一个 RealImage 类用来实现 Image 接口，另外还有一个 ProxyImage 类也实现了 Image 接口，它在显示图片之前先检查是否有权限。

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

在上面的例子中，当使用 ProxyImage 类显示图片时，它首先会检查是否有权限，如果没有权限，就不会显示图片；如果有权限，就会调用 RealImage 的 display() 方法来显示图片。这让客户端可以透明地访问图片，而不需要知道代理物件是否存在。