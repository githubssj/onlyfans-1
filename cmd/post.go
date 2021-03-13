package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/januairi/go-of/pkg/of"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// downloadPostcmd represents the download post command
var downloadPostcmd = &cobra.Command{
	Use:   "post",
	Short: "download post media from a user",
	Long:  `download post media from a user`,
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

		ps, err := c.ListPosts(ctx, fmt.Sprintf("%d", u.ID))
		if err != nil {
			log.Fatal(err)
		}

		saveDir := viper.GetString("save_dir")
		for _, p := range ps {
			c.DownloadContent(ctx, p.Media, u.Name, saveDir)
		}
	},
}

// downloadArchivedPostcmd represents the download post command
var downloadArchivedPostcmd = &cobra.Command{
	Use:   "post",
	Short: "download post media from a user",
	Long:  `download post media from a user`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		c := of.NewClient(viper.GetString("token"), viper.GetString("session"), viper.GetString("user_agent"), viper.GetString("auth_id"))
		username := args[0]
		ctx := context.Background()

		u, err := c.GetUser(ctx, username)
		if err != nil {
			log.Fatalf("unable to lookup user: %v", err)
		}

		ps, err := c.ListArchivedPosts(ctx, u.ID)
		if err != nil {
			log.Fatal(err)
		}

		saveDir := viper.GetString("save_dir")
		for _, p := range ps {
			c.DownloadContent(ctx, p.Media, u.Name, saveDir)
		}
	},
}

func init() {
	downloadCmd.AddCommand(downloadPostcmd)
	archivedCmd.AddCommand(downloadArchivedPostcmd)
}
