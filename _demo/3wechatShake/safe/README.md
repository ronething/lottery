### 加锁

```sh
# ronething @ ashings-macbook-pro in /tmp/log [12:07:37]
$ wc -l lottery_wechat.log
       0 lottery_wechat.log

# ronething @ ashings-macbook-pro in /tmp/log [12:07:40]
$ wrk -t10 -c10 -d5 http://127.0.0.1:8080/lucky

Running 5s test @ http://127.0.0.1:8080/lucky
  10 threads and 10 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     2.00ms    9.94ms 121.18ms   96.98%
    Req/Sec     2.49k   779.55     3.51k    70.40%
  124012 requests in 5.01s, 18.20MB read
Requests/sec:  24749.16
Transfer/sec:      3.63MB

# ronething @ ashings-macbook-pro in /tmp/log [12:07:47]
$ wc -l lottery_wechat.log
   20000 lottery_wechat.log
```