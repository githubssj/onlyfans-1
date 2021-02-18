package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/januairi/go-of/pkg/of"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// downloadContentCmd represents the content command
var downloadContentCmd = &cobra.Command{
	Use:   "content",
	Short: "download all onlyfans content from a user",
	Long:  `download all onlyfans content from a user`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		c := of.NewClient(viper.GetString("token"), viper.GetString("session"), viper.GetString("user_agent"), viper.GetString("auth_id"))
		ctx := context.Background()
		username := args[0]
		u, err := c.GetUser(ctx, username)
		if err != nil {
			log.Fatalf("unable to lookup user: %v", err)
		}

		saveDir := viper.GetString("save_dir")

		ps, err := c.ListPhotos(ctx, fmt.Sprintf("%d", u.ID))
		if err != nil {
			log.Fatal(err)
		}
		for _, p := range ps {
			go c.DownloadContent(ctx, p.Media, u.Name, saveDir)
		}

		hs, err := c.ListHighlights(ctx, u.ID)
		if err != nil {
			log.Fatal(err)
		}
		for _, h := range hs {
			for _, s := range h.Stories {
				go c.DownloadContent(ctx, s.Media, u.Name, saveDir)
			}
		}

		ms, err := c.ListMessages(ctx, u.ID)
		if err != nil {
			log.Fatal(err)
		}
		for _, m := range ms {
			go c.DownloadContent(ctx, m.Media, u.Name, saveDir)
		}

		posts, err := c.ListPosts(ctx, fmt.Sprintf("%d", u.ID))
		if err != nil {
			log.Fatal(err)
		}
		for _, p := range posts {
			go c.DownloadContent(ctx, p.Media, u.Name, saveDir)
		}

		vs, err := c.ListVideos(ctx, u.ID)
		if err != nil {
			log.Fatal(err)
		}
		for _, v := range vs {
			go c.DownloadContent(ctx, v.Media, u.Name, saveDir)
		}
	},
}

func init() {
	downloadCmd.AddCommand(downloadContentCmd)
}
