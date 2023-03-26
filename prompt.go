package anthropic

import "fmt"

func setPrompt(human, assistant string) (string, error) {
	if human == "" {
		return "", ErrPromptHumanEmpty
	}
	if assistant == "" {
		return fmt.Sprintf("\n\nHuman: %v\n\nAssistant:", human), nil
	}
	return fmt.Sprintf("%v%v", human, assistant), nil
}

func (opts *Opts) buildPrompts() (string, error) {
	if opts.Len() == 0 {
		return "", nil
	}

	prompts, _ := setPrompt(opts.Context[0].Human, opts.Context[0].Assistant)
	if opts.Len() < 2 {
		return prompts, nil
	}
	for _, d := range opts.Context[1:] {
		if d.Human == "" {
			return "", ErrPromptHumanEmpty
		}
		if d.Assistant == "" {
			return fmt.Sprintf("%v\n\nHuman: %v\n\nAssistant:", prompts, d.Human), nil
		}
		prompts = fmt.Sprintf("%v\n\nHuman: %v\n\nAssistant:%v", prompts, d.Human, d.Assistant)
	}
	return prompts, nil
}

func addPrompt(context, human string) (string, error) {
	if human == "" {
		return "", ErrPromptHumanEmpty
	}
	if context == "" {
		return "", ErrPromptCtxEmpty
	}
	return fmt.Sprintf("%v\n\nHuman: %v\n\nAssistant:", context, human), nil
}
