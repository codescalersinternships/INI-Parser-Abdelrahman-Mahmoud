package iniparser

import (
	"reflect"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

type LoadFromStringTestCase struct {
	desc          string
	input         string
	expectedError error
	expectedMap   map[string]map[string]string
}

type LoadFromFileTestCase struct {
	desc          string
	path          string
	expectedError error
	expectedMap   map[string]map[string]string
}

type StringTestCase struct {
	desc     string
	inputMap map[string]map[string]string
}

type SaveToFileTestCase struct {
	desc          string
	inputPath     string
	inputMap      map[string]map[string]string
	expectedError error
}

type GetSectionNamesTestCase struct {
	desc        string
	inputMap    map[string]map[string]string
	outputSlice []string
}

type GetSectionsTestCase struct {
	desc           string
	inputOutputMap map[string]map[string]string
}

type GetSetTestCase struct {
	desc          string
	inputMap      map[string]map[string]string
	section       string
	key           string
	value         string
	expectedError error
}

func sliceEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	a_copy := make([]string, len(a))
	b_copy := make([]string, len(b))

	copy(a_copy, a)
	copy(b_copy, b)

	sort.Strings(a_copy)
	sort.Strings(b_copy)

	return reflect.DeepEqual(a_copy, b_copy)
}

func TestINI_LoadFromString(t *testing.T) {
	parser := IniFile{}
	emptyMap := make(map[string]map[string]string)
	testCases := []LoadFromStringTestCase{
		{
			desc:          "Empty string as input",
			input:         "",
			expectedError: errFileIsEmpty,
			expectedMap:   emptyMap,
		},
		{
			desc:          "Only one comment line start with # as input",
			input:         "#This is a comment",
			expectedError: errFileIsEmpty,
			expectedMap:   emptyMap,
		},
		{
			desc:          "Only one comment line start with ; as input",
			input:         ";This is a comment",
			expectedError: errFileIsEmpty,
			expectedMap:   emptyMap,
		},
		{
			desc:          "Only one key value line as input",
			input:         "key = value",
			expectedError: errFileIsEmpty,
			expectedMap:   emptyMap,
		},
		{
			desc:          "Only one section line as input",
			input:         "[section]",
			expectedError: nil,
			expectedMap:   map[string]map[string]string{"section": make(map[string]string)},
		},
		{
			desc: "Only comments as input",
			input: "!This is a comment 1\n " +
				"!This is a comment 2\n" +
				"!This is a comment 3\n" +
				"!This is a comment 4",
			expectedError: errFileIsEmpty,
			expectedMap:   emptyMap,
		},
		{
			desc: "Only sections as input",
			input: "[section1]\n" +
				"[section2]\n" +
				"[section3]\n" +
				"[section4]",
			expectedError: nil,
			expectedMap: map[string]map[string]string{"section1": make(map[string]string),
				"section2": make(map[string]string),
				"section3": make(map[string]string),
				"section4": make(map[string]string)},
		},
		{
			desc: "Only keys values as input",
			input: "key1 = value1\n " +
				"key2 = value2\n" +
				"key3 = value3\n" +
				"key4 = value4",
			expectedError: errFileIsEmpty,
			expectedMap:   emptyMap,
		},
		{
			desc: "Normal case 1",
			input: `;This is a comment1
;This is a comment2

[section1]
key1 = value1
key2 = value2`,
			expectedError: nil,
			expectedMap: map[string]map[string]string{"section1": {
				"key1": "value1",
				"key2": "value2",
			},
			},
		},
		{
			desc: "Normal case 2",
			input: `;This is a comment1
;This is a comment2

[section1]
key1 = value1
key2 = value2
key3 = value3
key4 = value4

[section2]
key1 = value1
key2 = value2`,
			expectedError: nil,
			expectedMap: map[string]map[string]string{"section1": {
				"key1": "value1",
				"key2": "value2",
				"key3": "value3",
				"key4": "value4",
			},
				"section2": {
					"key1": "value1",
					"key2": "value2",
				},
			},
		},
		{
			desc: "Normal case 3",
			input: `[section1]
key1 = value1
key2 = value2
key3 = value3
key4 = value4

[section2]
key1 = value1
key2 = value2
key3 = value3

[section3]
key1 = value1
key2 = value2
key3 = value3
key4 = value4`,
			expectedError: nil,
			expectedMap: map[string]map[string]string{"section1": {
				"key1": "value1",
				"key2": "value2",
				"key3": "value3",
				"key4": "value4",
			},
				"section2": {
					"key1": "value1",
					"key2": "value2",
					"key3": "value3",
				},
				"section3": {
					"key1": "value1",
					"key2": "value2",
					"key3": "value3",
					"key4": "value4",
				},
			},
		},
		{
			desc: "Normal case 4",
			input: `[section1]
key1 = value1
key2 = value2
key3 = value3
key4 = value4
key5 = value5
key6 = value6
key7 = value7

[section2]
key1 = value1
key2 = value2
key3 = value3

[section3]
key1 = value1
key2 = value2
key3 = value3
key4 = value4`,
			expectedError: nil,
			expectedMap: map[string]map[string]string{"section1": {
				"key1": "value1",
				"key2": "value2",
				"key3": "value3",
				"key4": "value4",
				"key5": "value5",
				"key6": "value6",
				"key7": "value7",
			},
				"section2": {
					"key1": "value1",
					"key2": "value2",
					"key3": "value3",
				},
				"section3": {
					"key1": "value1",
					"key2": "value2",
					"key3": "value3",
					"key4": "value4",
				},
			},
		},
		{
			desc: "Normal case 5",
			input: `;This is a comment1
;This is a comment2

[section1]
key1 = value1
key2 = value2

[section2]
key1 = value1
key2 = value2
key3 = value3
key4 = value4
key5 = value5


[section3]
key1 = value1
key2 = value2
key3 = value3

[section4]
key1 = value1
key2 = value2
key3 = value3
`,
			expectedError: nil,
			expectedMap: map[string]map[string]string{"section1": {
				"key1": "value1",
				"key2": "value2",
			},
				"section2": {
					"key1": "value1",
					"key2": "value2",
					"key3": "value3",
					"key4": "value4",
					"key5": "value5",
				},
				"section3": {
					"key1": "value1",
					"key2": "value2",
					"key3": "value3",
				},
				"section4": {
					"key1": "value1",
					"key2": "value2",
					"key3": "value3",
				},
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			resultedMap, resultedError := parser.LoadFromString(test.input)

			assert.Equal(t, test.expectedError, resultedError)
			if !reflect.DeepEqual(test.expectedMap, resultedMap) {
				t.Fail()
			}

		})
	}

}

