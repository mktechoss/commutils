package servicenotify

type Notify interface {
	Send(title, content string) error
}
