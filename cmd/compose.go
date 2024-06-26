/*
Copyright © 2024 Ignacio Chalub <ignaciochalub@gmail.com> & Federico Pochat <federicopochat@gmail.com>
*/

package cmd

import (
	"github.com/chalet/cli/logger"
	"github.com/chalet/cli/utils"
	"github.com/spf13/cobra"
)

var componesCmd = &cobra.Command{
	Use:   "compose",
	Short: "",
	Long: `
	
	`,
	Run: func(cmd *cobra.Command, args []string) {
		err := composeHandler()
		if err != nil {
			logger.Error(err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(componesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// execCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// execCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

/*
	Compose command 
	1. Get root repository name: network name will be REPONAME_network
	2. Run app container connecting to the network
	3. If network doesn't exist, run docker compose and then run container app again
	4. Other case: docker compose sets different name for the network -> Option in chalet.yml to set network name
*/

func composeHandler() error {
	_, err := utils.ReadConfig()
	if err != nil {
		return err
	}
	return nil
}