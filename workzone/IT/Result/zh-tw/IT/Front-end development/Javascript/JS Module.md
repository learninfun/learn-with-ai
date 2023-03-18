+++
title = "JavaScript 模組"
+++
## 習題預習
{{< ask_chatgpt >}}
給我5題JS Module的問題
{{< /ask_chatgpt >}}



1. JS Module有什麼好處？它如何幫助我們在網頁開發中更好的管理和組織程式碼？
2. 在使用JS模組進行開發時，你如何保證模組與模組之間的依賴性很好的解決，並且減少模組之間可能的衝突？
3. 為什麼有些JS模組需要用babel轉換？babel轉換有哪些好處？
4. 如何编写可重用的JS模組，並且避免因為過多的封裝，使得模組的重量過重？
5. 在使用JS模組進行開發時，你可以使用哪些打包工具？它们各有什麼特点，如何选择？   

## 說明知識
{{< ask_chatgpt >}}
說明JS Module並舉例
{{< /ask_chatgpt >}}



JS Module是一種將複雜的代碼分解成獨立的功能單元，方便維護和協作的技術。在JS中，Module通常指代一個獨立的js文件，這個文件中將某個特定的功能進行了封裝，並提供了對外的接口。其他js文件可以通過這些暴露出來的接口來使用這個模塊。

舉例來說，有一個名為"utils.js"的模塊，裡面定義了一些函數用於輔助其他模塊的開發。

```javascript
// utils.js
export function add(a, b) {
  return a + b;
}

export function substract(a, b) {
  return a - b;
}

export function multiply(a, b) {
  return a * b;
}
```

其他模塊可以使用"utils.js"中提供的函數，只需要在文件中引入即可。

```javascript
// main.js
import { add, substract } from './utils';

console.log(add(1, 2)); // 3
console.log(substract(5, 3)); // 2
```

JS模組化能夠提高代碼的可讀性和可維護性，有助於團隊協作和代碼重複利用，是現代Web開發中不可缺少的一部分技術。   

## 彙總重點
{{< ask_chatgpt >}}
條列JS Module的重點
{{< /ask_chatgpt >}}



以下是JS Module的重點：

1. 分離與封裝程式碼 - 模組提供了一個方法來將程式碼結構化，並使其更易於管理和重複使用。

2. 降低全域命名衝突的風險 - 模組可讓您在不同的區域中使用相同的變量名稱，從而降低全域污染和命名衝突的風險。

3. 支援封裝私有函數和變數 - 模組可讓您將程式碼封裝起來，以防止外部訪問或更改私有函數和變數。

4. 增強代碼重用性 - 模組提供了一個方法來組織和重複使用功能，使代碼可讀性更高，並降低了代碼的複雜度。

5. 實現高質量的代碼 - 模組可使程式碼更容易調試和測試，從而提高代碼的品質。

6. 支援多版本和異步加載 - 模組可使複雜的代碼更容易管理，並支援多個異步模組的同步和非同步加載。

7. 支援模組清單和依賴關係管理 - 模組可使您輕鬆管理和維護相互依賴的模組，並確保它們正確加載。

8. 相容於不同環境 - 模組可在不同的 JavaScript 環境中運行，如瀏覽器、Node.js 和其他運行 JavaScript 的環境。   

## 知識測驗
{{< ask_chatgpt >}}
給我5題JS Module的中等難度問題，並在後面列出答案
{{< /ask_chatgpt >}}



1. 使用JS模塊實現一個計算機模塊，該模塊提供加、減、乘、除四種基本運算，以及計算幾次運算公式的方法。例如：
```js
const calculator = require('./calculator');

console.log(calculator.add(1, 2)); // 3
console.log(calculator.subtract(4, 2)); // 2
console.log(calculator.multiply(3, 4)); // 12
console.log(calculator.divide(10, 5)); // 2
console.log(calculator.calculate('1+2-3*4/5')); // -0.2
```

答案：
```js
// calculator.js
const add = (a, b) => a + b;
const subtract = (a, b) => a - b;
const multiply = (a, b) => a * b;
const divide = (a, b) => a / b;

const calculate = (formula) => {
  let numbers = formula.split(/[\+\-\*\/]/).map(Number);
  let operators = formula.split(/[1234567890\.]+/).filter(Boolean);

  for (let i = 0; i < operators.length; i++) {
    if (operators[i] === '*') {
      numbers.splice(i, 2, numbers[i] * numbers[i+1]);
      operators.splice(i, 1);
      i--;
    } else if (operators[i] === '/') {
      numbers.splice(i, 2, numbers[i] / numbers[i+1]);
      operators.splice(i, 1);
      i--;
    }
  }

  let result = numbers[0];
  for (let i = 0; i < operators.length; i++) {
    if (operators[i] === '+') {
      result += numbers[i+1];
    } else if (operators[i] === '-') {
      result -= numbers[i+1];
    }
  }

  return result;
};

module.exports = {
  add,
  subtract,
  multiply,
  divide,
  calculate
};
```

