# 简易区块链模型

## 环境搭建

## 测试
在需要测试的文件夹下运行:
go test -v .

-v 无论是否成功都显示输出

##Swagger 生成
在接口函数上 填写注释,
然后在根目录运行 swag init   更新文档

参数说明
```
@Summary 接口概要说明
@Description 接口详细描述信息
@Tags 用户信息   //swagger API分类标签, 同一个tag为一组
@accept json  //浏览器可处理数据类型，浏览器默认发 Accept: */*
@Produce  json  //设置返回数据的类型和编码
@Param id path int true "ID"    //url参数：（name；参数类型[query(?id=),path(/123)]；数据类型；required；参数描述）
@Param name query string false "name"
@Success 200 {object} Res {"code":200,"data":null,"msg":""}  //成功返回的数据结构， 最后是示例
@Failure 400 {object} Res {"code":200,"data":null,"msg":""}
@Router /test/{id} [get]    //路由信息，一定要写上
```


@Param
```参数，表示需要传递到服务器端的参数，有五列参数，使用空格或者 tab 分割，五个分别表示的含义如下
1.参数名
2.参数类型，可以有的值是 formData、query、path、body、header
        formData 表示是 post 请求的数据，
        query 表示带在 url 之后的参数，
        path 表示请求路径上得参数，例如上面例子里面的 key，
        body 表示是一个 raw 数据请求，
        header 表示带在 header 信息中得参数。

3.参数类型
4.是否必须
5.注释```


##mod 依赖包相关

下载包 : go get "包的github地址"

一般为了防止包的更新以及版本问题, 在编译之前把所有先关包都导入到当前目录中
go mod vendor
