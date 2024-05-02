package pesawat

// Pesawat adalah interface yang mendefinisikan perilaku yang diharapkan dari pesawat.
type Pesawat interface {
	Terbang()
	Landing()
	IzinTerbang()
}
