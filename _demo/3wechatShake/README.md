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