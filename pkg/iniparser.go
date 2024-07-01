package iniparser

import (
	"errors"
	"os"
	"reflect"
	"strings"
)

var (
	errReadingFile = errors.New("trying to read file")
	errWritingFile = errors.New("trying to write file")
	errFileIsEmpty = errors.New("ini file is empty or does not have section key value pair")
)

type IniFile struct {
	sectionKeyValuePairs map[string]map[string]string
}

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

func (ini *IniFile) LoadFromFile(fileName string) (map[string]map[string]string, error) {

	fileContent, err := os.ReadFile(fileName)

	if err != nil {
		emptyMap := make(map[string]map[string]string)
		return emptyMap, errReadingFile
	}

	return ini.LoadFromString(string(fileContent))
}

func (ini *IniFile) GetSectionNames() []string {
	var sectionNames []string
	for section := range ini.sectionKeyValuePairs {
		sectionNames = append(sectionNames, section)
	}
	return sectionNames
}

func (ini *IniFile) GetSections() map[string]map[string]string {
	return ini.sectionKeyValuePairs
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

func (ini *IniFile) SaveToFile(filePath string, sectionKeyValueMap map[string]map[string]string) error {

	fileContent := []byte(ini.String(sectionKeyValueMap))

	err := os.WriteFile(filePath, fileContent, 0644)

	if err != nil {
		return errWritingFile
	}

	return nil
}
