package main

import (
	"fmt"
	"log"
	"os"

	"github.com/fearful-symmetry/garlic"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

func handleArg(uargs []string) []garlic.EventType {

	events := make([]garlic.EventType, len(uargs))

	for i, evt := range uargs {
		switch evt {
		case "fork":
			events[i] = garlic.ProcEventFork
		case "exec":
			events[i] = garlic.ProcEventExec
		case "uid":
			events[i] = garlic.ProcEventUID
		case "gid":
			events[i] = garlic.ProcEventGID
		case "sid":
			events[i] = garlic.ProcEventSID
		case "ptrace":
			events[i] = garlic.ProcEventPtrace
		case "comm":
			events[i] = garlic.ProcEventPtrace
		case "coredump":
			events[i] = garlic.ProcEventCoredump
		case "exit":
			events[i] = garlic.ProcEventExit
		}
	}

	return events
}

func runMon(events []garlic.EventType, json bool) {

	le := log.New(os.Stderr, "procmon", 0)
	var ev garlic.CnConn
	var err error
	if events != nil {
		ev, err = garlic.DialPCNWithEvents(events)
		if err != nil {
			le.Fatalf("Could not dial proc connector: %s", err)
		}
	} else {
		ev, err = garlic.DialPCN()
		if err != nil {
			le.Fatalf("Could not dial proc connector: %s", err)
		}
	}

	for {
		evt, err := ev.ReadPCN()
		if err != nil {
			le.Printf("Error Reading Event: %s", err)
			continue
		}
		for _, singleEvt := range evt {
			evtString := formatEvtPretty(singleEvt)
			if json {
				json, err := formatEvtJSON(singleEvt)
				if err != nil {
					le.Fatalf("Error parsing to JSON: %s", err)
				}
				fmt.Println(json)
			} else {
				fmt.Printf("Got %s event on CPU %d at %s %s\n",
					formatEvtType(singleEvt.What),
					singleEvt.CPU,
					singleEvt.Timestamp.Local(),
					evtString)
			}

		}

	}

}

func main() {

	var (
		procCLI = kingpin.New("procmon", "Monitor process events from the command line")
		verbose = procCLI.Flag("verbose", "verbose mode").Short('v').Bool()
		events  = procCLI.Arg("event", "Event(s) to watch").Enums("fork", "exec", "uid", "gid", "sid", "ptrace", "comm", "coredump", "exit")
		isJSON  = procCLI.Flag("json", "output NDJSON").Bool()
	)

	kingpin.MustParse(procCLI.Parse(os.Args[1:]))

	var evtList []garlic.EventType
	if len(*events) == 0 {
		log.Printf("Reading all events with modes %v and %v", *isJSON, *verbose)
		evtList = nil
	} else {
		log.Printf("Reading events %v with modes %v and %v", *events, *isJSON, *verbose)
		evtList = handleArg(*events)
	}

	runMon(evtList, *isJSON)

}
