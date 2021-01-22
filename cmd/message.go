package cmd

import (
	"context"
	"log"

	"github.com/januairi/go-of/pkg/of"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// downloadMessageCmd represents the message command
var downloadMessageCmd = &cobra.Command{
	Use:   "message",
	Short: "get content from messages",
	Long:  `get content from messages`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		c := of.NewClient(viper.GetString("token"), viper.GetString("session"), viper.GetString("user_agent"), viper.GetString("auth_id"))
		username := args[0]
		ctx := context.Background()

		u, err := c.GetUser(ctx, username)
		if err != nil {
			log.Fatalf("unable to lookup user: %v", err)
		}

		ms, err := c.ListMessages(ctx, u.ID)
		if err != nil {
			log.Fatal(err)
		}

		for _, m := range ms {
			err = c.DownloadContent(ctx, m.Media, u.Name, viper.GetString("save_dir"))
			if err != nil {
				log.Fatal(err)
			}
		}
	},
}

func init() {
	downloadCmd.AddCommand(downloadMessageCmd)
}
