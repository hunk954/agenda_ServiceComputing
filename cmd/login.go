/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
    "encoding/json"
    "os"
	"github.com/spf13/cobra"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		username,_ := cmd.Flags().GetString("user")
        password,_ := cmd.Flags().GetString("password")
        
        f,err := os.OpenFile("./entity/curUser.txt",os.O_RDONLY,0060)
        if err != nil{
            fmt.Println(err)
            os.Exit(1)
        }
        buf := make([]byte,2048)
        n,_ := f.Read(buf)
        if n != 0{
            fmt.Println("已有用户登录！")
            os.Exit(1)
        }
        f.Close()

        f,err = os.OpenFile("./entity/user.json",os.O_RDONLY,0060)
        if err != nil{
            fmt.Println(err)
            os.Exit(1)
        }
        // 用户名密码检测
        var u UserList
        buf = make([]byte,2048)
        n,_ = f.Read(buf)
        flag := false
        if n == 0{
            //json中没有数据那么就是默认不可登录
            fmt.Println("用户名不存在或密码错误！")
            os.Exit(2)
        }else{
            jsonReadErr := json.Unmarshal(buf[:n], &u)
            if jsonReadErr != nil{
                fmt.Println(jsonReadErr)
                os.Exit(3)
            }
            for i := 0; i < len(u.Users); i++{
                if username == u.Users[i].Name{
                    if password == u.Users[i].Password{
                        fmt.Println("登陆成功！")
                        loginUser := u.Users[i]
                        writingBuffer, _ := json.Marshal(loginUser)
                        fCurUser, _ := os.OpenFile("./entity/curUser.txt",os.O_WRONLY|os.O_TRUNC,0060)
                        defer fCurUser.Close()
                        _, writingErr := fCurUser.Write(writingBuffer)
                        if writingErr != nil{
                            fmt.Println(writingErr)
                            os.Exit(3)
                        }
                        flag = true
                    }else{
                        fmt.Println("用户名不存在或密码错误！")
                        os.Exit(2)
                    }
                }
            }
        }
        if !flag{
            fmt.Println("用户名不存在或密码错误！")
            os.Exit(2)
        }

	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
	loginCmd.Flags().StringP("user","u","Anoymous","username of a user(required)")
    loginCmd.Flags().StringP("password","p","123456","password of a user(required)")
    loginCmd.MarkFlagRequired("user")
    loginCmd.MarkFlagRequired("password")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loginCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
