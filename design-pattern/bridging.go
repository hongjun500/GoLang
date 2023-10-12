// @author hongjun500
// @date 2023/10/12 15:30
// @tool ThinkPadX1隐士
// Created with 2022.2.Goland
// Description: 桥接者模式

package main

// IMsgSender 发送消息
type IMsgSender interface {
	Send(msg string) error
}

// EmailMsgSender 邮件发送
type EmailMsgSender struct {
	Email []string
}

func NewEmailMsgSender(emails []string) *EmailMsgSender {
	return &EmailMsgSender{Email: emails}
}

func (e *EmailMsgSender) Send(msg string) error {
	e.Email = append(e.Email, msg)
	return nil
}

// INotification 通知接口
type INotification interface {
	Notify(msg string) error
}

// ErrorNotification 错误通知
type ErrorNotification struct {
	MsgSender IMsgSender
}

func NewErrorNotification(sender IMsgSender) *ErrorNotification {
	return &ErrorNotification{MsgSender: sender}
}

func (e *ErrorNotification) Notify(msg string) error {
	return e.MsgSender.Send(msg)
}
