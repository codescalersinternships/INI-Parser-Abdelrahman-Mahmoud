package iniparser

import (
	"reflect"
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
	desc      string
	inputPath string
	inputMap  map[string]map[string]string
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
			path:          "testdata/test_01.txt",
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
			desc:      "Empty map as input",
			inputPath: "testdata/test_01.txt",
			inputMap:  make(map[string]map[string]string),
		},
		{
			desc:      "Only sections inside file",
			inputPath: "testdata/test_01.txt",
			inputMap: map[string]map[string]string{"section1": make(map[string]string),
				"section2": make(map[string]string),
				"section3": make(map[string]string),
				"section4": make(map[string]string)},
		},
		{
			desc:      "Normal case 1",
			inputPath: "testdata/test_01.txt",
			inputMap: map[string]map[string]string{"section1": {
				"key1": "value1",
				"key2": "value2",
			},
			},
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
		},
	}
	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			resultedError := parser.SaveToFile(test.inputPath, test.inputMap)
			resultedMap, _ := parser.LoadFromFile(test.inputPath)

			assert.Equal(t, nil, resultedError)
			if !reflect.DeepEqual(test.inputMap, resultedMap) {
				t.Fail()
			}

		})
	}
}
