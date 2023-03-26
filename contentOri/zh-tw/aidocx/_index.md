+++
title = "aidocx: 知識擷取小幫手"
weight = 2
+++

## 概念
樹狀排列的目錄 X 問題清單 ==> AI產生的書籍
![alt aidocx flow](https://github.com/learninfun/aidocx/blob/main/static/images/aidocx_flow_zh-tw.png?raw=true)

## 特點
- 自定義想了解的知識點或目錄
- 自定義想要批次詢問AI的問題
- 自動排程詢問AI並且產生電子書

## 用途
- 學生或自學者，希望使用新穎的學習工具來提高學習效果和效率。
- AI或機器學習專業人員，希望探索如何應用AI來創建教育訓練教材。
- 教育工作者或教育機構，尋求創新的教學方法和教材。

## 下載與使用說明
[Github](https://github.com/learninfun/aidocx)

## 範例結果
[AI理解的知識庫](https://learninfun.github.io/learn-with-ai/zh-tw/ai-knowledge-hub/)，就是以aidocx抽取的知識集

## 說明
隨著人工智慧技術的發展，AI可以自動從大量的資料中提取出相關的知識點體系。而結合這種人工智慧技術和樹狀排列的知識點，我們可以得到一個自動產生技術書籍的系統。

這種自動化生成技術書籍的系統，可以大大提高技術人員的學習效率。使用者可以通過瀏覽這些自動產生的書籍，快速地掌握所需的知識。同時，這種系統也可以快速地根據技術的變化和發展，自動更新知識庫和書籍內容，保持書籍的新鮮度和實用性。

這樣的技術也可以應用在企業內部的知識管理系統中，幫助企業快速建立完整的知識體系，並自動生成對應的培訓材料和教材。這樣可以大大減少企業內部培訓的時間和成本，提高員工的培訓效率和學習成效。

aidocx就是這樣一個工具，只要你輸入以樹狀排列的知識點清單(markdown格式)，加上每個知識點預計詢問的問題(yaml格式)，就可以自動產生一本技術書籍。

## 使用前提
經由使用各家廠商所提供的API介面呼叫AI，因此在使用前需要先申請相對應的API Key。
以OpenAI的ChatGPT為例，可使用此網址進行申請: [申請網址](https://openai.com/blog/openai-api)

## 使用方式
1. 下載aidocx 
[Github](https://github.com/learninfun/aidocx)

2. input.md: 輸入想要了解的知識點清單
```markdown
- 知識點1
  - 知識點1.1
  - 知識點1.2
  - 知識點1.3
- 知識點2
- 知識點3
```

2. 調整每個知識點想要問的問題

config.yaml:
```yaml
apiProvider: ChatGPT
apiModal: gpt-3.5-turbo-0301
initRole: Assuming you are an "IT" expert, answer my questions.
questions:
  - key: preview
    desc: 習題預習
    template: 給我5題{{ .keyword}}的中文問題
  - key: explain
    desc: 說明知識
    template: 以中文說明{{ .keyword}}並舉例
  - key: keypoint
    desc: 條列重點
    template: 以中文條列{{ .keyword}}的重點
  - key: test
    desc: 知識測驗
    template: 以中文給我5題{{ .keyword}}的中等難度問題，並在後面列出答案
```

3. 執行apidoc，獲得epub
注意: 按照知識點乘上問題數目，總問題越多，執行時間會越久。
```bash
aidocx -t epub \
       -o output.epub \
       -apikey-openai "-paste-your-api-key" \
        input.md
```

## 使用費用估計
aidocx本身是免費的，不過呼叫廠商(例如微軟、Google)的API，廠商本身會針對api key收費。
以ChatGPT3為例，1000個token收費0.002美金 [參考資料](https://openai.com/pricing#language-models)。
一個問題，題目和答案的字數合併計算，上限4096個token。
所以1個問題，最多4096/1000*0.002 = 0.008192美金。
如果有100個知識點，每個知識點問5個問題。
預計花費為 100*5*0.008192 = 4.096美金。

4096個是GPT-3的上限，正常回答的字數，不會這麼多。
按照經驗來說，平均使用會是上限值的1/4到1/2，具體數字按照問題複雜度而定。
相對使用GPT-4的費用就比較貴了，它可以回答的長度也會更長。
