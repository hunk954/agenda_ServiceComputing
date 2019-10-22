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
type User struct{
    Name string
    Password string
    Email string
    Phone string
}
type UserList struct{
    Users []User
}

// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		username,_ := cmd.Flags().GetString("user")
        password,_ := cmd.Flags().GetString("password")
        email,_ := cmd.Flags().GetString("email")
        phone,_ := cmd.Flags().GetString("phone")
        //打开json文件
        f,err := os.OpenFile("./entity/user.json",os.O_RDONLY,0060)
        if err != nil{
            fmt.Println(err)
            os.Exit(1)
        }

        // 用户名合法性检测: 判断用户名是否重复
        var u UserList
        buf := make([]byte,2048)
        n,_ := f.Read(buf)
        if n == 0{
            //json为空，不需要进行检测
        }else{
            jsonReadErr := json.Unmarshal(buf[:n], &u)
            if jsonReadErr != nil{
                fmt.Println(jsonReadErr)
                os.Exit(2)
            }
            for i := 0; i < len(u.Users); i++{
                if username == u.Users[i].Name{
                    fmt.Println("用户名已被注册")
                    os.Exit(3)
                }
            }
        }
        f.Close()

        //写入json
        f,err = os.OpenFile("./entity/user.json",os.O_WRONLY|os.O_TRUNC,0060)
        defer f.Close()
        if err != nil{
            fmt.Println(err)
            os.Exit(4)
        }
        u.Users = append(u.Users,User{Name:username,Password:password,Email:email,Phone:phone})
        buffer, jsonWriteErr := json.Marshal(u)
        if jsonWriteErr != nil{
            fmt.Println(jsonWriteErr)
            os.Exit(5)
        }
        _, writingErr := f.Write(buffer)
        if writingErr != nil{
            fmt.Println(writingErr)
            os.Exit(6)
        }
        fmt.Printf("%s 注册成功!\n",username)
        
	},
}

func init() {
	rootCmd.AddCommand(registerCmd)
	registerCmd.Flags().StringP("user","u","Anoymous","username of a user(required)")
    registerCmd.Flags().StringP("password","p","123456","password of a user(required)")
    registerCmd.Flags().StringP("email","e","default@email.com","email of a user(required)")
    registerCmd.Flags().StringP("phone","k","123456789","phone nummber of a user(required)")
    registerCmd.MarkFlagRequired("user")
    registerCmd.MarkFlagRequired("password")
    registerCmd.MarkFlagRequired("phone")
    registerCmd.MarkFlagRequired("email")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// registerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// registerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
