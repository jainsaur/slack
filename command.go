package main

// ID - command id type
type ID int

// CONST - chat commands
const (
	REG ID = iota
	JOIN
	LEAVE
	MSG
	CHNS
	USRS
	OK
	ERR
)
