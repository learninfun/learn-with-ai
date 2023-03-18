

1. 編寫一個具有3個線程的程序，實現有人去餐廳就餐。如果有一張桌子為空，顧客可以入座就餐，否則需要排隊等待。餐廳只有1個服務員，可以為並發的2張桌子提供服務。

Answer:

```
// Dining room class
public class DiningRoom {
    private static final int NUM_TABLES = 3;

    private final Semaphore waiter = new Semaphore(2); // only two tables can be served by the waiter at once
    private final Semaphore tables = new Semaphore(NUM_TABLES); // three tables available in the restaurant

    public void enterRestaurant(String name) throws InterruptedException {
        tables.acquire();
        System.out.println(name + " entered the restaurant.");
    }

    public void orderFood(String name) throws InterruptedException {
        waiter.acquire(); // wait for the waiter
        System.out.println(name + " is ordering food.");
        Thread.sleep(300); // simulate order process
    }

    public void eatFood(String name) throws InterruptedException {
        System.out.println(name + " is eating food.");
        Thread.sleep(1000); // simulate eating process
    }

    public void leaveRestaurant(String name) {
        waiter.release(); // free the waiter
        tables.release(); // free the table
        System.out.println(name + " left the restaurant.");
    }
}

// Customer class
public class Customer implements Runnable {
    private final String name;
    private final DiningRoom diningRoom;

    public Customer(String name, DiningRoom diningRoom) {
        this.name = name;
        this.diningRoom = diningRoom;
    }

    @Override
    public void run() {
        try {
            diningRoom.enterRestaurant(name);
            diningRoom.orderFood(name);
            diningRoom.eatFood(name);
            diningRoom.leaveRestaurant(name);
        } catch (InterruptedException ex) {
            ex.printStackTrace();
        }
    }
}

// Main class
public class Main {
    public static void main(String[] args) {
        DiningRoom diningRoom = new DiningRoom();

        Thread customer1 = new Thread(new Customer("Alice", diningRoom));
        Thread customer2 = new Thread(new Customer("Bob", diningRoom));
        Thread customer3 = new Thread(new Customer("Charlie", diningRoom));

        customer1.start();
        customer2.start();
        customer3.start();
    }
}
```

2. 編寫一個程序，使用Callable和Future實現從文件中查找指定字符串。程序應該支持多線程並發查找。

Answer:

```
public class FileSearcher {
    private final ExecutorService executorService;

    public FileSearcher(int numThreads) {
        executorService = Executors.newFixedThreadPool(numThreads);
    }

    public Future<Set<String>> search(Path directoryPath, String searchQuery) {
        return executorService.submit(() -> {
            Set<String> result = new HashSet<>();

            Files.walk(directoryPath)
                .filter(Files::isRegularFile)
                .parallel()
                .forEach(path -> {
                    try {
                        String content = new String(Files.readAllBytes(path));

                        if (content.contains(searchQuery)) {
                            result.add(path.toAbsolutePath().toString());
                        }
                    } catch (IOException ex) {
                        ex.printStackTrace();
                    }
                });

            return result;
        });
    }
}

// Main class
public class Main {
    public static void main(String[] args) throws Exception {
        FileSearcher fileSearcher = new FileSearcher(4);
        Path directoryPath = Paths.get("src");
        String searchQuery = "Lorem ipsum";

        System.out.println("Searching for \"" + searchQuery + "\" in directory: " + directoryPath);

        Future<Set<String>> resultFuture = fileSearcher.search(directoryPath, searchQuery);

        Set<String> result = resultFuture.get();

        System.out.println("Search results:");
        for (String filePath : result) {
            System.out.println(filePath);
        }

        fileSearcher.shutdown();
    }
}
```

3. 編寫一個程序，使用Lock和Condition實現無界阻塞隊列。該隊列應支持put()和take()操作，並嚴格按照先進先出的順序返回元素。

Answer:

