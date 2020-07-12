### 效果

```sh
$ ./3wechatShake
&{1 iphone4   4  [] 2 2 true 2 0 2}
&{2 apple   3  [] 5 5 true 100 2 102}
&{3 商城满2000元减50元优惠券   2 mall-coupon-2018 [] 50 50 true 2000 102 2102}
&{4 商城无门槛直降50元优惠券   1  [c01 c02 c03 c04 c05] 5 5 true 2000 2102 4102}
&{5 社区10个金币   0 10 [] 5 5 true 5000 4102 9102}
Now listening on: http://localhost:8080
Application started. Press CMD+C to shut down.


## curl
$ curl "http://127.0.0.1:8080/lucky"
{
 "data": "10",
 "id": 5,
 "link": "",
 "name": "社区10个金币",
 "success": true
}

# ronething @ ashings-macbook-pro in /tmp/log [11:46:32]
$ curl "http://127.0.0.1:8080/lucky"
{
 "data": "mall-coupon-2018",
 "id": 3,
 "link": "",
 "name": "商城满2000元减50元优惠券",
 "success": true
}

# ronething @ ashings-macbook-pro in /tmp/log [11:47:09]
$ curl "http://127.0.0.1:8080/lucky"
{
 "success": false
}
```

### wrk 压力测试

```
Usage: wrk <options> <url>
  Options:
    -c, --connections <N>  Connections to keep open
    -d, --duration    <T>  Duration of test
    -t, --threads     <N>  Number of threads to use

    -s, --script      <S>  Load Lua script file
    -H, --header      <H>  Add header to request
        --latency          Print latency statistics
        --timeout     <T>  Socket/request timeout
    -v, --version          Print version details
```

```sh
# ronething @ ashings-macbook-pro in /tmp/log [11:55:15]
$ wrk -t10 -c10 -d5 http://127.0.0.1:8080/lucky

Running 5s test @ http://127.0.0.1:8080/lucky
  10 threads and 10 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   495.01us  770.60us  32.06ms   94.13%
    Req/Sec     2.67k   505.13     3.64k    67.19%
  135316 requests in 5.10s, 19.76MB read
Requests/sec:  26516.00
Transfer/sec:      3.87MB

# ronething @ ashings-macbook-pro in /tmp/log [11:55:24]
$ wc -l lottery_wechat.log
   20051 lottery_wechat.log
```

giftList 中已经将 iphone4 奖品设置为 total: 20000, left: 20000, rate: 10000, 并且其他 gift 均设置 inuse = false

所以抽奖结果每次都是中奖 iphone4 并且最多只能有 20000 行获奖记录

但是日志记录超过了 20000 行，不符合我们的预期,存在线程不安全的情况