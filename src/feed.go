package main

import ()

const (
	Type_Success             = 1
	Type_Failed              = 0
	Action_FindSuccessor     = "FindSuccessor"
	Action_NotifyPredecessor = "NotifyPredecessor"
	Action_GetPredecessor    = "GetPredecessor"
	Action_CheckAlive        = "CheckAlive"
)

type Feed struct {
	Id      string `json:id`
	Title   string `json:title`
	Private bool   `json:private`
	Feed    string `json:feed`
}
