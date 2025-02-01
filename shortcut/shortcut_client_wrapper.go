package shortcut

import (
	"context"
	"fmt"
	"log/slog"

	sc "github.com/carpeliam/gitshorty/generated"
)

type Client interface {
	GetStory(storyID int) (sc.Story, error)
	UpdateTask(storyID int, tasks map[int64]sc.UpdateTask) error
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
		slog.Debug("Error received from Shortcut", "http response", resp)
		err = fmt.Errorf("received '%w' when fetching story sc-%d", err, publicID)
	}
	return story, err
}

func (shortcut ShortcutClient) UpdateTask(storyID int, tasks map[int64]sc.UpdateTask) error {
	for taskId, taskUpdates := range tasks {
		_, resp, err := shortcut.client.DefaultApi.UpdateTask(shortcut.auth, taskUpdates, int64(storyID), taskId)
		if err != nil {
			slog.Debug("Error received from Shortcut", "http response", resp)
			return err
		}
	}
	return nil
}
