package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"
)

var dataInput []byte
var dataSample1 []byte

func TestMain(m *testing.M) {
	var err error
	dataInput, err = os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	dataSample1, err = os.ReadFile("sample1")
	if err != nil {
		panic(err)
	}

	os.Exit(m.Run())
}

const debug = false

func Test_Part1(t *testing.T) {
	got := Solve(dataSample1, 1e3)
	want1 := uint64(32000000)
	if got != want1 {
		t.Fatalf("expected %d, got %d", want1, got)
	}

	fmt.Printf("Part 1 test succeeded.\n")
	got = Solve(dataInput, 1e3)
	fmt.Printf("Solution part 1: %d\n", got)
}

func Solve(data []byte, count int) uint64 {
	scanner := bufio.NewScanner(bytes.NewReader(data))
	for scanner.Scan() {
		ParseModule(scanner.Text())
	}

	for _, m := range modules {
		for _, o := range m.Outputs {
			dest, ok := modules[o]
			if !ok {
				continue
			}
			if dest.Type == ModuleTypeConjunction {
				dest.History[m.Name] = false
			}
		}
	}
	var totalLow, totalHigh uint64
	for i := 1; i <= count; i++ {
		incrementCounter(false)
		todo = []Schedule{
			{
				Module: "broadcaster",
				Input:  false,
				Source: "button",
			},
		}
		if debug {
			fmt.Printf("button false broadcaster\n")
		}
		for {
			s := todo[0]
			m := GetModule(s.Module)
			if m == nil {
				if len(todo) == 1 {
					break
				}
				todo = todo[1:]
				continue
			}
			(*m).Receive(s.Input, s.Source)
			if len(todo) == 1 {
				break
			}
			todo = todo[1:]
		}
		if debug {
			fmt.Printf("lfow=%d, high=%d\n", lowPulses, highPulses)
		}
		totalLow += lowPulses
		totalHigh += highPulses
		lowPulses, highPulses = 0, 0

	}
	return totalLow * totalHigh
}

func ParseModule(text string) {
	text = strings.Replace(text, " ", "", -1)
	if strings.HasPrefix(text, "&") {
		f := strings.Split(text[1:], "->")
		name := f[0]
		outputs := strings.Split(f[1], ",")
		modules[name] = NewModule(name, outputs, ModuleTypeConjunction)

	} else if strings.HasPrefix(text, "%") {
		f := strings.Split(text[1:], "->")
		name := f[0]
		outputs := strings.Split(f[1], ",")
		modules[name] = NewModule(name, outputs, ModuleTypeFlipFlop)

	} else {
		f := strings.Split(text, "->")
		name := f[0]
		outputs := strings.Split(f[1], ",")
		modules[name] = NewModule(name, outputs, ModuleTypeBroadcaster)
	}
}

var lowPulses, highPulses uint64

func incrementCounter(pulse bool) {
	if pulse {
		highPulses++
	} else {
		lowPulses++
	}
}

type Schedule struct {
	Module string
	Input  bool
	Source string
}

var todo []Schedule
var modules = make(map[string]*Module)

func GetModule(name string) *Module {
	m, ok := modules[name]
	if !ok {
		return nil
	}
	return m
}

type ModuleType int

const (
	ModuleTypeBroadcaster ModuleType = 0
	ModuleTypeFlipFlop    ModuleType = 1
	ModuleTypeConjunction ModuleType = 2
)

type Module struct {
	Type    ModuleType
	Name    string
	State   bool
	Outputs []string
	History map[string]bool
}

func NewModule(name string, outputs []string, Type ModuleType) *Module {
	return &Module{
		Type:    Type,
		Name:    name,
		State:   false,
		Outputs: outputs,
		History: make(map[string]bool),
	}
}

// Process - Toggles the switch on low Schedule
func (ff *Module) Receive(input bool, source string) {
	switch ff.Type {
	case ModuleTypeBroadcaster:
		for _, output := range ff.Outputs {
			todo = append(todo, Schedule{
				Module: output,
				Input:  input,
				Source: ff.Name,
			})

			if debug {
				fmt.Printf("%s %t %s\n", ff.Name, input, output)
			}
			incrementCounter(input)
		}
	case ModuleTypeFlipFlop:
		if !input {
			ff.State = !ff.State
			for _, output := range ff.Outputs {
				todo = append(todo, Schedule{
					Module: output,
					Input:  ff.State,
					Source: ff.Name,
				})

				if debug {
					fmt.Printf("%s %t %s\n", ff.Name, ff.State, output)
				}
				incrementCounter(ff.State)
			}
		}
	case ModuleTypeConjunction:
		ff.History[source] = input
		allHigh := true
		for _, state := range ff.History {
			if state == false {
				allHigh = false
				break
			}
		}
		for _, output := range ff.Outputs {
			todo = append(todo, Schedule{
				Module: output,
				Input:  !allHigh,
				Source: ff.Name,
			})

			if debug {
				fmt.Printf("%s %t %s\n", ff.Name, !allHigh, output)
			}
			incrementCounter(!allHigh)
		}
	default:
		panic("not implemented")
	}
}
