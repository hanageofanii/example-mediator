package pesawat

import "fmt"

type PesawatPenumpang struct {
	MenaraKontrol MenaraKontrol
}

// Metode Terbang untuk pesawat penumpang.
func (p *PesawatPenumpang) Terbang() {
	if !p.MenaraKontrol.BolehTerbang(p) {
		fmt.Println("PesawatPenumpang: Terbang ditolak, menunggu izin")
		return
	}
	fmt.Println("PesawatPenumpang: Sudah Terbang")
}

// Metode Landing untuk pesawat penumpang.
func (p *PesawatPenumpang) Landing() {
	fmt.Println("PesawatPenumpang: Mendarat")
	p.MenaraKontrol.NotifikasiLanding()
}

// Metode IzinTerbang untuk pesawat penumpang.
func (p *PesawatPenumpang) IzinTerbang() {
	fmt.Println("PesawatPenumpang: Izin terbang diberikan")
	p.Terbang()
}
