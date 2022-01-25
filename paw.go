package paw

import (
	"context"

	"github.com/johnhaha/echo"
	"github.com/johnhaha/nose"
)

func StartPaw(ctx context.Context, token string, pageID string) error {
	logPageID = pageID
	notionToken = token
	startListen(ctx)
	return nil
}

func Log(title string, msg string) error {
	return echo.PubJson(logChannel, NewLogMsg(title, msg))
}

func LogToDB(dbID string, data interface{}) error {
	db := nose.NewDBClient(notionToken, dbID)
	_, err := db.SaveData(data)
	return err
}
