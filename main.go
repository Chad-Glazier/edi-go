package main

import (
	"time"

	"github.com/edi/cli"
	"github.com/edi/eval"
	"github.com/edi/search/mm"
)

func main() {
	cli.RunGame(
		mm.AlphaBeta(eval.KMinDist),
		mm.HistoricAlphaBeta(eval.KMinDist),
		time.Second*5,
	)
}
