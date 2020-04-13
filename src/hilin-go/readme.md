Bee go 环境配置：
1.配置环境变量 GOPATH（代码路径，先在里面建立src,pkg,bin3个目录），GOROOT:go安装的目录，go安装目录下的bin目录放到Path环境变量
可执行文件默认存放在 $GOPATH/bin 里面，所以您需要把 $GOPATH/bin 添加到您的环境变量中。go version查看安装版本

2.下载beego包并安装
go get github.com/astaxie/beego 

没安装git也没关系，就是会麻烦点，得先去https://github.com/astaxie/beego这里下载beego包，然后放在你的GOPATH路径下的src里pkg\windows_386\github.com\astaxie\beego

go get github.com/beego/bee

如出现错误，那就先执行 git config --global http.sslVerify false
然后再执行上面那个命令，应该就能成功安装bee了。同样，成功安装后，会在go/bin下生成一个bee.exe文件。
打开cmd，输入bee version，输出如下图的相应信息就安装成功了。

在cmd命令下，到GOPATH目录的src下，输入下面的命令
bee new hello

在cmd命令下，到GOPATH目录的src/hello下，输入下面的命令
bee run
然后beego已经运行起来啦，在浏览器输入 127.0.0.1:8080 就看到了运行结果（welcome to beego的欢迎页）。

注意：可以在 conf 下的 app.conf 文件修改默认的端口