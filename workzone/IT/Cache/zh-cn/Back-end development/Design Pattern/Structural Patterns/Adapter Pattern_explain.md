

適配器模式是一種常見的設計模式，它通常是用來把一個類的接口轉換成客戶端所期待的另一種接口的模式。這意味著你可以使用原本不相容的類來一起工作，並且無需對現有的程式碼進行修改或重構。

舉例來說，在某個程式裡面有一個現成的組件可以呈現圓形，但你需要的是呈現正方形。這時，你可以使用適配器模式去轉換這兩種不相容的接口，並且達到你需要的結果。

以下是一個簡單的範例：

假設有一個工具箱，裡面有兩個工具：一把圓鉗和一把方鉗。現在，我們需要從這些鉗子當中選擇一把，並用它來剪開一個方形形狀的材料。

但問題是：圓鉗只適合剪開圓形形狀的材料，而方鉗只適合剪開方形形狀的材料。那該怎麼辦呢？我們可以使用適配器模式來轉換這兩種不相容的接口，並且讓它們一起工作。

我們可以實現一個適配器，它可以接收圓鉗的輸入，並將其轉換為正方形，最後再使用方鉗來剪開它。這樣，我們就可以用適配器來完成這個任務，而不必去重寫或重構原本的程式碼。

以下是示例程式碼：

```
// 鉗子介面
interface ITool {
    void cut();
}

// 圓形鉗子
class CircleTool implements ITool {
    @Override
    public void cut() {
        System.out.println("剪開圓形的材料");
    }
}

// 方形鉗子
class SquareTool implements ITool {
    @Override
    public void cut() {
        System.out.println("剪開方形的材料");
    }
}

// 適配器
class Adapter implements ITool {
    private CircleTool circleTool;

    public Adapter(CircleTool circleTool) {
        this.circleTool = circleTool;
    }

    @Override
    public void cut() {
        System.out.println("使用適配器進行轉換");
        System.out.println("將圓形材料轉換為方形材料");
        new SquareTool().cut();
    }
}

// 客户端
public class Client {
    public static void main(String[] args) {
        CircleTool circleTool = new CircleTool();
        ITool adapter = new Adapter(circleTool);
        adapter.cut();
    }
}
```

上述程式碼中，我們設計了一個 ITool 介面，裡面包含一個剪開的方法。並且定義了兩種鉗子：圓形鉗子和方形鉗子，分別實現了 ITool 介面的方法。

接著，我們又新增了一個 Adapter 類，它實現了 ITool 介面，但卻是通過圓形鉗子去轉換成方形鉗子。最後，在客戶端中，我們只需實例化 Adapter 類，就可以通過圓形鉗子去剪開方形形狀的材料了。

總結：適配器模式讓原本不相容的類可以共同協作，提高了程式的靈活性和可擴展性，使得系統的修改和維護變得更加方便。