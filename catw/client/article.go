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
	ArticleId	int 	`json:"article_id"`//检查文章是否存在
	ImgContent	string	`json:"img_content,omitempty"`
}

type RespondTwoJSON struct {
	TextContent	string	`json:"text_content"`
	RespondOne	int `json:"respond_one"`
	ImgContent	string	`json:"img_content,omitempty"`
}

type MessageJSON struct {
	ToUserID int	`json:"to_user_id,omitempty"`
	ArticleId int	`json:"article_id,omitempty"`
	RespondOneId	int	`json:"respond_one_id,omitempty"` //二级回复的id 从而获得用户Id
}