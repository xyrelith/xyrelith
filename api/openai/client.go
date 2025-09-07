package x_openai

import (
	"os"

	"github.com/sashabaranov/go-openai"
)

var Client *openai.Client

func Init() {

	Client = openai.NewClient(os.Getenv("OPENAI_KEY"))
}
