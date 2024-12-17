package main

import (
	_ "embed"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

var exampleRegisters = []int{729, 0, 0}
var exampleProgram = []int{0, 1, 5, 4, 3, 0}

var inputRegisters = []int{27334280, 0, 0}
var inputProgram = []int{2, 4, 1, 2, 7, 5, 0, 3, 1, 7, 4, 1, 5, 5, 3, 0}

func main() {
	fmt.Println("Part1 example:", solve(exampleRegisters, exampleProgram))
	fmt.Println("Part1:", solve(inputRegisters, inputProgram))
}

func solve(registers, program []int) string {
	c := Computer{
		program: program,
		A:       registers[0],
		B:       registers[1],
		C:       registers[2],
		ptr:     0,
	}
	for c.Run() {
	}

	return strings.Join(c.output, ",")
}

type Computer struct {
	program []int
	output  []string
	ptr     int
	A, B, C int
}

func (c *Computer) Run() bool {
	opcode := c.program[c.ptr]
	c.Execute(opcode)
	return c.ptr < len(c.program)
}

func (c *Computer) LiteralOperand() int {
	return c.program[c.ptr+1]
}

func (c *Computer) ComboOperand() int {
	v := c.program[c.ptr+1]
	switch v {
	case 0, 1, 2, 3:
		return v
	case 4:
		return c.A
	case 5:
		return c.B
	case 6:
		return c.C
	default:
		log.Panicf("reserved operand %v", v)
		return 0
	}
}

func (c *Computer) Execute(opcode int) {
	ptr := c.ptr
	switch opcode {
	case 0:
		c.A = int(math.Floor(float64(c.A) / math.Pow(2, float64(c.ComboOperand()))))
	case 1:
		c.B = c.LiteralOperand() ^ c.B
	case 2:
		c.B = c.ComboOperand() % 8
	case 3:
		if c.A != 0 {
			c.ptr = c.LiteralOperand()
		}
	case 4:
		c.B = c.B ^ c.C
	case 5:
		o := c.ComboOperand() % 8
		c.output = append(c.output, strconv.Itoa(o))
	case 6:
		c.B = int(math.Floor(float64(c.A) / math.Pow(2, float64(c.ComboOperand()))))
	case 7:
		c.C = int(math.Floor(float64(c.A) / math.Pow(2, float64(c.ComboOperand()))))
	}
	if c.ptr == ptr {
		c.ptr += 2
	}
}
