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
		log.Println("starting download process...")
		u, err := c.GetUser(ctx, username)
		if err != nil {
			log.Fatalf("unable to lookup user: %v", err)
		}

		media := make([]of.Media, 0)
		ps, err := c.ListPhotos(ctx, fmt.Sprintf("%d", u.ID))
		if err != nil {
			log.Fatal(err)
		}
		for _, p := range ps {
			media = append(media, p.Media...)
		}

		hs, err := c.ListHighlights(ctx, u.ID)
		if err != nil {
			log.Fatal(err)
		}
		for _, h := range hs {
			for _, s := range h.Stories {
				media = append(media, s.Media...)
			}
		}

		ms, err := c.ListMessages(ctx, u.ID)
		if err != nil {
			log.Fatal(err)
		}
		for _, m := range ms {
			media = append(media, m.Media...)
		}

		posts, err := c.ListPosts(ctx, fmt.Sprintf("%d", u.ID))
		if err != nil {
			log.Fatal(err)
		}
		for _, p := range posts {
			media = append(media, p.Media...)
		}

		vs, err := c.ListVideos(ctx, u.ID)
		if err != nil {
			log.Fatal(err)
		}
		for _, v := range vs {
			media = append(media, v.Media...)
		}

		c.DownloadContent(ctx, media, u.Name, viper.GetString("save_dir"))
		log.Printf("downloaded %d files", len(media))
	},
}

func init() {
	downloadCmd.AddCommand(downloadContentCmd)
}
