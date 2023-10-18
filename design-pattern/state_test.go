// @author hongjun500
// @date 2023/10/18 14:57
// @tool ThinkPadX1隐士
// Created with 2022.2.Goland
// Description:

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStateMachine_Approval(t *testing.T) {
	m := &Machine{
		state: &leaderApproveState{},
	}
	assert.Equal(t, "LeaderApproveState", m.GetState())
	m.Approval()
	assert.Equal(t, "FinanceApproveState", m.GetState())
	m.Reject()
	assert.Equal(t, "LeaderApproveState", m.GetState())
	m.Approval()
	assert.Equal(t, "FinanceApproveState", m.GetState())
	m.Approval()
}
