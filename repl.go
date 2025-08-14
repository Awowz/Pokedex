package main

// This func would have benifited from "strings" packaged. i just wanted to exercise my brain
func cleanInput(text string) []string {
	var stringArray []string
	wordSetup := ""
	for index, char := range text {
		if char == ' ' {
			if wordSetup != "" {
				stringArray = append(stringArray, wordSetup)
				wordSetup = ""
				continue
			}
		} else {
			if char >= 'A' && char <= 'Z' {
				wordSetup += string(char + ('a' - 'A'))
			} else {
				wordSetup += string(char)
			}

			if len(text)-1 == index {
				stringArray = append(stringArray, wordSetup)
				wordSetup = ""
				continue
			}
		}
	}
	return stringArray
}
