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
	"time"

	"github.com/olekukonko/tablewriter"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// categoryGetCmd represents the category command
var categoryGetCmd = &cobra.Command{
	Use:   "category",
	Short: "Get articles under Category",
	Long:  `TBD`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("category called")
	},
}

// categoryListCmd TBD
var categoryListCmd = &cobra.Command{
	Use:   "category",
	Short: "List categories",
	Long:  `TBD`,
	RunE:  runCategoryListCmd,
}

func init() {
	listCmd.AddCommand(categoryListCmd)
	getCmd.AddCommand(categoryGetCmd)

}

func runCategoryListCmd(cmd *cobra.Command, args []string) error {
	client, err := newDefaultClient()
	if err != nil {
		return errors.Wrap(err, "newClient failed:")
	}

	req := CategoryListRequest{
		Page:      o.Page,
		PerPage:   o.Limit,
		SortBy:    o.SortBy,
		SortOrder: o.SortOrder,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := client.CategoryList(ctx, req)
	if err != nil {
		return errors.Wrapf(err, "StackShow was failed: req = %+v, res = %+v", req, res)
	}

	categories := res.Categories
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"id", "title", "url", "updated at"})
	for _, v := range categories {
		table.Append([]string{
			fmt.Sprintf("%d", v.ID),
			v.Name,
			v.HTMLURL,
			v.UpdatedAt,
		})
	}
	table.Render()

	return nil
}

// CategoryList is TBD
func (client *Client) CategoryList(ctx context.Context, apiRequest CategoryListRequest) (*CategoryListResponse, error) {
	subPath := fmt.Sprintf("/%s/categories.json", client.Locale)

	httpRequest, err := client.newRequest(ctx, "GET", subPath, nil)
	if err != nil {
		return nil, err
	}

	q := httpRequest.URL.Query()
	q.Add("page", fmt.Sprintf("%d", apiRequest.Page))
	q.Add("per_page", fmt.Sprintf("%d", apiRequest.PerPage))
	q.Add("sort_by", apiRequest.SortBy)
	q.Add("sort_order", apiRequest.SortOrder)
	httpRequest.URL.RawQuery = q.Encode()
	fmt.Printf("Info: Request to %s\n", httpRequest.URL)

	httpResponse, err := client.HTTPClient.Do(httpRequest)
	if err != nil {
		return nil, err
	}

	// TODO: Check status code here…

	var apiResponse CategoryListResponse
	if err := decodeBody(httpResponse, &apiResponse); err != nil {
		return nil, err
	}
	return &apiResponse, nil
}
