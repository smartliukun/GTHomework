# GTHomework 训练营作业，请勿使用到实际项目

### 1. 我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？
结论：取决于业务期望的结果是怎样的

A）调用方希望返回一个集合（多条数据，或者空数据都可以的时候），这个时候，表里没有数据就不应该返回error

B）调用方希望返回至少一条数据的时候，这个时候，表里没有数据应该返回error。

如果需要传递这个错误给上层，则考虑error是否要wrap，建议由调用方wrap这个error


### 2. 基于 errgroup 实现一个 http server 的启动和关闭 ，以及 linux signal 信号的注册和处理，要保证能够一个退出，全部注销退出。


####关键知识点：

A) context 可以使用以下方法增加cancel方法 ，以在父子goroutine中安全的传递cancel的消息
ctx, cancel := context.WithCancel(context.Background())

B）signal.Notify 方法可以监听中断信号，并放在channel中，可以在不同的goroutine中安全传递信号

C) group, errCtx := errgroup.WithContext(ctx) 可以并发执行多个任务，一个任务失败则整个group执行结束，且只在第一次发生err的时候赋值给errCtx

D)



