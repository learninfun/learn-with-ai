

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