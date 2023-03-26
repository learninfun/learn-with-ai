+++
title = "aidocx: 知识撷取小帮手"
weight = 2
+++

## 概念
树状排列的目录 X 问题列表 ==> AI产生的书籍
![alt aidocx flow](https://github.com/learninfun/aidocx/blob/main/static/images/aidocx_flow_zh-tw.png?raw=true)

## 特点
- 自定义想了解的知识点或目录
- 自定义想要批次询问AI的问题
- 自动排程询问AI并且产生电子书

## 用途
- 学生或自学者，针对自己的需求客制化教材来提高学习效果和效率。
- 教育工作者或教育机构，快速编撰教材并方便更新内容。
- AI或机器学习专业人员，希望探索如何应用AI来创建教育训练教材。

## 下载与使用说明
[Github](https://github.com/learninfun/aidocx)

## 范例结果
[AI理解的知识库](https://learninfun.github.io/learn-with-ai/zh-tw/ai-knowledge-hub/)，就是以aidocx抽取的知识集

## 说明
随着人工智能技术的发展，AI可以自动从大量的资料中提取出相关的知识点体系。而结合这种人工智能技术和树状排列的知识点，我们可以得到一个自动产生技术书籍的系统。

这种自动化生成技术书籍的系统，可以大大提高技术人员的学习效率。使用者可以通过浏览这些自动产生的书籍，快速地掌握所需的知识。同时，这种系统也可以快速地根据技术的变化和发展，自动更新知识库和书籍内容，保持书籍的新鲜度和实用性。

这样的技术也可以应用在企业内部的知识管理系统中，帮助企业快速建立完整的知识体系，并自动生成对应的培训材料和教材。这样可以大大减少企业内部培训的时间和成本，提高员工的培训效率和学习成效。

aidocx就是这样一个工具，只要你输入以树状排列的知识点列表(markdown格式)，加上每个知识点预计询问的问题(yaml格式)，就可以自动产生一本技术书籍。

## 使用前提
经由使用各家厂商所提供的API接口呼叫AI，因此在使用前需要先申请相对应的API Key。
以OpenAI的ChatGPT为例，可使用此网址进行申请: [申请网址](https://openai.com/blog/openai-api)

## 使用方式
### 1. 下载aidocx 
[Github](https://github.com/learninfun/aidocx)

### 2. 新增input.md: 输入想要了解的知识点列表
```markdown
- 知识点1
  - 知识点1.1
  - 知识点1.2
  - 知识点1.3
- 知识点2
- 知识点3
```

**知识点列表范例**
```markdown
- 机器学习
  - 监督式学习
    - 分类
    - 回归
  - 非监督式学习
    - 聚类
    - 降维
- 深度学习
  - 神经网络
    - 启动函数
    - 损失函数
    - 反向传播
```

### 3. 新增config.yaml: 调整每个知识点想要问的问题

```yaml
apiProvider: OpenAI
apiModal: gpt-3.5-turbo-0301
initRole: Assuming you are an "IT" expert, answer my questions.
questions:
  - key: preview
    desc: 习题预习
    template: 给我5题{{ .keyword}}的中文问题
  - key: explain
    desc: 说明知识
    template: 以中文说明{{ .keyword}}并举例
  - key: keypoint
    desc: 条列重点
    template: 以中文条列{{ .keyword}}的重点
  - key: test
    desc: 知识测验
    template: 以中文给我5题{{ .keyword}}的中等难度问题，并在后面列出答案
```

4. 执行apidoc，获得epub  
注意: 按照知识点乘上问题数目，总问题越多，运行时间会越久。
```bash
aidocx -t epub \
       -o output.epub \
       -apikey-openai "-paste-your-api-key" \
        input.md
```

## 使用费用估计
aidocx 是一个免费的工具，但它需要呼叫厂商提供的 API，比如 Microsoft 或 Google。而厂商会针对相应的 API key 进行收费的。以 ChatGPT3 为例，1000 个 token 的费用为 0.002 美元，更多详情可参考 OpenAI 的[定价页面](https://openai.com/pricing)。

一个问题，包括题目和答案的字数总和最多为 4096 个 token，所以一个问题的最高成本是 0.008192 美元。如果您有 100 个知识点，并且每个知识点有 5 个问题，那么预计花费的金额为 4.096 美元。

需要注意的是，4096 是 GPT-3 一个问题的上限，实际上正常回答的字数并不会这么多。根据经验，平均而言，实际回答的字数可能在上限值的 1/4 到 1/2 之间，具体数字还取决于问题的复杂度。如果您使用 GPT-4，则费用会更高，但您可以得到更长的回答。