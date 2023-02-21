package _6_bridge

/*
桥接模式：将抽象部分与它的实现部分分离，使它们都可以独立地变化。
*/

// IMsgSender IMsgSender
type IMsgSender interface {
	Send(msg string) error
}

// EmailMsgSender 发送邮件
type EmailMsgSender struct {
	emails []string
}

func NewEmailMsgSender(emails []string) *EmailMsgSender {
	return &EmailMsgSender{emails: emails}
}

func (s *EmailMsgSender) Send(msg string) error {
	return nil
}

type INotification interface {
	Notify(msg string) error
}

type ErrorNotification struct {
	sender IMsgSender
}

func NewErrorNotification(sender IMsgSender) *ErrorNotification {
	return &ErrorNotification{sender: sender}
}

func (e *ErrorNotification) Notify(msg string) error {
	return e.sender.Send(msg)
}
