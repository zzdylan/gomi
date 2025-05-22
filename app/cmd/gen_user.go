package cmd

import (
	"gomi/app/models/user"

	"github.com/spf13/cobra"
)

var CmdGenUser = &cobra.Command{
	Use:   "gen_user",
	Short: "Generate a user",
	Run:   runGenUser,
}

// 调试完成后请记得清除测试代码
func runGenUser(cmd *cobra.Command, args []string) {
	user := user.User{
		Username: "test",
		Password: "123456",
	}
	user.Create()
}
