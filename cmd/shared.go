package cmd

import (
	"bufio"
	"log"
	"os"
	"regexp"
)

// Copy the src file to dst. Any existing file will be overwritten and will not
// copy file attributes.
func CopyRemovingMagicComment(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	line := 0
	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		if line > 0 {
			out.WriteString(scanner.Text() + "\n")
		}
		line += 1
	}
	return out.Close()
}

func ReadFirstLine(filename string) string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	var firstLine string
	firstLine, err = reader.ReadString('\n')

	return firstLine
}

func FetchUrlFromComment(comment string) string {
	re := regexp.MustCompile(`[^ ]+\s+goga\s+([^# ]+)$`)
	return re.ReplaceAllString(comment, `$1`)
}

func GetSubdirectoryFromUrl(url string) string {
	re := regexp.MustCompile(`^https://github.com/[^\/]+/[^\/]+/blob/([^\/]+)/(.+)\n?$`)
	// $1 - branch
	return re.ReplaceAllString(url, `$2`)
}

func GetRepoFromUrl(url string) string {
	re := regexp.MustCompile(`^https://github.com/([^\/]+/[^\/]+)/blob/([^\/]+)/(.+)\n?$`)
	return re.ReplaceAllString(url, `git@github.com:$1.git`)
}
