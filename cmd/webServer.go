/*
Copyright © 2024 Anurag Yadav <contact@anuragyadav.in>
*/

package cmd

import (
	"github.com/spf13/cobra"
	"go-jwt/pkg/web"
	"os"
)

// webServerCmd represents the webServer command
var webServerCmd = &cobra.Command{
	Use:   "webServer",
	Short: "Start the web-server",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		port := os.Getenv("WEB_SERVER_PORT")
		web.StartWebServer(port)
	},
}

func init() {
	rootCmd.AddCommand(webServerCmd)

}
