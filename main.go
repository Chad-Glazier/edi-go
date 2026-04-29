package main

import (
	"time"

	"github.com/edi/cli"
	"github.com/edi/eval"
	"github.com/edi/search/alpha_beta"
)

func main() {
	cli.RunGame(
		alpha_beta.HistoricAlphaBeta(eval.KMinDist),
		alpha_beta.HistoricAlphaBeta(eval.KMinDist),
		time.Second * 5,
	)
}
