package cmd

// Article https://developer.zendesk.com/rest_api/docs/help_center/articles#show-article
type Article struct {
	ID               int      `json:"id" yaml:"id"`
	URL              string   `json:"url" yaml:"url"`
	HTMLURL          string   `json:"html_url" yaml:"html_url"`
	Title            string   `json:"title" yaml:"title"`
	Body             string   `json:"body" yaml:"body"`
	Locale           string   `json:"locale" yaml:"locale"`
	SourceLocale     string   `json:"source_locale" yaml:"source_locale"`
	AuthorID         int      `json:"author_id" yaml:"author_id"`
	CommentsDisabled bool     `json:"comments_disabled" yaml:"comments_disabled"`
	OutdatedLocales  []string `json:"outdated_locales" yaml:"outdated_locales"`
	Outdated         bool     `json:"outdated" yaml:"outdated"`
	LabelNames       string   `json:"lable_names" yaml:"lable_names"`
	Draft            bool     `json:"draft" yaml:"draft"`
	Promoted         bool     `json:"promoted" yaml:"promoted"`
	Position         int      `json:"position" yaml:"position"`
	VoteSum          int      `json:"vote_sum" yaml:"vote_sum"`
	VoteCount        int      `json:"vote_count" yaml:"vote_count"`
	SectionID        int      `json:"section_id" yaml:"section_id"`
	CreatedAt        string   `json:"created_at" yaml:"created_at"`
	EditedAt         string   `json:"edited_at" yaml:"edited_at"`
	UpdatedAt        string   `json:"updated_at" yaml:"updated_at"`
}

// List options
type Options struct {
	Page, Limit, SectionID, CategoryID int
	SortBy, SortOrder                  string
}

// Configuration options
type Configuration struct {
	email, password, apikey, subdomain, locale string
}

var (
	o = &Options{}
	c = &Configuration{}
)

// ArticleGetRequest returns showArticle request
type ArticleGetRequest struct {
	ID int `json:"id"`
}

// ArticleGetResponse returns showArticle response
type ArticleGetResponse struct {
	Article Article `json:"article"`
}

// ArticleListRequest https://developer.zendesk.com/rest_api/docs/help_center/articles#list-articles
type ArticleListRequest struct {
	Page       int    `json:"page"`
	PerPage    int    `json:"per_page"`
	SortBy     string `json:"sort_by"`
	SortOrder  string `json:"sort_order"`
	CategoryID int    `json:"category_id"`
	SectionID  int    `json:"section_id"`
}

// ArticleListResponse https://developer.zendesk.com/rest_api/docs/help_center/articles#list-articles
type ArticleListResponse struct {
	Articles     []Article  `json:"articles"`
	Sections     []Section  `json:"sections"`
	Categories   []Category `json:"categories"`
	Count        int        `json:"count"`
	NextPage     string     `json:"next_page"`
	Page         int        `json:"page"`
	PageCount    int        `json:"page_count"`
	PerPage      int        `json:"per_page"`
	PreviousPage string     `json:"previous_page"`
	SortBy       string     `json:"sort_by"`
	SortOrder    string     `json:"sort_order"`
}

// Category https://developer.zendesk.com/rest_api/docs/help_center/categories#show-category
type Category struct {
	ID           int    `json:"id" yaml:"id"`
	Name         string `json:"name" yaml:"name"`
	Description  string `json:"description" yaml:"description"`
	Locale       string `json:"locale" yaml:"locale"`
	SourceLocale string `json:"source_locale" yaml:"source_locale"`
	URL          string `json:"url" yaml:"url"`
	HTMLURL      string `json:"html_url" yaml:"html_url"`
	Outdated     bool   `json:"outdated" yaml:"outdated"`
	Position     int    `json:"position" yaml:"position"`
	CreatedAt    string `json:"created_at" yaml:"created_at"`
	UpdatedAt    string `json:"updated_at" yaml:"updated_at"`
}

// CategoryListRequest
type CategoryListRequest struct {
	SortBy    string `json:"sort_by"`
	SortOrder string `json:"sort_order"`
	Page      int    `json:"page"`
	PerPage   int    `json:"per_page"`
}

// CategoryListResponse
type CategoryListResponse struct {
	Categories   []Category `json:"categories"`
	SortBy       string     `json:"sort_by"`
	SortOrder    string     `json:"sort_order"`
	Count        int        `json:"count"`
	NextPage     string     `json:"next_page"`
	Page         int        `json:"page"`
	PageCount    int        `json:"page_count"`
	PerPage      int        `json:"per_page"`
	PreviousPage string     `json:"previous_page"`
}

// Section https://developer.zendesk.com/rest_api/docs/help_center/categories#show-category
type Section struct {
	ID            int    `json:"id" yaml:"id"`
	Name          string `json:"name" yaml:"name"`
	Description   string `json:"description" yaml:"description"`
	Locale        string `json:"locale" yaml:"locale"`
	SourceLocale  string `json:"source_locale" yaml:"source_locale"`
	URL           string `json:"url" yaml:"url"`
	HTMLURL       string `json:"html_url" yaml:"html_url"`
	CategoryID    int    `json:"category_id" yaml:"category_id"`
	Outdated      bool   `json:"outdated" yaml:"outdated"`
	Position      int    `json:"position" yaml:"position"`
	ManageableBy  string `json:"manageable_by" yaml:"manageable_by"`
	UserSegmentID int    `json:"user_segment_id" yaml:"user_segment_id"`
	CreatedAt     string `json:"created_at" yaml:"created_at"`
	UpdatedAt     string `json:"updated_at" yaml:"updated_at"`
}

// SectionListRequest
type SectionListRequest struct {
	SortBy     string `json:"sort_by"`
	SortOrder  string `json:"sort_order"`
	Page       int    `json:"page"`
	PerPage    int    `json:"per_page"`
	CategoryID int    `json:"category_id"`
	SectionID  int    `json:"section_id"`
}

// SectionListResponse
type SectionListResponse struct {
	Sections     []Section  `json:"sections"`
	Categories   []Category `json:"categories"`
	SortBy       string     `json:"sort_by"`
	SortOrder    string     `json:"sort_order"`
	Count        int        `json:"count"`
	NextPage     string     `json:"next_page"`
	Page         int        `json:"page"`
	PageCount    int        `json:"page_count"`
	PerPage      int        `json:"per_page"`
	PreviousPage string     `json:"previous_page"`
}

// SectionGetRequest
type SectionGetRequest struct {
	ID int `json:"id"`
}

// SectionGetResponse
type SectionGetResponse struct {
	Articles     []Article  `json:"articles"`
	Sections     []Section  `json:"sections"`
	Categories   []Category `json:"categories"`
	Count        int        `json:"count"`
	NextPage     string     `json:"next_page"`
	Page         int        `json:"page"`
	PageCount    int        `json:"page_count"`
	PerPage      int        `json:"per_page"`
	PreviousPage string     `json:"previous_page"`
	SortBy       string     `json:"sort_by"`
	SortOrder    string     `json:"sort_order"`
}
