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

### 5. (A)使用 redis benchmark 工具, 测试 10 20 50 100 200 1k 5k 字节 value 大小，redis get set 性能。(B)写入一定量的 kv 数据, 根据数据大小 1w-50w 自己评估, 结合写入前后的 info memory 信息 , 分析上述不同 value 大小下，平均每个 key 的占用内存空间

A)在本人开发机读写性能，结论如下:value大小 10K字节的 get、set 读写性能大约 2万/秒，没有明显衰减，在value大小 100K字节开始，get/set 性能降低到 7K/秒;
另外对于 value 从10字节~5K字节过程中,lrange命令性能退化严重，当value为10字节的时候lrange(100) 为 1.3万/秒，当value为5K时为820/秒

B)内存占用分析结论：redis的内存占用基本和数据量增长成正比，每个key的占用空间，基本是value值字节的大小

![img_1.png](img_1.png)

### 6. 网络编程
#### (A)总结几种 socket 粘包的解包方式：fix length/delimiter based/length field based frame decoder。尝试举例其应用
#### 1.fix length:发送端将每个数据包封装为固定长度（不够的可以通过补0填充），这样接收端每次从接收缓冲区中读取固定长度的数据就自然而然的把每个数据包拆分开来。该协议实现方式简单，适合同一种数据格式大量传输。
#### 2.delimiter based:可以在数据包之间设置边界，如添加特殊符号，这样，接收端通过这个边界就可以将不同的数据包拆分开。比较适合长度不固定的数据结构，但是数据用途比较单一，或者要求传输速度数据量大。
#### 3.length field based frame decoder 将报文划分为报文头/报文体，根据报文头中,Length字段确定报文体的长度，因此报文提的长度是可变的.比较灵活的协议，适用于用途比较多样的场景，例如IM系统，

#### (B) 实现一个从 socket connection 中解码出 goim 协议的解码器。
#### 具体参考work6文件夹代码中的client.go 和server.go和common.go. 注意：goim协议参考的是github中的截图：https://github.com/Terry-Mao/goim/blob/master/docs/protocol.png
####  1. goim协议头约定了 PackageLength 4bytes, HeaderLength 2bytes，Protocol Version 2byes，Operation 4bytes，SequenceId 4bytes，报文Body 长度为 Package length - Header length。
####  2. 先启动 server.go ，然后再启动client.go   就可观察到数据正常传输到 server端了 ，而且是按照goim协议。


### 7. 毕业项目

本地启动redis
redis-server.exe redis.windows.conf

GET http://localhost:8080/api/user?userid=1
{"code":0,"data":{"Msg":"成功","UserId":1,"Name":"张三","Email":"zhangsan@qq.com","Age":18},"msg":"sucess"}

http://localhost:8080/api/product?productid=1
{"code":0,"data":{"Msg":"成功","ProductId":1,"Name":"衣服","Price":1000,"Stock":100},"msg":"sucess"}

