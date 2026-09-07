package model

type RiwayatPendidikan struct {
	Nama    string
	Jurusan string
}

type Biodata struct {
	Nama              string
	Foto              string
	Email             string
	Umur              int8
	NomorTelepon      string
	StatusPernikahan  string
	RiwayatPendidikan []RiwayatPendidikan
}
