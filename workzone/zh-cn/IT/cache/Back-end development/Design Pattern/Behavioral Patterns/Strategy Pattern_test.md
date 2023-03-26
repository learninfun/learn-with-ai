

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