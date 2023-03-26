

- Saga Pattern可以用来协调分散的事务，保证系统在失败的情况下能够恢复正常运作
- 单独的每个事务都是幂等的，即在重复执行时不会改变系统的状态
- Saga是一个长时间的交互过程，它管理一系列相关的事务
- Saga Pattern具有不可撤销性，一旦一个事务完成，它就无法被撤回
- Saga模式的实现方式有两种： Choreography-based saga和Orchestration-based saga