package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"

	"github.com/fearful-symmetry/garlic"
)

//functions to handle formatting events

//print the event based on the selected config
func printEvent(cfg garlicCfg, singleEvt garlic.ProcEvent) (string, error) {
	if cfg.IsJSON {
		return formatEvtJSON(singleEvt, cfg)

	}
	var ts time.Time
	if cfg.IsUTC {
		ts = singleEvt.Timestamp.UTC()
	} else {
		ts = singleEvt.Timestamp.Local()
	}
	return fmt.Sprintf("Got %s event on CPU %d at %s %s\n",
		singleEvt.WhatString,
		singleEvt.CPU,
		ts,
		formatEvtPretty(singleEvt)), nil

}

//json string out that goes to stdout
func formatEvtJSON(evt garlic.ProcEvent, cfg garlicCfg) (string, error) {

	//yah, this is...kinda hacky.
	if cfg.IsUTC {
		evt.Timestamp = evt.Timestamp.UTC()
	}

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
