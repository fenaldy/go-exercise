package main

import (
	"fmt"
)

type Kelas struct {
	KelasName string
	NamaMurid []string
}

func (k *Kelas) setKelasName(kelasName string) {
	k.KelasName = kelasName
}

func (k *Kelas) addMurid(murid string) {
	if k.getTotalMurid() >= 3 {
		fmt.Println("CLASS FULL !!!")
	} else {
		k.NamaMurid = append(k.NamaMurid, murid)
	}
}

func (k *Kelas) getKelasName() string {
	return k.KelasName
}

func (k *Kelas) getTotalMurid() int {
	return len(k.NamaMurid)
}

func (k *Kelas) showAllMuridName() {
	for x := 0; x < len(k.NamaMurid); x++ {
		fmt.Println(k.NamaMurid[x])
	}
}

func main() {
	kelas := Kelas{}
	kelas.setKelasName("X-IPA")
	fmt.Println(kelas.getKelasName())
	kelas.addMurid("Andrew1")
	kelas.addMurid("Andrew2")
	kelas.addMurid("Andrew3")
	kelas.addMurid("Andrew4")

	fmt.Println(kelas.getTotalMurid())
	kelas.showAllMuridName()
}
