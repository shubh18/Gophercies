package hyperlink

import (
	"net/url"
	"strings"
)

//CreateLinks builds hyperlinks to programs giving error
func CreateLinks(stacktrace string) string {
	lines := strings.Split(stacktrace, "\n")
	for index, line := range lines {

		if len(line) == 0 || line[0] != '\t' {
			continue
		}
		file := ""
		for i, ch := range line {
			if ch == ':' {
				file = line[1:i]
				break
			}
		}
		var lineNoBuilder strings.Builder
		for i := len(file) + 2; i < len(line); i++ {
			if line[i] < '0' || line[i] > '9' {
				break
			}
			lineNoBuilder.WriteByte(line[i])
		}
		fileURL := url.Values{}
		fileURL.Set("path", file)
		fileURL.Set("line", lineNoBuilder.String())
		lines[index] = "\t<a href=\"/debug/?" + fileURL.Encode() + "\">" + file + ":" + lineNoBuilder.String() + "</a>" + line[len(file)+2+len(lineNoBuilder.String()):]
	}
	hyperlink := strings.Join(lines, "\n")
	return hyperlink
}
