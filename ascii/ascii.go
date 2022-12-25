package ascii

import (
	"io/ioutil"
	"strings"
)

func Asciitext(input1, input2 string) string {
	if input2 != "shadow" && input2 != "tinkertoy" && input2 != "standard" {
		return "Incorrect template."
	}
	read := input2 + ".txt"
	template, err := ioutil.ReadFile(read)
	if err != nil {
		return "Internal error."
	}

	input := strings.ReplaceAll(input1, "\r", "")
	if input == "\n" {
		return " "
	}

	for i := 0; i < len(input); i++ {
		if (input[i] < 32 || input[i] > 127) && input[i] != 10 {
			return ("Incorrect input.")
		}
	}

	splitted := strings.Split(string(template)[1:], "\n\n")
	if len(splitted) != 95 {
		return ("Incorrect template.")
	}

	lines := strings.Split(input, "\n")
	res := ""

	for _, line := range lines {
		if line == "" && res != "" {
			res += string('\n')
			continue
		}
		for row := 0; row < 8; row++ {
			for i := 0; i < len(line); i++ {
				temp := strings.Split(splitted[line[i]-32], "\n")[row]
				for j := 0; j < len(temp); j++ {
					res += string(temp[j])
				}
			}
			res += string('\n')
		}
	}

	return res
}
