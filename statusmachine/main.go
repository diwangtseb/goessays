package main

import "fmt"

type State int

const (
	OFF State = iota
	ON
	STANDBY
)

func (s State) String() string {
	switch s {
	case OFF:
		return "OFF"
	case ON:
		return "ON"
	case STANDBY:
		return "STANDBY"
	}
	return ""
}

type Event int

const (
	PRESS_OFF Event = iota
	PRESS_ON
)

func (e Event) String() string {
	switch e {
	case PRESS_OFF:
		return "PRESS_OFF"
	case PRESS_ON:
		return "PRESS_ON"
	}
	return ""
}

type Transition func() (State, error)

type StateMachine struct {
	state       State
	transitions map[State]map[Event]Transition
}

func NewStateMachine() *StateMachine {
	sm := &StateMachine{
		state: OFF,
		transitions: map[State]map[Event]Transition{
			OFF: {
				PRESS_ON: func() (State, error) {
					return ON, nil
				},
			},
			ON: {
				PRESS_OFF: func() (State, error) {
					return OFF, nil
				},
				PRESS_ON: func() (State, error) {
					return STANDBY, nil
				},
			},
			STANDBY: {
				PRESS_OFF: func() (State, error) {
					return OFF, nil
				},
				PRESS_ON: func() (State, error) {
					return ON, nil
				},
			},
		},
	}
	return sm
}

func (sm *StateMachine) Transition(event Event) error {
	transition, ok := sm.transitions[sm.state][event]
	if !ok {
		return fmt.Errorf("invalid event %s for state %s", event, sm.state)
	}
	newState, err := transition()
	if err != nil {
		return err
	}
	sm.state = newState
	fmt.Printf("Event: %s, State: %s\n", event, sm.state)
	return nil
}

func main() {
	sm := NewStateMachine()
	sm.Transition(PRESS_ON)
	sm.Transition(PRESS_OFF)
	sm.Transition(PRESS_OFF)
	sm.Transition(PRESS_ON)
	sm.Transition(PRESS_ON)
	sm.Transition(PRESS_OFF)
}
