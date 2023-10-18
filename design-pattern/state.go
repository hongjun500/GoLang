// @author hongjun500
// @date 2023/10/18 14:31
// @tool ThinkPadX1隐士
// Created with 2022.2.Goland
// Description: 状态模式

package main

// IState 状态
type IState interface {
	Approval(m *Machine) // 审批
	Reject(m *Machine)   // 驳回
	GetState() string
}

// Machine 状态机
type Machine struct {
	state IState
}

// SetState 更新状态
func (m *Machine) SetState(state IState) {
	m.state = state
}

// GetState 获取状态
func (m *Machine) GetState() string {
	return m.state.GetState()
}

// Approval 审批
func (m *Machine) Approval() {
	m.state.Approval(m)
}

// Reject 驳回
func (m *Machine) Reject() {
	m.state.Reject(m)
}

type leaderApproveState struct {
}

func (leaderApproveState) Approval(m *Machine) {
	println("leader approve")
	m.SetState(&financeApproveState{})
}

func (leaderApproveState) GetState() string {
	return "LeaderApproveState"
}

func (leaderApproveState) Reject(m *Machine) {}

type financeApproveState struct {
}

func (financeApproveState) Approval(m *Machine) {
	println("finance approve")
}

func (financeApproveState) GetState() string {
	return "FinanceApproveState"
}

func (financeApproveState) Reject(m *Machine) {
	m.SetState(&leaderApproveState{})
}
