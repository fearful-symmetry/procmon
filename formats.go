package main

import (
	"bytes"
	"fmt"

	"github.com/fearful-symmetry/garlic"
)

//functions to handle formatting events

//string out that goes to stdout
func formatEvtStr(evt garlic.ProcEvent) string {

	var baseStr string
	switch evt.What {
	case garlic.ProcEventFork:
		p := evt.EventData.(garlic.Fork)
		baseStr = fmt.Sprintf("\tParent PID: %d\n\tParent TGID: %d\n\tChild PID: %d\n\tChild TGID: %d\n",
			p.ParentPid, p.ParentTgid, p.ChildPid, p.ChildTgid)
	case garlic.ProcEventExec:
		p := evt.EventData.(garlic.Exec)
		baseStr = fmt.Sprintf("\tProcess PID: %d\n\tProcess TGID: %d\n",
			p.ProcessPid, p.ProcessTgid)
	case garlic.ProcEventUID:
		p := evt.EventData.(garlic.ID)
		baseStr = fmt.Sprintf("\tProcess PID: %d\n\tProcess TGID: %d\n\tProcess RUID: %d\n\tProcess EUID: %d\n",
			p.ProcessPid, p.ProcessTgid, p.RealID, p.EffectiveID)
	case garlic.ProcEventGID:
		p := evt.EventData.(garlic.ID)
		baseStr = fmt.Sprintf("\tProcess PID: %d\n\tProcess TGID: %d\n\tProcess RGID: %d\n\tProcess EGID: %d\n",
			p.ProcessPid, p.ProcessTgid, p.RealID, p.EffectiveID)
	case garlic.ProcEventSID:
		p := evt.EventData.(garlic.Sid)
		baseStr = fmt.Sprintf("\tProcess PID: %d\n\tProcess TGID: %d\n",
			p.ProcessPid, p.ProcessTgid)
	case garlic.ProcEventPtrace:
		p := evt.EventData.(garlic.Ptrace)
		baseStr = fmt.Sprintf("\tProcess PID: %d\n\tProcess TGID: %d\n\tTracer PID: %d\n\tTracer TGID: %d\n",
			p.ProcessPid, p.ProcessTgid, p.TracerPid, p.TracerTgid)
	case garlic.ProcEventComm:
		p := evt.EventData.(garlic.Comm)
		//bless me father for I have sinned
		//Convert the comm array to a string while making sure we don't print null bytes
		baseStr = fmt.Sprintf("\tProcess PID: %d\n\tProcess TGID: %d\n\tCommand: %s\n",
			p.ProcessPid, p.ProcessTgid, string(p.Comm[:bytes.IndexByte(p.Comm[:], 0)]))
	case garlic.ProcEventCoredump:
		p := evt.EventData.(garlic.Coredump)
		baseStr = fmt.Sprintf("Process PID: %d\n\tProcess TGID: %d\n",
			p.ProcessPid, p.ProcessTgid)
	case garlic.ProcEventExit:
		p := evt.EventData.(garlic.Exit)
		baseStr = fmt.Sprintf("\tProcess PID: %d\n\tProcess TGID: %d\n\t Exit Code: %d\n\t Exit Signal: %d\n",
			p.ProcessPid, p.ProcessTgid, p.ExitCode, p.ExitSignal)
	}

	return baseStr
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