func TestINI_loadFromFile(t *testing.T) {
	parser := IniFile{}
	emptyMap := make(map[string]map[string]string)
	testCases := []LoadFromFileTestCase{
		{
			desc:          "Wrong file path",
			path:          "no path",
			expectedError: errReadingFile,
			expectedMap:   emptyMap,
		},
		{
			desc:          "Empty file as input",
			path:          "testdata/test_011.txt",
			expectedError: errFileIsEmpty,
			expectedMap:   emptyMap,
		},
		{
			desc:          "Only one comment line start with # inside file",
			path:          "testdata/test_02.txt",
			expectedError: errFileIsEmpty,
			expectedMap:   emptyMap,
		},
		{
			desc:          "Only one comment line start with ; inside file",
			path:          "testdata/test_03.txt",
			expectedError: errFileIsEmpty,
			expectedMap:   emptyMap,
		},
		{
			desc:          "Only one key value line inside file",
			path:          "testdata/test_04.txt",
			expectedError: errFileIsEmpty,
			expectedMap:   emptyMap,
		},
		{
			desc:          "Only one section line inside file",
			path:          "testdata/test_05.txt",
			expectedError: nil,
			expectedMap:   map[string]map[string]string{"section": make(map[string]string)},
		},
		{
			desc:          "Only comments inside file",
			path:          "testdata/test_06.txt",
			expectedError: errFileIsEmpty,
			expectedMap:   emptyMap,
		},
		{
			desc:          "Only sections inside file",
			path:          "testdata/test_07.txt",
			expectedError: nil,
			expectedMap: map[string]map[string]string{"section1": make(map[string]string),
				"section2": make(map[string]string),
				"section3": make(map[string]string),
				"section4": make(map[string]string)},
		},
		{
			desc:          "Only keys values inside file",
			path:          "testdata/test_08.txt",
			expectedError: errFileIsEmpty,
			expectedMap:   emptyMap,
		},
		{
			desc:          "Normal case 1",
			path:          "testdata/test_09.txt",
			expectedError: nil,
			expectedMap: map[string]map[string]string{"section1": {
				"key1": "value1",
				"key2": "value2",
			},
			},
		},
		{
			desc:          "Normal case 2",
			path:          "testdata/test_10.txt",
			expectedError: nil,
			expectedMap: map[string]map[string]string{"section1": {
				"key1": "value1",
				"key2": "value2",
				"key3": "value3",
				"key4": "value4",
			},
				"section2": {
					"key1": "value1",
					"key2": "value2",
				},
			},
		},
		{
			desc:          "Normal case 3",
			path:          "testdata/test_11.txt",
			expectedError: nil,
			expectedMap: map[string]map[string]string{"section1": {
				"key1": "value1",
				"key2": "value2",
				"key3": "value3",
				"key4": "value4",
			},
				"section2": {
					"key1": "value1",
					"key2": "value2",
					"key3": "value3",
				},
				"section3": {
					"key1": "value1",
					"key2": "value2",
					"key3": "value3",
					"key4": "value4",
				},
			},
		},
		{
			desc:          "Normal case 4",
			path:          "testdata/test_12.txt",
			expectedError: nil,
			expectedMap: map[string]map[string]string{"section1": {
				"key1": "value1",
				"key2": "value2",
				"key3": "value3",
				"key4": "value4",
				"key5": "value5",
				"key6": "value6",
				"key7": "value7",
			},
				"section2": {
					"key1": "value1",
					"key2": "value2",
					"key3": "value3",
				},
				"section3": {
					"key1": "value1",
					"key2": "value2",
					"key3": "value3",
					"key4": "value4",
				},
			},
		},
		{
			desc:          "Normal case 5",
			path:          "testdata/test_13.txt",
			expectedError: nil,
			expectedMap: map[string]map[string]string{"section1": {
				"key1": "value1",
				"key2": "value2",
			},
				"section2": {
					"key1": "value1",
					"key2": "value2",
					"key3": "value3",
					"key4": "value4",
					"key5": "value5",
				},
				"section3": {
					"key1": "value1",
					"key2": "value2",
					"key3": "value3",
				},
				"section4": {
					"key1": "value1",
					"key2": "value2",
					"key3": "value3",
				},
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			resultedMap, resultedError := parser.LoadFromFile(test.path)

			assert.Equal(t, test.expectedError, resultedError)
			if !reflect.DeepEqual(test.expectedMap, resultedMap) {
				t.Fail()
			}

		})
	}
}

