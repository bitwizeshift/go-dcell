package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	_ "embed"

	"rodusek.dev/pkg/dcell"
)

type Ref struct {
	Label string `json:"label" dcell:"label"`
	Ref   string `json:"ref" dcell:"ref"`
	SHA   string `json:"sha" dcell:"sha"`
	User  User   `json:"user" dcell:"user"`
}

type User struct {
	Login     string `json:"login" dcell:"login"`
	ID        int    `json:"id" dcell:"id"`
	AvatarURL string `json:"avatar_url" dcell:"avatar_url"`
	URL       string `json:"url" dcell:"url"`
	HTMLURL   string `json:"html_url" dcell:"html_url"`
}

type PullRequest struct {
	URL               string  `json:"url" dcell:"url"`
	ID                int     `json:"id" dcell:"id"`
	HTMLURL           string  `json:"html_url" dcell:"html_url"`
	DiffURL           string  `json:"diff_url" dcell:"diff_url"`
	PatchURL          string  `json:"patch_url" dcell:"patch_url"`
	IssueURL          string  `json:"issue_url" dcell:"issue_url"`
	Number            int     `json:"number" dcell:"number"`
	State             string  `json:"state" dcell:"state"`
	Locked            bool    `json:"locked" dcell:"locked"`
	Title             string  `json:"title" dcell:"title"`
	User              User    `json:"user" dcell:"user"`
	Body              string  `json:"body" dcell:"body"`
	CreatedAt         string  `json:"created_at" dcell:"created_at"`
	UpdatedAt         string  `json:"updated_at" dcell:"updated_at"`
	ClosedAt          *string `json:"closed_at" dcell:"closed_at"`
	MergedAt          *string `json:"merged_at" dcell:"merged_at"`
	MergeCommitSHA    *string `json:"merge_commit_sha" dcell:"merge_commit_sha"`
	Assignee          *string `json:"assignee" dcell:"assignee"`
	Milestone         *string `json:"milestone" dcell:"milestone"`
	CommitsURL        string  `json:"commits_url" dcell:"commits_url"`
	ReviewCommentsURL string  `json:"review_comments_url" dcell:"review_comments_url"`
	CommentsURL       string  `json:"comments_url" dcell:"comments_url"`
	StatusesURL       string  `json:"statuses_url" dcell:"statuses_url"`

	Head Ref `json:"head" dcell:"head"`
	Base Ref `json:"base" dcell:"base"`
}

type Repository struct {
	ID       int    `json:"id" dcell:"id"`
	Name     string `json:"name" dcell:"name"`
	FullName string `json:"full_name" dcell:"full_name"`
	Owner    User   `json:"owner" dcell:"owner"`
	Private  bool   `json:"private" dcell:"private"`
	HTMLURL  string `json:"html_url" dcell:"html_url"`
	URL      string `json:"url" dcell:"url"`
}

type Event struct {
	Action      string       `json:"action" dcell:"action"`
	Number      int          `json:"number" dcell:"number"`
	PullRequest *PullRequest `json:"pull_request" dcell:"pull_request"`
	Repository  Repository   `json:"repository" dcell:"repository"`
	Sender      User         `json:"sender" dcell:"sender"`
}

type GitHubContext struct {
	EventName string `json:"event_name" dcell:"event_name"`
	Ref       string `json:"ref" dcell:"ref"`
	Event     Event  `json:"event" dcell:"event"`
}

type Context struct {
	GitHub GitHubContext `json:"github" dcell:"github"`
}

//go:embed pull_request.payload.json
var eventPayload []byte

func LoadContext() (*Context, error) {
	var event Event
	if err := json.Unmarshal(eventPayload, &event); err != nil {
		return nil, fmt.Errorf("pull_request: %w", err)
	}
	ctx := &Context{
		GitHub: GitHubContext{
			EventName: "pull_request",
			Ref:       event.PullRequest.Head.Ref,
			Event:     event,
		},
	}

	return ctx, nil
}

func main() {
	log.SetFlags(0)
	if len(os.Args) < 2 {
		log.Fatalf("Usage: go-dcell <expression>")
	}
	expr, err := dcell.Compile(os.Args[1])
	if err != nil {
		log.Fatalf("error: compilation failed: %v", err)
	}
	ctx, err := LoadContext()
	if err != nil {
		log.Fatalf("error: loading context failed: %v", err)
	}
	result, err := expr.Eval(ctx)
	if err != nil {
		log.Fatalf("error: evaluation failed: %v", err)
	}
	out := result.Interface()
	bytes, err := json.Marshal(out)
	if err != nil {
		log.Fatalf("error: marshalling result failed: %v", err)
	}
	fmt.Println(string(bytes))
}
