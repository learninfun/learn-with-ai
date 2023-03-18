

Primality testing是一种判断一个数字是否为质数的方法。质数是仅能被1和自身整除的正整数，例如2、3、5、7、11等。

Primality testing方法有很多种，下面列出几种常见的方法：

1. 费马小定理：对于一个质数p和任意不是p的整数a，a^(p-1) mod p = 1，如果结果不等于1，则a不是质数。这种方法非常快，但存在伪质数的情况，即一些合数通过此测试。

2. Miller-Rabin测试：使用随机算法来测试质数，并且能够检测伪质数，并且可以设置错误率。这是一种广泛使用的primality testing方法。

3. AKS算法：这是一种相对较新的方法，可以在多项式时间内确定一个数字是否为质数，但是速度非常慢，难以处理大数据。

下面以示例展示Miller-Rabin测试如何检测质数。

假设我们要判断n=17是否为质数。

1. 将n-1表示为2^r*d(d是奇数)，因为16=2^4*1，所以r=4，d=1。

2. 随机选择一个a，并且取a^d mod n的值(这里取a=2)。由于d=1，所以a^d mod n=2^1 mod 17=2。

3. 只要a^(d*2^j) mod n不等于1且不等于n-1(j=0,1,...,r-1)，则n不是质数。这里需要计算2^2 mod 17=4，即a^(d*2^1) mod n，由于4不等于1且不等于16(即n-1)，所以n不是质数。

所以17不是质数。

总之，Primality testing是寻找质数的重要方法，不同方法的速度和精度各异，可以根据需要选择适当的方法。