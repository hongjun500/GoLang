// @author hongjun500
// @date 2023/10/16 13:31
// @tool ThinkPadX1隐士
// Created with 2022.2.Goland
// Description: 观察者模式

package main

type ISubject interface {
	Register(observer IObserver)
	Remove(observer IObserver)
	Notify(msg string)
}

// IObserver 观察者
type IObserver interface {
	Update(msg string)
}

// Subject 主题
type Subject struct {
	observers []IObserver
}

func (s *Subject) Register(observer IObserver) {
	s.observers = append(s.observers, observer)
}

func (s *Subject) Remove(observer IObserver) {
	for i, o := range s.observers {
		if o == observer {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
		}
	}
}

func (s *Subject) Notify(msg string) {
	for _, o := range s.observers {
		o.Update(msg)
	}
}

type Observer1 struct {
}

func (Observer1) Update(msg string) {
	println("Observer1: " + msg)
}

type Observer2 struct {
}

func (Observer2) Update(msg string) {
	println("Observer2: " + msg)
}
