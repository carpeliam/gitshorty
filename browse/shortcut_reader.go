package browse

import (
	"context"
	"fmt"

	sc "github.com/carpeliam/gitshorty/generated"
)

type Backlog interface {
	GetStory(storyID int) (sc.Story, error)
}

type ShortcutReader struct {
	client *sc.APIClient
	auth   context.Context
}

func NewShortcutReader(apiToken string) Backlog {
	cfg := sc.NewConfiguration()
	client := sc.NewAPIClient(cfg)
	auth := context.WithValue(context.Background(), sc.ContextAPIKey, sc.APIKey{Key: apiToken})
	return ShortcutReader{client: client, auth: auth}
}

func (shortcut ShortcutReader) GetStory(publicID int) (sc.Story, error) {
	story, resp, err := shortcut.client.DefaultApi.GetStory(shortcut.auth, int64(publicID))
	if err != nil {
		fmt.Printf("Error received; Shortcut Response: %+v", resp)
	}
	return story, err
}
