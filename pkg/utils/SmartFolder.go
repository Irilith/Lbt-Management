package utils

import (
	"regexp"
)

func GenerateFolder(file string) string {
	re := regexp.MustCompile(`\.[a-zA-Z0-9]+$`)
	file = re.ReplaceAllString(file, "")
	patterns := []string{
		`\s*-\s*(Vol(?:ume)?(?: \d+)?)\s*$`,
		`\s*(volume|Vol|vol|Volume|v\d+)\s*.*`, // Remove volume and everything after it
		`\s*\[\s*[^\]]*\s*\]\s*`,               // Light novel usually have the Publisher and Uploader in the square bracket
		`\s*\(\s*[^\)]*\s*\)\s*`,               // Manga usually have the Publisher and Uploader in the round bracket
		`\s*\{\s*[^\}]*\s*\}\s*`,               // There are some file that contain {r} which i assume its reprint or something (its when the book is reprinted with some changes)
	}

	for _, pattern := range patterns {
		re := regexp.MustCompile(pattern)
		file = re.ReplaceAllString(file, " ")
	}

	return file
}
