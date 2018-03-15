package printer

import (
	"fmt"
	"io"
	"os"
)

type color string
type chars string

const (
	pref = "\x1b[1;"
	suff = "m"
	ERROR = 2

	// Black color = "\x1b[1;30m"
	red color = pref+"31"+suff
	green color = pref+"32"+suff
	yellow color = pref+"33"+suff
	blue color = pref+"34"+suff
	// Magenta color = "\x1b[1;35m"
	// Cyan color = "\x1b[1;36m"
	white color = pref+"37"+suff
)

const (
	delim chars = ": "
	cr chars = "\r"
  lf chars = "\n"
  tab chars = "\t"
)

const (
	debugMsg = "Debug"
	assertMsg = "Assertion failed"
	errorMsg = "Error"
	fatalMsg = "Fatal error"
	noteMsg = "Note"
	emptyMsg = ""
)

const logEveryN int = 5000
var last_in_place bool = false

func generic_print (who string, msg string, in_place bool, 
	params map[string]string, who_color color, msg_color color,
	log_every_n bool, line_no int, stderr bool) {

	var out io.Writer
	if stderr == true {
		out = os.Stderr
	} else {
		out = os.Stdout
	}
	if !log_every_n || log_every_n && line_no % logEveryN == 0 {
		if last_in_place && !in_place {
      fmt.Fprint(out,cr+lf)
      last_in_place = false
		} else if in_place {
			last_in_place = true
		}
    if who == "" {
      fmt.Fprint(out,msg_color,msg)
    } else {
		  fmt.Fprint(out,who_color,who,delim,msg_color,msg)
		}
		if in_place {
			fmt.Fprint(out,cr)
		} else {
			fmt.Fprint(out,cr+lf, white)
			for s1,s2 := range params {
				fmt.Fprint(out,tab, s1, delim, s2, cr+lf)
			}
		}
	}
}

func Debug (args ...interface{}) {
	var (
		who string = debugMsg
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
			msg = fmt.Sprintf("%s",args[0])
		case 2:
			who = fmt.Sprintf("%s",args[1])
			msg = fmt.Sprintf("%s",args[0])
		case 3:
			who = fmt.Sprintf("%s",args[1])
			msg = fmt.Sprintf("%s",args[0])
			params = args[2].(map[string]string)
		case 4:
			who = fmt.Sprintf("%s",args[1])
			msg = fmt.Sprintf("%s",args[0])
			params = args[2].(map[string]string)
			in_place = args[3].(bool)
		case 6:
			who = fmt.Sprintf("%s",args[1])
			msg = fmt.Sprintf("%s",args[0])
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
		green,
		white,
		log_every_n,
		line_no,
		stderr,
	)
}

func Note (args ...interface{}) {
	var (
		who string = noteMsg
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
			msg = fmt.Sprintf("%s",args[0])
		case 2:
			who = fmt.Sprintf("%s",args[1])
			msg = fmt.Sprintf("%s",args[0])
		case 3:
			who = fmt.Sprintf("%s",args[1])
			msg = fmt.Sprintf("%s",args[0])
			params = args[2].(map[string]string)
		case 4:
			who = fmt.Sprintf("%s",args[1])
			msg = fmt.Sprintf("%s",args[0])
			params = args[2].(map[string]string)
			in_place = args[3].(bool)
		case 6:
			who = fmt.Sprintf("%s",args[1])
			msg = fmt.Sprintf("%s",args[0])
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
		yellow,
		white,
		log_every_n,
		line_no,
		stderr,
	)
}

func Error (args ...interface{}) {
	var (
		who string = errorMsg
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
			msg = fmt.Sprintf("%s",args[0])
		case 2:
			who = fmt.Sprintf("%s",args[1])
			msg = fmt.Sprintf("%s",args[0])
		case 3:
			who = fmt.Sprintf("%s",args[1])
			msg = fmt.Sprintf("%s",args[0])
			params = args[2].(map[string]string)
	}
	generic_print(
		who,
		msg,
		in_place,
		params,
		red,
		white,
		log_every_n,
		line_no,
		stderr,
	)
}

func Fatal (args ...interface{}) {
	var (
		who string = fatalMsg
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
			msg = fmt.Sprintf("%s",args[0])
		case 2:
			who = fmt.Sprintf("%s",args[1])
			msg = fmt.Sprintf("%s",args[0])
		case 3:
			who = fmt.Sprintf("%s",args[1])
			msg = fmt.Sprintf("%s",args[0])
			params = args[2].(map[string]string)
	}
	generic_print(
		who,
		msg,
		in_place,
		params,
		red,
		white,
		log_every_n,
		line_no,
		stderr,
	)
	os.Exit(ERROR)
}
