package model

type NewsFromPerm59 struct {
	ResultData struct {
		StatusCode int    `json:"statusCode"`
		Error      string `json:"error"`
		Type       string `json:"type"`
		Data       []struct {
			ID                int           `json:"id"`
			CommentsCount     int           `json:"commentsCount"`
			Formats           []interface{} `json:"formats"`
			Header            string        `json:"header"`
			IsCommentsAllowed bool          `json:"isCommentsAllowed"`
			IsCommercial      bool          `json:"isCommercial"`
			MainPhoto         struct {
				ID          string `json:"id"`
				Author      string `json:"author"`
				Description string `json:"description"`
				FileMask    string `json:"fileMask"`
				Height      int    `json:"height"`
				SourceName  string `json:"sourceName"`
				URL         string `json:"url"`
				Width       int    `json:"width"`
			} `json:"mainPhoto"`
			PublishAt string `json:"publishAt"`
			Rubrics   []struct {
				ID        int    `json:"id"`
				BuildPath string `json:"buildPath"`
				Name      string `json:"name"`
				Path      string `json:"path"`
			} `json:"rubrics"`
			Subheader string        `json:"subheader"`
			Themes    []interface{} `json:"themes"`
			Urls      struct {
				URL          string `json:"url"`
				URLCanonical string `json:"urlCanonical"`
				URLComments  string `json:"urlComments"`
			} `json:"urls"`
			ViewsCount int `json:"viewsCount"`
		} `json:"data"`
		Meta struct {
			Themes []interface{} `json:"themes"`
			Total  int           `json:"total"`
		} `json:"meta"`
	} `json:"result.data"`
	Metatags struct {
		Canonical   string `json:"canonical"`
		Description string `json:"description"`
		H1          string `json:"h1"`
		Keywords    string `json:"keywords"`
		Title       string `json:"title"`
	} `json:"metatags"`
	Notifications []struct {
		ID        int    `json:"id"`
		Datetime  string `json:"datetime"`
		Link      string `json:"link"`
		Text      string `json:"text"`
		Timestamp int    `json:"timestamp"`
		Title     string `json:"title"`
	} `json:"notifications"`
}

func (news NewsFromPerm59) ConvertToSliceOneNews() (*[]OneNews) {
	result := new([]OneNews)
	for _, n := range news.ResultData.Data {
		*result = append(*result, OneNews{
			Header: n.Header,
			URL:    n.Urls.URL,
			Site:   "https://59.ru",
		})
	}
	return result
}
