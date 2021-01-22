package cmd

import (
	"context"
	"encoding/json"
	"log"

	"github.com/januairi/go-of/pkg/of"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// getUserCmd represents the user command
var getUserCmd = &cobra.Command{
	Use:   "user",
	Short: "Get user information",
	Long:  `Get user information`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		c := of.NewClient(viper.GetString("token"), viper.GetString("session"), viper.GetString("user_agent"), viper.GetString("auth_id"))
		ctx := context.Background()

		user, err := c.GetUser(ctx, args[0])
		if err != nil {
			log.Fatal(err)
		}

		b, err := json.Marshal(user)
		if err != nil {
			log.Fatal(err)
		}

		log.Println(string(b))
	},
}

func init() {
	getCmd.AddCommand(getUserCmd)
}
