package options

import (
	"fmt"
	"sort"
	"strings"

	"github.com/JintaTechx/DDG/tree/main/pkg/services"
	"github.com/JintaTechx/DDG/tree/main/pkg/services/writer"
)

// Options for output format
var Options = map[string]services.Printer{
	"DEFAULT": &writer.DEFAULT{},
	"MD":      &writer.MD{},
	"TXT":     &writer.TXT{},
	"HTML":    &writer.HTML{},
	"JSON":    &writer.JSON{},
}

// GetKeys returns an array of string containing the output possibilites
func GetKeys() []string {
	keys := []string{}

	for k := range Options {
		keys = append(keys, strings.ToLower(k))
	}

	sort.Strings(keys)
	return keys
}

// Message returns the message that says the output possibilities
func Message() string {
	return fmt.Sprintf("the output types %v", GetKeys())
}
