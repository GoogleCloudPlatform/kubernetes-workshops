// Copyright 2016 Google, Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package main

import (
	"bufio"
	"errors"
	"io"
	"strings"
)

// An Action is one instance of a START/END tag pair. ActionType is
// the word following START. Lines are the lines between the tags.
type Action struct {
	ActionType string
	Lines      []string
	closed     bool
}

type stateFn func(string, *[]Action) (stateFn, error)

func neutralState(line string, actions *[]Action) (stateFn, error) {
	words := strings.Fields(line)
	for index, word := range words {
		if word == "END" {
			return nil, errors.New("Unexpected END tag outside action")
		}
		if word == "START" {
			action := words[index+1]
			newAction(action, actions)
			return insideActionState, nil
		}
	}
	return neutralState, nil
}

func insideActionState(line string, actions *[]Action) (stateFn, error) {
	words := strings.Fields(line)
	if strings.HasPrefix(line, "```") || words[0] == "```" {
		// Ignore line.
		return insideActionState, nil
	}
	for _, word := range words {
		if word == "END" {
			closeAction(*actions)
			return neutralState, nil
		}
		if word == "START" {
			return nil, errors.New("Unexpected START tag inside action")
		}
	}
	actionLine(line, *actions)
	return insideActionState, nil
}

func newAction(aType string, actions *[]Action) {
	aa := Action{
		ActionType: aType,
		Lines:      make([]string, 0),
		closed:     false,
	}
	*actions = append(*actions, aa)
}

func actionLine(line string, actions []Action) {
	actions[len(actions)-1].Lines =
		append(actions[len(actions)-1].Lines, line)
}

func closeAction(actions []Action) {
	actions[len(actions)-1].closed = true
}

// GetActions parses the input for START/END tags and compiles a list
// of actions, comprising of the lines between the lines with
// START/END. Lines that begin with ``` are ignored. The word
// immediately after START is saved as the action type.
func GetActions(reader io.Reader) ([]Action, error) {
	scanner := bufio.NewScanner(reader)
	actions := make([]Action, 0)

	for state := neutralState; scanner.Scan(); {
		var err error
		state, err = state(scanner.Text(), &actions)
		if err != nil {
			return nil, err
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return actions, nil
}
