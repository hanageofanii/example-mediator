package menara_kontrol

// MenaraKontrol adalah interface yang mendefinisikan metode yang diharapkan dari menara kontrol.
type MenaraKontrol interface {
	BolehTerbang(Pesawat) bool
	NotifikasiLanding()
}
