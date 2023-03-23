+++
title = "結構型模式"
weight = "2"
+++
## 習題預習
{{< ask_chatgpt >}}
給我5題Structural Patterns的中文問題
{{< /ask_chatgpt >}}



1. 什麼是Structural Patterns？它們如何區分於其他設計模式？

2. 何時應該使用Adapter Pattern？它如何協助解決系統中的問題？

3. Decorator Pattern和Proxy Pattern的區別是什麼？這些模式如何用於提高代碼的可讀性和可維護性？

4. 在使用Facade Pattern時，如何設計接口以支持系統的簡化和壓縮？你可以舉例說明嗎？

5. Bridge Pattern如何協助開發人員更好地設計具有多層次抽像的軟件系統，並優化代碼的可維護性和擴展性？   

## 說明知識
{{< ask_chatgpt >}}
以中文說明Structural Patterns並舉例
{{< /ask_chatgpt >}}



Structural Patterns是指將類和對像組合在一起形成更大的結構，用於解決對象的組合，繼承等結構性問題。

以下是幾個常見的Structural Patterns：

1. Adapter Pattern：將一個類的接口轉換成客戶希望的另一個接口。例如，將老版本的類適配成新版本的類。

2. Bridge Pattern：將抽像部分與它的實現部分分離，使得它們可以獨立地變化。例如，將操作系統與應用程序分離，這樣就可以實現不同操作系統的應用程序了。

3. Composite Pattern：用於管理樹形結構的類，把相同的方法應用到每個元素上。例如，一個樹形菜單，每個菜單項都有相同的方法。

4. Decorator Pattern：動態地給一個對像添加一些額外的職責，同時又不改變其結構。例如，裝飾者模式可以增強一個類的功能，但不需要修改這個類本身。

5. Facade Pattern：為一組複雜的子系統提供一個簡單的接口，通過這個接口來訪問這個子系統。例如，一個文件系統的Facade可以讓用戶通過一個簡單的接口來訪問所有的文件。

6. Flyweight Pattern：運用共享技術有效地支持大量細粒度的對象。例如，一個網站可以使用Flyweight模式來減少服務器的壓力。

這些模式可以提高代碼的復用性，使代碼更加靈活和可擴展。   

## 彙總重點
{{< ask_chatgpt >}}
以中文條列Structural Patterns的重點
{{< /ask_chatgpt >}}



1.提供對系統內部結構的抽像表示

2.允許隨著時間的推移進行系統內部的變化，同時不影響系統外部的用戶界面

3.支持系統內部的外部化和內部化，使應用程序更加靈活

4.簡化系統內部結構的操作，提高系統的可讀性和可維護性

5.提供對系統中核心組件（例如，UI組件、數據組件等）的解耦，增強系統的可擴展性

6.提高系統的性能和效率，降低系統運行時所需的資源成本

7.提供對系統中不同部分的通信方式的抽像表示，使得系統之間的交互更加簡單和易於管理

8.在軟件設計過程中，提供一個共通的框架，方便開發人員進行協作、集成和測試   

## 知識測驗
{{< ask_chatgpt >}}
以中文給我5題Structural Patterns的中等難度問題，並在後面列出答案
{{< /ask_chatgpt >}}



1. 請實作一個簡單的裝飾器模式，可以讓使用者在商品價格上添加不同的折扣。

答案：

class Product {
  constructor(name, price) {
    this.name = name;
    this.price = price;
  }
}

function saleDiscount(product, percentage) {
  return {
    getPrice: function() {
      return product.price * (1 - percentage);
    },
    getName: function() {
      return product.name + " (" + percentage * 100 + "% off)";
    }
  };
}

// 使用範例
const product = new Product("iPhone", 1000);
const discountProduct = saleDiscount(product, 0.2); // 打80折

console.log(discountProduct.getName()); // iPhone (20% off)
console.log(discountProduct.getPrice()); // 800


2. 請實作一個中介者模式，讓使用者可以協調不同的元件之間的互動。

答案：

class Mediator {
  constructor() {
    this.components = [];
  }

