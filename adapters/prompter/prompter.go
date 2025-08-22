package prompter

import (
	"github.com/manifoldco/promptui"
)

type Prompter struct{}

func NewPrompter() *Prompter {
	return new(Prompter)
}

func (p *Prompter) Confirm(message string) (bool, error) {
	prompt := promptui.Prompt{
		Label:     message + " (Y/n)",
		IsConfirm: true,
	}

	result, err := prompt.Run()
	if err != nil {
		return false, err
	}

	return result == "y" || result == "Y" || result == "", nil
}
