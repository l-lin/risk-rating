package risk

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

// Value represents a value selected by the user in the terminal
type Value struct {
	Weight int
	Label  string
}

// NewValue creates a new risk value
func NewValue(value int, label string) *Value {
	return &Value{value, label}
}

// Parse given string to a risk.Value
func Parse(val string) *Value {
	rWeight := regexp.MustCompile("[0-9]")
	weight, err := strconv.Atoi(rWeight.FindString(val))
	if err != nil {
		log.Fatalln(fmt.Sprintf("Could not parse %s to integer. Erro was %s", val, err))
	}

	rLabel := regexp.MustCompile("\\([a-z -]+")
	label := strings.ReplaceAll(rLabel.FindString(val), "(", "")
	return NewValue(weight, label)
}

func (v Value) String() string {
	if v.Label == "" {
		return strconv.Itoa(v.Weight)
	}
	return fmt.Sprintf("%d (%s)", v.Weight, v.Label)
}
