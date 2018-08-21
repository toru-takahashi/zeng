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
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/olekukonko/tablewriter"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	yaml "gopkg.in/yaml.v2"
)

// sectionCmd represents the section command
var sectionGetCmd = &cobra.Command{
	Use:   "section",
	Short: "Get articles under Section",
	Long:  `TBD.`,
	Args:  cobra.MinimumNArgs(1),
	RunE:  runSectionGetCmd,
}

var sectionListCmd = &cobra.Command{
	Use:   "section",
	Short: "List sections",
	Long:  `TBD`,
	RunE:  runSectionListCmd,
}

func init() {
	sectionListCmd.Flags().IntVarP(&o.CategoryID, "category-id", "", 0, "Filter by Category ID")

	listCmd.AddCommand(sectionListCmd)
	getCmd.AddCommand(sectionGetCmd)
}

func runSectionListCmd(cmd *cobra.Command, args []string) error {
	client, err := newDefaultClient()
	if err != nil {
		return errors.Wrap(err, "newClient failed:")
	}

	req := SectionListRequest{
		Page:      o.Page,
		PerPage:   o.Limit,
		SortBy:    o.SortBy,
		SortOrder: o.SortOrder,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := client.SectionList(ctx, req)
	if err != nil {
		return errors.Wrapf(err, "StackShow was failed: req = %+v, res = %+v", req, res)
	}

	sections := res.Sections
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"category id", "category", "section id", "title", "url", "updated at"})
	for _, v := range sections {
		category := ""
		categoryID := 0

		for _, c := range res.Categories {
			if v.CategoryID == c.ID {
				category = v.Name
				categoryID = v.CategoryID
			}
		}

		table.Append([]string{
			fmt.Sprintf("%d", categoryID),
			category,
			fmt.Sprintf("%d", v.ID),
			v.Name,
			v.HTMLURL,
			v.UpdatedAt,
		})
	}
	table.Render()

	return nil
}

// SectionList
func (client *Client) SectionList(ctx context.Context, apiRequest SectionListRequest) (*SectionListResponse, error) {
	subPath := fmt.Sprintf("/%s/sections.json", client.Locale)

	if apiRequest.CategoryID != 0 {
		subPath = fmt.Sprintf("/%s/categories/%d/sections.json", client.Locale, apiRequest.CategoryID)
	}

	httpRequest, err := client.newRequest(ctx, "GET", subPath, nil)
	if err != nil {
		return nil, err
	}

	q := httpRequest.URL.Query()
	q.Add("page", fmt.Sprintf("%d", apiRequest.Page))
	q.Add("per_page", fmt.Sprintf("%d", apiRequest.PerPage))
	q.Add("sort_by", apiRequest.SortBy)
	q.Add("sort_order", apiRequest.SortOrder)
	q.Add("include", "categories")
	httpRequest.URL.RawQuery = q.Encode()
	fmt.Printf("Info: Request to %s\n", httpRequest.URL)

	httpResponse, err := client.HTTPClient.Do(httpRequest)
	if err != nil {
		return nil, err
	}

	// Check status code here…

	var apiResponse SectionListResponse
	if err := decodeBody(httpResponse, &apiResponse); err != nil {
		return nil, err
	}
	return &apiResponse, nil
}

// SectionGet
func (client *Client) SectionGet(ctx context.Context, apiRequest SectionGetRequest) (*SectionGetResponse, error) {
	subPath := fmt.Sprintf("/%s/sections/%d/articles.json", client.Locale, apiRequest.ID)

	httpRequest, err := client.newRequest(ctx, "GET", subPath, nil)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Info: Request to %s\n", httpRequest.URL)

	httpResponse, err := client.HTTPClient.Do(httpRequest)
	if err != nil {
		return nil, err
	}

	// Check status code here…

	var apiResponse SectionGetResponse
	if err := decodeBody(httpResponse, &apiResponse); err != nil {
		return nil, err
	}
	return &apiResponse, nil
}

func runSectionGetCmd(cmd *cobra.Command, args []string) error {
	client, err := newDefaultClient()
	if err != nil {
		return errors.Wrap(err, "newClient failed:")
	}

	sectionID, err := strconv.Atoi(args[0])
	if err != nil {
		return errors.Wrapf(err, "failed to parse SectionID: %s", args[0])
	}

	req := SectionGetRequest{
		ID: sectionID,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := client.SectionGet(ctx, req)
	if err != nil {
		return errors.Wrapf(err, "StackShow was failed: req = %+v, res = %+v", req, res)
	}

	articles := res.Articles

	if res.Count == 0 {
		fmt.Printf("%d articles in Section %d\n", res.Count, sectionID)
		return nil
	}

	// Structure
	// s_<section_id>_<locale>/
	//		a_<article id>/
	//				a_<article_id>_<locale>_<title>.html
	//				meta_<article_id>_<locale>.yaml
	//		a_<article id>/
	// 				a_<article_id>_<locale>_<title>.html
	//				meta_<article_id>.yaml
	dirPath := fmt.Sprintf("./s_%d_%s", sectionID, client.Locale)
	if err := os.MkdirAll(dirPath, 0777); err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Exported Section %d (Total %d articles) to %s\n", sectionID, res.Count, dirPath)

	for _, v := range articles {
		subPath := fmt.Sprintf("%s/a_%d", dirPath, v.ID)
		if err := os.Mkdir(subPath, 0777); err != nil {
			panic(err)
		}

		bodyPath := fmt.Sprintf("%s/a_%d_%s_%s.html", subPath, v.ID, v.Locale, v.Title)
		metaPath := fmt.Sprintf("%s/meta_%d_%s.html", subPath, v.ID, v.Locale)

		file, err := os.Create(bodyPath)
		if err != nil {
			return err
		}
		defer file.Close()
		file.Write(([]byte)(v.Body))
		fmt.Printf("Exported article to %s\n", bodyPath)

		metafile, err := os.Create(metaPath)
		if err != nil {
			return err
		}
		defer metafile.Close()
		yml, err := yaml.Marshal(v)
		if err != nil {
			return err
		}
		metafile.Write(([]byte)(string(yml)))
	}
	return nil
}
