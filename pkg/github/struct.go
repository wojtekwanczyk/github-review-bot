package github

type Hook struct {
	Action      string      `json:"action"`
	Review      Review      `json:"review"`
	PullRequest PullRequest `json:"pull_request"`
}

type Review struct {
	ID   int  `json:"id"`
	User User `json:"user"`
}

type PullRequest struct {
	Url                string `json:"html_url"`
	Title              string `json:"title"`
	User               User   `json:"user"`
	RequestedReviewers []User `json:"requested_reviewers"`
}

type User struct {
	Login string `split_words:"true"`
}
