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
		u, err := c.GetUser(ctx, username)
		if err != nil {
			log.Fatalf("unable to lookup user: %v", err)
		}

		ps, err := c.ListPhotos(ctx, fmt.Sprintf("%d", u.ID))
		if err != nil {
			log.Fatal(err)
		}

		base := viper.GetString("save_dir")
		for _, p := range ps {
			for _, m := range p.PhotoMedia {
				err = c.DownloadFile(ctx, m, u.Username, base)
				if err != nil {
					log.Println(err)
				}
			}
		}
	},
}

func init() {
	downloadCmd.AddCommand(downloadPhotoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// downloadPhotoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// downloadPhotoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