func TestINI_String(t *testing.T) {
	parser := IniFile{}
	testCases := []StringTestCase{
		{
			desc:     "Empty map as input",
			inputMap: make(map[string]map[string]string),
		},
		{
			desc: "Only sections inside file",
			inputMap: map[string]map[string]string{"section1": make(map[string]string),
				"section2": make(map[string]string),
				"section3": make(map[string]string),
				"section4": make(map[string]string)},
		},
		{
			desc: "Normal case 1",
			inputMap: map[string]map[string]string{"section1": {
				"key1": "value1",
				"key2": "value2",
			},
			},
		},
		{
			desc: "Normal case 2",
			inputMap: map[string]map[string]string{"section1": {
				"key1": "value1",
				"key2": "value2",
				"key3": "value3",
				"key4": "value4",
			},
				"section2": {
					"key1": "value1",
					"key2": "value2",
				},
			},
		},
		{
			desc: "Normal case 3",
			inputMap: map[string]map[string]string{"section1": {
				"key1": "value1",
				"key2": "value2",
				"key3": "value3",
				"key4": "value4",
			},
				"section2": {
					"key1": "value1",
					"key2": "value2",
					"key3": "value3",
				},
				"section3": {
					"key1": "value1",
					"key2": "value2",
					"key3": "value3",
					"key4": "value4",
				},
			},
		},
		{
			desc: "Normal case 4",
			inputMap: map[string]map[string]string{"section1": {
				"key1": "value1",
				"key2": "value2",
				"key3": "value3",
				"key4": "value4",
				"key5": "value5",
				"key6": "value6",
				"key7": "value7",
			},
				"section2": {
					"key1": "value1",
					"key2": "value2",
					"key3": "value3",
				},
				"section3": {
					"key1": "value1",
					"key2": "value2",
					"key3": "value3",
					"key4": "value4",
				},
			},
		},
		{
			desc: "Normal case 5",
			inputMap: map[string]map[string]string{"section1": {
				"key1": "value1",
				"key2": "value2",
			},
				"section2": {
					"key1": "value1",
					"key2": "value2",
					"key3": "value3",
					"key4": "value4",
					"key5": "value5",
				},
				"section3": {
					"key1": "value1",
					"key2": "value2",
					"key3": "value3",
				},
				"section4": {
					"key1": "value1",
					"key2": "value2",
					"key3": "value3",
				},
			},
		},
	}
	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			resultedOutput := parser.String(test.inputMap)
			resultedMap, _ := parser.LoadFromString(resultedOutput)

			if !reflect.DeepEqual(test.inputMap, resultedMap) {
				t.Fail()
			}

		})
	}
}

