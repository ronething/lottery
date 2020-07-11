### 效果

```sh
# ronething @ ashings-macbook-pro in ~/Documents/lottery/_demo/2ticket on git:master x [11:10:32] 
$ curl "http://127.0.0.1:8080"                                                             
尾号为1获得一等奖<br/>尾号为2或3获得二等奖<br/>尾号为4/5/6获得三等奖<br/>code=3<br/>恭喜你获得: 二等奖%                                                                                                                    

# ronething @ ashings-macbook-pro in ~/Documents/lottery/_demo/2ticket on git:master x [11:10:51] 
$ curl "http://127.0.0.1:8080/prize"                                                       
今日开奖号码是: [20 7 18 24 25 31 12]%   
```