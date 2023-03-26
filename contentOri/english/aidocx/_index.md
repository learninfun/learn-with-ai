+++
title = "aidocx: AI-Powered Knowledge Books"
weight = 2
+++

## Concept
Tree-structured directory X List of questions to learn ==> AI-generated book
![alt aidocx flow](https://github.com/learninfun/aidocx/blob/main/static/images/aidocx_flow_en.png?raw=true)

## Features
- Customizable knowledge points or directory to learn
- Customizable batch of questions to ask AI
- Automated scheduling to query AI and generate e-books.

## Usage
- Students or self-learners who want to customize learning materials based on their own needs to improve learning effectiveness and efficiency.
- Educators or educational institutions who want to quickly create teaching materials and easily update content.
- AI or machine learning professionals who want to explore how to use AI to create educational and training materials.

## Download and Usage Instructions
[Github](https://github.com/learninfun/aidocx)

## Example Results
[AI Knowledge Hub](https://learninfun.github.io/learn-with-ai/ai-knowledge-hub/) refers to the knowledge set extracted by aidocx.



## Explain
With the development of artificial intelligence technology, AI can automatically extract a relevant knowledge base from a large amount of data. By combining this AI technology with a tree-structured knowledge base, we can create a system that automatically generates technical books.

This automated system for generating technical books can greatly improve the learning efficiency of technical personnel. Users can quickly grasp the required knowledge by browsing these automatically generated books. At the same time, this system can quickly update the knowledge base and book content according to changes and developments in technology, maintaining the freshness and practicality of the books.

Such technology can also be applied in enterprise internal knowledge management systems to help companies quickly establish a complete knowledge system and automatically generate corresponding training materials and textbooks. This can greatly reduce the time and cost of internal training in companies, and improve the training efficiency and learning effectiveness of employees.

Aidocx is such a tool. By inputting a tree-structured knowledge list (in markdown format) and the expected questions to ask for each knowledge point (in yaml format), you can automatically generate a technical book.

## Prerequisites
Obtain an API key from OpenAI (Bard still does not provide a web API interface.)
[OpenAI](https://openai.com/blog/openai-api)

## Introduction to Usage
### 1. Downland aidocx 
[Github](https://github.com/learninfun/aidocx)

### 2. Add input.md: Enter the list of knowledge points you want to learn.
```markdown
- Knowledge point 1
  - Knowledge point 1.1
  - Knowledge point 1.2
  - Knowledge point 1.3
- Knowledge point 2
- Knowledge point 3
```

**知識點清單範例**
```markdown
- Machine Learning
  - Supervised Learning
    - Classification
    - Regression
  - Unsupervised Learning
    - Clustering
    - Dimensionality Reduction
- Deep Learning
  - Neural Networks
    - Activation Functions
    - Loss Functions
    - Backpropagation
```

### 3. Add config.yaml: adjust the questions to be asked for each knowledge point.

```yaml
apiProvider: OpenAI
apiModal: gpt-3.5-turbo-0301
initRole: Assuming you are an Machine Learning expert, answer my questions.
questions:
  - key: preview
    desc: Preview
    template: give me five question about {{ .keyword}}
  - key: explain
    desc: Explain
    template: Explain {{ .keyword}} and give an example
  - key: keypoint
    desc: Keypoint
    template: List the key points of {{ .keyword}}
  - key: review
    desc: Review
    template: Give me 5 medium-difficulty questions with answers about {{ .keyword}}
```

4. Run aidocx to generate the epub.
Note: The total execution time will be longer as the total number of questions increases, based on the multiplication of the number of questions per knowledge point.
```bash
aidocx -t epub \
       -o output.epub \
       -apikey-openai "-paste-your-api-key" \
        input.md
```

## Cost estimation
Aidocx is a free tool, but it requires calling an API provided by a vendor, such as Microsoft or Google, and the vendor will charge accordingly for the corresponding API key. For example, for ChatGPT3, the cost of 1000 tokens is 0.002 USD. For more details, please refer to OpenAI's [pricing page]((https://openai.com/pricing)).

A question, including the total number of characters in the question and answer, can have a maximum of 4096 tokens, so the maximum cost of a question is 0.008192 USD. If you have 100 knowledge points and each knowledge point has 5 questions, the expected cost will be 4.096 USD.

It should be noted that 4096 is the upper limit of one question for GPT-3, and the actual number of characters in the answer will not be that high. Based on experience, on average, the actual number of characters in the answer may be between 1/4 to 1/2 of the upper limit value, depending on the complexity of the question. If you use GPT-4, the cost will be higher, but you can get longer answers.