2. 使用JS模塊實現一個點擊某個按鈕，就會觸發一種隨機動畫的模塊。該模塊提供三種隨機動畫：漸隱消失、循環變換顏色、抖動。例如：
```js
const randomAnimation = require('./randomAnimation');

document.querySelector('#button1').addEventListener('click', () => {
  randomAnimation.fadeOut('box1');
});

document.querySelector('#button2').addEventListener('click', () => {
  randomAnimation.colorCycle('box2');
});

document.querySelector('#button3').addEventListener('click', () => {
  randomAnimation.shake('box3');
});
```

答案：
```js
// randomAnimation.js
const fadeOut = (elementId) => {
  document.querySelector(`#${elementId}`).style.opacity = 1;

  let opacity = 1;
  let intervalId = setInterval(() => {
    opacity -= 0.1;
    document.querySelector(`#${elementId}`).style.opacity = opacity;

    if (opacity <= 0) {
      clearInterval(intervalId);
    }
  }, 50);
};

const colorCycle = (elementId) => {
  let currentColorIndex = 0;
  let colors = ['#FF0000', '#00FF00', '#0000FF', '#FFFF00', '#00FFFF', '#FF00FF'];

  let intervalId = setInterval(() => {
    document.querySelector(`#${elementId}`).style.backgroundColor = colors[currentColorIndex];
    currentColorIndex = (currentColorIndex + 1) % colors.length;
  }, 200);
};

const shake = (elementId) => {
  let positions = [{x:0, y:0}, {x:5, y:5}, {x:-5, y:-5}, {x:5, y:-5}, {x:-5, y:5}, {x:5, y:0}, {x:-5, y:0}, {x:0, y:5}, {x:0, y:-5}];

  let originalPosition = {x: document.querySelector(`#${elementId}`).offsetLeft, y: document.querySelector(`#${elementId}`).offsetTop};
  let currentPositionIndex = 0;

  let intervalId = setInterval(() => {
    let newPosition = {x: originalPosition.x + positions[currentPositionIndex].x, y: originalPosition.y + positions[currentPositionIndex].y};
    document.querySelector(`#${elementId}`).style.left = `${newPosition.x}px`;
    document.querySelector(`#${elementId}`).style.top = `${newPosition.y}px`;

    currentPositionIndex = (currentPositionIndex + 1) % positions.length;
  }, 50);
};

module.exports = {
  fadeOut,
  colorCycle,
  shake
};
```

3. 使用JS模塊實現一個簡單的購物車模塊。該模塊提供添加商品、刪除商品、計算總價格、清空購物車四種方法。例如：
```js
const cart = require('./cart');

cart.add({name: 'shirt', price: 20});
cart.add({name: 'pants', price: 30});

console.log(cart.totalPrice()); // 50

cart.remove('shirt');

console.log(cart.totalPrice()); // 30

cart.clear();

console.log(cart.totalPrice()); // 0
```

答案：
```js
// cart.js
let items = [];

const add = (item) => {
  items.push(item);
};

const remove = (itemName) => {
  items = items.filter((item) => item.name !== itemName);
};

const totalPrice = () => {
  return items.reduce((sum, item) => sum + item.price, 0);
};

const clear = () => {
  items = [];
};

module.exports = {
  add,
  remove,
  totalPrice,
  clear
};
```

4. 使用JS模塊實現一個模擬視頻播放器的模塊。該模塊提供設置視頻、播放、暫停、快進/退、全屏等功能。例如：
```js
const videoPlayer = require('./videoPlayer');

videoPlayer.setVideo('path/to/video.mp4');

document.querySelector('#playButton').addEventListener('click', () => {
  videoPlayer.play();
});

document.querySelector('#pauseButton').addEventListener('click', () => {
  videoPlayer.pause();
});

document.querySelector('#skipBackButton').addEventListener('click', () => {
  videoPlayer.skip(-10);
});

document.querySelector('#skipForwardButton').addEventListener('click', () => {
  videoPlayer.skip(10);
});

