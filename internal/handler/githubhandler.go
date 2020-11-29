package handler

import (
	"encoding/json"
	"net/http"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

// GithubHandler test
type GithubHandler struct {
	Token  string
	Client *github.Client
}

// TrafficHandle test
func (gh *GithubHandler) TrafficHandle(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: gh.Token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)
	opts := github.TrafficBreakdownOptions{
		Per: "day",
	}
	g, _, err := client.Repositories.ListTrafficViews(ctx, "dgduncan", "SevenSegment", &opts)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(g)

}
