# web-template-ddd
## 1.简述
基于DDD架构的Web服务快速构建模板。

>DDD领域驱动设计是一个很宽泛的方法论，涉及到的概念也很多，DDD并没有给出标准的代码模型，不同的人可能会有不同理解。
> 
>该模板项目尝试通过不断的理解领域驱动，结合DDD的思想和规范，构建出一个代码目录结构，将DDD落地。

## 2.DDD分层架构
分层架构有一个重要的原则：每层只能与位于其下方的层发生耦合。具体又可以分为：严格分层架构和松散分层架构。

- 严格分层架构：某层只能与直接位于其下方的层发生耦合；
- 松散分层架构：允许任意上方层与任意下方层发生耦合。

严格分层，自然是最理想化的，但这样肯定会导致大量繁琐的适配代码出现，故在严格与松散之间，一般追寻和把握一下平衡。

DDD包含4层，将领域模型和业务逻辑分离出来，并减少对基础设施、用户界面甚至应用层逻辑的依赖，因为它们不属业务逻辑。将一个复杂 的系统分为不同的层，每层都应该具有良好的内聚性，并且只依赖于比其自身更低的层。

![img.png](https://img-blog.csdnimg.cn/1bdb7c94da80488db5b1af723675b665.png?x-oss-process=image/watermark,type_ZHJvaWRzYW5zZmFsbGJhY2s,shadow_50,text_Q1NETiBATXJLb3JiaW4=,size_15,color_FFFFFF,t_70,g_se,x_16)


### 2.1 用户接口层
该层在基于Gin框架的实践中，我更偏向于将其命名为路由层，因为GIN处理了这层的绝大部分逻辑。这层一般包括如下内容：

- Web服务和中间件
- 对外暴露的API接口，接受用户或者外部系统的请求，响应必要的数据信息
- DTO数据传输对象，即请求和响应的数据对象
- 数据安全性校验，比如：id不为空

用户接口层我将其命名为`server`，因为我觉得它包含了构建一个Web服务的大部分内容，以此命名更容易让人一看到就知道这是微服务的入口，目录结构如下：
```
.
├─server          # 用户接口层
│  ├─handler      # 路由
│  ├─middleware   # 中间件，如CROS，认证拦截器，过滤器等
│  └─dto        # DTO数据传输对象
.
```


### 2.2 应用层

应用层关心处理完一个完整的业务逻辑，该层只负责业务编排，对象转换，实际业务逻辑由领域层完成。应用层不关心【请求从何处来】，但是关心【谁来做、做什么、有没有权限做】。该层非常适合处理事务，日志和安全等。相对于领域层，应用层应该是很薄的一层。它只是协调领域层对象执行实际的工作。

```
.
├─application # 应用层
.
```
>这一层的目录我趋向于让其简洁一点，当然如果想拆分得更加详细，像处理消息发布订阅这部分逻辑的，可以再单独在`application`包下建类似`publish`和`subscribe`这样的包。
> 
>但我觉得这也是属于被编排的一部分逻辑，直接调用领域层或基础设施层的相关处理即可。


### 2.3 领域层
领域层主要包含聚合根、实体、值对象、领域服务等领域模型中的领域对象；领域层主要负责表达业务概念，业务状态信息和业务规则。领域层是整个系统的核心层，几乎全部的业务逻辑会在该层实现。领域模型层主要包含以下的内容：

- 实体(Entities):具有唯一标识的对象, 如：商品
- 值对象(Value Objects): 无需唯一标识, 如：商品快照
- 领域服务(Domain): 与业务逻辑相关的，具有属性和行为的对象
- 聚合/聚合根(Aggregates & Aggregate Roots): 聚合是指一组具有内聚关系的相关对象的集合
- 仓储(Repository): 提供持久化数据和操作数据库的方法

```
.
├─domain         # 领域层
│  ├─entity      # 实体
|  ├─vo          # 值对象
│  ├─repository  # 仓储
│  └─service     # 服务OR聚合
.
```
>这里比较特殊的是`service`，我将`service`理解为当处理单一实体甚至聚合不能很好解决的场景时，为了保持实体本身自己的内聚，此时才新建`service`做处理；并且，我发现`service`实际上本身也是一种聚合的形式，包含了`repository`和`entity`等，所以目前我将聚合也放到`service`当中。

### 2.4 基础设施层
基础设施层为上面各层提供通用的技术能力：为应用层传递消息，为领域层提供持久化机制，为用户界面层提供通用组件等。基础设施层以不同的方式支持所有三个层，促进层之间的通信。

```
.
├─component
│  ├─cache
│  ├─doc
│  ├─pubsub
│  ├─storage
│  └─util
│      └─constant
.
```
如果你不太确定某一部分组件属于那一层，那么我觉得你都可以将其放到基础设施层，因为该层支撑着其他三层，任何一层从该层使用某一组件，都是合理的，也是符合DDD思想的。

所以我也很自然的将其取名为`component`，这样很容易让我一眼就看出来这里放了支撑全局的各种通用组件或者工具。

## 3. 目录结构
```
./
├─application
├─cmd
├─component
│  ├─cache
│  ├─doc
│  ├─pubsub
│  ├─storage
│  └─util
│      └─constant
├─config
├─domain
│  ├─entity
│  ├─repository
│  ├─vo
│  └─service
├─log
├─script
├─server
│  ├─handler
│  ├─middleware
│  └─proto
└─worker
```
除了领域驱动中常见的四层目录，对于一个完整的项目，和一些常见的项目，我又补充了一些专门用于描述这部分需求和逻辑的目录：

```
./
├─cmd       # 命令行程序
├─config    # 配置文件
├─log       # 日志文件（通常只适用于本机调试时的日志输出）
├─script    # 脚本，如：数据库索引可以记录在一个index.js脚本文件中
└─worker    # 工作进程，包括独立部署的进程和集成在服务中的守护进程等，如：定期清理数据库软删除数据的定时器
```
## 4.实践简例
详细案例请自行查看模板项目代码。

## 5.编码风格
可以发现所有的包名都采用的单数形式，主要参考于该规范：https://rakyll.org/style-packages/

## 6.生成项目
从`Release`中下载可执行程序`ddd`，执行下述命令，将自动拉去模版项目并初始化：
```bash
ddd [项目名]
```
然后初始化为git仓库，并自行关联远程仓库即可：
```bash
git init
```
Example：
```bash
hezebin@MacBookPro go-projects % ./ddd test-ddd

Start to init project: test-ddd

Wait for the project template to be pulled from Git...
Cloning into '/Users/hezebin/Develop/go-projects/test-ddd'...

Organizing project files...
[Success]  test-ddd/README.md
[Success]  test-ddd/application/test.go
[Success]  test-ddd/cmd/root.go
[Success]  test-ddd/component/cache/memory.go
[Success]  test-ddd/component/cache/redis.go
[Success]  test-ddd/component/constant/commom.go
[Success]  test-ddd/component/doc/doc.go
[Success]  test-ddd/component/doc/swagger.json
[Success]  test-ddd/component/email/email.go
[Success]  test-ddd/component/pubsub/pulsar.go
[Success]  test-ddd/component/sms/sms.go
[Success]  test-ddd/component/storage/mongo.go
[Success]  test-ddd/component/storage/mysql.go
[Success]  test-ddd/config/config.go
[Success]  test-ddd/config/config.json
[Success]  test-ddd/domain/entity/test.go
[Success]  test-ddd/domain/repository/impl/mongo/base.go
[Success]  test-ddd/domain/repository/impl/mongo/test.go
[Success]  test-ddd/domain/repository/impl/redis/test.go
[Success]  test-ddd/domain/repository/test.go
[Success]  test-ddd/domain/service/test.go
[Success]  test-ddd/go.mod
[Success]  test-ddd/main.go
[Success]  test-ddd/script/index.js
[Success]  test-ddd/script/test.js
[Success]  test-ddd/script/test.py
[Success]  test-ddd/server/handler/test.go
[Success]  test-ddd/server/middleware/cors.go
[Success]  test-ddd/server/dto/test.go
[Success]  test-ddd/server/server.go
[Success]  test-ddd/worker/timer.go

Init project success!
```