  register(component) {
    this.components.push(component);
    component.setMediator(this);
  }

  send(message, sender) {
    for (let component of this.components) {
      if (component !== sender) {
        component.receive(message);
      }
    }
  }
}

class Component {
  constructor(name) {
    this.name = name;
    this.mediator = null;
  }

  setMediator(mediator) {
    this.mediator = mediator;
  }

  send(message) {
    this.mediator.send(message, this);
  }

  receive(message) {
    console.log(`${this.name} received message: ${message}`);
  }
}

// 使用範例
const mediator = new Mediator();
const component1 = new Component("Component 1");
const component2 = new Component("Component 2");

mediator.register(component1);
mediator.register(component2);

component1.send("Hello, Component 2");
component2.send("Hi, Component 1");


3. 請實作一個代理模式，讓使用者可以透過代理物件存取需要較長時間才能處理的資料。

答案：

class RealObject {
  process() {
    console.log("Processing...");
  }
}

class ProxyObject {
  constructor() {
    this.realObject = new RealObject();
    this.isProcessing = false;
  }

  process() {
    if (!this.isProcessing) {
      console.log("ProxyObject: Delegating request to RealObject...");
      this.isProcessing = true;
      setTimeout(() => {
        this.realObject.process();
        this.isProcessing = false;
      }, 3000); // 模擬需要較長時間處理的情況
    } else {
      console.log("ProxyObject: RealObject is currently busy.");
    }
  }
}

// 使用範例
const proxyObject = new ProxyObject();
proxyObject.process();
proxyObject.process(); // 第二次呼叫時會直接回傳 RealObject is currently busy.


4. 請實作一個享元模式，讓使用者可以存儲和重複使用相同或相似的物件。

答案：

class Flyweight {
  constructor(sharedState) {
    this.sharedState = sharedState;
  }

  operation(uniqueState) {
    const shared = JSON.stringify(this.sharedState);
    const unique = JSON.stringify(uniqueState);
    console.log(`Flyweight: Displaying shared (${shared}) and unique (${unique}) state.`);
  }
}

class FlyweightFactory {
  constructor() {
    this.flyweights = {};
  }

  getFlyweight(sharedState) {
    if (!(sharedState in this.flyweights)) {
      this.flyweights[sharedState] = new Flyweight(sharedState);
    }
    return this.flyweights[sharedState];
  }

  listFlyweights() {
    const count = Object.keys(this.flyweights).length;
    console.log(`FlyweightFactory: I have ${count} flyweights:`);
    for (const key in this.flyweights) {
      console.log(key);
    }
  }
}

// 使用範例
const factory = new FlyweightFactory();

const flyweight1 = factory.getFlyweight({ size: "small", color: "red" });
flyweight1.operation({ x: 10, y: 10 });

const flyweight2 = factory.getFlyweight({ size: "small", color: "red" });
flyweight2.operation({ x: 20, y: 20 });

const flyweight3 = factory.getFlyweight({ size: "small", color: "green" });
flyweight3.operation({ x: 30, y: 30 });

factory.listFlyweights();


5. 請實作一個外觀模式，讓使用者可以透過顯示簡單介面來操作複雜的程式邏輯。

答案：

class SubSystemA {
  operationA() {
    console.log("SubSystemA.operationA()");
  }
}

class SubSystemB {
  operationB() {
    console.log("SubSystemB.operationB()");
  }
}

class SubSystemC {
  operationC() {
    console.log("SubSystemC.operationC()");
  }
}

class Facade {
  constructor() {
    this.subSystemA = new SubSystemA();
    this.subSystemB = new SubSystemB();
    this.subSystemC = new SubSystemC();
  }

  operation1() {
    console.log("Operation 1 ----");
    this.subSystemA.operationA();
    this.subSystemB.operationB();
    console.log("Operation 1 completed ----");
  }

  operation2() {
    console.log("Operation 2 ----");
    this.subSystemC.operationC();
    console.log("Operation 2 completed ----");
  }
}

// 使用範例
const client = new Facade();
client.operation1();
client.operation2();   

