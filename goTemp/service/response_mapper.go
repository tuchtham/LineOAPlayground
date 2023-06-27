package service

import (
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

type MessageHandler struct {
	//TODO: Do we need anything here?
}

func NewHandler() *MessageHandler {
	return &MessageHandler{}
}

func (h *MessageHandler) MapResponse(input string) *linebot.SendingMessage {
	switch input {
	case "ping":
		return h.newTextMessage("pong")
	case "askQr":
		return h.newTextMessageWithQuickReplies("ok?", "yes", "no")
	case "block":
		return h.newBlock(
			"header",
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aenean euismod neque tincidunt, facilisis eros vel, bibendum urna. Cras dignissim ultrices facilisis. Donec nec magna at nisi pellentesque semper. Sed eu dui tortor. Sed accumsan libero ut enim imperdiet ornare. Suspendisse metus ipsum, suscipit et nisl nec, mattis consequat lectus. Nullam at dui eu ipsum laoreet vehicula nec et elit. Vestibulum tempus tortor ut orci fringilla commodo.",
			"footer lol",
			"block preview",
		)
	case inquireRefund:
	default:
		return h.newTextMessage(input)
	}
}

func (h *MessageHandler) newTextMessage(text string) *linebot.SendingMessage {
	textMessage := linebot.NewTextMessage(text).WithQuickReplies(nil)
	return &textMessage
}

func (h *MessageHandler) newTextMessageWithQuickReplies(text string, choice1 string, choice2 string) *linebot.SendingMessage {
	quickReplyItems := linebot.NewQuickReplyItems(
		linebot.NewQuickReplyButton("", linebot.NewMessageAction(choice1, choice1)),
		linebot.NewQuickReplyButton("", linebot.NewMessageAction(choice2, choice2)),
	)
	textMessage := linebot.NewTextMessage(text).WithQuickReplies(quickReplyItems)
	return &textMessage
}

func (h *MessageHandler) newBlock(header string, body string, footer string, previewMsg string) *linebot.SendingMessage {
	block := linebot.NewFlexMessage(previewMsg,
		&linebot.BubbleContainer{
			Type: linebot.FlexContainerTypeBubble,
			Styles: &linebot.BubbleStyle{
				Header: &linebot.BlockStyle{
					BackgroundColor: "#00ffff",
					Separator:       true,
					SeparatorColor:  "#ff4400",
				},
				Body: &linebot.BlockStyle{
					//BackgroundColor: "#00ffff",
					Separator:      true,
					SeparatorColor: "#ff4400",
				},
				Footer: &linebot.BlockStyle{
					BackgroundColor: "#00ffff",
					Separator:       true,
					SeparatorColor:  "#ff4400",
				},
			},
			Header: &linebot.BoxComponent{
				Type:   linebot.FlexComponentTypeBox,
				Layout: linebot.FlexBoxLayoutTypeVertical,
				Contents: []linebot.FlexComponent{
					&linebot.TextComponent{
						Type: linebot.FlexComponentTypeText,
						Text: header,
					},
				}},
			Hero: &linebot.ImageComponent{
				Type:        "image",
				URL:         "https://www.linefriends.com/img/img_sec.jpg",
				Size:        "full",
				AspectRatio: "2:1",
			},
			Body: &linebot.BoxComponent{
				Type:   linebot.FlexComponentTypeBox,
				Layout: linebot.FlexBoxLayoutTypeVertical,
				Contents: []linebot.FlexComponent{
					&linebot.TextComponent{
						Type: linebot.FlexComponentTypeText,
						Text: body,
						Wrap: true,
					},
				}},
			Footer: &linebot.BoxComponent{
				Type:   "box",
				Layout: "vertical",
				Contents: []linebot.FlexComponent{
					&linebot.TextComponent{
						Type: linebot.FlexComponentTypeText,
						Text: footer,
					},
				},
			},
		})
	sendingMessage := block.WithQuickReplies(nil)
	return &sendingMessage
}
