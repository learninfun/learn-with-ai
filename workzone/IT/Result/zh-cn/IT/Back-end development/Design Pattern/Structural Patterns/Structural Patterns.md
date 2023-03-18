+++
title = "结构型模式"
+++
## 习题预习
{{< ask_chatgpt >}}
给我5题Structural Patterns的问题
{{< /ask_chatgpt >}}



1. 什么是Structural Patterns？它们如何区分于其他设计模式？

2. 何时应该使用Adapter Pattern？它如何协助解决系统中的问题？

3. Decorator Pattern和Proxy Pattern的区别是什么？这些模式如何用于提高代码的可读性和可维护性？

4. 在使用Facade Pattern时，如何设计接口以支持系统的简化和压缩？你可以举例说明吗？

5. Bridge Pattern如何协助开发人员更好地设计具有多层次抽象的软件系统，并优化代码的可维护性和扩展性？   

## 说明知识
{{< ask_chatgpt >}}
说明Structural Patterns并举例
{{< /ask_chatgpt >}}



Structural Patterns是指将类和对象组合在一起形成更大的结构，用于解决对象的组合，继承等结构性问题。

以下是几个常见的Structural Patterns：

1. Adapter Pattern：将一个类的接口转换成客户希望的另一个接口。例如，将老版本的类适配成新版本的类。

2. Bridge Pattern：将抽象部分与它的实现部分分离，使得它们可以独立地变化。例如，将操作系统与应用程序分离，这样就可以实现不同操作系统的应用程序了。

3. Composite Pattern：用于管理树形结构的类，把相同的方法应用到每个元素上。例如，一个树形菜单，每个菜单项都有相同的方法。

4. Decorator Pattern：动态地给一个对象添加一些额外的职责，同时又不改变其结构。例如，装饰者模式可以增强一个类的功能，但不需要修改这个类本身。

5. Facade Pattern：为一组复杂的子系统提供一个简单的接口，通过这个接口来访问这个子系统。例如，一个文件系统的Facade可以让用户通过一个简单的接口来访问所有的文件。

6. Flyweight Pattern：运用共享技术有效地支持大量细粒度的对象。例如，一个网站可以使用Flyweight模式来减少服务器的压力。

这些模式可以提高代码的复用性，使代码更加灵活和可扩展。   

## 汇总重点
{{< ask_chatgpt >}}
条列Structural Patterns的重点
{{< /ask_chatgpt >}}



1.提供对系统内部结构的抽象表示

2.允许随着时间的推移进行系统内部的变化，同时不影响系统外部的用户界面

3.支持系统内部的外部化和内部化，使应用程序更加灵活

4.简化系统内部结构的操作，提高系统的可读性和可维护性

5.提供对系统中核心组件（例如，UI组件、数据组件等）的解耦，增强系统的可扩展性

6.提高系统的性能和效率，降低系统运行时所需的资源成本

7.提供对系统中不同部分的通信方式的抽象表示，使得系统之间的交互更加简单和易于管理

8.在软件设计过程中，提供一个共通的框架，方便开发人员进行协作、集成和测试   

## 知识测验
{{< ask_chatgpt >}}
给我5题Structural Patterns的中等难度问题，并在后面列出答案
{{< /ask_chatgpt >}}



1. 请实作一个简单的装饰器模式，可以让使用者在商品价格上添加不同的折扣。

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

// 使用范例
const product = new Product("iPhone", 1000);
const discountProduct = saleDiscount(product, 0.2); // 打80折

console.log(discountProduct.getName()); // iPhone (20% off)
console.log(discountProduct.getPrice()); // 800


2. 请实作一个中介者模式，让使用者可以协调不同的元件之间的互动。

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

// 使用范例
const mediator = new Mediator();
const component1 = new Component("Component 1");
const component2 = new Component("Component 2");

mediator.register(component1);
mediator.register(component2);

component1.send("Hello, Component 2");
component2.send("Hi, Component 1");


3. 请实作一个代理模式，让使用者可以透过代理物件存取需要较长时间才能处理的资料。

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
      }, 3000); // 模拟需要较长时间处理的情况
    } else {
      console.log("ProxyObject: RealObject is currently busy.");
    }
  }
}

// 使用范例
const proxyObject = new ProxyObject();
proxyObject.process();
proxyObject.process(); // 第二次呼叫时会直接回传 RealObject is currently busy.


4. 请实作一个享元模式，让使用者可以存储和重复使用相同或相似的物件。

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

// 使用范例
const factory = new FlyweightFactory();

const flyweight1 = factory.getFlyweight({ size: "small", color: "red" });
flyweight1.operation({ x: 10, y: 10 });

const flyweight2 = factory.getFlyweight({ size: "small", color: "red" });
flyweight2.operation({ x: 20, y: 20 });

const flyweight3 = factory.getFlyweight({ size: "small", color: "green" });
flyweight3.operation({ x: 30, y: 30 });

factory.listFlyweights();


5. 请实作一个外观模式，让使用者可以透过显示简单介面来操作复杂的程式逻辑。

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

// 使用范例
const client = new Facade();
client.operation1();
client.operation2();   

## 网络数据
{{< ask_chatgpt >}}
给我5篇Structural Patterns的网络数据
{{< /ask_chatgpt >}}



1. "Structural Design Patterns in Java" by TutorialsPoint: This article provides a comprehensive overview of the various Structural Design Patterns in Java, including the Adapter, Bridge, Composite, Decorator, and Facade patterns. It explains the purpose and implementation of each pattern, along with code examples for better understanding.

2. "Structural Patterns" by SourceMaking: This website offers a detailed explanation of Structural Patterns, including the Adapter, Bridge, Composite, Decorator, Façade and Flyweight design patterns. The site also provides code examples in Java to demonstrate how to implement each pattern.

3. "The Top 5 Structural Design Patterns for Beginners" by DZone: This article is intended for beginners and provides an overview of five of the most commonly used Structural Design Patterns. It explains each pattern with examples and how to use it in different scenarios.

4. "Structural Design Patterns: A Guide to Decoding Software Architecture" by Stackify: This article provides an in-depth explanation of Structural Design Patterns, including information about their purpose and how to apply them in real-world situations. Additionally, it has code examples in C# to help readers better understand this programming concept.

5. "Structural Design Patterns" by GeeksforGeeks: This website provides an easy-to-understand explanation of Structural Design Patterns, including the Adapter, Bridge, Composite, Decorator, Flyweight, and Facade patterns. It includes diagrams and examples in Java to help readers understand the concept behind each pattern.   

