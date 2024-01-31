package state

var currentState ServerState = Unknown

type ServerState string

const (
	Unknown   ServerState = "unknwon"
	Running   ServerState = "running"
	Rebooting ServerState = "rebooting"
)

func Update(newState ServerState) {
	currentState = newState
}

func Current() ServerState {
	return currentState
}

func IsRunning() bool {
	return currentState == Running
}
