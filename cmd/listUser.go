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
    "os"
    "encoding/json"
	"github.com/spf13/cobra"
)

// listUserCmd represents the listUser command
var listUserCmd = &cobra.Command{
	Use:   "listUser",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		f,err := os.OpenFile("./entity/curUser.txt", os.O_RDONLY,0060)
        if err != nil{
            fmt.Println(err)
            os.Exit(1)
        }
        buf := make([]byte,1024)
        n, _ := f.Read(buf)
        if n == 0{
            fmt.Println("请先登录！")
            os.Exit(2)
        }
        f.Close()
        
        //打开json打印信息
        f,err = os.OpenFile("./entity/user.json", os.O_RDONLY, 0060)
        if err != nil{
            fmt.Println(err)
            os.Exit(1)
        }
        n,_ = f.Read(buf)
        
        var u UserList
        jsonReadErr := json.Unmarshal(buf[:n], &u)
        if jsonReadErr != nil{
            fmt.Println(err)
            os.Exit(2)
        }
        for i:=0; i < len(u.Users); i++{
            fmt.Printf("#%d {Name:%s, Email:%s, Phone:%s}\n",i+1,u.Users[i].Name, u.Users[i].Email, u.Users[i].Phone)
        }
        

        
	},
}

func init() {
	rootCmd.AddCommand(listUserCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listUserCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listUserCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
