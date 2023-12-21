package day19

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/advent-of-code/2023/util"
)

func Run() {
	input, err := util.ReadLines("./day19/input/input.txt")
	if err != nil {
		fmt.Printf("error reading file: %s", err.Error())
		return
	}

	var start time.Time

	start = time.Now()
	util.PrettyPrint(19, 1, Part1(input), start)

	start = time.Now()
	util.PrettyPrint(19, 2, Part2(input), start)
}

type WorkflowRule struct {
	Rule        string
	Destination string
}

type Part struct {
	X int
	M int
	A int
	S int
}

func Part1(input []string) int {
	var totalRating int
	workflows, parts := parseInput(input)
	for _, part := range parts {
		if applyRule(workflows, "in", part) {
			totalRating += part.X + part.M + part.A + part.S
		}
	}
	return totalRating
}

func Part2(input []string) int {

	return 0
}

func applyRule(workflows map[string][]WorkflowRule, currentFlow string, part Part) bool {
	for _, wfr := range workflows[currentFlow] {
		if validateRule(wfr.Rule, part) {
			if wfr.Destination == "R" {
				return false
			}
			if wfr.Destination == "A" {
				return true
			}
			return applyRule(workflows, wfr.Destination, part)
		}
	}
	return false
}

func validateRule(rule string, part Part) bool {
	if rule == "" {
		return true
	}
	n := 0
	o, _ := strconv.Atoi(rule[2:])

	switch rule[0] {
	case 'x':
		n = part.X
	case 'm':
		n = part.M
	case 'a':
		n = part.A
	case 's':
		n = part.S
	}
	switch rule[1] {
	case '>':
		return n > o
	case '<':
		return n < o
	}

	return false
}

func parseInput(input []string) (map[string][]WorkflowRule, []Part) {
	workflows := make(map[string][]WorkflowRule)
	parts := make([]Part, 0)

	for _, l := range input {
		pattern := regexp.MustCompile("{|}$")
		ss := pattern.Split(l, -1)

		if len(ss) > 1 {
			if ss[0] != "" {
				// Workflow found
				name := ss[0]
				workflows[name] = parseWorkFlow(ss[1])
			}

			if ss[0] == "" {
				// Part found
				parts = append(parts, parseRatings(ss[1]))
			}

		}

	}

	fmt.Println(workflows)

	return workflows, parts
}

func parseRatings(s string) Part {
	var part Part

	for _, r := range strings.Split(s, ",") {
		ss := strings.Split(r, "=")
		switch ss[0] {
		case "x":
			part.X, _ = strconv.Atoi(ss[1])
		case "m":
			part.M, _ = strconv.Atoi(ss[1])
		case "a":
			part.A, _ = strconv.Atoi(ss[1])
		case "s":
			part.S, _ = strconv.Atoi(ss[1])
		}
	}

	return part
}

func parseWorkFlow(s string) []WorkflowRule {
	var rules []WorkflowRule

	for _, r := range strings.Split(s, ",") {
		var rule WorkflowRule
		ss := strings.Split(r, ":")
		if len(ss) == 1 {
			rule = WorkflowRule{"", ss[0]}
		} else if len(ss) == 2 {
			rule = WorkflowRule{ss[0], ss[1]}
		}
		rules = append(rules, rule)
	}

	return rules
}
