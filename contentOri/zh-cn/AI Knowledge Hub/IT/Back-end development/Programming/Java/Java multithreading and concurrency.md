+++
title = "Java多线程和并发"
weight = "3"
+++
## 习题预习
{{< ask_chatgpt >}}
给我5题Java multithreading and concurrency的问题
{{< /ask_chatgpt >}}



1. 什麼是Multithreading? 請解釋為什麼Java中的Multithreading對於編程是很重要的。

2. 請簡單說明Java中的Synchronized關鍵字是什麼？ 它是如何防止多個線程同時訪問同一個資源的？

3. 理解死鎖是什麼，以及在Java中避免死鎖的方法。

4. 如何使用Java中的ThreadLocal類？ 它有什麼用途？

5. 解釋Java中的BlockingQueue是什麼？ 它是如何在多線程應用中使用的？   

## 说明知识
{{< ask_chatgpt >}}
说明Java multithreading and concurrency并举例
{{< /ask_chatgpt >}}



Java multithreading 是指在一个 Java 程序中同时执行多个线程的技术。它可以提高程序的性能和效率，特别是对于需要同时处理多个任务的程序来说，非常重要。

Java concurrency 是指在多个线程之间可以共享资源的技术。当多个线程同时获取和修改同一个资源时，就需要使用 Java concurrency 技术来确保资源的安全性和一致性。

以下是 Java multithreading 和 concurrency 的一些例子：

1. 多个线程同时执行不同的任务。例如，一个 Web 服务器可以同时处理多个请求，并将请求分配到不同的线程中处理，提高处理效率和速度。

2. 多个线程同时更新同一个资料库。当多个线程同时读取和修改同一个资料库时，就需要使用 Java concurrency 技术来确保资源的安全性和一致性。

3. 多个线程同时执行相同的任务，但使用不同的资源。例如，一个文件塞满了大量的单词，我们可以使用多个线程同时计算每个单词的出现次数，进而提高运算效率。

总体来说，Java multithreading 和 concurrency 技术可以大大提高程序的效率和效果，使程序更加高效和稳定。   

## 汇总重点
{{< ask_chatgpt >}}
条列Java multithreading and concurrency的重点
{{< /ask_chatgpt >}}



1. Multithreading是指一个程序可以同时运行多个线程，每个线程都可以执行不同的任务。

2. 在Java中，创建一个线程有两种方式：Runnable接口和Thread类。

3. 使用synchronized关键字可以实现线程同步，避免多线程同时访问共享变量的问题。

4. 在Java中可以使用Lock和Condition类实现比synchronized更细粒度的线程同步。

5. 使用volatile关键字可以保证多线程之间共享变量的可见性。

6. 在Java中可以使用Atomic类进行原子性操作，避免多线程操作共享变量时出现的竞争和冲突。

7. 使用ThreadLocal类可以实现线程本地变量，每个线程都有自己的变量副本，避免多线程之间的访问冲突。

8. 使用Executor框架可以实现线程池管理多线程，提高多线程的执行效率和性能。

9. 在Java中可以使用ForkJoin框架实现分治算法，将大问题分为若干个小问题，并行求解提高效率。

10. 在Java 8中，引入了Stream API，可以使用多线程来实现对集合的操作，提高程序的执行效率。   

## 知识测验
{{< ask_chatgpt >}}
给我5题Java multithreading and concurrency的中等难度问题，并在后面列出答案
{{< /ask_chatgpt >}}



1. 编写一个具有3个线程的程序，实现有人去餐厅就餐。如果有一张桌子为空，顾客可以入座就餐，否则需要排队等待。餐厅只有1个服务员，可以为并发的2张桌子提供服务。

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

2. 编写一个程序，使用Callable和Future实现从文件中查找指定字符串。程序应该支持多线程并发查找。

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

3. 编写一个程序，使用Lock和Condition实现无界阻塞队列。该队列应支持put()和take()操作，并严格按照先进先出的顺序返回元素。

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

4. 编写一个程序，使用ReadWriteLock实现一个缓存器。缓存器应支持put()和get()操作，put()操作应该写入一个key/value对，get()操作应该返回一个key所对应的value。

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

5. 编写一个程序，使用Semaphore和Barrier实现一个流水线，有多个工人在流水线上工作。流水线分为3个阶段，每个阶段只能由一个工人完成，一个工人完成当前阶段后才能进入下一阶段，所有工人完成第一阶段后，才能继续进行第二阶段。所有工人完成第二阶段后，才能继续进行第三阶段。

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

