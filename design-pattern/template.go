// @author hongjun500
// @date 2023/10/17 14:28
// @tool ThinkPadX1隐士
// Created with 2022.2.Goland
// Description: 模板模式

package main

import "fmt"

type ISMS interface {
	send(content string, phone string) error
}

type sms struct {
	ISMS
}

func (s *sms) Valid(content string) error {
	if len(content) > 100 {
		return fmt.Errorf("content too long")
	}
	return nil
}

func (s *sms) Send(content string, phone string) error {
	if err := s.Valid(content); err != nil {
		return err
	}
	// 调用子类的实现
	return s.ISMS.send(content, phone)
}

type TelecomSms struct {
	*sms
}

func NewTelecomSms() *TelecomSms {
	tel := &TelecomSms{}
	tel.sms = &sms{ISMS: tel}
	return tel
}

func (t *TelecomSms) send(content string, phone string) error {
	fmt.Printf("telecom send sms, content: %s, phone: %s\n", content, phone)
	return nil
}
