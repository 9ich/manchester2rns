package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	flag.Parse()
	if flag.NArg() != 1 {
		fmt.Println("usage: manchester2rns string")
		os.Exit(2)
	}
	msg := strings.ToLower(flag.Arg(0))
	fmt.Println(manchester2rns(msg))
}

func manchester2rns(msg string) string {
	out := beginning
	m := manc(msg)
	for _, c := range m {
		if c == '0' {
			out += mkline("C-3")
		} else {
			out += mkline("C-4")
		}
	}
	return out + end
}

func manc(s string) string {
	out := ""
	for _, c := range s {
		b := byte(c)
		for i := uint(0); i != 8; i++ {
			if (b>>i)&1 == 0 {
				out += "01"
			} else {
				out += "10"
			}
		}
	}
	return out
}

var lineIndex = 0

func mkline(note string) string {
	lineIndex++

	var instrument = "00"
	if note == "" || note == "OFF" {
		instrument = ".."
	}

	if note == "" {
		return fmt.Sprintf("\n        <Line index=\"%d\"/>	", lineIndex-1)
	}
	return fmt.Sprintf(`
          <Line index="%d">
            <NoteColumns>
              <NoteColumn>
                <Note>%s</Note>
                <Instrument>%s</Instrument>
                <Volume>..</Volume>
                <Panning>..</Panning>
                <Delay>..</Delay>
              </NoteColumn>
            </NoteColumns>
          </Line>`, lineIndex-1, note, instrument)
}

const beginning = `<?xml version="1.0" encoding="UTF-8"?>
<PatternClipboard.BlockBuffer doc_version="0">
  <TrackColumns>
    <TrackColumn>
      <TrackColumn>
        <Lines>`

const end = `
        </Lines>
        <ColumnType>NoteColumn</ColumnType>
      </TrackColumn>
    </TrackColumn>
  </TrackColumns>
</PatternClipboard.BlockBuffer>`
