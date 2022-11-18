package main

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
)


func main(){

	os.Setenv("CHANNEL_ID", "C04C837LDQQ")
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-4394866151444-4405073684257-eOPq72dDwzfB8HYhJZSlCQ1N")
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))

	channelArr := []string{os.Getenv("CHANNEL_ID")}
	fileArr := []string{"image.jpg"}

	for i := 0; i < len(fileArr); i++{
		params := slack.FileUploadParameters{
			Channels :channelArr,
			File: fileArr[i],
		}

		file, err := api.UploadFile(params)

		if err != nil {
			fmt.Printf("Error: %v", err)
			return
		}

		fmt.Printf("Name: %v Url: %v",file.Name, file.URL)

	}


}