package model

import "fmt"

func GetBio() {
	biodata := Biodata{
		Nama:             "Ucup",
		Foto:             "ucup-profile.jpg",
		Email:            "ucup@mail.com",
		Umur:             15,
		NomorTelepon:     "0812345678910",
		StatusPernikahan: "Duda",
		RiwayatPendidikan: []RiwayatPendidikan{
			{
				Nama:    "Universitas Indonesia",
				Jurusan: "Sastra Jawa",
			},
		},
	}

	fmt.Println(biodata)

}