```
public class UnboundedBlockingQueue<T> {
    private final Queue<T> queue = new LinkedList<>();
    private final Lock lock = new ReentrantLock();
    private final Condition notEmpty = lock.newCondition();

    public void put(T item) throws InterruptedException {
        lock.lockInterruptibly();
        try {
            queue.add(item);
            notEmpty.signal();
        } finally {
            lock.unlock();
        }
    }

    public T take() throws InterruptedException {
        lock.lockInterruptibly();
        try {
            while (queue.isEmpty()) {
                notEmpty.await();
            }

            return queue.poll();
        } finally {
            lock.unlock();
        }
    }
}

// Main class
public class Main {
    public static void main(String[] args) {
        UnboundedBlockingQueue<Integer> queue = new UnboundedBlockingQueue<>();

        // Producer thread
        new Thread(() -> {
            for (int i = 1; i <= 10; i++) {
                try {
                    queue.put(i);
                    System.out.println("Produced " + i + ".");
                    Thread.sleep(500);
                } catch (InterruptedException ex) {
                    ex.printStackTrace();
                }
            }
        }).start();

        // Consumer thread
        new Thread(() -> {
            while (true) {
                try {
                    int item = queue.take();
                    System.out.println("Consumed " + item + ".");
                    Thread.sleep(1000);
                } catch (InterruptedException ex) {
                    ex.printStackTrace();
                }
            }
        }).start();
    }
}
```

4. 編寫一個程序，使用ReadWriteLock實現一個緩存器。緩存器應支持put()和get()操作，put()操作應該寫入一個key/value對，get()操作應該返回一個key所對應的value。

Answer:

```
public class Cache<K, V> {
    private final Map<K, V> cache = new HashMap<>();
    private final ReadWriteLock lock = new ReentrantReadWriteLock();

    public void put(K key, V value) {
        lock.writeLock().lock();
        try {
            cache.put(key, value);
        } finally {
            lock.writeLock().unlock();
        }
    }

    public V get(K key) {
        lock.readLock().lock();
        try {
            return cache.get(key);
        } finally {
            lock.readLock().unlock();
        }
    }
}

// Main class
public class Main {
    public static void main(String[] args) {
        Cache<String, Integer> cache = new Cache<>();

        // Producer thread
        new Thread(() -> {
            for (int i = 1; i <= 10; i++) {
                cache.put("key" + i, i);
                System.out.println("Put key" + i + " with value: " + i + ".");
                try {
                    Thread.sleep(500);
                } catch (InterruptedException ex) {
                    ex.printStackTrace();
                }
            }
        }).start();

        // Consumer thread
        new Thread(() -> {
            for (int i = 1; i <= 10; i++) {
                Integer value = cache.get("key" + i);
                System.out.println("Got key" + i + " with value: " + value + ".");
                try {
                    Thread.sleep(500);
                } catch (InterruptedException ex) {
                    ex.printStackTrace();
                }
            }
        }).start();
    }
}
```

5. 編寫一個程序，使用Semaphore和Barrier實現一個流水線，有多個工人在流水線上工作。流水線分為3個階段，每個階段只能由一個工人完成，一個工人完成當前階段後才能進入下一階段，所有工人完成第一階段後，才能繼續進行第二階段。所有工人完成第二階段後，才能繼續進行第三階段。

Answer:

```
public class Pipeline {
    private static final int NUM_WORKERS = 4;

    private final Semaphore[] semaphores;
    private final CyclicBarrier barrier;

    public Pipeline() {
        semaphores = new Semaphore[NUM_WORKERS];
        for (int i = 0; i < NUM_WORKERS; i++) {
            semaphores[i] = new Semaphore(1);
        }

        barrier = new CyclicBarrier(NUM_WORKERS);
    }

    public void start() {
        for (int i = 0; i < NUM_WORKERS; i++) {
            new Thread(new Worker(i, this)).start();
        }
    }

    public void doStage(int stageIndex, int workerIndex) throws InterruptedException {
        semaphores[workerIndex].acquire();

        System.out.println("Worker " + workerIndex + " is doing stage " + stageIndex + ".");
        Thread.sleep(500); // simulate work

        if (stageIndex == 2) {
            barrier.await(); // wait for all workers to complete stage 2
        }

        semaphores[(workerIndex + 1) % NUM_WORKERS].release();
    }
}

// Worker class
public class Worker implements Runnable {
    private final int index;
    private final Pipeline pipeline;

    public Worker(int index, Pipeline pipeline) {
        this.index = index;
        this.pipeline = pipeline;
    }

    @Override
    public void run() {
        try {
            pipeline.doStage(1, index);
            pipeline.doStage(2, index);
            pipeline.doStage(3, index);
        } catch (InterruptedException | BrokenBarrierException ex) {
            ex.printStackTrace();
        }
    }
}

// Main class
public class Main {
    public static void main(String[] args) {
        Pipeline pipeline = new Pipeline();
        pipeline.start();
    }
}
```