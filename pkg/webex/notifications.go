package webex

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/wojtekwanczyk/github-review-bot/pkg/config"
	"github.com/wojtekwanczyk/github-review-bot/pkg/github"
)

func NotifyComment(h *github.Hook) {
	if roomID, ok := config.RoomMapping[h.PullRequest.User.Login]; ok {
		text := fmt.Sprintf("User %s updated your pull request: [%s](%s)",
			h.Review.User.Login, h.PullRequest.Title, h.PullRequest.Url)
		SendMessage(text, roomID)
	} else {
		log.Errorf("User not found in mapping: %s", h.PullRequest.User.Login)
	}
}

func NotifyNewPr(pr *github.PullRequest) {
	template := "User %s assigned new pull request to you: [%s](%s)"
	notifyReviewers(pr, template)
}

func NotifyPrUpdated(pr *github.PullRequest) {
	template := "User %s updated pull request: [%s](%s)"
	notifyReviewers(pr, template)
}

func notifyReviewers(pr *github.PullRequest, template string) {
	for _, user := range pr.RequestedReviewers {
		if roomID, ok := config.RoomMapping[user.Login]; ok {
			text := fmt.Sprintf(template, pr.User, pr.Title, pr.Url)
			SendMessage(text, roomID)
		} else {
			log.Errorf("User not found in mapping: %s", pr.User.Login)
		}
	}
}
