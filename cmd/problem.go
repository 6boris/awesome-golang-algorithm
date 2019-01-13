package cmd

type Problem struct {
	Stat       Stat       `json:"stat"`
	Status     Status     `json:"status"`
	Difficulty Difficulty `json:"difficulty"`
	PaidOnly   bool       `json:"paid_only"`
	IsFavor    bool       `json:"is_favor"`
	Frequency  int        `json:"frequency"`
	Progress   int        `json:"progress"`
}
type Stat struct {
	QuestionID          int    `json:"question_id"`
	QuestionArticleLive bool   `json:"question_article_live"`
	QuestionArticleSlug string `json:"question_article_slug"`
	QuestionTitle       string `json:"question_title"`
	QuestionTitleSlup   string `json:"question_title_slup"`
	QuestionHide        bool   `json:"question_hide"`
	TotalAcs            int    `json:"total_acs"`
	TotalSubmitted      int    `json:"total_submitted"`
	FrontendQuestion_ID int    `json:"frontend_question_id"`
	IsNewQuestion       bool   `json:"is_new_question"`
}

type Status struct {
}
type Difficulty struct {
	Level int
}
