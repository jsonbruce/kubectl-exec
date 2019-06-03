package main

import (
	"os"
	"flag"
	"fmt"
	"strings"
)

type StringSliceVar []string

func (ss *StringSliceVar) Set(value string) error {
	*ss = StringSliceVar(strings.Split(value, " "))
	return nil
}

func (ss *StringSliceVar) String() string {
	s := fmt.Sprintf("%s", *ss)
	return s[1:len(s)-1]
}

var (
	pod string
	container string
	namespace string
	tty bool
	stdin bool
	commands StringSliceVar
)

func usage() {
	s := "kubectl ws exec POD_NAME -i -t -m "
	fmt.Printf("%s\n", s)
}

func init() {
	flag.StringVar(&container, "c", "", "name of container")
	flag.StringVar(&namespace, "n", "default", "name of namespace")
	flag.BoolVar(&tty, "t", false, "use TTY")
	flag.BoolVar(&stdin, "i", false, "use Stdin")
	flag.Var(&commands, "m", "list of commands seprated by blank")
	flag.Usage = usage
}

func main() {
	if len(os.Args) == 1 {
		flag.PrintDefaults()
		os.Exit(0)
	}

	pod = os.Args[1]

	os.Args = os.Args[1:]
	flag.Parse()

	fmt.Printf("Pod=%s\nContainer=%s\nNamespace=%s\nTTY=%v\nStdin=%v\nCommands=%s\n", pod, container, namespace, tty, stdin, commands)
}
