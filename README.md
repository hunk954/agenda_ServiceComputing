# CLI 命令行实用程序开发实战 - Agenda
17级软件工程服务计算 17308154 王俊焕
## 一、实验目的
1. 熟悉 go 命令行工具管理项目
2. 综合使用 go 的函数、数据结构与接口，编写一个简单命令行应用 agenda
3. 使用面向对象的思想设计程序，使得程序具有良好的结构命令，并能方便修改、扩展新的命令,不会影响其他命令的代码

## 二、实验环境
1. 虚拟机CentOS-7: go version go1.11.5 linux/amd64
2. [go-online](http://www.go-online.org.cn:8080): go version go1.12.7 linux/amd64

## 三、实验步骤
### 1. 安装cobra
- 使用命令`go get -v github.com/spf13/cobra/cobra`
  - **虚拟机**上使用出现报错。
  ```
  Fetching https://golang.org/x/sys/unix?go-get=1
  https fetch failed: Get https://golang.org/x/sys/unix?go-get=1: dial tcp 216.239.37.1:443: i/o timeout
  ``````
  - 根据课程指引，在 $GOPATH/src/golang.org/x 目录下用git clone下载sys和text项目。即
  ```
  git clone https://github.com/golang/sys
  git clone https://github.com/golang/text
  ``````
  - 然后使用 go install github.com/spf13/cobra/cobra, 安装后在 $GOBIN 下出现了 cobra 可执行程序。  
  ![虚拟机安装成功](/img/虚拟机安装cobra.png)
  - **go-online**上可以直接使用该命令完成安装
  
### 2. 简单使用cobra
1. 使用命令`cobra -help`可以打印使用指南
2. 进入到工作目录，创建我们的应用程序`cobra init --pkg-name=agenda`  
- 新版本的cobra中，--pkg-name已经不再是可缺省的参数，需要我们手动加入。使用该命令后，我们就会看到目录下多了LICENSE、cmd文件夹以及main.go。 

  ![init](/img/init.png)
- **在go-online中**由于存在多个GOPATH的问题，安装完之后我们输入cobra会显示没有该命令，这是因为我们此处调用cobra命令时，系统去寻找的程序bin目录不是cobra安装的目录。为了减少不必要的麻烦，所以此处我的方法是首先进入工作目录，然后使用相对路径到cobra所在的bin目录进行使用（或者我们可以定义环境变量$BIN，指向我们需要的bin目录，以使之后更加顺畅的调用cobra）**需要先进入工作目录，因为cobra是自动将文件生成在当前所在目录**  
![cobra1](/img/online-cobra1.png)  
![cobra2](/img/online-cobra2.png)

3. 创建子命令`cobra add [func]`，之后就会在/cmd下自动生成对应的[func].go  
4. 之后在/cmd/[func].go中的init()设置参数flag，在var [func]Cmd的Run属性中的匿名函数中加入对应实现即可  
![register1](/img/register1.png)
5. 使用`go run main.go [func] [temp]`即可实现对应的功能(当然也可以先go install，之后`agenda [func] [temp]`调用)  
![register2](/img/register2.png)  

参考文档：
- [官方文档](https://github.com/spf13/cobra#overview)
- [【中文】golang命令行库cobra的使用](https://www.cnblogs.com/borey/p/5715641.html)

### 3. 逐步完成需求
#### 1. 用户注册`/cmd/register.go`
- 参数（必填）：用户名、密码、邮箱、电话信息
- 需求：
    - 用户名具有唯一性，而且用户名对应唯一的密码
    - 注册成功和失败均有反馈的消息
- 存储用户信息使用User信息结构体，为了在json中得到保存，将用户信息存储在UserList结构体中
```go
    type User struct{
        Name string
        Password string
        Email string
        Phone string
    }
    type UserList struct{
        Users []User
    }
``````
- 设置参数：init()中使用`registerCmd.Flags().StringP([temp]...)`,其中StringP表示的是读取的参数类型是string，其中参数分别表示参数名字、参数缩写、缺省值、提示信息。
```go
	registerCmd.FLags().StringP("user","u","Anoymous","username of a user(required)")
	//...
``````
- 设置参数为必须填的属性：init()中使用`registerCmd.MarkFlagRequired([temp]...)`，其中参数填入参数名字
```go
	registerCmd.MarkFlagRequired("user")
	//...
``````
- 参数读取：在`var registerCmd = $cobra.Command{}`的Run属性中的匿名函数中使用`[name],err := cmd.Flags().GetString([temp]...)`。name为变量名称， 方便之后调用。参数中则填入需要读取的参数名称。
```go
    username,_ := cmd.Flags().GetString("user")
    //...
``````
- 代码逻辑：
	- 打开user.json，将json转化为所需结构体
    ```go
    	f,err := os.OpenFile("./entity/user.json",os.O_RDONLY,0060)
    	//...对err进行nil判断（略）
    	var u UserList
        buf := make([]byte,2048)
        n,_ := f.Read(buf)
        if n == 0{
            //json为空，不需要进行检测
        }else{
            jsonReadErr := json.Unmarshal(buf[:n], &u)
            //开始处理
        }
    ```
	- 用户名合法性检测：判断所输入的用户名与json文件中用户名是否重复
	```go
		 for i := 0; i < len(u.Users); i++{
        	 if username == u.Users[i].Name{
         		fmt.Println("用户名已被注册")
         		os.Exit(3)
         	}
		}
	```
	- 用户名合法，更新json：此处实现是重新打开user.json，此时读写属性是os.O_TRUNC(表示覆盖写入)。将新用户加入到维护的UserList中，之后转化为json序列，写入json文件中
	```go
        f,err = os.OpenFile("./entity/user.json",os.O_WRONLY|os.O_TRUNC,0060)
        //...err是否nil的检测（略）
        u.Users = append(u.Users,User{Name:username,Password:password,Email:email,Phone:phone})
        buffer, jsonWriteErr := json.Marshal(u)
        //...jsonWriteErr是否nil的检测（略）
        _, writingErr := f.Write(buffer)
        //...writingErr是否nil的检测（略）
	```
#### 2. 用户登录 `/cmd/login.go`
- 参数（必填）：用户名、密码
- 需求：
    - 用户名和登录需要匹配才可以登录
    - 如果已有用户登录，则无法登陆
    - 无论登录是否成功均反馈信息
- 代码逻辑：
	- 通过curUser.txt是否为空，判断是否已有用户登录：
	- 用户与密码是否匹配检测：
	- 写入curUser.txt表示登录
#### 3. 用户登出 `/cmd/logout.go`
- 参数：无
- 需求：
	- 若没有用户登录则无法进行登出
	- 必要的反馈信息
- 代码逻辑：
	- 通过curUser.txt是否为空，判断是否有用户在登录
	- 登出：清空curUser.txt
#### 4. 用户查询 `listUser.go`
- 参数：无
- 需求：
	- 只有用户登录了才可以进行查询
	- 打印所有用户的用户名、邮箱、电话
- 代码逻辑：
	- 通过curUser.txt是否为空，判断是否用户在登录
	- 打印信息：user.json中除了password的信息打印
#### 5. 用户注销 `/cmd/logoff.go'
- 参数：无
- 需求：
	- 只有登录的用户才可以删除
	- 删除后同时登出
	- 必要的反馈信息
- 代码逻辑：
	- 判断是否用户在登陆
	- 在user.json中删除该用户
	- 清空curUser.txt，登出

## 四、实验结果
### 1. 用户注册

### 2. 用户登录
### 3. 用户登出
### 4. 用户查询
### 5. 用户注销

## 五、实验感想与思考
这个实验建立在cobra库下，使命令行应用程序的实现变得更加简单，让我们可以把重心放在用户需求的代码实现上。除此之外，还让我对golang的文件读写更加娴熟，并初步了解了json序列化和反序列化的方法，让我能够简单的实现用户数据的存储。当然，我也能感受到我现在所实现的程序的局限性，比如输入密码时明文回显，数据存储中密码也是明文存储。在邮箱、电话的输入中没有正则表达式的检验。现实生活中，我们肯定不会遇到这样做的系统应用。还有就是，这次所做的应用有许多冗余的代码（比如检验用户是否已经登录的检测上反复调用相同的代码，其实应该可以考虑直接设置一个BOOL变量进行多个go文件之间的信息传输，不知道可不可行），以及许多err的变量命名不规范等。对于会议方面的命令，因为时间不够充分就没有进行实现。其实在软件工程初级实训中就实现了类似的功能，不过相比起c++实现，golang中应该有许多方便的库对时间进行判断，减少代码书写的时间。
