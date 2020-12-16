
# Kit Project Layout
kit 项目必须具备的特点:
- 统一
- 标准库方式布局
- 高度抽象
- 支持插件

```
cache
    memcache
        test
    redis
        test
conf
    dsn
    env
    flagvar
    paladin
        appollo
            internal
                mockserver
container
    group
    pool
    queue
        aqm
database
    hbase
    sql
    tidb
ecode
    types
log
    internal
        core
        filewritee
```


# Service Application Project
- VO 视图对象，用于展示层，它的作用是把某个指定页面（或组件）的所有数据封装起来。
- DTO(Data Transfer Object) ：数据传输对象
- DO(Domain Object):  领域对象
- PO(Persistent Object):  持久化对象

app 目录下有 api、cmd、configs、internal 目录，目录里一般还会放置 README、CHANGELOG、OWNERS。
- internal:  是为了避免有同业务下有人跨目录引用了内部的 biz 、data 、 service  等内部 struct 。
- biz:  业务逻辑的组装层，类似 DDD  的 domain  层， data  类似 DDD的 repo ， repo  接口在这里定义，使用依赖倒置的原则。
- data:  业务数据访问，包含 cache 、 db  等封装，实现了 biz  的 repo接口。我们可能会把 data  与 dao  混淆在一起， data  偏重业务的含义，它所要做的是将领域对象重新拿出来，我们去掉了 DDD  的 infra 层。
- service:  实现了 api  定义的服务层，类似 DDD  的 application  层，处理 DTO  到 biz  领域实体的转换 (DTO -> DO) ，同时协同各类 biz  交互，但是不应处理复杂逻辑。

DDD 贫血模型

https://github.com/facebook/ent

```
CHANGELOG
OWENERS
README
api
cmd
    myapp1-admin
    myapp1-interface
    myapp1-job
    myapp1-service
    myapp1-task
configs
go.mod
internal
    biz
    data
    pkg
    service
```

# Lifecycle
Lifecycle 需要考虑服务应用的对象初始化以及生命周期的管理，所有 HTTP/gRPC 依赖的前置资源初始化，包括data、biz、service，之后再启动监听服务。我们使用https://github.com/google/wire ，来管理所有资源的依赖注入。


核心是为了： 1 、方便测试； 2 、单次初始化和复用；


# grpc

IDL
元数据

不要过早关注性能问题, 先标准化

lgtm
https://github.com/protocolbuffers/protobuf/blob/master/src/google/protobuf/wrappers.proto
https://www.bookstack.cn/read/API-design-guide/API-design-guide-04-%E6%A0%87%E5%87%86%E6%96%B9%E6%B3%95.md


# 配置管理
配置初始化 与 对象初始化

```
func main() {
  // load config file from yaml.
  c := new(redis.Config)
  _ = ApplyYAML(c, loadConfig())
  r, _ := redis.Dial(c.Network, c.Address, Options(c)...)
}
```


# test

mock fake stub 混沌测试

Google 测试之道
微软 测试

blog.golang.org/subtests

yapi

CICD

架构整洁之道
领域驱动设计


做好防御

简单善良
没有永远的敌人没有永远的朋友
做了什么说了什么
向上汇报 换位思考
对事不对人


# 作业

按照自己的构想，写一个项目满足基本的目录结构和工程，
代码需要包含对数据层、业务层、API 注册，以及 main 函数对于服务的注册和启动，信号处理，使用 Wire 构建依赖。
可以使用自己熟悉的框架。

https://github.com/Go-000/Go-000/issues/76