func TestINI_SaveToFile(t *testing.T) {
	parser := IniFile{}
	testCases := []SaveToFileTestCase{
		{
			desc:          "Empty map as input",
			inputPath:     "testdata/test_01.txt",
			inputMap:      make(map[string]map[string]string),
			expectedError: nil,
		},
		{
			desc:      "Only sections inside file",
			inputPath: "testdata/test_01.txt",
			inputMap: map[string]map[string]string{"section1": make(map[string]string),
				"section2": make(map[string]string),
				"section3": make(map[string]string),
				"section4": make(map[string]string)},
			expectedError: nil,
		},
		{
			desc:      "Normal case 1",
			inputPath: "testdata/test_01.txt",
			inputMap: map[string]map[string]string{"section1": {
				"key1": "value1",
				"key2": "value2",
			},
			},
			expectedError: nil,
		},
		{
			desc:      "Normal case 2",
			inputPath: "testdata/test_01.txt",
			inputMap: map[string]map[string]string{"section1": {
				"key1": "value1",
				"key2": "value2",
				"key3": "value3",
				"key4": "value4",
			},
				"section2": {
					"key1": "value1",
					"key2": "value2",
				},
			},
			expectedError: nil,
		},
		{
			desc:      "Normal case 3",
			inputPath: "testdata/test_01.txt",
			inputMap: map[string]map[string]string{"section1": {
				"key1": "value1",
				"key2": "value2",
				"key3": "value3",
				"key4": "value4",
			},
				"section2": {
					"key1": "value1",
					"key2": "value2",
					"key3": "value3",
				},
				"section3": {
					"key1": "value1",
					"key2": "value2",
					"key3": "value3",
					"key4": "value4",
				},
			},
		},
		{
			desc:      "Normal case 4",
			inputPath: "testdata/test_01.txt",
			inputMap: map[string]map[string]string{"section1": {
				"key1": "value1",
				"key2": "value2",
				"key3": "value3",
				"key4": "value4",
				"key5": "value5",
				"key6": "value6",
				"key7": "value7",
			},
				"section2": {
					"key1": "value1",
					"key2": "value2",
					"key3": "value3",
				},
				"section3": {
					"key1": "value1",
					"key2": "value2",
					"key3": "value3",
					"key4": "value4",
				},
			},
		},
		{
			desc:      "Normal case 5",
			inputPath: "testdata/test_01.txt",
			inputMap: map[string]map[string]string{"section1": {
				"key1": "value1",
				"key2": "value2",
			},
				"section2": {
					"key1": "value1",
					"key2": "value2",
					"key3": "value3",
					"key4": "value4",
					"key5": "value5",
				},
				"section3": {
					"key1": "value1",
					"key2": "value2",
					"key3": "value3",
				},
				"section4": {
					"key1": "value1",
					"key2": "value2",
					"key3": "value3",
				},
			},
			expectedError: nil,
		},
	}
	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			resultedError := parser.SaveToFile(test.inputPath, test.inputMap)
			resultedMap, _ := parser.LoadFromFile(test.inputPath)

			assert.Equal(t, test.expectedError, resultedError)
			if !reflect.DeepEqual(test.inputMap, resultedMap) {
				t.Fail()
			}

		})
	}
}

