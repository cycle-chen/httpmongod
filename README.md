# httpmongod
httpmongod

http方式的mongod服务,提供增删改查的仿mongoshell命令接口

启动mongod，再启动httpmongod

然后在你的浏览器地址栏试试输入
``` 
http://localhost:8090/mongo.show dbs

```
![](https://github.com/golangdeveloper/httpmongod/blob/master/img/example.png?raw=true)

打开chrome的postman工具
``` 
post http://localhost:8090/mongo/DB/C.save(&httprequestbody)
```
![](https://github.com/golangframework/httpmongo/blob/master/exampleimg/savebybody.png)
## 引用库来源

[https://github.com/golangframework/httpmongo](https://github.com/golangframework/httpmongo)
