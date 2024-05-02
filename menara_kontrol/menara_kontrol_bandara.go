package menara_kontrol

import "fmt"

// MenaraKontrolBandara adalah implementasi konkret dari menara kontrol.
type MenaraKontrolBandara struct {
	IsRunwayFree bool
}

// NewMenaraKontrolBandara membuat instance baru dari MenaraKontrolBandara.
func NewMenaraKontrolBandara() *MenaraKontrolBandara {
	return &MenaraKontrolBandara{
		IsRunwayFree: true,
	}
}

// BolehTerbang memeriksa apakah pesawat boleh terbang.
func (m *MenaraKontrolBandara) BolehTerbang(p Pesawat) bool {
	if m.IsRunwayFree {
		m.IsRunwayFree = false
		return true
	}
	return false
}

// NotifikasiLanding memberitahu menara kontrol tentang pendaratan pesawat.
func (m *MenaraKontrolBandara) NotifikasiLanding() {
	m.IsRunwayFree = true
	fmt.Println("MenaraKontrol: Pemberitahuan - Pesawat telah mendarat")
}
