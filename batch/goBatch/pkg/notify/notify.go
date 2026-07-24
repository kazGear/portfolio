package notify

type Notify interface {
	Notify(content string) error
}