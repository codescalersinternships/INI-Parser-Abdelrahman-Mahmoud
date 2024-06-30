package iniparser

import (
	"fmt"
	"os"
	"strings"
)

type IniFile struct {
	sectionKeyValuePairs map[string]map[string]string
	comments             []string
	fileName             string
}

func (ini *IniFile) loadFromString(iniText string) error {

	lines := strings.Split(iniText, "\n")

	if len(lines) == 0 {
		return fmt.Errorf("Error: INI file is empty!")
	}

	var section string
	var key string
	var value string
	ini.sectionKeyValuePairs = make(map[string]map[string]string)

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		if string(line[0]) == ";" || string(line[0]) == "#" {
			ini.comments = append(ini.comments, line)
			continue
		}
		if string(line[0]) == "[" {
			section = strings.TrimSpace(line[1 : len(line)-1])
			ini.sectionKeyValuePairs[section] = make(map[string]string)
		} else {
			s := strings.Split(line, "=")
			key, value = strings.TrimSpace(s[0]), strings.TrimSpace(s[1])
			ini.sectionKeyValuePairs[section][key] = value
		}
	}
	return nil
}
func (ini *IniFile) LoadFromFile(fileName string) error {

	ini.fileName = fileName
	fileContent, err := os.ReadFile(fileName)

	if err != nil {
		return fmt.Errorf("Error: trying to read file!")
	}

	err = ini.loadFromString(string(fileContent))

	fmt.Println(ini.sectionKeyValuePairs)

	return err
}

func (ini *IniFile) GetSectionNames() []string {
	var sectionNames []string
	for section := range ini.sectionKeyValuePairs {
		sectionNames = append(sectionNames, section)
	}
	return sectionNames
}

func (ini *IniFile) GetSections() string {
	sectionNames := ini.GetSectionNames()
	sections := "{ "
	for i, sectionName := range sectionNames {
		sections += sectionName
		sections += ": {"
		j := 0
		for key, value := range ini.sectionKeyValuePairs[sectionName] {
			sections += key
			sections += ": "
			sections += value
			if j != len(ini.sectionKeyValuePairs[sectionName])-1 {
				sections += ", "
			}
			j++
		}
		if i != len(sectionNames)-1 {
			sections += "}, "
		}
	}
	sections += "} }"
	return sections
}

func (ini *IniFile) Get(section string, key string) string {
	return ini.sectionKeyValuePairs[section][key]
}

func (ini *IniFile) Set(section string, key string, value string) {
	sectionNames := ini.GetSectionNames()
	for _, sectionName := range sectionNames {
		if sectionName == section {
			ini.sectionKeyValuePairs[section][key] = value
			return
		}
	}
	ini.sectionKeyValuePairs[section] = make(map[string]string)
	ini.sectionKeyValuePairs[section][key] = value

}

func (ini *IniFile) toString() string {

	iniText := ""

	for _, comment := range ini.comments {
		iniText += comment
		iniText += "\n"
	}

	for sectionName := range ini.sectionKeyValuePairs {
		iniText += "\n["
		iniText += sectionName
		iniText += "]\n"
		for key, value := range ini.sectionKeyValuePairs[sectionName] {
			iniText += key
			iniText += "="
			iniText += value
			iniText += "\n"
		}
	}
	fmt.Print(iniText)
	return iniText
}

func (ini *IniFile) SaveToFile(arguments ...string) error {
	var fileName string
	if len(arguments) == 0 {
		fileName = ini.fileName
	} else {
		fileName = arguments[0]
	}
	fileContent := []byte(ini.toString())

	err := os.WriteFile(fileName, fileContent, 0644)

	if err != nil {
		return fmt.Errorf("Error: trying to write file!")
	}

	return nil
}
