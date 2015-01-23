package alert

import (
	"fmt"
	"os"
)

const (
	Clear = iota
	Green
	Yellow
	Red
)

type Level int

type Alert interface {
	Level() Level
	Reset(newLevel Level, newMessage string)
	Clear()
}

var LevelName = map[Level]string{
	Clear:  "OK",
	Green:  "GREEN",
	Yellow: "YELLOW",
	Red:    "RED",
}

type alert struct {
	id      int
	level   Level
	message string
}

var lastID int

func (this *alert) print() {
	fmt.Fprintf(os.Stdout, "%s\t%d\t%s\n", LevelName[this.level], this.id, this.message)
}

func (this *alert) Level() Level {
	return this.level
}

func (this *alert) Reset(newLevel Level, newMessage string) {
	if newMessage != this.message {
		this.message = newMessage

	}
	if newLevel != this.level {
		fmt.Fprintf(os.Stdout, "%s->%s\t%d\t%s\n",
			LevelName[this.level], LevelName[newLevel],
			this.id,
			this.message)
		this.level = newLevel
	}
}

func (this *alert) Clear() {
	this.level = Clear
}

var New = func(level Level, message string) Alert {
	lastID++
	this := &alert{id: lastID, level: level, message: message}
	this.print()
	return this
}
