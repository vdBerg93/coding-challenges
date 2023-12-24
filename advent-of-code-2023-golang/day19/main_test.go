package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"
)

var data []byte
var testData []byte

func TestMain(m *testing.M) {
	var err error
	data, err = os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	testData, err = os.ReadFile("sample")
	if err != nil {
		panic(err)
	}

	os.Exit(m.Run())
}

func Test_Part1(t *testing.T) {
	got := Solve(testData)
	want1 := 62
	if got != want1 {
		t.Fatalf("expected %d, got %d", want1, got)
	}
	fmt.Printf("Part 1 test succeeded.\n")
	fmt.Printf("Solution part 1: %d\n", Solve(data))
}

func Solve(data []byte) int {
	blocks := bytes.Split(data, []byte("\r\n\r\n"))

	wf := parseWorkflows(blocks[0])

	parts := parseParts(blocks[1])

	var accepted []Part
	for _, part := range parts {
		wf.PartAccepted(part)
	}
}

func (wf *Workflows) PartAccepted(part Part) bool {
	i := 0
	for {
		for _, name := range wf.order {
			rules, ok := wf.op[name]
			if !ok {
				panic("missing operation")
			}
			if rules[0].Active(part){

			}

		}
		i++
	}
}

type Workflows struct {
	order []string
	op    map[string][]Operation
}

type Operation struct {
	Source string
	Type   OperationType
	Value  int
}

func (o *Operation) Active(part Part) bool {
	switch o.Type {
	case OperationTypeNone:
		return true
	case OperationTypeGreater:
		return part[o.Source] > o.Value
	case OperationTypeLess:
		return part[o.Source] < o.Value
	}
	panic("impossible")
}

type OperationType int

const (
	OperationTypeNone                  = 0
	OperationTypeGreater OperationType = 1
	OperationTypeLess                  = 2
)

func parseWorkflows(data []byte) Workflows {
	scanner := bufio.NewScanner(bytes.NewReader(data))

	wf := Workflows{
		order: nil,
		op:    make(map[string][]Operation),
	}
	for scanner.Scan() {
		row := scanner.Text()
		fields := strings.Split(row, "{")

		name := fields[0]

		var ops []Operation
		fields = strings.Split(fields[1][:len(fields[1])-1], ",")
		for _, f := range fields {
			ops = append(ops, parseOperation(f))
		}

		wf.op[name] = ops
		wf.order = append(wf.order, name)
	}
	return wf
}
func parseOperation(text string) Operation {
	if strings.Contains(text, ">") {
		f := strings.Split(text, ">")
		val, err := strconv.Atoi(f[2])
		if err != nil {
			panic(err)
		}
		return Operation{
			Source: f[0],
			Type:   OperationTypeGreater,
			Value:  val,
		}
	} else if strings.Contains(text, "<") {
		f := strings.Split(text, "<")
		val, err := strconv.Atoi(f[2])
		if err != nil {
			panic(err)
		}
		return Operation{
			Source: f[0],
			Type:   OperationTypeLess,
			Value:  val,
		}
	} else {
		val, err := strconv.Atoi(text)
		if err != nil {
			panic(err)
		}
		return Operation{
			Type:  OperationTypeNone,
			Value: val,
		}
	}
}

type Part map[string]int

func parseParts(data []byte) []Part {
	scanner := bufio.NewScanner(bytes.NewReader(data))
	var parts []Part
	for scanner.Scan() {
		parts = append(parts, parsePart(scanner.Text()))
	}
	return parts
}

func parsePart(text string) Part {
	part := Part{}
	fields := strings.Split(strings.Trim(text, "{}"), ",")
	for _, f := range fields {
		sf := strings.Split(f, "=")
		val, err := strconv.Atoi(sf[2])
		if err != nil {
			panic(err)
		}
		part[sf[0]] = val
	}
	return part
}
