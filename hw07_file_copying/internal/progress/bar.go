package progress

import (
	"fmt"
	"strings"
)

var (
	emptyLetter = "."
	fullLetter  = "█"
	maxProgress = 10
)

type Bar struct {
	name       string
	total      int64
	current    int64
	applicator float32
}

func NewBar(name string, total int64) *Bar {
	return &Bar{
		total:      total,
		applicator: float32(total) / float32(maxProgress),
		name:       fmt.Sprintf("[%s]: ", name),
	}
}

func (b *Bar) Add(value int64) {
	b.current += value

	b.printString(fmt.Sprintf("%d/%d", b.current, b.total))
}

func (b *Bar) Current() int64 {
	return b.current
}

func (b *Bar) Finish() {
	if b.current < b.total {
		b.printString("FAILED  \n")
	} else {
		b.printString("DONE    \n")
	}
}

func (b *Bar) printString(postfix string) {
	fmt.Print(
		"\r", b.name,
		progressString(b.current, b.applicator),
		" ", postfix,
	)
}

func progressString(curr int64, applicator float32) string {
	var result strings.Builder
	fullCount := int(float32(curr) / applicator)
	for i := 0; i < maxProgress; i++ {
		if i <= fullCount {
			result.WriteString(fullLetter)
		} else {
			result.WriteString(emptyLetter)
		}
	}
	return result.String()
}
