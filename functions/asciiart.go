package asciiartweb

import (
	"strings" // Import for string manipulation
)

func PrintArt(input string, lines []string) string {
	// Check if the input is not ASCII and not a newline
	for _, char := range input {
		if (char < 32 || char > 126) && !(char == 10 || char == 13) {
			return "Error: Input accepts English letters, Arabic numerals, spaces, newlines, and regular special characters only."
		}
	}

	// Create a map to store ASCII art representations for each character
	asciiArtRunes := make(map[rune][]string)

	// Initialize an index to track characters (starting from ASCII code 32 - space)
	charIndex := 32

	// Iterate over the input lines, parsing ASCII art blocks
	for i := 1; i < len(lines); i += 9 {
		char := rune(charIndex)   // Get the current character
		art := lines[i : i+9]     // Extract the 9-line art block for the character
		asciiArtRunes[char] = art // Store the art block in the map
		charIndex++               // Increment index for the next character
	}

	// Split the input into individual lines
	splitText := strings.Split(input, "\n")

	// Initialize an empty output
	var output strings.Builder

	// Process each line of the input text
	for i, char := range splitText {
		if i > 0 {
			output.WriteByte('\n')
		}

		if char != "" { // Skip empty lines within the input
			// Process each character in the line
			for line := 0; line < 8; line++ {
				for _, letter := range char {
					// Retrieve the ASCII art for the current letter
					art, found := asciiArtRunes[letter]

					if found {
						output.WriteString(strings.TrimRight(art[line], "\r\n"))
					} else {
						output.WriteByte(' ')
					}
				}
				output.WriteByte('\n')
			}
		}
	}

	return output.String()
}
