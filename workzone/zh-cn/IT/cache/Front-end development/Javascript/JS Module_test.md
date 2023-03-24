

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