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
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	input "github.com/tcnksm/go-input"
)

// configureCmd represents the configure command
var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "Configure your Zendesk Account",
	Long:  `TBD`,
	Run: func(cmd *cobra.Command, args []string) {

		ui := &input.UI{
			Writer: os.Stdout,
			Reader: os.Stdin,
		}

		subdomain, err := ui.Ask("What is Zendesk Subdomain (subdomain.zendesk.com)?", &input.Options{
			Default:  "",
			Required: true,
			Loop:     true,
		})
		if err != nil {
			log.Fatal(err)
		}
		viper.Set("subdomain", subdomain)

		email, err := ui.Ask("What is your Zendesk Email?", &input.Options{
			Default:  "",
			Required: true,
			Loop:     true,
		})

		if err != nil {
			log.Fatal(err)
		}
		viper.Set("email", email)

		locale, err := ui.Ask("What is your default locale?", &input.Options{
			Default:  "en-us",
			Required: true,
			Loop:     true,
		})

		if err != nil {
			log.Fatal(err)
		}
		viper.Set("locale", locale)

		method, err := ui.Select("How do you prefer to acccess to Zendesk?", []string{"password", "apikey"}, &input.Options{
			Default:  "password",
			Required: true,
			Loop:     true,
		})

		if err != nil {
			log.Fatal(err)
		}

		if method == "password" {
			password, err := ui.Ask("What is your Zendesk password?", &input.Options{
				Default:     "",
				Required:    true,
				Mask:        true,
				MaskDefault: true,
				Loop:        true,
			})
			if err != nil {
				log.Fatal(err)
			}
			viper.Set("password", password)
		} else if method == "apikey" {
			apikey, err := ui.Ask("What is your Zendesk APIKEY?", &input.Options{
				Default:     "",
				Required:    true,
				Mask:        true,
				MaskDefault: true,
				Loop:        true,
			})
			if err != nil {
				log.Fatal(err)
			}
			viper.Set("apikey", apikey)
		}

		viper.WriteConfig()
	},
}

func init() {
	rootCmd.AddCommand(configureCmd)
}
