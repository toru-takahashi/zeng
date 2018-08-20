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
	"fmt"
	"log"
	"os/exec"
	"runtime"
	"strconv"

	"github.com/spf13/cobra"
)

// previewCmd represents the preview command
var previewCmd = &cobra.Command{
	Use:   "preview",
	Short: "Preview your article on Zendesk",
	Long:  `TBD`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		articleID, err := strconv.Atoi(args[0])

		client, err := newDefaultClient()
		if err != nil {
			return
		}

		fmt.Println("Please push your article before preview")
		subPath := fmt.Sprintf("knowledge/api/articles/%d/translations/%s/preview/head", articleID, client.Locale)
		url := client.previewPath(subPath)
		openbrowser(url)
		return
	},
}

func init() {
	rootCmd.AddCommand(previewCmd)
}

func openbrowser(url string) {
	var err error
	fmt.Printf("Open: %s\n", url)

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}

}
