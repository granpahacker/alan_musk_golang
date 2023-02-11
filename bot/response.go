// BY GRANDPA HACKER 2023
package bot

import (
	"context"
	"fmt"

	gpt3 "github.com/PullRequestInc/go-gpt3"
	"github.com/bwmarrin/discordgo"
)

// get response without stream
func GetResponse(client gpt3.Client, ctx context.Context, quesiton string, s *discordgo.Session, m *discordgo.MessageCreate) {
	resp, err := client.CompletionWithEngine(ctx, gpt3.TextDavinci003Engine, gpt3.CompletionRequest{
		Prompt: []string{
			quesiton,
		},
		MaxTokens:   gpt3.IntPtr(3000),
		Temperature: gpt3.Float32Ptr(0),
		Echo:        true,
	})

	// send response message
	_, _ = s.ChannelMessageSend(m.ChannelID, resp.Choices[0].Text)

	if err != nil {
		fmt.Println(err)
	}
}
