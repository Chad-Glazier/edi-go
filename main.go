package main

import (
	"time"

	"github.com/edi/cli"
	"github.com/edi/vi"
)

func main() {
	cli.RunGame(
		&vi.EDI{},
		&vi.EDI{},
		time.Second*5,
	)
}
