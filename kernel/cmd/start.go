package cmd

import (
	"batkor/restapp/routes"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

// startCmd Run web server.
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the server",
	Long:  `Run web server.`,
	Run: func(cmd *cobra.Command, args []string) {
		routerCollection := gin.Default()
		routes.UserRoutes(routerCollection)
		routerCollection.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
		err := routerCollection.Run()

		if err != nil {
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(startCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
