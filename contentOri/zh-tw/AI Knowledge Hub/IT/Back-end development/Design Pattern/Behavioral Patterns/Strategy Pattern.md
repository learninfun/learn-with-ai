+++
title = "策略模式"
weight = "9"
+++
## 習題預習
{{< ask_chatgpt >}}
給我5題Strategy Pattern的中文問題
{{< /ask_chatgpt >}}



1. 策略模式是什麼？
2. 策略模式與其他設計模式有什麼區別？
3. 策略模式的優勢是什麼？
4. 策略模式如何應用到實際開發中？
5. 如何在實現策略模式時確保代碼的維護性和可擴展性？   

## 說明知識
{{< ask_chatgpt >}}
以中文說明Strategy Pattern並舉例
{{< /ask_chatgpt >}}



策略模式（Strategy Pattern）是一種行為型模式，用以定義不同的算法和演算法，並在需要時 dynamically interchange during runtime for different scenarios。

使用策略模式的目的是將相關的算法和演算法封裝成不同的策略，讓程式能夠更加靈活和易於維護。此外，策略模式遵循開放封閉原則（Open-Closed Principle, OCP），允許新增新的策略而不會修改原有的程式碼。

以下為舉例：

例如，假設我們有一個以動物為主題的遊戲，裡面有不同類型的動物，如狗、貓、鳥等，每種動物都有自己的攻擊技能。當我們在遊戲中使用不同的動物時，會需要對應不同的攻擊技能。

為此，我們可以使用策略模式。我們可以定義一個名為 AttackStrategy 的介面，並在其中定義攻擊行為。接著，對每個動物實現一個不同的 AttackStrategy 策略，以封裝不同的攻擊技能。當使用不同的動物時，我們只需要動態選擇對應的策略即可。

AttackStrategy.java

```
public interface AttackStrategy {
   public void attack();
}
```

DogAttackStrategy.java 做了咬的動作

```
public class DogAttackStrategy implements AttackStrategy {
   @Override
   public void attack() {
      System.out.println("Dog attacks by biting!");
   }
}
```

CatAttackStrategy.java 做了抓的動作

```
public class CatAttackStrategy implements AttackStrategy {
   @Override
   public void attack() {
      System.out.println("Cat attacks by scratching!");
   }
}
```

BirdAttackStrategy.java 做了咬和飛的動作

```
public class BirdAttackStrategy implements AttackStrategy {
   @Override
   public void attack() {
      System.out.println("Bird attacks by biting and flying!");
   }
}
```

Animal.java

```
public class Animal {
   private String name;
   private AttackStrategy attackStrategy;

   public Animal(String name, AttackStrategy attackStrategy) {
      this.name = name;
      this.attackStrategy = attackStrategy;
   }

   public void attack() {
      System.out.println(name + " attacks.");
      attackStrategy.attack();
   }
}
```

Client.java

```
public class Client {
   public static void main(String[] args) {
      AttackStrategy dogStrategy = new DogAttackStrategy();
      AttackStrategy catStrategy = new CatAttackStrategy();
      AttackStrategy birdStrategy = new BirdAttackStrategy();

      Animal dog = new Animal("Dog", dogStrategy);
      Animal cat = new Animal("Cat", catStrategy);
      Animal bird = new Animal("Bird", birdStrategy);

      dog.attack();    // Dog attacks by biting!
      cat.attack();    // Cat attacks by scratching!
      bird.attack();   // Bird attacks by biting and flying!
   }
}
```   

## 彙總重點
{{< ask_chatgpt >}}
以中文條列Strategy Pattern的重點
{{< /ask_chatgpt >}}



1. 定義了一個族群的相關演算法，並讓它們可以互相替換，使得程式更具有彈性和擴展性。 
2. 當程式有多種可行的演算法時，使用者可以透過策略模式動態地選擇所需的演算法，而不必修改原始碼。 
3. 策略模式將算法具體實現與使用分離，使得算法可以獨立於使用者而維護，從而提高了程式的可讀性和可維護性。
4. 在策略模式中，每個演算法都是一個獨立的類別，它們之間可以替換、擴充。 
5. 策略模式可以應用於各種場景，例如排序、搜尋、運算、加密等任意場景。   

