

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