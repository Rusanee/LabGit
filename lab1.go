package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	cpu1  string
	cpu2  string
	ready []string
	io1   []string
	io2   []string
	io3   []string
	io4   []string
)

func main() {
	initialized()
	for {
		showProcess()
		command, err := getCommand()
		if err != nil {
			fmt.Println("เกิดข้อผิดพลาดในการอ่านข้อมูล:", err)
			return
		}
		switch strings.TrimSpace(command) {
		case "exit":
			fmt.Println("---Exit Progarm---")
			return
		case "new":
			fmt.Print("Enter name -> ")
			name, _ := getCommand()
			new_p(name)
		case "expire":
			fmt.Print("Which one ? -> ")
			name, _ := getCommand()
			expire_p(name)
		case "terminate":
			fmt.Print("Which one ? -> ")
			name, _ := getCommand()
			terminate_p(name)
		case "io1":
			fmt.Print("Which one ? -> ")
			name, _ := getCommand()
			io1_p(name)
		case "io2":
			fmt.Print("Which one ? -> ")
			name, _ := getCommand()
			io2_p(name)
		case "io3":
			fmt.Print("Which one ? -> ")
			name, _ := getCommand()
			io3_p(name)
		case "io4":
			fmt.Print("Which one ? -> ")
			name, _ := getCommand()
			io4_p(name)
		case "io1p_x":
			iop_x(io1)
		case "io2p_x":
			iop_x(io2)
		case "io3p_x":
			iop_x(io3)
		case "io4p_x":
			iop_x(io4)
		default:
			fmt.Printf("\n!!-Command Error-!!\n")
		}
	}
}

func initialized() {
	cpu1 = ""
	cpu2 = ""
	ready = make([]string, 10)
	io1 = make([]string, 10)
	io2 = make([]string, 10)
	io3 = make([]string, 10)
	io4 = make([]string, 10)
}

func showProcess() {
	fmt.Printf("\n-----------\n")
	fmt.Printf("CPU1   -> %s\n", cpu1)
	fmt.Printf("CPU2  -> %s\n", cpu2)

	fmt.Printf("Ready -> ")
	for i := range ready {
		if ready[i] != "" {
			fmt.Printf("%s ", ready[i])
		}
	}
	fmt.Printf("\n")

	fmt.Printf("I/O 1 -> ")
	for i := range io1 {
		fmt.Printf("%s ", io1[i])
	}
	fmt.Printf("\n")

	fmt.Printf("I/O 2 -> ")
	for i := range io2 {
		fmt.Printf("%s ", io2[i])
	}
	fmt.Printf("\n")

	fmt.Printf("I/O 3 -> ")
	for i := range io3 {
		fmt.Printf("%s ", io3[i])
	}
	fmt.Printf("\n")

	fmt.Printf("I/O 4 -> ")
	for i := range io4 {
		fmt.Printf("%s ", io4[i])
	}

	fmt.Printf("\n\nCommand > ")
}

func getCommand() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	data, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	data = strings.Trim(data, "\n")
	return strings.TrimSpace(strings.ToLower(data)), nil
}

func insertQueue(que []string, data string) {
	for i := range que {
		if que[i] == "" {
			que[i] = strings.TrimSpace(data)
			//fmt.Printf("--add data--")
			break
		}
	}
}

func deleteQueue(q []string) string {
	result := q[0]
	for i := range q {
		if i == 0 {
			continue
		}
		q[i-1] = q[i]
	}
	q[9] = ""
	return result
}

func new_p(data string) {
	if cpu1 == "" {
		cpu1 = data
	} else if cpu2 == "" {
		cpu2 = data
	} else {
		insertQueue(ready, data)
	}
}

func expire_p(data string) {
	switch strings.TrimSpace(data) {
	case "cpu1":
		q := deleteQueue(ready)
		if q == "" {
			return
		}
		insertQueue(ready, cpu1)
		cpu1 = q
	case "cpu2":
		q := deleteQueue(ready)
		if q == "" {
			return
		}
		insertQueue(ready, cpu2)
		cpu2 = q
	default:
		fmt.Printf("\n!!-Command expire Error-!!\n")
	}
}

func terminate_p(data string) {
	switch strings.TrimSpace(data) {
	case "cpu1":
		if cpu1 != "" {
			cpu1 = deleteQueue(ready)
		} else {
			fmt.Println("No program in process")
		}
	case "cpu2":
		if cpu2 != "" {
			cpu2 = deleteQueue(ready)
		} else {
			fmt.Println("No program in process")
		}
	default:
		fmt.Printf("\n!!-Command Error-!!\n")
	}
}

func io1_p(data string) {
	switch strings.TrimSpace(data) {
	case "cpu1":
		insertQueue(io1, cpu1)
		cpu1 = ""
		expire_p("cpu1")
	case "cpu2":
		insertQueue(io1, cpu2)
		cpu2 = ""
		expire_p("cpu2")
	default:
		fmt.Printf("\n!!-Command Error-!!\n")
	}
}

func io2_p(data string) {
	switch strings.TrimSpace(data) {
	case "cpu1":
		insertQueue(io2, cpu1)
		cpu1 = ""
		expire_p("cpu1")
	case "cpu2":
		insertQueue(io2, cpu2)
		cpu2 = ""
		expire_p("cpu2")
	default:
		fmt.Printf("\n!!-Command Error-!!\n")
	}
}

func io3_p(data string) {
	switch strings.TrimSpace(data) {
	case "cpu1":
		insertQueue(io3, cpu1)
		cpu1 = ""
		expire_p("cpu1")
	case "cpu2":
		insertQueue(io3, cpu2)
		cpu2 = ""
		expire_p("cpu2")
	default:
		fmt.Printf("\n!!-Command Error-!!\n")
	}
}

func io4_p(data string) {
	switch strings.TrimSpace(data) {
	case "cpu1":
		insertQueue(io4, cpu1)
		cpu1 = ""
		expire_p("cpu1")
	case "cpu2":
		insertQueue(io4, cpu2)
		cpu2 = ""
		expire_p("cpu2")
	default:
		fmt.Printf("\n!!-Command Error-!!\n")
	}
}

func iop_x(q []string) {
	p := deleteQueue(q)
	if p == "" {
		fmt.Printf("\nNo program in io\n")
		return
	}
	if cpu1 == "" {
		cpu1 = p
	} else if cpu2 == "" {
		cpu2 = p
	} else {
		insertQueue(ready, p)
	}
}
