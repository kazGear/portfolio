package notify

import (
	"fmt"
	"strings"

	"github.com/kazGear/portfolio/goBatch/internal/batchMonitor/model"
	"github.com/kazGear/portfolio/goBatch/pkg/utils"
)

type MessageNotify struct {
	Content strings.Builder
}

func NewMessageDiscord() Message {
	return &MessageNotify{
		Content: strings.Builder{},
	}
}

func (m *MessageNotify) CreateMessage(source *model.BatchExecution) string {
	m.Content.WriteString("🚨 Portfolio batch notify.\n\n")
	m.Content.WriteString(fmt.Sprintf("Exec date: %v\n\n", source.StartAt))
	m.Content.WriteString(fmt.Sprintf("Batch name: %v\n\n", source.BatchName))
	m.Content.WriteString(fmt.Sprintf("Status: %v\n\n", source.Status))
	m.Content.WriteString(fmt.Sprintf("Log id: %v\n\n", source.LogId))

	if source.Message != nil {
		m.Content.WriteString(
			utils.TruncateString(fmt.Sprintf("Message: %v\n\n", *source.Message), 200),
		)
	} else {
		m.Content.WriteString("Message:")
	}

	return m.Content.String()
}