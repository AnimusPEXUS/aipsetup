package main

type FileOperation uint

const (
	OperationNone FileOperation = iota
	OperationCut
	OperationCopy
)
