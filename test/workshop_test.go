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
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

// TestWorkshops searches the parent directory for directories that
// contain a README.md file. It then parses those for START/END tags
// and runs the bash scripts between those tags. It only supports bash
// actions at this time.
func TestWorkshops(t *testing.T) {
	workshops, err := findWorkshops()
	if err != nil {
		t.Fatalf("Unexpected error finding workshop READMEs: %v", err)
	}

	for _, ws := range workshops {
		file, err := os.Open(ws)
		if err != nil {
			t.Fatalf("[%v] Unexpected error opening workshop: %v", ws, err)
		}

		actions, err := GetActions(file)
		if err != nil {
			t.Fatalf("[%v] Unexpected error parsing actions: %v", ws, err)
		}

		if err := os.Chdir(filepath.Dir(ws)); err != nil {
			t.Fatalf("[%v] Unexpected error changing to workshop dir: %v",
				ws, err)
		}

		t.Logf("[%v] Running workshop scripts", ws)
	Loop:
		for _, action := range actions {
			switch action.ActionType {
			case "bash":
				if runBash(t, action.Lines) == false {
					// Error in script, skip rest
					// of workshop, move to next
					// workshop file.
					break Loop
				}
			default:
				t.Errorf("[%v] Invalid action type found in workshop: %v",
					ws, action.ActionType)
			}
		}
	}
}

// Run a series of bash commands, returns true if successful.
func runBash(t *testing.T, commands []string) bool {
	for _, line := range commands {
		args := []string{"-c", line}
		cmd := exec.Command("bash", args...)
		output, err := cmd.CombinedOutput()
		if err != nil {
			t.Errorf("Command failed: %v", line)
			t.Errorf("Error: %v", err)
			t.Errorf("Combined Out: %v", string(output))
			return false
		}
		t.Logf("Command succeeded: %v", line)
		t.Logf("Combined Out: %v", string(output))
	}
	return true
}

// Find all README.md files in neighboring directories.
func findWorkshops() ([]string, error) {
	readmes := make([]string, 0)

	directory, err := os.Open("..")
	if err != nil {
		return nil, err
	}

	contents, err := directory.Readdir(0)
	if err != nil {
		return nil, err
	}

	for _, entry := range contents {
		if !entry.IsDir() {
			continue
		}
		readme := filepath.Join("..", entry.Name(), "README.md")
		_, err := os.Stat(readme)
		if os.IsNotExist(err) {
			continue
		}
		if err != nil {
			return nil, err
		}
		readmes = append(readmes, readme)
	}
	return readmes, nil
}
