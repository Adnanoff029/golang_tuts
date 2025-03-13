package main

import (
	"fmt"
	"strings"
)

type sms struct {
	id      string
	content string
	tags    []string
}

func tagMessages(messages []sms, tagger func(sms) []string) []sms {
	for i, value := range messages {
		t := tagger(value)
		messages[i].tags = t
	}
	return messages
}

func main() {
	messages := []sms{
		{id: "001", content: "Urgent! Last chance to see!"},
		{id: "002", content: "Big sale on all items!"},
		// Additional messages...
	}
	taggedMessages := tagMessages(messages, func(a sms) []string {
		var tags []string
		msg := a.content
		msg = strings.ToLower(msg)
		if strings.Contains(msg, "urgent") {
			tags = append(tags, "Urgent")
		}
		if strings.Contains(msg, "sale") {
			tags = append(tags, "Promo")
		}
		return tags
	})
	fmt.Println(taggedMessages)
}
