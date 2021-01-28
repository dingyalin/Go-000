
# 隔离

# 超时
快速失败

# 过载保护
令牌桶  允许爆发  
漏桶    流量整形  流量控制

https://pkg.go.dev/golang.org/x/time/rate

Codel + tcp bbr

cpu tick 

滑动均值 [权重]  EMV算法

强化学习

# 限流
分布式限流

异步批量获取quota

Max-Min Fairness

重要性服务等级

# 熔断

客户端

Google SRE

K 默认 2
熔断概率 max(0, (requests -K*accepts)/(requests + 1))

positive feedback

backoff & jitter

Gutter kafka 双熔断


scaling memcached at facebook


# 关键字
false sharing

Head-of-line blocking  http2.0 --> http3.0 quic

h256 webp ijkplayer sharpP  编节码

tengine + cornet quic库

docker CPU Throttleing 隔离性 安全容器

cpu set

58 沈剑

pre-request

doorman

load bound hash

Consistent Hashing with Bounded Loads  有限负载一致性哈希

ketama hash


interface mock

logging

https://dave.cheney.net/2015/11/05/lets-talk-about-logging
https://www.ardanlabs.com/blog/2013/11/using-log-package-in-go.html
https://www.ardanlabs.com/blog/2017/05/design-philosophy-on-logging.html
https://dave.cheney.net/2017/01/23/the-package-level-logger-anti-pattern

@黄乃文 可以看看这个，https://itician.org/pages/viewpage.action?pageId=1114345 
https://pkg.go.dev/github.com/jinzhu/copier


# 参考
- [DDD 中的那些模式 — CQRS](https://zhuanlan.zhihu.com/p/115685384)
- [false sharing](https://www.cnblogs.com/cyfonly/p/5800758.html)
- [吴恩达机器学习]https://www.bilibili.com/video/BV164411S78V?p=6&t=12
- [cpu过载限流 RollingCounter](https://github.com/go-kratos/kratos/tree/master/pkg/stat/metric)
- [backoff and jitter](https://aws.amazon.com/blogs/architecture/exponential-backoff-and-jitter/)
- [sre](https://landing.google.com/sre/books/)
- [www.letmeread](https://www.letmeread.net/)
- [logger](https://dave.cheney.net/2017/01/23/the-package-level-logger-anti-pattern)
- [Hystrix 滑动窗口](https://blog.csdn.net/liubenlong007/article/details/86613317)


# 图书
- SRE google运维解密
- SRE google工作手册
- the-site-reliability-workbook 2
- 代码简洁之道
- Google软件测试之道


红黑树
动态规划
最小堆
