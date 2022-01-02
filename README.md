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

C) group, errCtx := ~errgroup.WithContext(ctx) 可以并发执行多个任务，一个任务失败则整个group执行结束，且只在第一次发生err的时候赋值给errCtx

### 3. 按照自己的构想，写一个项目满足基本的目录结构和工程，代码需要包含对数据层、业务层、API 注册，以及 main 函数对于服务的注册和启动，信号处理，使用 Wire 构建依赖。可以使用自己熟悉的框架。


####关键知识点：

A) 采用Gin网络框架，Gin是一个golang的微框架，封装比较优雅，API友好

B) 数据ORM框架，采用的是 Gorm，是一个全功能的ORM框架，扩展性良好，社区活跃，文档齐全，使用体验类似java的 mybatis

C）通过wire框架，实现依赖注入功能

D) 整体项目采用类似MVC架构，分别是 数据访问层dao包(对象是dto),业务服务层service包(对象是entity),展示层 handler包(对象是vo)

E) 代码组织参考  https://github.com/golang-standards/project-layout ，本项目主要有 api包,errors包，configs包，internal包，biz包等

### 4.参考 Hystrix 实现一个滑动窗口计数器。

A)  采用 container/ring 循环链表保存窗口的总访问次数，链表的每一个节点代表一秒内的访问次数

B) 采用 time.NewTicker 在异步goroutine内，每秒滑动一次窗口，扔掉最老的循环链表的节点和访问次数

C) 一旦窗口内访问次数超限，则限流，返回调用方http 429错误码和提示

D) 每次请求进来，如果处理成功则增加访问次数，如果被限流了则不增加访问次数。

