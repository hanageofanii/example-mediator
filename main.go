package main

import (
	"github.com/example-mediator/menara_kontrol"
	"github.com/example-mediator/pesawat"
)

func main() {
	menaraKontrol := menara_kontrol.NewMenaraKontrolBandara()

	pesawatPenumpang := &pesawat.PesawatPenumpang{
		MenaraKontrol: menaraKontrol,
	}
	pesawatBarang := &pesawat.PesawatBarang{
		MenaraKontrol: menaraKontrol,
	}

	pesawatPenumpang.IzinTerbang()
	pesawatBarang.IzinTerbang()
	pesawatPenumpang.Landing()
}
