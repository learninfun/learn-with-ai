+++
title = "策略模式"
weight = "9"
+++
## 习题预习
{{< ask_chatgpt >}}
给我5题Strategy Pattern的中文问题
{{< /ask_chatgpt >}}



1. 策略模式是什么？
2. 策略模式与其他设计模式有什么区别？
3. 策略模式的优势是什么？
4. 策略模式如何应用到实际开发中？
5. 如何在实现策略模式时确保代码的维护性和可扩展性？   

## 说明知识
{{< ask_chatgpt >}}
以中文说明Strategy Pattern并举例
{{< /ask_chatgpt >}}



策略模式（Strategy Pattern）是一种行为型模式，用以定义不同的算法和演算法，并在需要时 dynamically interchange during runtime for different scenarios。

使用策略模式的目的是将相关的算法和演算法封装成不同的策略，让程式能够更加灵活和易于维护。此外，策略模式遵循开放封闭原则（Open-Closed Principle, OCP），允许新增新的策略而不会修改原有的程式码。

以下为举例：

例如，假设我们有一个以动物为主题的游戏，里面有不同类型的动物，如狗、猫、鸟等，每种动物都有自己的攻击技能。当我们在游戏中使用不同的动物时，会需要对应不同的攻击技能。

为此，我们可以使用策略模式。我们可以定义一个名为 AttackStrategy 的介面，并在其中定义攻击行为。接着，对每个动物实现一个不同的 AttackStrategy 策略，以封装不同的攻击技能。当使用不同的动物时，我们只需要动态选择对应的策略即可。

AttackStrategy.java

```
public interface AttackStrategy {
   public void attack();
}
```

DogAttackStrategy.java 做了咬的动作

```
public class DogAttackStrategy implements AttackStrategy {
   @Override
   public void attack() {
      System.out.println("Dog attacks by biting!");
   }
}
```

CatAttackStrategy.java 做了抓的动作

```
public class CatAttackStrategy implements AttackStrategy {
   @Override
   public void attack() {
      System.out.println("Cat attacks by scratching!");
   }
}
```

BirdAttackStrategy.java 做了咬和飞的动作

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

## 条列重点
{{< ask_chatgpt >}}
以中文条列Strategy Pattern的重点
{{< /ask_chatgpt >}}



1. 定义了一个族群的相关演算法，并让它们可以互相替换，使得程式更具有弹性和扩展性。 
2. 当程式有多种可行的演算法时，使用者可以透过策略模式动态地选择所需的演算法，而不必修改原始码。 
3. 策略模式将算法具体实现与使用分离，使得算法可以独立于使用者而维护，从而提高了程式的可读性和可维护性。
4. 在策略模式中，每个演算法都是一个独立的类别，它们之间可以替换、扩充。 
5. 策略模式可以应用于各种场景，例如排序、搜寻、运算、加密等任意场景。   

## 知识测验
{{< ask_chatgpt >}}
以中文给我5题Strategy Pattern的中等难度问题，并在后面列出答案
{{< /ask_chatgpt >}}



1. 问题：在一个游戏中，不同种类的角色有不同的攻击方式，请使用Strategy Pattern设计一个攻击接口，然后让各种角色使用不同的攻击策略。

```java
public interface AttackStrategy {
    void attack();
}

public class SwordAttackStrategy implements AttackStrategy {
    @Override
    public void attack() {
        System.out.println("使用剑进行攻击！");
    }
}

public class MagicAttackStrategy implements AttackStrategy {
    @Override
    public void attack() {
        System.out.println("使用魔法攻击！");
    }
}

public class ArcherAttackStrategy implements AttackStrategy {
    @Override
    public void attack() {
        System.out.println("使用弓箭进行攻击！");
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

2. 问题：你正在开发一个搜寻引擎，使用者可以根据不同的搜寻条件来查找结果，例如按相关度排序、按价格排序等等。请使用Strategy Pattern设计一个搜寻接口，然后让使用者可以根据自己的喜好自由选择排序方式。

```java
public interface SearchStrategy {
    void search(String keyword);
}

public class RelevanceSearchStrategy implements SearchStrategy {
    @Override
    public void search(String keyword) {
        System.out.println("按相关度排序，搜寻结果为：" + keyword);
    }
}

public class PriceSearchStrategy implements SearchStrategy {
    @Override
    public void search(String keyword) {
        System.out.println("按价格排序，搜寻结果为：" + keyword);
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

3. 问题：你正在开发一个网路游戏，每个玩家加入后都会被分配到一个不同的房间，每个房间的玩法和配置都不同。请使用Strategy Pattern设计一个房间接口，然后让不同的玩家加入不同的房间享受不同的游戏体验。

```java
public interface RoomStrategy {
    void play();
}

public class NormalRoomStrategy implements RoomStrategy {
    @Override
    public void play() {
        System.out.println("普通房间：双方PK，攻防双方各有优势");
    }
}

public class BeginnerRoomStrategy implements RoomStrategy {
    @Override
    public void play() {
        System.out.println("新手房间：双方PK，攻方攻击力降低，防方防御力提高");
    }
}

public class VIPRoomStrategy implements RoomStrategy {
    @Override
    public void play() {
        System.out.println("VIP房间：双方PK，攻方攻击力提高，防方防御力提高");
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

4. 问题：你正在设计一个图像处理软件，使用者可以选择不同的图像处理策略来处理他们的图像，例如灰度化、缩放、旋转等等。请使用Strategy Pattern设计一个图像处理接口，然后让使用者根据需要选择不同的处理策略。

```java
public interface ImageProcessStrategy {
    void process(Image image);
}

public class GrayScaleImageProcessStrategy implements ImageProcessStrategy {
    @Override
    public void process(Image image) {
        System.out.println("灰度处理图像：" + image);
    }
}

public class ScaleImageProcessStrategy implements ImageProcessStrategy {
    @Override
    public void process(Image image) {
        System.out.println("缩放图像：" + image);
    }
}

public class RotateImageProcessStrategy implements ImageProcessStrategy {
    @Override
    public void process(Image image) {
        System.out.println("旋转图像：" + image);
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

5. 问题：你正在设计一个音乐软件，使用者可以选择不同的音效处理策略来聆听歌曲，例如增强低音、增强高音等等。请使用Strategy Pattern设计一个音效处理接口，然后让使用者根据需要选择不同的处理策略。

```java
public interface MusicEffectStrategy {
    void apply(Music music);
}

public class BassBoostMusicEffectStrategy implements MusicEffectStrategy {
    @Override
    public void apply(Music music) {
        System.out.println("增强低音效果：" + music);
    }
}

public class TrebleBoostMusicEffectStrategy implements MusicEffectStrategy {
    @Override
    public void apply(Music music) {
        System.out.println("增强高音效果：" + music);
    }
}

public class SurroundMusicEffectStrategy implements MusicEffectStrategy {
    @Override
    public void apply(Music music) {
        System.out.println("环绕声效果：" + music);
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

