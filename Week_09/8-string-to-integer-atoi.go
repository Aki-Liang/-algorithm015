package Week_09

import "math"

func myAtoi(s string) int {
	a := &Automaton{
		sign:  1,
		ans:   0,
		state: "start",
		table: map[string][]string{
			"start":     []string{"start", "signed", "in_number", "end"},
			"signed":    {"end", "end", "in_number", "end"},
			"in_number": {"end", "end", "in_number", "end"},
			"end":       {"end", "end", "end", "end"},
		},
	}

	for _, c := range s {
		a.get(c)
	}

	return a.sign * a.ans
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type Automaton struct {
	sign  int
	ans   int
	state string
	table map[string][]string
}

func (a *Automaton) get(c rune) {
	a.state = a.table[a.state][a.getCol(c)]
	if a.state == "in_number" {
		a.ans = a.ans*10 + int(c-'0')
		if a.sign == 1 {
			a.ans = min(a.ans, math.MaxInt32)
		} else if a.ans*a.sign < math.MinInt32 {
			a.ans = int(math.Abs(math.MinInt32))
		}
	} else if a.state == "signed" {
		if c == '+' {
			a.sign = 1
		} else {
			a.sign = -1
		}
	}
}

func (a *Automaton) getCol(c rune) int {
	switch c {
	case ' ':
		return 0
	case '+', '-':
		return 1
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return 2
	}
	return 3
}
