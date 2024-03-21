package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "generate responses",
	Short: "give me responses for pretty much everything",
	Long:  `A longer description of your application...`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello from YourProject!")
	},
}

func Execute() error {
	return rootCmd.Execute()
}
