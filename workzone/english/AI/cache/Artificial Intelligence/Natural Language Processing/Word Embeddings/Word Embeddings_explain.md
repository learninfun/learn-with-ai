Word embeddings are a type of natural language processing technique in which words are represented as dense vectors, typically in a high-dimensional space (100–300 dimensions). Each word in a corpus (large body of text) is assigned a distinct vector, with words that have similar meanings or are contextually related being closer together in that space.

For example, with Word2Vec, we can represent the word “king” as the vector:

[0.2, -0.4, 0.1, …, 0.9]

And the word “queen” as:

[0.3, -0.5, 0.2, …, 0.9]

Where each number represents a characteristic feature of the word. In this way, we can use these vectors to measure relationships between words and even perform algebraic operations on them, like:

king - man + woman = queen

This equation compares the vector for "king" with the vector for "man," subtracts the vector for "man," and adds the vector for "woman" to find the most similar vector, which is the vector for "queen." Word embeddings are used in many natural language processing tasks, including sentiment analysis, machine translation, and document classification.