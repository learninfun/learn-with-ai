1) What is the main difference between an LSTM and a GRU?
Answer: The main difference is that a GRU has two gates (reset and update) while an LSTM has three gates (input, output, and forget).

2) What is the purpose of the reset gate in a GRU?
Answer: The reset gate determines how much of the previous hidden state should be forgotten.

3) How does a GRU prevent vanishing gradients?
Answer: A GRU uses a direct path from the current input to the current hidden state, making it easier for gradients to flow during backpropagation.

4) What is the role of the update gate in a GRU?
Answer: The update gate determines how much of the previous hidden state should be retained and how much of the new input should be added.

5) How can a GRU be used for time series prediction?
Answer: A GRU can be trained on a time series dataset to predict the next value in the series, and can be used to generate forecasts and make predictions about future trends.