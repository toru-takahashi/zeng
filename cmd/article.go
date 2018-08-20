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
)

// articleGetCmd represents the article command
var articleGetCmd = &cobra.Command{
	Use:   "article [ArticleID]",
	Short: "Get article",
	Long:  `TBD`,
	Args:  cobra.MinimumNArgs(1),
	RunE:  runArticleGetCmd,
}

// articleListCmd represents the article command
var articleListCmd = &cobra.Command{
	Use:   "article",
	Short: "List articles",
	Long:  `TBD`,
	RunE:  runArticleListCmd,
}

func init() {
	articleListCmd.Flags().IntVarP(&o.CategoryID, "category-id", "", 0, "Filter by Category ID")
	articleListCmd.Flags().IntVarP(&o.SectionID, "section-id", "", 0, "Filter by Section ID")

	getCmd.AddCommand(articleGetCmd)
	listCmd.AddCommand(articleListCmd)
}

func runArticleGetCmd(cmd *cobra.Command, args []string) error {
	client, err := newDefaultClient()
	if err != nil {
		return errors.Wrap(err, "newClient failed:")
	}

	articleID, err := strconv.Atoi(args[0])

	if err != nil {
		return errors.Wrapf(err, "failed to parse ArticleID: %s", args[0])
	}

	req := ArticleGetRequest{
		ID: articleID,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := client.ArticleGet(ctx, req)
	if err != nil {
		return errors.Wrapf(err, "StackShow was failed: req = %+v, res = %+v", req, res)
	}

	article := res.Article

	if article.ID == 0 {
		fmt.Printf("Loading article failed. Please confirm the article id.\n")
		return nil
	}

	outPath := fmt.Sprintf("%d.html", article.ID)
	file, err := os.Create(outPath)
	if err != nil {
		// File Error
	}
	defer file.Close()
	file.Write(([]byte)(article.Body))
	fmt.Printf("Exported %s (%s)to %s\n", article.Title, article.Locale, outPath)
	return nil
}

func runArticleListCmd(cmd *cobra.Command, args []string) error {
	client, err := newDefaultClient()
	if err != nil {
		return errors.Wrap(err, "newClient failed:")
	}

	req := ArticleListRequest{
		Page:       o.Page,
		PerPage:    o.Limit,
		SortBy:     o.SortBy,
		SortOrder:  o.SortOrder,
		CategoryID: o.CategoryID,
		SectionID:  o.SectionID,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := client.ArticleList(ctx, req)
	if err != nil {
		return errors.Wrapf(err, "StackShow was failed: req = %+v, res = %+v", req, res)
	}

	articles := res.Articles

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{
		"category id",
		"category",
		"section id",
		"section",
		"article id",
		"title",
		"url",
		"edited at",
	})
	for _, v := range articles {

		category := ""
		categoryID := 0
		section := ""

		for _, s := range res.Sections {
			if v.SectionID == s.ID {
				section = s.Name
				categoryID = s.CategoryID
			}
		}

		for _, c := range res.Categories {
			if categoryID == c.ID {
				category = c.Name
			}
		}

		table.Append([]string{
			fmt.Sprintf("%d", categoryID),
			category,
			fmt.Sprintf("%d", v.SectionID),
			section,
			fmt.Sprintf("%d", v.ID),
			v.Title,
			v.HTMLURL,
			v.EditedAt,
		})
	}
	table.Render()
	// fmt.Println(articles)

	return nil
}

// ArticleGet is
func (client *Client) ArticleGet(ctx context.Context, apiRequest ArticleGetRequest) (*ArticleGetResponse, error) {
	subPath := fmt.Sprintf("/%s/articles/%d.json", client.Locale, apiRequest.ID)

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

	var apiResponse ArticleGetResponse
	if err := decodeBody(httpResponse, &apiResponse); err != nil {
		return nil, err
	}
	return &apiResponse, nil
}

// ArticleList is
func (client *Client) ArticleList(ctx context.Context, apiRequest ArticleListRequest) (*ArticleListResponse, error) {
	subPath := fmt.Sprintf("/%s/articles.json", client.Locale)

	if apiRequest.CategoryID != 0 {
		subPath = fmt.Sprintf("/%s/categories/%d/articles.json", client.Locale, apiRequest.CategoryID)
	}

	if apiRequest.SectionID != 0 {
		subPath = fmt.Sprintf("/%s/sections/%d/articles.json", client.Locale, apiRequest.SectionID)
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
	q.Add("include", "categories,sections")
	httpRequest.URL.RawQuery = q.Encode()
	fmt.Printf("Info: Request to %s\n", httpRequest.URL)

	httpResponse, err := client.HTTPClient.Do(httpRequest)
	if err != nil {
		return nil, err
	}

	// Check status code here…

	var apiResponse ArticleListResponse
	if err := decodeBody(httpResponse, &apiResponse); err != nil {
		return nil, err
	}
	return &apiResponse, nil
}
