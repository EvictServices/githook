package events

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/go-github/v61/github"
	"github.com/shi-gg/githook/discord"
	"github.com/shi-gg/githook/utils"
)

func Ping(w http.ResponseWriter, r *http.Request, url string) {
	var body github.PingEvent
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&body)

	username := "GitHub"
	avatarURL := ""
	repoName := "Repository"
	repoURL := ""

	if body.Sender != nil && body.Sender.Login != nil {
		username = *body.Sender.Login
	}
	if body.Sender != nil && body.Sender.AvatarURL != nil {
		avatarURL = *body.Sender.AvatarURL
	}
	if body.Repo != nil && body.Repo.FullName != nil {
		repoName = *body.Repo.FullName
	}
	if body.Repo != nil && body.Repo.HTMLURL != nil {
		repoURL = *body.Repo.HTMLURL
	}

	discord.SendWebhook(
		url,
		discord.WebhookPayload{
			Username:  username,
			AvatarURL: avatarURL,
			Embeds: []discord.Embed{
				{
					Title:       fmt.Sprintf("%s: Ping", repoName),
					URL:         repoURL,
					Description: "🏓 Ping! Pong!",
					Color:       utils.GetColors().Default,
				},
			},
		},
	)
}