## 知識測驗
{{< ask_chatgpt >}}
以中文給我5題Strategy Pattern的中等難度問題，並在後面列出答案
{{< /ask_chatgpt >}}



1. 問題：在一個遊戲中，不同種類的角色有不同的攻擊方式，請使用Strategy Pattern設計一個攻擊接口，然後讓各種角色使用不同的攻擊策略。

```java
public interface AttackStrategy {
    void attack();
}

public class SwordAttackStrategy implements AttackStrategy {
    @Override
    public void attack() {
        System.out.println("使用劍進行攻擊！");
    }
}

public class MagicAttackStrategy implements AttackStrategy {
    @Override
    public void attack() {
        System.out.println("使用魔法攻擊！");
    }
}

public class ArcherAttackStrategy implements AttackStrategy {
    @Override
    public void attack() {
        System.out.println("使用弓箭進行攻擊！");
    }
}

public class Character {
    private AttackStrategy attackStrategy;

    public void setAttackStrategy(AttackStrategy attackStrategy) {
        this.attackStrategy = attackStrategy;
    }

    public void attack() {
        attackStrategy.attack();
    }
}

public class Main {
    public static void main(String[] args) {
        Character knight = new Character();
        knight.setAttackStrategy(new SwordAttackStrategy());
        knight.attack();

        Character mage = new Character();
        mage.setAttackStrategy(new MagicAttackStrategy());
        mage.attack();

        Character archer = new Character();
        archer.setAttackStrategy(new ArcherAttackStrategy());
        archer.attack();
    }
}
```

2. 問題：你正在開發一個搜尋引擎，使用者可以根據不同的搜尋條件來查找結果，例如按相關度排序、按價格排序等等。請使用Strategy Pattern設計一個搜尋接口，然後讓使用者可以根據自己的喜好自由選擇排序方式。

```java
public interface SearchStrategy {
    void search(String keyword);
}

public class RelevanceSearchStrategy implements SearchStrategy {
    @Override
    public void search(String keyword) {
        System.out.println("按相關度排序，搜尋結果為：" + keyword);
    }
}

public class PriceSearchStrategy implements SearchStrategy {
    @Override
    public void search(String keyword) {
        System.out.println("按價格排序，搜尋結果為：" + keyword);
    }
}

public class SearchEngine {
    private SearchStrategy searchStrategy;

    public void setSearchStrategy(SearchStrategy searchStrategy) {
        this.searchStrategy = searchStrategy;
    }

    public void search(String keyword) {
        searchStrategy.search(keyword);
    }
}

public class Main {
    public static void main(String[] args) {
        SearchEngine engine = new SearchEngine();

        engine.setSearchStrategy(new RelevanceSearchStrategy());
        engine.search("Java");

        engine.setSearchStrategy(new PriceSearchStrategy());
        engine.search("MacBook");
    }
}
```

3. 問題：你正在開發一個網路遊戲，每個玩家加入後都會被分配到一個不同的房間，每個房間的玩法和配置都不同。請使用Strategy Pattern設計一個房間接口，然後讓不同的玩家加入不同的房間享受不同的遊戲體驗。

```java
public interface RoomStrategy {
    void play();
}

public class NormalRoomStrategy implements RoomStrategy {
    @Override
    public void play() {
        System.out.println("普通房間：雙方PK，攻防雙方各有優勢");
    }
}

public class BeginnerRoomStrategy implements RoomStrategy {
    @Override
    public void play() {
        System.out.println("新手房間：雙方PK，攻方攻擊力降低，防方防禦力提高");
    }
}

public class VIPRoomStrategy implements RoomStrategy {
    @Override
    public void play() {
        System.out.println("VIP房間：雙方PK，攻方攻擊力提高，防方防禦力提高");
    }
}

public class Player {
    private RoomStrategy roomStrategy;

    public Player(RoomStrategy roomStrategy) {
        this.roomStrategy = roomStrategy;
    }

    public void play() {
        roomStrategy.play();
    }
}

public class Main {
    public static void main(String[] args) {
        Player player1 = new Player(new NormalRoomStrategy());
        player1.play();

        Player player2 = new Player(new BeginnerRoomStrategy());
        player2.play();

        Player player3 = new Player(new VIPRoomStrategy());
        player3.play();
    }
}
```

