package models

import "github.com/Matltin/git-cli-ai/utils"

type AIModel interface {
	Call(utils.Config) error
}