func TestINI_GetSections(t *testing.T) {
	parser := IniFile{}
	testCases := []GetSectionsTestCase{
		{
			desc:           "Empty map as input",
			inputOutputMap: make(map[string]map[string]string),
		},
		{
			desc: "Only sections inside file",
			inputOutputMap: map[string]map[string]string{"section1": make(map[string]string),
				"section2": make(map[string]string),
				"section3": make(map[string]string),
				"section4": make(map[string]string)},
		},
		{
			desc: "Normal case 1",
			inputOutputMap: map[string]map[string]string{"section1": {
				"key1": "value1",
				"key2": "value2",
			},
			},
		},
		{
			desc: "Normal case 2",
			inputOutputMap: map[string]map[string]string{"section1": {
				"key1": "value1",
				"key2": "value2",
				"key3": "value3",
				"key4": "value4",
			},
				"section2": {
					"key1": "value1",
					"key2": "value2",
				},
			},
		},
		{
			desc: "Normal case 3",
			inputOutputMap: map[string]map[string]string{"section1": {
				"key1": "value1",
				"key2": "value2",
				"key3": "value3",
				"key4": "value4",
			},
				"section2": {
					"key1": "value1",
					"key2": "value2",
					"key3": "value3",
				},
				"section3": {
					"key1": "value1",
					"key2": "value2",
					"key3": "value3",
					"key4": "value4",
				},
			},
		},
		{
			desc: "Normal case 4",
			inputOutputMap: map[string]map[string]string{"section1": {
				"key1": "value1",
				"key2": "value2",
				"key3": "value3",
				"key4": "value4",
				"key5": "value5",
				"key6": "value6",
				"key7": "value7",
			},
				"section2": {
					"key1": "value1",
					"key2": "value2",
					"key3": "value3",
				},
				"section3": {
					"key1": "value1",
					"key2": "value2",
					"key3": "value3",
					"key4": "value4",
				},
			},
		},
		{
			desc: "Normal case 5",
			inputOutputMap: map[string]map[string]string{"section1": {
				"key1": "value1",
				"key2": "value2",
			},
				"section2": {
					"key1": "value1",
					"key2": "value2",
					"key3": "value3",
					"key4": "value4",
					"key5": "value5",
				},
				"section3": {
					"key1": "value1",
					"key2": "value2",
					"key3": "value3",
				},
				"section4": {
					"key1": "value1",
					"key2": "value2",
					"key3": "value3",
				},
			},
		},
	}
	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			resultedOutput := parser.String(test.inputOutputMap)
			_, _ = parser.LoadFromString(resultedOutput)
			resultedMap := parser.GetSections()

			if !reflect.DeepEqual(test.inputOutputMap, resultedMap) {
				t.Fail()
			}

		})
	}
}

