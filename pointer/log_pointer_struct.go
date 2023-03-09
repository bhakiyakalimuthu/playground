package main

import (
	"fmt"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/log"
)

func main() {
	h := handler{serviceName: "playground"}
	log.Root().SetHandler(h)
	log.Root().SetHandler(log.LvlFilterHandler(log.LvlInfo, log.StreamHandler(os.Stderr, log.TerminalFormat(true))))
	log.Root().New()

	l := log.New(log.Ctx{"service": "playground"})
	l.Info("before log pointer")
	x := &pointer{
		a: "test",
		b: 10,
		c: big.NewInt(30),
	}
	ll := l.New(log.Ctx{"service-1": "groundplay"})
	ll.Info("before log pointer", "var", fmt.Sprintf("%p", x))
}

type handler struct {
	serviceName string
}

func (h handler) Log(r *log.Record) error {
	r.Msg = h.serviceName
	return nil
}

type pointer struct {
	a string
	b int
	c *big.Int
}
