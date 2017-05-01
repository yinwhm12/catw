package client

type ArticleJSON struct {
	ArticleRoot2 int `json:"article_root_2"`
	ArticleTitle	string `json:"article_title"`
	ArticleContent	string 	`json:"article_content"`
	ArticleRoot1	int	`json:"article_root_1"`
	ArticleLevel	int 	`json:"article_level"`
}

type RespondJSON struct {
	TextContent	string	`json:"text_content"`
	ArticleId	int 	`json:"article_id"`
}