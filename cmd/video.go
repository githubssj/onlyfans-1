package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/januairi/go-of/pkg/of"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// downloadVideoCmd represents the video command
var downloadVideoCmd = &cobra.Command{
	Use:   "video",
	Short: "download videos from a user",
	Long:  `download videos from a user`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		c := of.NewClient(viper.GetString("token"), viper.GetString("session"), viper.GetString("user_agent"), viper.GetString("auth_id"))
		username := args[0]
		ctx := context.Background()
		u, err := c.GetUser(ctx, username)
		if err != nil {
			log.Fatalf("unable to lookup user: %v", err)
		}

		vs, err := c.ListVideos(ctx, fmt.Sprintf("%d", u.ID))
		if err != nil {
			log.Fatal(err)
		}

		base := viper.GetString("save_dir")
		for _, p := range vs {
			for _, m := range p.VideoMedia {
				err = c.DownloadFile(ctx, m, u.Username, base)
				if err != nil {
					log.Println(err)
				}
			}
		}
	},
}

func init() {
	downloadCmd.AddCommand(downloadVideoCmd)
}
