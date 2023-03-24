

1. 在Online Convex Optimization的框架下考虑线性回归问题，若目标函数为$w\in W\mapsto \sum_{i=1}^n(w\cdot x_i-y_i)^2$，其中$x_i\in R^d,y_i\in R$，请问所采用的算法应该是哪种？

答案：梯度下降法。


2. 在Online Convex Optimization的框架下考虑逻辑回归问题，若目标函数为$w\in W\mapsto \sum_{i=1}^n\log(1+e^{-y_iw\cdot x_i})$，其中$x_i\in R^d,y_i\in \{-1,1\}$，请问所采用的算法应该是哪种？

答案：随机梯度下降法。


3. 在Online Convex Optimization的框架下考虑最小化加权分类误差问题，若目标函数为$w\in W\mapsto \sum_{i=1}^n\varphi(\langle w,x_i\rangle,y_i)$，其中$\varphi:\mathbb{R}\times \{-1,1\}\rightarrow \mathbb{R}$为下凸函数，$x_i\in R^d,y_i\in \{-1,1\}$且有相应的权重$w_i$，请给出一个求解该问题的随机算法并进行推导。

答案：Stochastic Subgradient Descent (SSGD)。算法流程：
1. 初始化$\theta_1$为0。
2. 对于$t=2,\ldots,T$：
   a. 选择一个样本$(x_i,y_i)$，计算梯度下降方向$g_t$
   $$g_t=\theta_t'-w_i\nabla_1\varphi(\langle \theta_t,x_i\rangle,y_i)x_i$$
   b. 更新$\theta_t$
   $$\theta_{t+1}=\frac{1}{\sqrt{t}}\sum_{i=1}^tg_i$$
   其中$\theta_t'$为$\theta_t$的一个随机修正，用来解决部分精度梯度问题。

4. 在Online Convex Optimization的框架下考虑最小化$L_1$正则化的线性回归问题，若目标函数为$w\in W\mapsto \sum_{i=1}^n(w\cdot x_i-y_i)^2+\lambda \Vert w\Vert_1$，其中$x_i\in R^d,y_i\in R$，请问所采用的算法应该是哪种？

答案：Subgradient Descent。


5. 在Online Convex Optimization的框架下考虑最小化Huber损失函数的线性回归问题，若目标函数为$w\in W\mapsto \sum_{i=1}^n\varphi(w\cdot x_i-y_i)$，其中$\varphi:\mathbb{R}\rightarrow \mathbb{R}$为下凸函数，并且$\varphi(z)=z^2/2$，当$|z|\leq \delta$，且$\varphi(z)=\delta|z|-\delta^2/2$，当$|z|>\delta$，请问所采用的算法应该是哪种？

答案：Subgradient Descent。