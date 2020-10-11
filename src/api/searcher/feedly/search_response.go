package feedly

type SearchResponse struct {
	Results []SearchResult `json:"results"`
}

type SearchResult struct {
	FeedID              string         `json:"feedId"`
	Score               float64        `json:"score"`
	LastUpdated         int64          `json:"lastUpdated"`
	Coverage            float64        `json:"coverage"`
	AverageReadTime     float64        `json:"averageReadTime"`
	CoverageScore       float64        `json:"coverageScore"`
	EstimatedEngagement int            `json:"estimatedEngagement"`
	TagCounts           map[string]int `json:"tagCounts"`
	TotalTagCount       int            `json:"totalTagCount"`
	WebsiteTitle        string         `json:"websiteTitle"`
	ID                  string         `json:"id"`
	Title               string         `json:"title"`
	Topics              []string       `json:"topics"`
	Velocity            float64        `json:"velocity"`
	Subscribers         int            `json:"subscribers"`
	Updated             int64          `json:"updated"`
	Website             string         `json:"website"`
	IconURL             string         `json:"iconUrl"`
	Partial             bool           `json:"partial"`
	VisualURL           string         `json:"visualUrl"`
	Language            string         `json:"language"`
	ContentType         string         `json:"contentType"`
	Description         string         `json:"description"`
	DeliciousTags       []string       `json:"deliciousTags"`
}