4. 問題：你正在設計一個圖像處理軟件，使用者可以選擇不同的圖像處理策略來處理他們的圖像，例如灰度化、縮放、旋轉等等。請使用Strategy Pattern設計一個圖像處理接口，然後讓使用者根據需要選擇不同的處理策略。

```java
public interface ImageProcessStrategy {
    void process(Image image);
}

public class GrayScaleImageProcessStrategy implements ImageProcessStrategy {
    @Override
    public void process(Image image) {
        System.out.println("灰度處理圖像：" + image);
    }
}

public class ScaleImageProcessStrategy implements ImageProcessStrategy {
    @Override
    public void process(Image image) {
        System.out.println("縮放圖像：" + image);
    }
}

public class RotateImageProcessStrategy implements ImageProcessStrategy {
    @Override
    public void process(Image image) {
        System.out.println("旋轉圖像：" + image);
    }
}

public class ImageProcessor {
    private ImageProcessStrategy imageProcessStrategy;

    public void setImageProcessStrategy(ImageProcessStrategy imageProcessStrategy) {
        this.imageProcessStrategy = imageProcessStrategy;
    }

    public void process(Image image) {
        imageProcessStrategy.process(image);
    }
}

public class Main {
    public static void main(String[] args) {
        Image image = new Image("image.jpg");

        ImageProcessor processor = new ImageProcessor();

        processor.setImageProcessStrategy(new GrayScaleImageProcessStrategy());
        processor.process(image);

        processor.setImageProcessStrategy(new ScaleImageProcessStrategy());
        processor.process(image);

        processor.setImageProcessStrategy(new RotateImageProcessStrategy());
        processor.process(image);
    }
}
```

5. 問題：你正在設計一個音樂軟件，使用者可以選擇不同的音效處理策略來聆聽歌曲，例如增強低音、增強高音等等。請使用Strategy Pattern設計一個音效處理接口，然後讓使用者根據需要選擇不同的處理策略。

```java
public interface MusicEffectStrategy {
    void apply(Music music);
}

public class BassBoostMusicEffectStrategy implements MusicEffectStrategy {
    @Override
    public void apply(Music music) {
        System.out.println("增強低音效果：" + music);
    }
}

public class TrebleBoostMusicEffectStrategy implements MusicEffectStrategy {
    @Override
    public void apply(Music music) {
        System.out.println("增強高音效果：" + music);
    }
}

public class SurroundMusicEffectStrategy implements MusicEffectStrategy {
    @Override
    public void apply(Music music) {
        System.out.println("環繞聲效果：" + music);
    }
}

public class MusicPlayer {
    private MusicEffectStrategy musicEffectStrategy;

    public void setMusicEffectStrategy(MusicEffectStrategy musicEffectStrategy) {
        this.musicEffectStrategy = musicEffectStrategy;
    }

    public void play(Music music) {
        musicEffectStrategy.apply(music);
    }
}

public class Main {
    public static void main(String[] args) {
        Music music = new Music("song.mp3");

        MusicPlayer player = new MusicPlayer();

        player.setMusicEffectStrategy(new BassBoostMusicEffectStrategy());
        player.play(music);

        player.setMusicEffectStrategy(new TrebleBoostMusicEffectStrategy());
        player.play(music);

        player.setMusicEffectStrategy(new SurroundMusicEffectStrategy());
        player.play(music);
    }
}
```   

