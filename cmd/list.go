// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
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
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List articles/sections/categories",
	Long:  `TBD`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	listCmd.PersistentFlags().IntVarP(&o.Page, "page", "p", 1, "Page")
	listCmd.PersistentFlags().IntVarP(&o.Limit, "limit", "l", 20, "Per page limit")
	listCmd.PersistentFlags().StringVarP(&o.SortBy, "sortby", "", "position", "Sort By (position, title, created_at, updated_at)")
	listCmd.PersistentFlags().StringVarP(&o.SortOrder, "orderby", "", "asc", "Sort order by (asc, desc)")
	rootCmd.AddCommand(listCmd)
}
