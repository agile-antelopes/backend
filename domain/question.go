package domain

type Question struct {
	Id         int    `json:"id"`
	QuestionId string `json:"question"`
	Answer     string `json:"answer"`
	Country    string `json:"country"`
	UserName   string `json:"username"`
}
