package reply

import "github.com/nlopes/slack"

func CancelButton() slack.BlockElement {

	return slack.ButtonBlockElement{
		ActionID: CancelValue,
		Value:    CancelValue,
		Text:     slack.NewTextBlockObject("plain_text", cancelText, false, false),
		Type:     "button",
		Style:    slack.StyleDanger,
	}
}

func PrimaryButton(text string) slack.BlockElement {

	return slack.ButtonBlockElement{
		ActionID: AcceptValue,
		Value:    AcceptValue,
		Text:     slack.NewTextBlockObject("plain_text", text, false, false),
		Type:     "button",
		Style:    slack.StylePrimary,
	}
}