func TestINI_Get(t *testing.T) {
	parser := IniFile{}
	testCases := []GetSetTestCase{
		{
			desc:          "Empty map as input",
			inputMap:      make(map[string]map[string]string),
			section:       "section1",
			key:           "key1",
			value:         "",
			expectedError: errMissingValue,
		},
		{
			desc: "Only sections inside file",
			inputMap: map[string]map[string]string{"section1": make(map[string]string),
				"section2": make(map[string]string),
				"section3": make(map[string]string),
				"section4": make(map[string]string)},
			section:       "section1",
			key:           "key1",
			value:         "",
			expectedError: errMissingValue,
		},
		{
			desc: "Normal case 1",
			inputMap: map[string]map[string]string{"section1": {
				"key1": "value1",
				"key2": "value2",
			},
			},
			section:       "section1",
			key:           "key1",
			value:         "value1",
			expectedError: nil,
		},
		{
			desc: "Normal case 2",
			inputMap: map[string]map[string]string{"section1": {
				"key1": "value1",
				"key2": "value2",
				"key3": "value3",
				"key4": "value4",
			},
				"section2": {
					"key1": "value1",
					"key2": "value2",
				},
			},
			section:       "section2",
			key:           "key1",
			value:         "value1",
			expectedError: nil,
		},
		{
			desc: "Normal case 3",
			inputMap: map[string]map[string]string{"section1": {
				"key1": "value1",
				"key2": "value2",
				"key3": "value3",
				"key4": "value4",
			},
				"section2": {
					"key1": "value1",
					"key2": "value2",
					"key3": "value3",
				},
				"section3": {
					"key1": "value1",
					"key2": "value2",
					"key3": "value3",
					"key4": "value4",
				},
			},
			section:       "section3",
			key:           "key5",
			value:         "",
			expectedError: errMissingValue,
		},
		{
			desc: "Normal case 4",
			inputMap: map[string]map[string]string{"section1": {
				"key1": "value1",
				"key2": "value2",
				"key3": "value3",
				"key4": "value4",
				"key5": "value5",
				"key6": "value6",
				"key7": "value7",
			},
				"section2": {
					"key1": "value1",
					"key2": "value2",
					"key3": "value3",
				},
				"section3": {
					"key1": "value1",
					"key2": "value2",
					"key3": "value3",
					"key4": "value4",
				},
			},
			section:       "section1",
			key:           "key1",
			value:         "value1",
			expectedError: nil,
		},
		{
			desc: "Normal case 5",
			inputMap: map[string]map[string]string{"section1": {
				"key1": "value1",
				"key2": "value2",
			},
				"section2": {
					"key1": "value1",
					"key2": "value2",
					"key3": "value3",
					"key4": "value4",
					"key5": "value5",
				},
				"section3": {
					"key1": "value1",
					"key2": "value2",
					"key3": "value3",
				},
				"section4": {
					"key1": "value1",
					"key2": "value2",
					"key3": "value3",
				},
			},
			section:       "section3",
			key:           "key4",
			value:         "",
			expectedError: errMissingValue,
		},
	}
	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			resultedOutput := parser.String(test.inputMap)
			_, _ = parser.LoadFromString(resultedOutput)
			resultedValue, resultedError := parser.Get(test.section, test.key)

			assert.Equal(t, test.expectedError, resultedError)
			if resultedValue != test.value {
				t.Fail()
			}

		})
	}
}

func TestINI_Set(t *testing.T) {
	parser := IniFile{}
	testCases := []GetSetTestCase{
		{
			desc:          "Empty map as input",
			inputMap:      make(map[string]map[string]string),
			section:       "section1",
			key:           "key1",
			value:         "value1",
			expectedError: nil,
		},
		{
			desc: "Only sections inside file",
			inputMap: map[string]map[string]string{"section1": make(map[string]string),
				"section2": make(map[string]string),
				"section3": make(map[string]string),
				"section4": make(map[string]string)},
			section:       "section1",
			key:           "key1",
			value:         "value1",
			expectedError: nil,
		},
		{
			desc: "Normal case 1",
			inputMap: map[string]map[string]string{"section1": {
				"key1": "value1",
				"key2": "value2",
			},
			},
			section:       "section1",
			key:           "key1",
			value:         "value1",
			expectedError: errAlreadyExists,
		},
		{
			desc: "Normal case 2",
			inputMap: map[string]map[string]string{"section1": {
				"key1": "value1",
				"key2": "value2",
				"key3": "value3",
				"key4": "value4",
			},
				"section2": {
					"key1": "value1",
					"key2": "value2",
				},
			},
			section:       "section2",
			key:           "key1",
			value:         "value1",
			expectedError: errAlreadyExists,
		},
		{
			desc: "Normal case 3",
			inputMap: map[string]map[string]string{"section1": {
				"key1": "value1",
				"key2": "value2",
				"key3": "value3",
				"key4": "value4",
			},
				"section2": {
					"key1": "value1",
					"key2": "value2",
					"key3": "value3",
				},
				"section3": {
					"key1": "value1",
					"key2": "value2",
					"key3": "value3",
					"key4": "value4",
				},
			},
			section:       "section3",
			key:           "key5",
			value:         "key5",
			expectedError: nil,
		},
		{
			desc: "Normal case 4",
			inputMap: map[string]map[string]string{"section1": {
				"key1": "value1",
				"key2": "value2",
				"key3": "value3",
				"key4": "value4",
				"key5": "value5",
				"key6": "value6",
				"key7": "value7",
			},
				"section2": {
					"key1": "value1",
					"key2": "value2",
					"key3": "value3",
				},
				"section3": {
					"key1": "value1",
					"key2": "value2",
					"key3": "value3",
					"key4": "value4",
				},
			},
			section:       "section4",
			key:           "key1",
			value:         "value1",
			expectedError: nil,
		},
		{
			desc: "Normal case 5",
			inputMap: map[string]map[string]string{"section1": {
				"key1": "value1",
				"key2": "value2",
			},
				"section2": {
					"key1": "value1",
					"key2": "value2",
					"key3": "value3",
					"key4": "value4",
					"key5": "value5",
				},
				"section3": {
					"key1": "value1",
					"key2": "value2",
					"key3": "value3",
				},
				"section4": {
					"key1": "value1",
					"key2": "value2",
					"key3": "value3",
				},
			},
			section:       "section3",
			key:           "key4",
			value:         "key4",
			expectedError: nil,
		},
	}
	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			resultedOutput := parser.String(test.inputMap)
			_, _ = parser.LoadFromString(resultedOutput)
			setError := parser.Set(test.section, test.key, test.value)
			resultedValue, _ := parser.Get(test.section, test.key)

			assert.Equal(t, test.expectedError, setError)
			if resultedValue != test.value {
				t.Fail()
			}

		})
	}
}

