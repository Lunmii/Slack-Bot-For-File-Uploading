package main

import (
	"fmt"
	"github.com/slack-go/slack"
	"os"
)

func main() {
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-4052568962279-4068539812182-mXvC2XHbsbiQH5aH7O3MOkHq")
	os.Setenv("CHANNEL_ID", "C041ND9DAAK")
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channelArr := []string{os.Getenv("CHANNEL_ID")}
	fileArr := []string{"package main.docx"}

	for i := 0; i < len(fileArr); i++ {
		params := slack.FileUploadParameters{
			Channels: channelArr,
			File:     fileArr[i],
		}
		file, err := api.UploadFile(params)
		if err != nil {
			fmt.Printf("%\n", err)
			return
		}
		fmt.Printf("Name: %s, URL:%s\n", file.Name, file.URL)
	}
}
