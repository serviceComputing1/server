# server
server's code 

## 安装指南
```
go get github.com/serviceComputing1/server
```
您可能需要需改包的路径以运行项目


## 项目小结
1. 我认为go 的json处理比起js差，可能是js原生支持json访问，在go里面都要经过结构体转换再解码或编码，我觉得很麻烦。
2.项目后端采用了 MVC模式，model部分负责处理query数据，controller部分负责处理路由请求，从model中获得数据，并根据不同的路由请求返回不同的数据。
3. 我觉得mux的路径访问对于optional paramater的参数不方便，比如路径中带问号的，就不是作为路径输入，而是作为Queries的参数输入，也可能是官方给的文档不够清楚，没有详细的说明可选的参数该怎么解析。
4. 后来看到一个很好的框架echo,觉得很不错，里面有jsonp和JWT验证，很方便。
5. 我觉得interface 的多态和继承还要好好学学，想根据传入的结构体类型的interface，传出对应的interface ，再断言为struct，但是不成功。
6. boltdb真的很简单，所以他没有sql的query功能。


## 复制swapi
1. 获取数据，预先人工获取了每个资源你的数量，然后每个循环请求，得到的json数据再存入boltdb中。
2. 由于只是设计name和 json整体数据，因此就没有再对json里面的字段单独用key存储，而是整个json为值，name为key。

