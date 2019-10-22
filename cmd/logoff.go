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

// logoffCmd represents the logoff command
var logoffCmd = &cobra.Command{
	Use:   "logoff",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		 //检验是否登录
		 f,err := os.OpenFile("./entity/curUser.txt", os.O_RDONLY,0060)
		 if err != nil{
			 fmt.Println(err)
			 os.Exit(1)
		 }
		 buf := make([]byte,1024)
		 n, _ := f.Read(buf)
		 var username string
		 var curUser User
		 if n == 0{
			 fmt.Println("请先登录！")
			 os.Exit(2)
		 }else{
			 jsonReadErr := json.Unmarshal(buf[:n], &curUser)
			 if jsonReadErr != nil{
				 fmt.Println(jsonReadErr)
				 os.Exit(3)
			 }
			 username = curUser.Name
		 }
		 f.Close()
 
		 //注销账户
		 //从user.json中删除对应的条目
		 var u UserList
		 f,err = os.OpenFile("./entity/user.json", os.O_RDONLY,0060)
		 n, _=f.Read(buf)
		 jsonReadErr := json.Unmarshal(buf[:n], &u)
		 if jsonReadErr != nil{
			 fmt.Println(jsonReadErr)
			 os.Exit(3)
		 }
		 for i := 0; i < len(u.Users); i++{
			 if username == u.Users[i].Name{
				 u.Users = append(u.Users[:i],u.Users[i+1:]...)
				 break
			 }
		 }
		 f.Close()
		 f,err = os.OpenFile("./entity/user.json", os.O_WRONLY|os.O_TRUNC,0060)
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
		 //登出：清空curUser.txt
		 f.Close()
		 f,err = os.OpenFile("./entity/curUser.txt", os.O_TRUNC, 0060)
		 emptyStr := ""
		 emptyStrBuffer := []byte(emptyStr)
		 _,_ = f.Write(emptyStrBuffer)
		 f.Close()
		 fmt.Println("注销成功！")
	},
}

func init() {
	rootCmd.AddCommand(logoffCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// logoffCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// logoffCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
