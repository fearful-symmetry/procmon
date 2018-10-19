package main

import (
	"encoding/json"
	"fmt"

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

// func formatEvtJSON(evt garlic.ProcEvent) string {

// 	var jsonMap map[string]interface{}
// 	switch evt.What {
// 	case garlic.ProcEventFork:
// 		p := evt.EventData.(garlic.Fork)
// 		jsonMap = map[string]interface{}{
// 			"parent_pid":  p.ParentPid,
// 			"parent_tgid": p.ParentTgid,
// 			"child_pid":   p.ChildPid,
// 			"child_tgid":  p.ChildTgid,
// 		}
// 	case garlic.ProcEventExec:
// 		p := evt.EventData.(garlic.Exec)
// 		jsonMap = map[string]interface{}{
// 			"process_pid":  p.ProcessPid,
// 			"process_tgid": p.ProcessTgid,
// 		}
// 	case garlic.ProcEventUID:
// 		p := evt.EventData.(garlic.ID)
// 		jsonMap = map[string]interface{}{
// 			"process_pid":  p.ProcessPid,
// 			"process_tgid": p.ProcessTgid,
// 			"process_ruid": p.RealID,
// 			"process_euid": p.EffectiveID,
// 		}
// 	case garlic.ProcEventGID:
// 		p := evt.EventData.(garlic.ID)
// 		jsonMap = map[string]interface{}{
// 			"process_pid":  p.ProcessPid,
// 			"process_tgid": p.ProcessTgid,
// 			"process_rgid": p.RealID,
// 			"process_egid": p.EffectiveID,
// 		}
// 	case garlic.ProcEventSID:
// 		p := evt.EventData.(garlic.Sid)
// 		jsonMap = map[string]interface{}{
// 			"process_pid":  p.ProcessPid,
// 			"process_tgid": p.ProcessTgid,
// 		}
// 	case garlic.ProcEventPtrace:
// 		p := evt.EventData.(garlic.Ptrace)
// 		jsonMap = map[string]interface{}{
// 			"process_pid":  p.ProcessPid,
// 			"process_tgid": p.ProcessTgid,
// 			"tracer_pid":   p.TracerPid,
// 			"tracer_tgid":  p.TracerTgid,
// 		}
// 	case garlic.ProcEventComm:
// 		p := evt.EventData.(garlic.Comm)
// 		jsonMap = map[string]interface{}{
// 			"process_pid":  p.ProcessPid,
// 			"process_tgid": p.ProcessTgid,
// 			"command":      string(p.Comm[:bytes.IndexByte(p.Comm[:], 0)]),
// 		}
// 	case garlic.ProcEventCoredump:
// 		p := evt.EventData.(garlic.Coredump)
// 		jsonMap = map[string]interface{}{
// 			"process_pid":  p.ProcessPid,
// 			"process_tgid": p.ProcessTgid,
// 		}
// 	case garlic.ProcEventExit:
// 		p := evt.EventData.(garlic.Exit)
// 		jsonMap = map[string]interface{}{
// 			"process_pid":  p.ProcessPid,
// 			"process_tgid": p.ProcessTgid,
// 			"exit_code":    p.ExitCode,
// 			"exit_signal":  p.ExitSignal,
// 		}
// 	}

// }

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
			p.ProcessPid, p.ProcessTgid, p.Comm)
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
