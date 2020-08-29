package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "serve",
	Short: "Starts golang-api server",
	Long: `serve' Starts golang-api server, serving 
	 my sample API`,
	Run: runServer,
}

func init() {
	rootCmd.AddCommand(serverCmd)
}

func runServer(cmd *cobra.Command, args []string) {
	fmt.Print("Hello world")
}
