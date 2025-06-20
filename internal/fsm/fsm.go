package fsm

import "fmt"

type State string
type Event string

const (
	Idle       State = "idle"
	Preparing  State = "preparing"
	Processing State = "processing"
	Exporting  State = "exporting"
	Success    State = "success"
	Error      State = "error"

	EventStart   Event = "start"
	EventFail    Event = "fail"
	EventSuccess Event = "success"
	EventFinish  Event = "finish"
)

type EventPayload struct {
	Event   Event
	Message string
}

type FSM struct {
	state State
}

func New() *FSM {
	return &FSM{
		state: Idle,
	}
}

func (f *FSM) Dispatch(actions map[State]func() EventPayload) {
	if fn, ok := actions[f.state]; ok {
		event := fn()
		f.Send(event)
	}
}

func (f *FSM) IsTerminal() bool {
	return f.state == Success || f.state == Error
}

func (f *FSM) Send(payload EventPayload) {
	switch f.state {
	case Idle:
		if payload.Event == EventStart {
			say("Starting...", payload.Message)
			f.state = Preparing
		}
	case Preparing:
		if payload.Event == EventSuccess {
			say("Files loaded", payload.Message)
			f.state = Processing
		} else if payload.Event == EventFail {
			say("Failed to load files", payload.Message)
			f.state = Error
		}
	case Processing:
		if payload.Event == EventSuccess {
			say("Assembly completed", payload.Message)
			f.state = Exporting
		} else if payload.Event == EventFail {
			say("Assembly failed", payload.Message)
			f.state = Error
		}
	case Exporting:
		if payload.Event == EventSuccess {
			say("Files exported", payload.Message)
			f.state = Success
		} else if payload.Event == EventFail {
			say("Export failed", payload.Message)
			f.state = Error
		}
	case Success:
		if payload.Event == EventFinish {
			say("Done — everything succeeded!", payload.Message)
		}
	case Error:
		if payload.Event == EventFinish {
			say("Stopped due to an error", "")
		}
	}
}

func say(base, extra string) {
	if len(extra) > 0 {
		fmt.Printf("FSM: %s. %s\n", base, extra)
	} else {
		fmt.Println("FSM:", base)
	}
}

func (f *FSM) CurrentState() State {
	return f.state
}
