Recurrent Neural Networks (RNNs) are a type of neural network specifically designed to process sequential data. Unlike traditional neural networks that are feedforward, RNNs have loops that allow information to be carried from one step of the network to the next.

In an RNN, nodes receive inputs at each time step, along with the output from the previous time step. This input, together with the information carried from the previous time step, is then fed into the same set of node weights to generate the output for the current time step. The output from each time step can be treated as a function of the inputs at all previous time steps combined.

A common example of RNNs is language modeling. In this case, the input sequence is a sentence (or a collection of sentences), and the output is the probability distribution over all possible next words. RNNs can use the context of previous words to generate predictions for the next word in the sequence. For instance, when predicting the next word in the sentence "I want to eat some ____", the RNN would predict "food" as a more likely candidate than "paint". 

Another example of RNNs can be found in sentiment analysis, where the input sequence is a user's text message and the output is the sentiment label (positive, negative or neutral) of the message. RNNs can learn from the sequence of words and their corresponding sentiment labels to generate a sentiment label for a new text message.