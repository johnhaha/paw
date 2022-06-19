package paw

import (
	"context"

	"github.com/johnhaha/nose"
)

func StartPaw(ctx context.Context, token string, pageID string) error {
	logPageID = pageID
	notionToken = token
	startListen(ctx)
	return nil
}

func Log(title string, msg string) error {
	return logPubSub.Pub(LogMsg{
		Time:    getNowTimeString(),
		Title:   title,
		Content: msg,
	})
}

func LogToDB(dbID string, data any) error {
	db := nose.NewDBClient(notionToken, dbID)
	_, err := db.SaveData(data)
	return err
}
