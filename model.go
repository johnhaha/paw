package paw

import (
	"github.com/johnhaha/nose"
)

type LogMsg struct {
	Time    string `json:"time" nose:"text-code"`
	Title   string `json:"title" nose:"blue,bold"`
	Content string `json:"content"`
}

func NewLogMsg(title string, content string) *LogMsg {
	return &LogMsg{
		Time:    getNowTimeString(),
		Title:   " " + title + " ",
		Content: content,
	}
}

func (msg *LogMsg) Log() error {
	n := nose.NewPageClient(notionToken, logPageID)
	return n.AppendDataTextBlock(*msg)
}
