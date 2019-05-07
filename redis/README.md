# redis 的一些用法(github.com/go-redis)

### 1. redis 叠加 lua 脚本

- 简介

通过内嵌对 Lua 环境的支持， Redis 解决了长久以来不能高效地处理 CAS （check-and-set）命令的缺点， 并且可以通过组合使用多个命令， 轻松实现以前很难实现或者不能高效实现的模式。

### 2. lua 脚本示例

```go
// lua 脚本里默认会传 KEY、ARGV(value) 数组
// redis.call() 是调用 redis 脚本
// zrangebyscore 是根据指定的条件来设置排序规则,见下
// ZRANGEBYSCORE key min max [WITHSCORES] [LIMIT offset count]
// -inf 是从负无穷大开始(指底部)
func init() {
	popJobsLuaScript = redis.NewScript(`
		local name = ARGV[1]							# value 数组第一个
		local timestamp = ARGV[2]					# value 数组第二个
    local limit = ARGV[3]					    # value 数组第三个
		local results = redis.call('zrangebyscore', name, '-inf', timestamp, 'LIMIT', 0, limit)
		if table.getn(results) > 0 then
			redis.call('zrem', name, unpack(results))
		end
        return results`)
}
// 功能释义：通过排序返回第一个key value，然后unpack删除得到的key
```



 

