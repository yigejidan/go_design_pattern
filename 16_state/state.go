package _6_state

import "fmt"

/*
状态模式：允许一个对象在其内部状态改变时改变它的行为。对象看起来似乎修改了它的类。
*/

// IState 状态接口
type IState interface {
	Approval(m *Machine)
	Reject(m *Machine)
	GetName() string
}

// Machine 状态机
type Machine struct {
	state IState
}

func (m *Machine) SetState(state IState) {
	m.state = state
}

func (m *Machine) GetStateName() string {
	return m.state.GetName()
}

func (m *Machine) Approval() {
	m.state.Approval(m)
}

func (m *Machine) Reject() {
	m.state.Reject(m)
}

// leaderApproveState 直属领导审批
type leaderApproveState struct {
}

func GetLeaderApproveState() IState {
	return &leaderApproveState{}
}

func (leaderApproveState) Approval(m *Machine) {
	fmt.Println("直属 leader 审批成功！")
	m.SetState(GetFinanceApproveState())
}

func (leaderApproveState) Reject(m *Machine) {

}

// GetName 获取状态名字
func (leaderApproveState) GetName() string {
	return "LeaderApproveState"
}

// financeApproveState 财务审批
type financeApproveState struct{}

// Approval 审批通过
func (f financeApproveState) Approval(m *Machine) {
	fmt.Println("财务审批成功")
	fmt.Println("出发打款操作")
}

// Reject 拒绝
func (f financeApproveState) Reject(m *Machine) {
	m.SetState(GetLeaderApproveState())
}

// GetName 获取名字
func (f financeApproveState) GetName() string {
	return "FinanceApproveState"
}

// GetFinanceApproveState GetFinanceApproveState
func GetFinanceApproveState() IState {
	return &financeApproveState{}
}
