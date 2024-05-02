package pesawat

import "fmt"

// PesawatBarang adalah implementasi konkret dari pesawat barang.
type PesawatBarang struct {
	MenaraKontrol MenaraKontrol
}

// Metode Terbang untuk pesawat barang.
func (p *PesawatBarang) Terbang() {
	if !p.MenaraKontrol.BolehTerbang(p) {
		fmt.Println("PesawatBarang: Terbang ditolak, menunggu izin")
		return
	}
	fmt.Println("PesawatBarang: Sudah Terbang")
}

// Metode Landing untuk pesawat barang.
func (p *PesawatBarang) Landing() {
	fmt.Println("PesawatBarang: Mendarat")
	p.MenaraKontrol.NotifikasiLanding()
}

// Metode IzinTerbang untuk pesawat barang.
func (p *PesawatBarang) IzinTerbang() {
	fmt.Println("PesawatBarang: Izin terbang diberikan")
	p.Terbang()
}
