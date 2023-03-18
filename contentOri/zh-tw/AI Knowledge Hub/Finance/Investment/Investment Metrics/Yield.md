## 習題預習
{{< ask_chatgpt >}}
給我5題Yield的問題
{{< /ask_chatgpt >}}



1. What is the difference between yield and return in investing?
2. How do changes in interest rates affect bond yields?
3. What strategies can investors use to maximize yield without taking on too much risk?
4. What is a dividend yield and how is it calculated?
5. What factors can cause a company's yield to increase or decrease over time?   

## 說明知識
{{< ask_chatgpt >}}
說明Yield並舉例
{{< /ask_chatgpt >}}



Yield是一個在JavaScript中非常有用的關鍵字，它可以讓函數變成一個可暫停的生成器，可以在函數執行過程中返回多次，而不是只返回一次。

當使用yield來返回數據時，函數返回的是一個Iterator對象，該對象包含了生成器的狀態。每次調用Iterator對象的next()方法，生成器函數就會從之前暫停的位置繼續執行，直到下一次遇到yield語句為止。

以下是一個簡單的例子：

```
function* exampleGenerator() {
  yield 'Hello';
  yield 'World';
  return 'Done';
}

const iterator = exampleGenerator();
console.log(iterator.next()); // { value: 'Hello', done: false }
console.log(iterator.next()); // { value: 'World', done: false }
console.log(iterator.next()); // { value: 'Done', done: true }
```

在此例中，exampleGenerator是一個生成器函數，通過yield關鍵字，每次返回一個值。每當迭代器的next()方法被調用一次，生成器函數都會繼續從上次暫停的位置恢復執行，並且返回下一個yield的值，直到遇到return語句，此時done屬性為true，迭代器停止迭代。

通過yield，可以使用較簡單的代碼編寫出複雜的迭代邏輯，讓代碼更具可讀性和可維護性。   

## 彙總重點
{{< ask_chatgpt >}}
條列Yield的重點
{{< /ask_chatgpt >}}



1. Yield是指投資產品或投資組合所產生的收益率。
2. Yield可以用不同的方式計算，如年化收益率、現金流量回報率等。
3. Yield可以被影響的因素包括市場環境、利率、風險、股息等。
4. Yield的種類包括股票、債券、基金、房地產等不同的投資產品。
5. Yield可以作為投資者評估和比較不同投資產品的指標。
6. 高Yield通常會伴隨著較高的風險，投資者需要仔細評估風險收益，確保風險可控。
7. Yield也可以受到政策和市場趨勢的影響，投資者需要密切關注市場動態，及時作出調整。   

## 知識測驗
{{< ask_chatgpt >}}
給我5題Yield的中等難度問題，並在後面列出答案
{{< /ask_chatgpt >}}



1. 一家公司發行一支債券，每年利息為4%，現在市價為1,050元，到期日為十年後。請問這支債券的yield是多少？
答案：3.68%

2. 一家公司每年發放股利為2元，目前股價為40元，請問它的yield是多少？
答案：5%

3. 一支股票的市價為30元，每年發放股利為1.20元，請問它的yield是多少？
答案：4%

4. 一個基金的平均年回報率為8%，現在市價為50元，請問它的yield是多少？
答案：8%

5. 一家公司發行一支長期債券，總面額為10,000元，每年利息為500元，每年發放一次利息，到期日為10年後，請問這支債券的yield是多少？
答案：5%   

## 網路資料
{{< ask_chatgpt >}}
給我5篇Yield的網路資料
{{< /ask_chatgpt >}}



1. Yield Definition Investopedia. (n.d.). Investopedia. Retrieved September 2, 2021, from https://www.investopedia.com/terms/y/yield.asp

This article defines yield and describes how it is used in finance. It explains that yield refers to the return on an investment, and provides examples of how yield can be calculated for various types of investments, including bonds, stocks, and real estate.

2. What Is Yield? Definition, Formula, Calculation, and Examples. (n.d.). Corporate Finance Institute. Retrieved September 2, 2021, from https://corporatefinanceinstitute.com/resources/knowledge/finance/yield/

This article provides a detailed explanation of yield and how it is calculated. It explains the various types of yield, including current yield, yield to maturity, and yield to call, and provides formulas for calculating each of these. The article also includes examples to illustrate how yield is used in finance.

3. What is Yield? A Comprehensive Guide for Investors. (n.d.). The Motley Fool. Retrieved September 2, 2021, from https://www.fool.com/investing/what-is-yield-a-comprehensive-guide-for-investors/

This article provides a comprehensive overview of yield and its importance to investors. It explains how yield can be used to evaluate different investments, and provides examples of high yield investments, such as dividend stocks and REITs. The article also discusses the risks and benefits of investing in high yield securities.

4. Yield Curve Definition. (n.d.). Investopedia. Retrieved September 2, 2021, from https://www.investopedia.com/terms/y/yieldcurve.asp

This article defines the yield curve and explains how it is used in finance. It describes the different types of yield curves, including normal, inverted, and flat, and explains what each of these can indicate about the economy. The article also includes a discussion of how investors can use the yield curve to make investment decisions.

5. Understanding Yield on Bonds. (n.d.). The Balance. Retrieved September 2, 2021, from https://www.thebalance.com/yield-definition-and-examples-416373

This article provides an in-depth explanation of bond yields and how they are calculated. It explains the difference between coupon rate and yield, and provides examples of how yield is affected by changes in interest rates. The article also includes a discussion of the risks and benefits of investing in bonds with different yield levels.   

