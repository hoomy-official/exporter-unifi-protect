package cmd

import (
	"fmt"
)

// Licence holds the licence content of an application.
// It is constructed in a manner that allows Kong CLI library to recognize
// the licence command and display the content accordingly.
type Licence struct {
	content string
}

// NewLicence creates a new Licence instance with the provided licence content.
func NewLicence(s string) Licence {
	return Licence{content: s}
}

// Run prints the licence content to the console when the licence command is invoked.
// This method satisfies the Kong interface contract for command execution.
func (l Licence) Run() error {
	//nolint:forbidigo // using a custom writer is not necessary here
	fmt.Printf("%s", l.content)
	return nil
}
