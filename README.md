# GolangStudy #
学习golang的每一步

本人电脑系统：windows10


# Windows #

## golang安装开发 ##

### 安装golang  ###
[下载golang](https://studygolang.com/dl)

[本文选择下载的版本>>>](https://studygolang.com/dl/golang/go1.19.3.windows-amd64.msi)

### 配置环境 ###
如果选择的msi方式安装则自动会配置环境，源码安装需配置环境  

配置GOROOT和GOPATH环境变量,GOROOT是go的安装目录,GOPATH是代码的位置  
1. 右键点击我的电脑 => 高级系统配置 => 环境变量配置  
2. 添加GOPATH参数，值为代码位置  
3. 将GOROOT添加到Path中  

### 验证 ###
win+R，输入cmd打开窗口

使用go env如果有输出则安装成功


## 学习教程 ##  
[golang文档教程](http://docscn.studygolang.com/doc/)
会对其中的内容进行一些讲解，自己看文档也可以  


本人会在基础教程上进行基础+底层的模式进行学习教程，希望可以更好的进行学习  

### 第一章 认识Golang ###  
历史人物：Go是从2007年末由Robert Griesemer, Rob Pike, Ken Thompson主持开发，后来还加入了Ian Lance Taylor, Russ Cox等人，并最终于2009年11月开源，在2012年早些时候发布了Go 1稳定版本。现在Go的开发已经是完全开放的，并且拥有一个活跃的社区。


编辑工具：VSCODE，GoLand和subline Text  

如何执行go程序？  
如：test.go文件  
第一种： go run命令执行  
go run test.go  获取执行结果  

第二种：go build 生成编译文件，然后执行编译文件，获取执行结果  
go build test.go  
./test 获取执行结果  

[第一课：Hello World](https://github.com/zggxuexihao/GolangStudy/blob/main/HelloWorld.go)  

[第二课：基础语法](https://github.com/zggxuexihao/GolangStudy/blob/main/Grammar.go)  

[第三课：基础数据类型](https://github.com/zggxuexihao/GolangStudy/blob/main/DataType.go)  

[第四课：数据类型底层结构]()  

[第五课：Go错误处理]()  

[第六课：Go并发]()