package shortcut

import (
	"context"
	"fmt"

	sc "github.com/carpeliam/gitshorty/generated"
)

type Client interface {
	GetStory(storyID int) (sc.Story, error)
}

type ShortcutClient struct {
	client *sc.APIClient
	auth   context.Context
}

func NewShortcutClient(apiToken string) Client {
	cfg := sc.NewConfiguration()
	client := sc.NewAPIClient(cfg)
	auth := context.WithValue(context.Background(), sc.ContextAPIKey, sc.APIKey{Key: apiToken})
	return ShortcutClient{client: client, auth: auth}
}

func (shortcut ShortcutClient) GetStory(publicID int) (sc.Story, error) {
	story, resp, err := shortcut.client.DefaultApi.GetStory(shortcut.auth, int64(publicID))
	if err != nil {
		fmt.Printf("Error received; Shortcut Response: %+v", resp)
	}
	return story, err
}
