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
	var sections []string
    for section := range ini.sectionKeyValuePairs {
        sections = append(sections, section)
    }
	return sections
}