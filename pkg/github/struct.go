package github

type Hook struct {
	PullRequest PullRequest `json:"pull_request"`
}

type PullRequest struct {
	RequestedReviewers []RequestedReviewers `json:"requested_reviewers"`
}

type RequestedReviewers struct {
	Login string `split_words:"true"`
}
