// Package iniparser provide methods to use on ini files.
package iniparser

import (
	"errors"
	"os"
	"reflect"
	"strings"
)

var (
	errReadingFile   = errors.New("trying to read file")
	errFileIsEmpty   = errors.New("ini file is empty or does not have section key value pair")
	errMissingValue  = errors.New("section and key pair does not exist")
	errAlreadyExists = errors.New("section key value pair already exists")
)

type IniFile struct {
	sectionKeyValuePairs map[string]map[string]string
}

// LoadFromString loads the content of ini file from multi-lined string.
func (ini *IniFile) LoadFromString(iniText string) (map[string]map[string]string, error) {

	var section string
	var key string
	var value string
	emptyMap := make(map[string]map[string]string)

	sectionAvailable := false
	ini.sectionKeyValuePairs = emptyMap

	if iniText == "" {
		return emptyMap, errFileIsEmpty
	}

	lines := strings.Split(iniText, "\n")

	for _, line := range lines {

		if len(line) == 0 {
			continue
		} else if string(line[0]) == ";" || string(line[0]) == "#" {
			continue
		} else if string(line[0]) == "[" {
			section = strings.TrimSpace(line[1 : len(line)-1])
			ini.sectionKeyValuePairs[section] = make(map[string]string)
			sectionAvailable = true
		} else if sectionAvailable {

			s := strings.Split(line, "=")
			key, value = strings.TrimSpace(s[0]), strings.TrimSpace(s[1])
			ini.sectionKeyValuePairs[section][key] = value
		}
	}

	if reflect.DeepEqual(ini.sectionKeyValuePairs, make(map[string]map[string]string)) {
		return emptyMap, errFileIsEmpty
	}
	return ini.sectionKeyValuePairs, nil
}

// String converts the ini file contents into a mult-lined string
func (ini *IniFile) String(sectionKeyValueMap map[string]map[string]string) string {

	iniText := ""

	for sectionName := range sectionKeyValueMap {
		iniText = iniText + "\n[" + sectionName + "]\n"
		for key, value := range sectionKeyValueMap[sectionName] {
			iniText = iniText + key + " = " + value + "\n"
		}
	}

	return iniText
}

// LoadFromFile loads the content of a given ini file
func (ini *IniFile) LoadFromFile(fileName string) (map[string]map[string]string, error) {

	fileContent, err := os.ReadFile(fileName)

	if err != nil {
		emptyMap := make(map[string]map[string]string)
		return emptyMap, errReadingFile
	}

	return ini.LoadFromString(string(fileContent))
}

// GetSectionNames retrives the sections of the ini file
func (ini *IniFile) GetSectionNames() []string {
	var sectionNames []string
	for section := range ini.sectionKeyValuePairs {
		sectionNames = append(sectionNames, section)
	}
	return sectionNames
}

// GetSections retrives the section key value pairs of the ini file
func (ini *IniFile) GetSections() map[string]map[string]string {
	return ini.sectionKeyValuePairs
}

// Get retrives a value for a specific key inside a section
func (ini *IniFile) Get(section string, key string) (string, error) {
	value := ini.sectionKeyValuePairs[section][key]
	if value == "" {
		return value, errMissingValue
	}
	return value, nil
}

// Set adds a key value pair inside a section
func (ini *IniFile) Set(section string, key string, value string) error {
	sectionNames := ini.GetSectionNames()
	for _, sectionName := range sectionNames {
		if sectionName == section {
			oldValue := ini.sectionKeyValuePairs[section][key]
			if oldValue == "" {
				ini.sectionKeyValuePairs[section][key] = value
				return nil
			} else {
				return errAlreadyExists
			}
		}
	}
	ini.sectionKeyValuePairs[section] = make(map[string]string)
	ini.sectionKeyValuePairs[section][key] = value
	return nil
}

// SaveToFile save the section key value pairs to a given file
func (ini *IniFile) SaveToFile(filePath string, sectionKeyValueMap map[string]map[string]string) error {

	fileContent := []byte(ini.String(sectionKeyValueMap))

	_ = os.WriteFile(filePath, fileContent, 0644)

	return nil
}
