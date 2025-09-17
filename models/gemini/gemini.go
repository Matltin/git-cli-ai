package gemini

import (
	"github.com/Matltin/git-cli-ai/models"
	"github.com/Matltin/git-cli-ai/utils"
)

var _ models.AIModel = (*Gemini)(nil)

type Gemini struct {
}

func (g *Gemini) Call(config utils.Config) error {
	return nil
}
