package iniparser

import (
	"strings"
	"bufio"
	"os"
	"fmt"
)

type IniFile struct {
	sectionKeyValuePairs map[string]map[string]string
}

func (ini *IniFile) LoadFromFile(file *os.File) {

	fileScanner := bufio.NewScanner(file)

	fileScanner.Split(bufio.ScanLines)

	var section string
	var key string
	var value string
	ini.sectionKeyValuePairs = make(map[string]map[string]string) 

	for fileScanner.Scan() {
		line := fileScanner.Text()
		line = strings.TrimSpace(line)
		if len(line) == 0 || string(line[0]) == ";"  || string(line[0]) == "#"{
			continue
		}			
		if string(line[0]) == "[" {
			section = strings.TrimSpace(line[1:len(line)-1])
			ini.sectionKeyValuePairs[section] = make(map[string]string)
		} else {
			s := strings.Split(line, "=")
			key, value = strings.TrimSpace(s[0]), strings.TrimSpace(s[1])
			ini.sectionKeyValuePairs[section][key] = value
		}
	}

	fmt.Println(ini.sectionKeyValuePairs)

	defer file.Close()
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
    for i ,section := range sectionNames {
		sections += section
		sections += ": {"
		j :=0
		for key ,value := range ini.sectionKeyValuePairs[section] {
			sections += key
			sections += ": "
			sections += value
			if j != len(ini.sectionKeyValuePairs[section]) - 1 {
				sections += ", "
			}
			j++
		}
		if i != len(sectionNames) - 1 {
        	sections +="}, "
		}
    }
	sections +="} }"
	return sections
}

func (ini *IniFile) Get(section string, key string) string {
	return ini.sectionKeyValuePairs[section][key]
}