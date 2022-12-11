package main

import (
	"context"
	"fmt"
	"github.com/PullRequestInc/go-gpt3"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	r := gin.Default()
	godotenv.Load()
	api_key := os.Getenv("TOKEN")
	if api_key == "" {
		log.Fatal("Fatal .env token key")
	}
	f, _ := os.OpenFile("chatGPT.txt", 777, 0)
	ctx := context.Background()
	client := gpt3.NewClient(api_key)
	r.POST("/", func(c *gin.Context) {
		msg := c.PostForm("message")
		err := client.CompletionStreamWithEngine(ctx, gpt3.TextDavinci003Engine, gpt3.CompletionRequest{
			Prompt:    []string{msg},
			MaxTokens: gpt3.IntPtr(2500),
			Echo:      true,
		}, func(resp *gpt3.CompletionResponse) {
			c.JSON(200, resp.Choices[0].Text)
			f.Write([]byte(resp.Choices[0].Text))
			fmt.Printf("%s\n", resp.Choices[0].Text)
		})
		if err != nil {
			log.Fatal(err)
		}
	})

	r.Run(":80")
}
