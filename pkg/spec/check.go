package spec

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/apigear-io/cli/pkg/model"

	"github.com/apigear-io/cli/pkg/idl"

	"github.com/gocarina/gocsv"
)

type Result struct {
	File   string        `json:"file"`
	Errors []ErrorResult `json:"errors"`
}

func (r *Result) Valid() bool {
	return len(r.Errors) == 0
}

type ErrorResult struct {
	Field       string `json:"field"`
	Description string `json:"description"`
	Related     string `json:"related"`
}

func (e ErrorResult) String() string {
	str := ""
	if e.Description != "" {
		str += fmt.Sprintf("error: %s\n", e.Description)
	}
	if e.Field != "" {
		str += fmt.Sprintf("field: %s\n", e.Field)
	}
	if e.Related != "" {
		str += fmt.Sprintf("--- related value\n%s\n---", e.Related)
	}
	return str
}

func CheckFileAndType(file string, t DocumentType) (*Result, error) {
	dt, error := GetDocumentType(file)
	if error != nil {
		return nil, error
	}
	if dt != t {
		return nil, fmt.Errorf("file %s is not a %s file", file, t)
	}
	return CheckFile(file)
}

func CheckFile(file string) (*Result, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}
	if filepath.Ext(file) == ".yaml" || filepath.Ext(file) == ".yml" {
		data, err = YamlToJson(data)
		if err != nil {
			return nil, err
		}
	}
	typ, err := GetDocumentType(file)
	if err != nil {
		return nil, err
	}
	return CheckJson(typ, data)
}

func CheckNdjsonFile(name string) error {
	// read file line by line with scanner
	file, err := os.Open(name)
	if err != nil {
		return fmt.Errorf("open file %s: %w", name, err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// decode each line into an empty interface using json
		var event any
		err := json.NewDecoder(bufio.NewReader(strings.NewReader(line))).Decode(&event)
		if err != nil {
			return fmt.Errorf("decode line %s: %w", line, err)
		}
	}
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("read file %s: %w", name, err)
	}
	return nil
}

func CheckCsvFile(name string) error {
	// read file line by line using scanner
	file, err := os.Open(name)
	if err != nil {
		return fmt.Errorf("open file %s: %w", name, err)
	}
	defer file.Close()
	var data []any
	err = gocsv.UnmarshalFile(file, &data)
	if err != nil {
		return fmt.Errorf("unmarshal file %s: %w", name, err)
	}
	return nil
}

func CheckIdlFile(name string) error {
	s := model.NewSystem("check")
	parser := idl.NewParser(s)
	err := parser.ParseFile(name)
	if err != nil {
		return fmt.Errorf("parse file %s: %w", name, err)
	}
	return s.ResolveAll()
}
