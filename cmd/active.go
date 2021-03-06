// Copyright © 2016 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"

	"github.com/chanwit/belt/util"
	"github.com/spf13/cobra"
)

// activeCmd represents the active command
var activeCmd = &cobra.Command{
	Use:   "active",
	Short: "set or show active node",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		cluster, err := util.GetActiveCluster()
		if err != nil {
			return err
		}

		if len(args) == 0 {
			// GET
			node, err := util.GetActive()
			if err != nil {
				fmt.Println(err.Error())
				return nil
			}
			fmt.Println(cluster + "/" + node)
		} else if len(args) > 0 {
			// SET
			node := args[0]

			err := util.SetActive(node)
			if err != nil {
				fmt.Println(err.Error())
				return nil
			}
			fmt.Println(cluster + "/" + node)
		}

		return nil
	},
}

func init() {
	RootCmd.AddCommand(activeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// activeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// activeCmd.Flags().BoolP("cluster", "c", false, "set or get active cluster")

}
