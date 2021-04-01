package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/januairi/go-of/pkg/of"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// downloadPhotoCmd represents the photo command
var downloadPhotoCmd = &cobra.Command{
	Use:   "photo",
	Short: "download photos from a user",
	Long:  `download photos from a user`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		c := of.NewClient(viper.GetString("token"), viper.GetString("session"), viper.GetString("user_agent"), viper.GetString("auth_id"))
		username := args[0]
		ctx := context.Background()
		log.Println("starting download process...")
		u, err := c.GetUser(ctx, username)
		if err != nil {
			log.Fatalf("unable to lookup user: %v", err)
		}
		ps, err := c.ListPhotos(ctx, fmt.Sprintf("%d", u.ID))
		if err != nil {
			log.Fatal(err)
		}

		media := make([]of.Media, len(ps))
		for _, p := range ps {
			media = append(media, p.Media...)
		}

		c.DownloadContent(ctx, media, u.Name, viper.GetString("save_dir"))
	},
}

func init() {
	downloadCmd.AddCommand(downloadPhotoCmd)
}
