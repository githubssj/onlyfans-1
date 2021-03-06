package cmd

import (
	"context"
	"log"

	"github.com/januairi/go-of/pkg/of"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// downloadHighlightCmd represents the highlights command
var downloadHighlightCmd = &cobra.Command{
	Use:   "highlight",
	Short: "download highlights from a user",
	Long:  `download highlights from a user`,
	Run: func(cmd *cobra.Command, args []string) {
		c := of.NewClient(viper.GetString("token"), viper.GetString("session"), viper.GetString("user_agent"), viper.GetString("auth_id"))
		username := args[0]
		ctx := context.Background()
		log.Println("starting download process...")
		u, err := c.GetUser(ctx, username)
		if err != nil {
			log.Fatalf("unable to lookup user: %v", err)
		}

		hs, err := c.ListHighlights(ctx, u.ID)
		if err != nil {
			log.Fatal(err)
		}

		media := make([]of.Media, 0)
		for _, h := range hs {
			for _, s := range h.Stories {
				media = append(media, s.Media...)
			}
		}

		c.DownloadContent(ctx, media, u.Name, viper.GetString("save_dir"))
	},
}

func init() {
	downloadCmd.AddCommand(downloadHighlightCmd)
}
