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
  - 虚拟机上使用出现报错。
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
  - go-online上可以直接使用该命令完成安装
  
参考文档：
- [官方文档] 
- [【中文】golang命令行库cobra的使用]


