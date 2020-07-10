### 效果

```sh
# ronething @ ashings-macbook-pro in ~/Documents/lottery on git:master x [23:36:13]
$ curl http://127.0.0.1:8080
当前总共参与抽奖的用户数：0

# ronething @ ashings-macbook-pro in ~/Documents/lottery on git:master x [23:36:18]
$ curl -X POST "http://127.0.0.1:8080/import?users=ashing,panda"
当前总共参与抽奖的用户数：2，成功导入用户数：2

# ronething @ ashings-macbook-pro in ~/Documents/lottery on git:master x [23:37:18]
$ curl -X POST "http://127.0.0.1:8080/import?users=xzx"
当前总共参与抽奖的用户数：3，成功导入用户数：1

# ronething @ ashings-macbook-pro in ~/Documents/lottery on git:master x [23:37:29]
$ curl "http://127.0.0.1:8080/lucky"
当前中奖用户：xzx，剩余用户数：2

# ronething @ ashings-macbook-pro in ~/Documents/lottery on git:master x [23:38:15]
$ curl "http://127.0.0.1:8080/lucky"
当前中奖用户：panda，剩余用户数：1

# ronething @ ashings-macbook-pro in ~/Documents/lottery on git:master x [23:38:18]
$ curl "http://127.0.0.1:8080/lucky"
当前中奖用户：ashing，剩余用户数：0

# ronething @ ashings-macbook-pro in ~/Documents/lottery on git:master x [23:38:19]
$ curl "http://127.0.0.1:8080/lucky"
已经没有参与用户，请先通过 /import 导入用户
```