package main

import (
	"fmt"
)

// 基础事件结构体
type BasicEvent struct {
	name  string
	value float64
}

// 创建一个新的基础事件
func NewBasicEvent(name string, value float64) BasicEvent {
	return BasicEvent{name: name, value: value}
}

// 组合规则接口
type CombiningRule interface {
	Combine(events []BasicEvent) BasicEvent
}

// 与逻辑组合规则
type AndRule struct{}

// 实现与逻辑组合规则
func (r AndRule) Combine(events []BasicEvent) BasicEvent {
	var result float64 = 1.0
	for _, event := range events {
		result *= event.value
	}
	return BasicEvent{name: "AND", value: result}
}

// 或逻辑组合规则
type OrRule struct{}

// 实现或逻辑组合规则
func (r OrRule) Combine(events []BasicEvent) BasicEvent {
	var result float64 = 0.0
	for _, event := range events {
		result += event.value
	}
	return BasicEvent{name: "OR", value: result}
}

func main() {
	// 创建两个基础事件
	event1 := NewBasicEvent("Event 1", 0.5)
	event2 := NewBasicEvent("Event 2", 0.7)

	// 使用与逻辑组合规则
	andRule := AndRule{}
	andEvent := andRule.Combine([]BasicEvent{event1, event2})

	// 使用或逻辑组合规则
	orRule := OrRule{}
	orEvent := orRule.Combine([]BasicEvent{event1, event2})

	// 输出结果
	fmt.Printf("AND Event: %s, value: %f\n", andEvent.name, andEvent.value)
	fmt.Printf("OR Event: %s, value: %f\n", orEvent.name, orEvent.value)
}
