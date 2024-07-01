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

// type LoadFromFileTestCase struct {
// 	desc     string
// 	path     string
// 	expected error
// }

func TestINI_loadFromString(t *testing.T) {
	parser := IniFile{}
	emptyMap := make(map[string]map[string]string)
	testCases := []LoadFromStringTestCase{
		{
			desc:          "Empty string as input",
			input:         "",
			expectedError: errorFileIsEmpty,
			expectedMap:   emptyMap,
		},
		{
			desc:          "Only one comment line start with # as input",
			input:         "#This is a comment",
			expectedError: errorFileIsEmpty,
			expectedMap:   emptyMap,
		},
		{
			desc:          "Only one comment line start with ; as input",
			input:         ";This is a comment",
			expectedError: errorFileIsEmpty,
			expectedMap:   emptyMap,
		},
		{
			desc:          "Only one key value line as input",
			input:         "key = value",
			expectedError: errorFileIsEmpty,
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
			expectedError: errorFileIsEmpty,
			expectedMap:   emptyMap,
		},
		{
			desc: "Only sections as input",
			input: "[section1]\n " +
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
			expectedError: errorFileIsEmpty,
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
			expectedMap: map[string]map[string]string{"section1": map[string]string{
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
			expectedMap: map[string]map[string]string{"section1": map[string]string{
				"key1": "value1",
				"key2": "value2",
				"key3": "value3",
				"key4": "value4",
			},
				"section2": map[string]string{
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
			expectedMap: map[string]map[string]string{"section1": map[string]string{
				"key1": "value1",
				"key2": "value2",
				"key3": "value3",
				"key4": "value4",
			},
				"section2": map[string]string{
					"key1": "value1",
					"key2": "value2",
					"key3": "value3",
				},
				"section3": map[string]string{
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
			expectedMap: map[string]map[string]string{"section1": map[string]string{
				"key1": "value1",
				"key2": "value2",
				"key3": "value3",
				"key4": "value4",
				"key5": "value5",
				"key6": "value6",
				"key7": "value7",
			},
				"section2": map[string]string{
					"key1": "value1",
					"key2": "value2",
					"key3": "value3",
				},
				"section3": map[string]string{
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
			expectedMap: map[string]map[string]string{"section1": map[string]string{
				"key1": "value1",
				"key2": "value2",
			},
				"section2": map[string]string{
					"key1": "value1",
					"key2": "value2",
					"key3": "value3",
					"key4": "value4",
					"key5": "value5",
				},
				"section3": map[string]string{
					"key1": "value1",
					"key2": "value2",
					"key3": "value3",
				},
				"section4": map[string]string{
					"key1": "value1",
					"key2": "value2",
					"key3": "value3",
				},
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			resultedMap, resultedError := parser.loadFromString(test.input)

			assert.Equal(t, test.expectedError, resultedError)
			if !reflect.DeepEqual(test.expectedMap, resultedMap) {
				t.Fail()
			}

		})
	}

}