document.querySelector('#fullScreenButton').addEventListener('click', () => {
  videoPlayer.fullScreen();
});
```

答案：
```js
// videoPlayer.js
let videoUrl = '';
let isPlaying = false;
let currentTime = 0;
let volume = 0.5;

const setVideo = (url) => {
  videoUrl = url;
  document.querySelector('#videoElement').setAttribute('src', url);
};

const play = () => {
  if (isPlaying) return;

  document.querySelector('#videoElement').play();
  isPlaying = true;
};

const pause = () => {
  if (!isPlaying) return;

  document.querySelector('#videoElement').pause();
  isPlaying = false;
};

const skip = (seconds) => {
  let newTime = document.querySelector('#videoElement').currentTime + seconds;
  if (newTime < 0) {
    newTime = 0;
  } else if (newTime > document.querySelector('#videoElement').duration) {
    newTime = document.querySelector('#videoElement').duration;
  }

  document.querySelector('#videoElement').currentTime = newTime;
};

const fullScreen = () => {
  document.querySelector('#videoElement').requestFullscreen();
};

const setVolume = (value) => {
  volume = value;
  document.querySelector('#videoElement').volume = volume;
};

module.exports = {
  setVideo,
  play,
  pause,
  skip,
  fullScreen,
  setVolume
};
```

5. 使用JS模塊實現一個簡單的瀑布流圖片展示模塊。該模塊可以自動加載更多圖片，直到所有圖片都加載完。例如：
```js
const waterfall = require('./waterfall');

const images = ['path/to/image1.jpg', 'path/to/image2.jpg', 'path/to/image3.jpg', ...];

let currentPage = 0;
let loading = false;
let loadedAll = false;

window.addEventListener('scroll', () => {
  if (loading || loadedAll) return;
  
  if (window.innerHeight + window.scrollY >= document.body.offsetHeight) {
    loadNextPage();
  }
});

const loadNextPage = async () => {
  currentPage++;

  loading = true;

  const newImages = [];

  // 模擬從API中加載圖片
  for (let i = 0; i < 10; i++) {
    const imageIndex = (currentPage - 1) * 10 + i;
    if (imageIndex >= images.length) {
      loadedAll = true;
      break;
    }

    newImages.push(await loadImage(images[imageIndex]));
  }

  waterfall.addImages(newImages);

  loading = false;
}

const loadImage = (url) => {
  return new Promise((resolve, reject) => {
    const image = new Image();
    image.addEventListener('load', () => {
      resolve(image);
    });

    image.src = url;
  });
};
```

答案：
```js
// waterfall.js
let columns = [];
const columnCount = 3;
const margin = 10;

const container = document.querySelector('#container');

const addImages = (images) => {
  images.forEach((image) => {
    const columnIndex = getLowestColumnIndex();
    const left = columnIndex * (image.width + margin);
    const top = columns[columnIndex].height + margin;

    const imageElement = document.createElement('img');
    imageElement.setAttribute('src', image.src);
    imageElement.style.position = 'absolute';
    imageElement.style.left = `${left}px`;
    imageElement.style.top = `${top}px`;

    container.appendChild(imageElement);

    columns[columnIndex].height = top + image.height;
    columns[columnIndex].images.push(imageElement);
  });
};

const getLowestColumnIndex = () => {
  let columnIndex = 0;
  for (let i = 1; i < columns.length; i++) {
    if (columns[i].height < columns[columnIndex].height) {
      columnIndex = i;
    }
  }

  return columnIndex;
};

const initialize = () => {
  columns = [];
  for (let i = 0; i < columnCount; i++) {
    columns.push({height: 0, images: []});
  }
};

initialize();

module.exports = {
  addImages
};
```   

## 網路資料
{{< ask_chatgpt >}}
給我5篇JS Module的網路資料
{{< /ask_chatgpt >}}



1. "What is a JavaScript Module?" by Mozilla
https://developer.mozilla.org/en-US/docs/Web/JavaScript/Guide/Modules

2. "ES6 Modules in Depth" by Pony Foo
https://ponyfoo.com/articles/es6-modules-in-depth

3. "Mastering Modular JavaScript" by SitePoint
https://www.sitepoint.com/mastering-modular-javascript/

4. "JavaScript Modules: A Beginner’s Guide" by Toptal
https://www.toptal.com/javascript/javascript-modules-a-beginner-s-guide

5. "Creating and using JavaScript modules" by David Walsh
https://davidwalsh.name/javascript-modules   

