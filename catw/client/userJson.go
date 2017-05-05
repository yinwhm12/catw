package client

type UserJSON struct {
	OldPWD	string	`json:"old_pwd,omitempty"`
	NewPWD	string	`json:"new_pwd,omitempty"`
	NewPWDMore	string	`json:"new_pwd_more,omitempty"`
	Name	string	`json:"name,omitempty"`
	
}

type UserNotKeyJSON struct {
	Id	int	`json:"id,omitempty"`
	Name	string	`json:"name,omitempty"`

	Motto string	`json:"motto,omitempty"`
	City	string	`json:"city,omitempty"`
	Describe	string	`json:"describe,omitempty"`
	School	string	`json:"school,omitempty"`
	UpArticles string	`json:"up_articles,omitempty"`
	CollectArticles	string	`json:"collect_articles,omitempty"`
}