func TestINI_GetSectionNames(t *testing.T) {
	parser := IniFile{}
	testCases := []GetSectionNamesTestCase{
		{
			desc:        "Empty map as input",
			inputMap:    make(map[string]map[string]string),
			outputSlice: []string{},
		},
		{
			desc: "Only sections inside file",
			inputMap: map[string]map[string]string{"section1": make(map[string]string),
				"section2": make(map[string]string),
				"section3": make(map[string]string),
				"section4": make(map[string]string)},
			outputSlice: []string{"section1", "section2", "section3", "section4"},
		},
		{
			desc: "Normal case 1",
			inputMap: map[string]map[string]string{"section1": {
				"key1": "value1",
				"key2": "value2",
			},
			},
			outputSlice: []string{"section1"},
		},
		{
			desc: "Normal case 2",
			inputMap: map[string]map[string]string{"section1": {
				"key1": "value1",
				"key2": "value2",
				"key3": "value3",
				"key4": "value4",
			},
				"section2": {
					"key1": "value1",
					"key2": "value2",
				},
			},
			outputSlice: []string{"section1", "section2"},
		},
		{
			desc: "Normal case 3",
			inputMap: map[string]map[string]string{"section1": {
				"key1": "value1",
				"key2": "value2",
				"key3": "value3",
				"key4": "value4",
			},
				"section2": {
					"key1": "value1",
					"key2": "value2",
					"key3": "value3",
				},
				"section3": {
					"key1": "value1",
					"key2": "value2",
					"key3": "value3",
					"key4": "value4",
				},
			},
			outputSlice: []string{"section1", "section2", "section3"},
		},
		{
			desc: "Normal case 4",
			inputMap: map[string]map[string]string{"section1": {
				"key1": "value1",
				"key2": "value2",
				"key3": "value3",
				"key4": "value4",
				"key5": "value5",
				"key6": "value6",
				"key7": "value7",
			},
				"section2": {
					"key1": "value1",
					"key2": "value2",
					"key3": "value3",
				},
				"section3": {
					"key1": "value1",
					"key2": "value2",
					"key3": "value3",
					"key4": "value4",
				},
			},
			outputSlice: []string{"section1", "section2", "section3"},
		},
		{
			desc: "Normal case 5",
			inputMap: map[string]map[string]string{"section1": {
				"key1": "value1",
				"key2": "value2",
			},
				"section2": {
					"key1": "value1",
					"key2": "value2",
					"key3": "value3",
					"key4": "value4",
					"key5": "value5",
				},
				"section3": {
					"key1": "value1",
					"key2": "value2",
					"key3": "value3",
				},
				"section4": {
					"key1": "value1",
					"key2": "value2",
					"key3": "value3",
				},
			},
			outputSlice: []string{"section1", "section2", "section3", "section4"},
		},
	}
	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			resultedOutput := parser.String(test.inputMap)
			_, _ = parser.LoadFromString(resultedOutput)
			resultedSlice := parser.GetSectionNames()

			if !sliceEqual(test.outputSlice, resultedSlice) {
				t.Fail()
			}

		})
	}
}
