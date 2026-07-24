package notify

import "github.com/kazGear/portfolio/goBatch/internal/batchMonitor/model"

type Message interface {
	CreateMessage(source *model.BatchExecution) string
}