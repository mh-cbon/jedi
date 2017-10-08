package runtime

import (
	"fmt"
	"io"
)

// EventReceiver writes on W
type EventReceiver struct {
	W io.Writer
}

// Event receives a simple notification when various events occur
func (n *EventReceiver) Event(eventName string) {
	fmt.Fprintf(n.W, "eventName %q\n", eventName)
}

// EventKv receives a notification when various events occur along with
// optional key/value data
func (n *EventReceiver) EventKv(eventName string, kvs map[string]string) {
	fmt.Fprintf(n.W, "eventName %q kvs %#v\n", eventName, kvs)
}

// EventErr receives a notification of an error if one occurs
func (n *EventReceiver) EventErr(eventName string, err error) error {
	fmt.Fprintf(n.W, "eventName %q err %v\n", eventName, err)
	return err
}

// EventErrKv receives a notification of an error if one occurs along with
// optional key/value data
func (n *EventReceiver) EventErrKv(eventName string, err error, kvs map[string]string) error {
	fmt.Fprintf(n.W, "eventName %q kvs %#v err %v\n", eventName, kvs, err)
	return err
}

// Timing receives the time an event took to happen
func (n *EventReceiver) Timing(eventName string, nanoseconds int64) {}

// TimingKv receives the time an event took to happen along with optional key/value data
func (n *EventReceiver) TimingKv(eventName string, nanoseconds int64, kvs map[string]string) {}
