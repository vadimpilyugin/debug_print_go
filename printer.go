package main

import (
	"fmt"
	"io"
	"os"
)

type Color string
type Chars string
type Messages string

const (
	Black Color = "\x1b[1;30m"
	Red Color = "\x1b[1;31m"
	Green Color = "\x1b[1;32m"
	Yellow Color = "\x1b[1;33m"
	Blue Color = "\x1b[1;34m"
	// Magenta Color = "\x1b[1;35m"
	// Cyan Color = "\x1b[1;36m"
	White Color = "\x1b[1;37m"
)

const (
	Delim Chars = ": "
	CR Chars = "\r"
  LF Chars = "\n"
  Tab Chars = "\t"
)

const (
	DebugMsg = "Debug"
	AssertMsg = "Assertion failed"
	ErrorMsg = "Error"
	FatalMsg = "Fatal error"
	NoteMsg = "Note"
	EmptyMsg = ""
)

const LOG_EVERY_N int = 5000
var last_in_place bool = false

func generic_print (who string, msg string, in_place bool, 
	params map[string]string, who_color Color, msg_color Color,
	log_every_n bool, line_no int, stderr bool) {

	var out io.Writer
	if stderr == true {
		out = os.Stderr
	} else {
		out = os.Stdout
	}
	if !log_every_n || log_every_n && line_no % LOG_EVERY_N == 0 {
		if last_in_place && !in_place {
      fmt.Fprint(out,LF)
      last_in_place = false
		} else if in_place {
			last_in_place = true
		}

		fmt.Fprint(out,who_color,who,Delim,msg_color,msg)
		if in_place {
			fmt.Fprint(out,CR)
		} else {
			fmt.Fprint(out,LF, White)
			for s1,s2 := range params {
				fmt.Fprint(out,Tab, s1, Delim, s2, LF)
			}
		}
	}
}

func Debug (args ...interface{}) {
	var (
		who string = DebugMsg
		msg string
		params map[string]string
		in_place bool = false
		log_every_n bool = false
		line_no int = 0 
		stderr bool = false
	)

	switch len(args) {
		case 0: return
		case 1: 
			msg = args[0].(string)
		case 2:
			who = args[1].(string)
			msg = args[0].(string)
		case 3:
			who = args[1].(string)
			msg = args[0].(string)
			params = args[2].(map[string]string)
		case 4:
			who = args[1].(string)
			msg = args[0].(string)
			params = args[2].(map[string]string)
			in_place = args[3].(bool)
		case 6:
			who = args[1].(string)
			msg = args[0].(string)
			params = args[2].(map[string]string)
			in_place = args[3].(bool)
			log_every_n = args[4].(bool)
			line_no = args[5].(int)
	}
	generic_print(
		who,
		msg,
		in_place,
		params,
		Green,
		White,
		log_every_n,
		line_no,
		stderr,
	)
}

func Note (args ...interface{}) {
	var (
		who string = NoteMsg
		msg string
		params map[string]string
		in_place bool = false
		log_every_n bool = false
		line_no int = 0 
		stderr bool = false
	)

	switch len(args) {
		case 0: return
		case 1: 
			msg = args[0].(string)
		case 2:
			who = args[1].(string)
			msg = args[0].(string)
		case 3:
			who = args[1].(string)
			msg = args[0].(string)
			params = args[2].(map[string]string)
		case 4:
			who = args[1].(string)
			msg = args[0].(string)
			params = args[2].(map[string]string)
			in_place = args[3].(bool)
		case 6:
			who = args[1].(string)
			msg = args[0].(string)
			params = args[2].(map[string]string)
			in_place = args[3].(bool)
			log_every_n = args[4].(bool)
			line_no = args[5].(int)
	}
	generic_print(
		who,
		msg,
		in_place,
		params,
		Yellow,
		White,
		log_every_n,
		line_no,
		stderr,
	)
}

func Error (args ...interface{}) {
	var (
		who string = ErrorMsg
		msg string
		params map[string]string
		in_place bool = false
		log_every_n bool = false
		line_no int = 0 
		stderr bool = true
	)

	switch len(args) {
		case 0: return
		case 1: 
			msg = args[0].(string)
		case 2:
			who = args[1].(string)
			msg = args[0].(string)
		case 3:
			who = args[1].(string)
			msg = args[0].(string)
			params = args[2].(map[string]string)
	}
	generic_print(
		who,
		msg,
		in_place,
		params,
		Red,
		White,
		log_every_n,
		line_no,
		stderr,
	)
}

func Fatal (args ...interface{}) {
	var (
		who string = FatalMsg
		msg string
		params map[string]string
		in_place bool = false
		log_every_n bool = false
		line_no int = 0 
		stderr bool = true
	)

	switch len(args) {
		case 0: panic(msg)
		case 1: 
			msg = args[0].(string)
		case 2:
			who = args[1].(string)
			msg = args[0].(string)
		case 3:
			who = args[1].(string)
			msg = args[0].(string)
			params = args[2].(map[string]string)
	}
	generic_print(
		who,
		msg,
		in_place,
		params,
		Red,
		White,
		log_every_n,
		line_no,
		stderr,
	)
	panic(msg)
}

func fatal_func() {
	defer func() {
		err := recover()
		Debug(err,"Recovered from fatal error")
	}()
	Fatal("Fatal error","Fatal msg")
}

func main () {
	
	Debug("Hello, world!","Debug msg")
	Note("Hello, world!","Note msg")
	Error("Hello, world!","Error msg")
	fatal_func()
}