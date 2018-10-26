package main

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/fearful-symmetry/garlic"
)

//functions to handle formatting events

//json string out that goes to stdout
func formatEvtJSON(evt garlic.ProcEvent) (string, error) {

	j, err := json.Marshal(evt)
	if err != nil {
		return "", err
	}

	return string(j), nil

}

//use reflection to pretty-print
func formatEvtPretty(evt garlic.ProcEvent) string {
	t := reflect.TypeOf(evt.EventData)
	v := reflect.ValueOf(evt.EventData)
	var out string

	for index := 0; index < t.NumField(); index++ {
		field := t.Field(index)
		tag := field.Tag.Get("pretty")
		//this comm type is always making my life harder
		if fv, ok := v.Field(index).Interface().(string); ok {
			out = fmt.Sprintf("%s\n\t %s: %s", out, tag, fv)
		} else {
			out = fmt.Sprintf("%s\n\t %s: %d", out, tag, v.Field(index).Interface())
		}

	}

	return out
}

//Turn the event into a human-readable string
func formatEvtType(evt garlic.EventType) string {

	var event string

	switch evt {
	case garlic.ProcEventFork:
		event = "Fork"
	case garlic.ProcEventExec:
		event = "Exec"
	case garlic.ProcEventUID:
		event = "UID"
	case garlic.ProcEventGID:
		event = "GID"
	case garlic.ProcEventSID:
		event = "SID"
	case garlic.ProcEventPtrace:
		event = "Ptrace"
	case garlic.ProcEventComm:
		event = "Command"
	case garlic.ProcEventCoredump:
		event = "Core Dump"
	case garlic.ProcEventExit:
		event = "Exit"
	}

	return event
}
