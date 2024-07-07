# INI-Parser-Abdelrahman-Mahmoud

## Introduction

An INI file is a configuration file for computer software that consists of a text-based content with a structure and syntax comprising key–value pairs for properties, and sections that organize the properties.[1] The name of these configuration files comes from the filename extension INI, for initialization, used in the MS-DOS operating system which popularized this method of software configuration. The format has become an informal standard in many contexts of configuration, but many applications on other operating systems use different file name extensions, such as conf and cfg. for more visit [INI File](https://en.wikipedia.org/wiki/INI_file) 

In programming, "parsing" means analyzing and breaking down structured data, like text or code, into its individual components to understand or process it. It's like taking a sentence and identifying each word or analyzing a mathematical expression to evaluate it step by step. for more visit [Parser](https://www.quora.com/What-exactly-does-parsing-mean-in-programming)

This Package allows the user to read, write and Modify INI files easily.

## Features

### Read INI File 

To read an INI file LoadFromFile function is used to scan the file line by line extract comments, sections and keyValue pairs. A private function loadFromString is used to extract comments, sections and keyValue pairs and save them for later usage.

### Retrive section names

To retrieve INI file section names GetSectionNames function is used.

### Retrive sections and keyValue pairs

To retrieve INI file sections and keyValue pairs GetSections function is used to return a string that consists of a unique structure that describes the contents of the INI file.

### Retrive value for a certain key inside a section

To retrieve value for a certain key inside a section Get function is used.

### Add key value pair 

To add a key value pair to INI file Set function is used to add to existing section or to create a new section and add the new key value pair.

### Write INI File 

To read an INI file SaveToFile function is used to write the INI file comments, sections and keyValue pairs. A private function toString is used to compose comments, sections and keyValue pairs into multi-lined string in a similar structure to the initial INI file.

## Setup

1. Clone the Repository to a directory of your choice.
2. Make sure you have go version 1.22.4 installed on your device
3. Create demo.go file inside the working directory
4. import the package using
   ```GO
	  import "github.com/codescalersinternships/INI-Parser-Abdelrahman-Mahmoud/pkg"
   ```
5. Finish writing your code
6. Add the INI file to the working directory
7. Open terminal
8. Build the project using
   ```console
   user@user-VirtualBox:~$ go build demo.go
   ```
9. Run the project using
   ```console
   user@user-VirtualBox:~$ ./demo
   ```

## Demo
- Code:
```GO
	parser := iniparser.IniFile{}

	_ = parser.LoadFromFile("test.ini")

	fmt.Println(parser.GetSectionNames())
	fmt.Println(parser.GetSections())
	fmt.Println(parser.Get("section 1", "key1"))

	parser.Set("section 3", "key1", "value1")
	parser.Set("section 1", "key3", "value3")

	
	fmt.Println(parser.GetSections())

	_ = parser.SaveToFile()
```

- Output:
```console
user@user-VirtualBox:~$ ./demo
[section 3 section 1 section 2]
{ section 1: {key3: value3, key1: value1, key2: value2}, section 2: {key3: value3}, section 3: {key1: value1} }
value1
{ section 1: {key1: value1, key2: value2, key3: value3}, section 2: {key3: value3}, section 3: {key1: value1} }
#This is a comment

[section 1]
key1=value1
key2=value2
key3=value3

[section 2]
key3=value3

[section 3]
key1=value1
```

## Tests

--- PASS: TestINI_LoadFromString (0.00s)
    --- PASS: TestINI_LoadFromString/Empty_string_as_input (0.00s)
    --- PASS: TestINI_LoadFromString/Only_one_comment_line_start_with_#_as_input (0.00s)
    --- PASS: TestINI_LoadFromString/Only_one_comment_line_start_with_;_as_input (0.00s)
    --- PASS: TestINI_LoadFromString/Only_one_key_value_line_as_input (0.00s)
    --- PASS: TestINI_LoadFromString/Only_one_section_line_as_input (0.00s)
    --- PASS: TestINI_LoadFromString/Only_comments_as_input (0.00s)
    --- PASS: TestINI_LoadFromString/Only_sections_as_input (0.00s)
    --- PASS: TestINI_LoadFromString/Only_keys_values_as_input (0.00s)
    --- PASS: TestINI_LoadFromString/Normal_case_1 (0.00s)
    --- PASS: TestINI_LoadFromString/Normal_case_2 (0.00s)
    --- PASS: TestINI_LoadFromString/Normal_case_3 (0.00s)
    --- PASS: TestINI_LoadFromString/Normal_case_4 (0.00s)
    --- PASS: TestINI_LoadFromString/Normal_case_5 (0.00s)
--- PASS: TestINI_loadFromFile (0.00s)
    --- PASS: TestINI_loadFromFile/Wrong_file_path (0.00s)
    --- PASS: TestINI_loadFromFile/Empty_file_as_input (0.00s)
    --- PASS: TestINI_loadFromFile/Only_one_comment_line_start_with_#_inside_file (0.00s)
    --- PASS: TestINI_loadFromFile/Only_one_comment_line_start_with_;_inside_file (0.00s)
    --- PASS: TestINI_loadFromFile/Only_one_key_value_line_inside_file (0.00s)
    --- PASS: TestINI_loadFromFile/Only_one_section_line_inside_file (0.00s)
    --- PASS: TestINI_loadFromFile/Only_comments_inside_file (0.00s)
    --- PASS: TestINI_loadFromFile/Only_sections_inside_file (0.00s)
    --- PASS: TestINI_loadFromFile/Only_keys_values_inside_file (0.00s)
    --- PASS: TestINI_loadFromFile/Normal_case_1 (0.00s)
    --- PASS: TestINI_loadFromFile/Normal_case_2 (0.00s)
    --- PASS: TestINI_loadFromFile/Normal_case_3 (0.00s)
    --- PASS: TestINI_loadFromFile/Normal_case_4 (0.00s)
    --- PASS: TestINI_loadFromFile/Normal_case_5 (0.00s)
--- PASS: TestINI_String (0.00s)
    --- PASS: TestINI_String/Empty_map_as_input (0.00s)
    --- PASS: TestINI_String/Only_sections_inside_file (0.00s)
    --- PASS: TestINI_String/Normal_case_1 (0.00s)
    --- PASS: TestINI_String/Normal_case_2 (0.00s)
    --- PASS: TestINI_String/Normal_case_3 (0.00s)
    --- PASS: TestINI_String/Normal_case_4 (0.00s)
    --- PASS: TestINI_String/Normal_case_5 (0.00s)
--- PASS: TestINI_SaveToFile (0.01s)
    --- PASS: TestINI_SaveToFile/Empty_map_as_input (0.00s)
    --- PASS: TestINI_SaveToFile/Only_sections_inside_file (0.00s)
    --- PASS: TestINI_SaveToFile/Normal_case_1 (0.00s)
    --- PASS: TestINI_SaveToFile/Normal_case_2 (0.00s)
    --- PASS: TestINI_SaveToFile/Normal_case_3 (0.00s)
    --- PASS: TestINI_SaveToFile/Normal_case_4 (0.00s)
    --- PASS: TestINI_SaveToFile/Normal_case_5 (0.00s)
--- PASS: TestINI_GetSections (0.00s)
    --- PASS: TestINI_GetSections/Empty_map_as_input (0.00s)
    --- PASS: TestINI_GetSections/Only_sections_inside_file (0.00s)
    --- PASS: TestINI_GetSections/Normal_case_1 (0.00s)
    --- PASS: TestINI_GetSections/Normal_case_2 (0.00s)
    --- PASS: TestINI_GetSections/Normal_case_3 (0.00s)
    --- PASS: TestINI_GetSections/Normal_case_4 (0.00s)
    --- PASS: TestINI_GetSections/Normal_case_5 (0.00s)
--- PASS: TestINI_Get (0.00s)
    --- PASS: TestINI_Get/Empty_map_as_input (0.00s)
    --- PASS: TestINI_Get/Only_sections_inside_file (0.00s)
    --- PASS: TestINI_Get/Normal_case_1 (0.00s)
    --- PASS: TestINI_Get/Normal_case_2 (0.00s)
    --- PASS: TestINI_Get/Normal_case_3 (0.00s)
    --- PASS: TestINI_Get/Normal_case_4 (0.00s)
    --- PASS: TestINI_Get/Normal_case_5 (0.00s)
--- PASS: TestINI_Set (0.00s)
    --- PASS: TestINI_Set/Empty_map_as_input (0.00s)
    --- PASS: TestINI_Set/Only_sections_inside_file (0.00s)
    --- PASS: TestINI_Set/Normal_case_1 (0.00s)
    --- PASS: TestINI_Set/Normal_case_2 (0.00s)
    --- PASS: TestINI_Set/Normal_case_3 (0.00s)
    --- PASS: TestINI_Set/Normal_case_4 (0.00s)
    --- PASS: TestINI_Set/Normal_case_5 (0.00s)
--- PASS: TestINI_GetSectionNames (0.00s)
    --- PASS: TestINI_GetSectionNames/Empty_map_as_input (0.00s)
    --- PASS: TestINI_GetSectionNames/Only_sections_inside_file (0.00s)
    --- PASS: TestINI_GetSectionNames/Normal_case_1 (0.00s)
    --- PASS: TestINI_GetSectionNames/Normal_case_2 (0.00s)
    --- PASS: TestINI_GetSectionNames/Normal_case_3 (0.00s)
    --- PASS: TestINI_GetSectionNames/Normal_case_4 (0.00s)
    --- PASS: TestINI_GetSectionNames/Normal_case_5 (0.00s)
PASS
coverage: 100.0% of statements
ok      github.com/codescalersinternships/INI-Parser-Abdelrahman-Mahmoud/pkg    0.028s  coverage: 100.0% of statements
