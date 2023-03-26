Markov Decision Process (MDP) is a mathematical framework used for decision-making problems that involve stochastic processes. It is a model defined by a set of states, actions, and rewards that have a probability distribution over possible outcomes, which satisfies the Markov property.

In MDP, an agent takes actions in a particular state and receives a reward based on the action taken and the current state of the system. The goal is to maximize the cumulative rewards received over a sequence of actions, considering the probabilities associated with each state transition.

For instance, we can take the example of a robot cleaning a room as a Markov Decision Process. The robot has to take actions, such as moving left, right, or forward, to clean the room. The robot receives positive rewards for cleaning a dirty area in the room and negative rewards for hitting an obstacle or leaving the room uncleaned.

The process can be modeled by states, actions, and rewards. The states can be the different positions the robot can take in the room, and the actions can be moving left, right, or forward. The rewards can be defined as +1 for cleaning a dirty area, -1 for hitting an obstacle, and 0 for being in a clean area.

The probability of transitioning from one state to another depends on the action taken by the robot. If the robot moves left, the probability of reaching the left adjacent state is higher, whereas the probability of reaching a right adjacent state is lower.

By solving this MDP, the robot will learn the optimal sequence of actions it should take to clean the room efficiently while receiving the maximum possible reward.