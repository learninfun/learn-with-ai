1. What is Deep Q-Network (DQN)?
Answer: DQN is a reinforcement learning algorithm that combines deep neural networks with Q-learning, allowing for the use of high-dimensional input spaces such as raw pixel data in video games.

2. How does DQN work?
Answer: DQN uses a deep neural network to approximate the Q-function, which predicts the value of taking an action in a given state. The algorithm iteratively updates the weights of the neural network to improve the accuracy of the Q-function and optimize the policy.

3. What are the key components of DQN?
Answer: The key components of DQN include the replay buffer, which stores experiences for learning from past transitions, and the target network, which provides stable Q-value targets by periodically copying the weights of the online network.

4. What are some limitations of DQN?
Answer: DQN can suffer from overestimation of Q-values due to the use of a maximum operator in the Q-learning update rule. Additionally, it may struggle in domains where exploration is difficult, as it relies on past experiences stored in the replay buffer.

5. How have researchers extended DQN to improve performance?
Answer: Researchers have extended DQN through various techniques such as prioritized experience replay, dueling networks, and distributional reinforcement learning, which aim to address challenges such as sample efficiency, stability, and generalization to different task distributions.