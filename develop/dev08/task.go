package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"

	ps "github.com/mitchellh/go-ps"
)

func main() {
	shell()
}

func shell() {
	sc := bufio.NewScanner(os.Stdin)
	header()
	for sc.Scan() {
		str := sc.Text()
		if str == "exit" {
			return
		}

		parsePipeline(os.Stdout, str)
		header()
	}
}

func parsePipeline(out io.Writer, line string) {
	cmnds := strings.Split(line, "|")
	var res string
	var err error
	for _, v := range cmnds {
		res, err = parseCommand(res, v)
		if err != nil {
			fmt.Fprintln(out, err)
		}
	}

	if res != "" {
		fmt.Fprintln(out, res)
	}
}

func parseCommand(in string, cmd string) (string, error) {
	str := strings.Fields(cmd)
	if len(str) == 0 {
		return "", errors.New("empty command")
	}
	if in != "" {
		str = append(str, in)
	}
	switch str[0] {
	case "cd":
		if len(str) != 2 {
			return "", errors.New("cd: must be 1 parametr")
		}
		if err := os.Chdir(str[1]); err != nil {
			return "", err
		}
	case "pwd":
		if len(str) != 1 {
			return "", errors.New("pwd: unused paramters")
		}
		res, err := os.Getwd()
		if err != nil {
			return "", err
		}
		return res, nil
	case "echo":
		if len(str) != 2 {
			return "", errors.New("echo: must be 1 parametr")
		}
		return str[1], nil
	case "kill":
		if len(str) < 2 {
			return "", errors.New("kill: not enough parametrs")
		}
		for i := 1; i < len(str); i++ {
			pid, err := strconv.Atoi(str[i])
			if err != nil {
				return "", err
			}
			if err = syscall.Kill(pid, syscall.SIGINT); err != nil {
				return "", err
			}
		}
		return "", nil
	case "ps":
		procs, err := ps.Processes()
		if err != nil {
			return "", err
		}
		var strBldr strings.Builder
		strBldr.WriteString("\tPID\tCMD\n")
		for _, proc := range procs {
			strBldr.WriteString(fmt.Sprintf("\t%v\t%v\n", proc.Pid(), proc.Executable()))
		}
		return strBldr.String(), nil
	case "fork()":
		id, _, _ := syscall.Syscall(syscall.SYS_FORK, 0, 0, 0)
		return strconv.Itoa(int(id)), nil
	case "exec":
		if len(str) < 2 {
			return "", errors.New("exec: not enough parameters")
		}
		cmd := exec.Command(str[1], str[2:]...)
		stdout, err := cmd.Output()
		if err != nil {
			return "", err
		}
		return string(stdout), err
	case "netcat":
		if len(str) == 2 {
			res, err := netcat(str[1], false)
			if err != nil {
				return "", err
			}
			return res, err
		} else if len(str) == 3 {
			if str[1] == "-u" {
				res, err := netcat(str[2], true)
				if err != nil {
					return "", err
				}
				return res, err
			}
		}
	}
	return "", nil
}

func netcat(addr string, UDP bool) (string, error) {
	network := "tcp"
	if UDP {
		network = "udp"
	}
	con, err := net.Dial(network, addr)
	if err != nil {
		return "", err
	}
	defer con.Close()

	err = stdinToConn(con)
	return "", err
}

func stdinToConn(con net.Conn) error {
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		txt := sc.Text()
		if txt == "exit" {
			return nil
		}
		_, err := con.Write([]byte(txt))
		if err != nil {
			return err
		}
	}
	return nil
}

func header() {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Printf("[shell] %v$ ", wd)
}
