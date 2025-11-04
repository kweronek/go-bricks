package main

import (
	"fmt"
	"strconv"
	"time"
)

type laenderCode = [2]byte
type pruefZiffern = [2]byte
type bankLeitZahl = [8]byte
type kontoNummer = [10]byte
type iban = [22]byte

type ibanLog struct {
	landesKenner [2]byte
	pruefSumme   [2]byte
	bankLeitZahl [8]byte
	kontoNummer  [10]byte
}

type buchung struct {
	vonIBAN        iban
	nachIBAN       iban
	buchungsBetrag int64 // cents
	buchungsZeit   time.Time
	valutaZeit     time.Time
}

type konto struct {
	kontoBesitzerID uint64
	kontoIBAN       iban
	buchungen       []buchung
}

func decomposeIBAN(myIBAN iban) (ok bool, myLandesKenner laenderCode, myPruefSumme pruefZiffern, myBLZ bankLeitZahl, myKtoNr kontoNummer) {
	ok = true
	copy(myLandesKenner[0:2], myIBAN[0:2])
	copy(myPruefSumme[0:2], myIBAN[2:4])
	copy(myBLZ[0:8], myIBAN[4:12])
	copy(myKtoNr[0:10], myIBAN[12:22])
	return
}

func parseIbanString(myIBAN iban) (landesKenner string, pruefSumme string, bankLeitZahl string, kontoNummer string) {
	landesKenner = string(myIBAN[0:2])
	pruefSumme = string(myIBAN[2:4])
	bankLeitZahl = string(myIBAN[4:12])
	kontoNummer = string(myIBAN[12:22])
	return
}

func convertIbanS2L(myIbanString string) (myIbanLog ibanLog) {
	copy(myIbanLog.landesKenner[:], []byte(myIbanString[0:2]))
	copy(myIbanLog.pruefSumme[:], []byte(myIbanString[2:4]))
	copy(myIbanLog.bankLeitZahl[:], []byte(myIbanString[4:12]))
	copy(myIbanLog.kontoNummer[:], []byte(myIbanString[12:22]))
	return
}

func main() {
	var myIbanString = "DE12123456781234567890"
	fmt.Println(myIbanString)

	if myIbanString[0:2] == "DE" {
		var myIbanLog ibanLog = convertIbanS2L(myIbanString)
		fmt.Println("Landeskenner:", string(myIbanLog.landesKenner[:]))
		if s, err := strconv.Atoi(string(myIbanLog.pruefSumme[:])); err == nil {
			fmt.Println("Prüfsumme:", s)
		}
		fmt.Println("Bankleitzahl: ", string(myIbanLog.bankLeitZahl[:]))
		fmt.Println("Kontonummer:", string(myIbanLog.kontoNummer[:]))

		fmt.Println("---")
		fmt.Print(myIbanString[0:2])
		fmt.Println(myIbanString[2:4])
		fmt.Println(myIbanString[4:12])
		fmt.Println(myIbanString[12:22])

		actLandesKenner := myIbanString[0:2]
		actPruefZiffern := myIbanString[2:4]
		actBankLeitZahl := myIbanString[4:12]
		actKontoNummer  := myIbanString[12:22]

		fmt.Println("Landeskenner:",actLandesKenner,actPruefZiffern,"BLZ:",actBankLeitZahl,"Kto.-Nr.:",actKontoNummer)
	}

}
