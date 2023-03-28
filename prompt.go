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

func buildPrompts(module any) (string, error) {
	switch r := module.(type) {
	case MessageModule:
		return setPrompt(r.Human, r.Assistant)
	case []MessageModule:
		var prompts string
		for _, d := range r {
			if d.Human == "" {
				return "", ErrPromptHumanEmpty
			}
			if d.Assistant == "" {
				return fmt.Sprintf("%v\n\nHuman: %v\n\nAssistant:", prompts, d.Human), nil
			}
			prompts = fmt.Sprintf("%v\n\nHuman: %v\n\nAssistant:%v", prompts, d.Human, d.Assistant)
		}
		return prompts, nil
	default:
		return "", fmt.Errorf("unknown module type: %T", module)
	}
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
