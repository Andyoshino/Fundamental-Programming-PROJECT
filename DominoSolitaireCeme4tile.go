package main
/*	DATA KELOMPOK
	Judul	         : Domino Solitaire Ceme-4tile
	Anggota Kelompok : 1301190206 Butrahandisya
					   1301190396 Qomarudin Sifak

*/
import ("fmt"
		"math/rand"
		"time"
		"os"
		"os/exec")

type Pemain struct {
	namaPemain string
	matchmakingPoints int
	gameYangDimenangkan int
	gameYangSudahDimainkan int
	nomerPemain int
}

const pemainMax int = 100
type himpunanPemain [pemainMax]Pemain


type Domino struct {
	sisiKanan int
	sisiKiri int
}

//LoadingScreen
func loadingScreen() {
	var i int
	for i <= 100 {
		fmt.Print("\n\n\n\n")
		fmt.Print("\t\t\t\t\tLoading..            ",i,"%")
		clearScreen()
		i = i + 5
	}
}

//fungsi clear screen
func clearScreen() {
	cmd := exec.Command("cmd", "/c", "cls")
    cmd.Stdout = os.Stdout
    cmd.Run()
}
//Tampilan saat game dijalankan
func mainScreen() {
	fmt.Println("\t\t\t\t\t  _______________________________________________________________________")
	fmt.Println("\t\t\t\t\t||                                                                     ||")
	fmt.Println("\t\t\t\t\t||                             Welcome to                              ||")
	fmt.Println("\t\t\t\t\t||                         Domino Solitaire                            ||")
	fmt.Println("\t\t\t\t\t||                             Ceme-4tile                              ||")
	fmt.Println("\t\t\t\t\t||             A game by Butrahandisya & Qomarudin Sifak               ||")
	fmt.Println("\t\t\t\t\t[]_____________________________________________________________________[]")
	fmt.Println("\t\t\t\t\t||                          M a i n  M e n u                           ||")
	fmt.Println("\t\t\t\t\t||                              [1] Play                               ||")
	fmt.Println("\t\t\t\t\t||                              [2] Score                              ||")
	fmt.Println("\t\t\t\t\t||                              [3] Instruction                        ||")
	fmt.Println("\t\t\t\t\t||                              [4] Stat User                          ||")
	fmt.Println("\t\t\t\t\t||                              [5] Exit                               ||")
	fmt.Println("\t\t\t\t\t[]_____________________________________________________________________[]")
	fmt.Print("\t\t\t\t\t--= What are you going to do? : ")
}

//
func login(sedangBermain *Pemain, nomerYangSedangBermain *int, totalPemain int, listPemain himpunanPemain, statusKeberhasilan *bool) {
	/*	I.S Terdefinisi totalPemain bernilai bilangan bulat
		F.S Mengembalikan pemain yang akan bermain
	*/
	var inputanLogin Pemain
	var statusLog bool
	var indeksFound int
	var konfirmasi string
	fmt.Print("\t\t\t\t\tMasukkan username anda (case sensitive) : ")
	fmt.Scan(&inputanLogin.namaPemain)
	searchingNama(inputanLogin.namaPemain, totalPemain, listPemain, &statusLog, &indeksFound)
	if statusLog {
		fmt.Println("\t\t\t\t\tLog-in berhasil!")
		*sedangBermain = listPemain[indeksFound]
		*nomerYangSedangBermain = sedangBermain.nomerPemain
	} else {
		fmt.Println("\t\t\t\t\tLog-in gagal...")
		fmt.Print("\t\t\t\t\tUsername '", inputanLogin.namaPemain, "' tidak terdaftar.\n")
		fmt.Print("\t\t\t\t\tKembali? [Y] : ")
		fmt.Scan(&konfirmasi)
		for konfirmasi != "Y" && konfirmasi != "y" {
			fmt.Print("\t\t\t\t\tInputan hanya menerima 'Y' & 'y'!\n")
			fmt.Print("\t\t\t\t\tKembali? [Y] : ")
			fmt.Scan(&konfirmasi)
		}
	}
	*statusKeberhasilan = statusLog
}

// Register
func register(totalPemain *int, listPemain himpunanPemain, statusUsernameSama *bool, mendaftar *Pemain) {
	/*	I.S Terdefinisi array listPemain
		F.S Registasi pemain,jika statusUsernameSama / sudah pernah terdaftar maka akan terjadi login
	*/
	var registering Pemain
	var indeksKetemu int
	var statusKetemu bool
	fmt.Println("\t\t\t\t\tSebelum bermain, Anda perlu meregistrasikan diri anda sebagai pemain.")
	fmt.Print("\t\t\t\t\tNama : ")
	fmt.Scan(&registering.namaPemain)
	searchingNama(registering.namaPemain, *totalPemain, listPemain, &statusKetemu, &indeksKetemu)
	if statusKetemu {
		fmt.Print("\t\t\t\t\tApakah anda bermaksud untuk log-in dengan username '", listPemain[indeksKetemu].namaPemain, "' ?\n")
		fmt.Println("\t\t\t\t\tLog-in berhasil!")
		*mendaftar = listPemain[indeksKetemu]
	} else {
		registering.gameYangDimenangkan = 0
		registering.matchmakingPoints = 0
		registering.gameYangSudahDimainkan = 0
		registering.nomerPemain = *totalPemain+1
		*totalPemain = *totalPemain+1
		*mendaftar = registering
	}
	*statusUsernameSama = statusKetemu
}

func searchingNama(namaInputan string, totalPemain int, listPemain himpunanPemain, statusFinding *bool, indeksKetemu *int) {
	/*	I.S Terdefinisi namaInputan,totalPemain,array listPemain
		F.S Mencari pemain menggunakan metode sequentinal search, dan mengembalikan statusFinding,indeks pemain yang dicari
	*/
	var indeksSearchingNama int
	indeksSearchingNama = 0
	*statusFinding = false
	for indeksSearchingNama < totalPemain && !*statusFinding {
		*statusFinding = (namaInputan == listPemain[indeksSearchingNama].namaPemain)
		if *statusFinding {
			*indeksKetemu = indeksSearchingNama
		}
		indeksSearchingNama = indeksSearchingNama + 1
	}
}


func loginRegistScreen(optionPlaying *int) {
	/*	I.S Terdefinisi optionPlaying bernilai bilangan bulat
		F.S mengembalikan optionPlaying, yang akan menentukan apa yang akan dilakukan program selanjutnya.
	*/
	var opsiPemainLokal int
	fmt.Println("\t\t\t\t\t[Choose one before playing]")
	fmt.Println("\t\t\t\t\t[1] Log-in")
	fmt.Println("\t\t\t\t\t[2] Register")
	fmt.Print("\t\t\t\t\tWhat are you going to do? ")
	fmt.Scan(&opsiPemainLokal)
	*optionPlaying = opsiPemainLokal
	for *optionPlaying != 1 && *optionPlaying != 2 {
		fmt.Println("\t\t\t\t\tInputan hanya ada 1 dan 2!")
		fmt.Print("\t\t\t\t\tPilihan :")
		fmt.Scan(&opsiPemainLokal)
		*optionPlaying = opsiPemainLokal
	}
}

//
func createDominoSets (dominoClean *[28]Domino) {
	/*	I.S Terdefinisi array DominoClean
		F.S Mengisi array DominoClean
	*/
	var Indeks, sisiKanan, sisiKiri, det int
	Indeks = 0
	sisiKanan = 6
	sisiKiri = 0
	det = 7
	for Indeks < 28 {
		dominoClean[Indeks].sisiKiri = sisiKiri
		dominoClean[Indeks].sisiKanan = sisiKanan
		sisiKiri = sisiKiri + 1
		if sisiKiri == det {
			sisiKiri = 0
			sisiKanan = sisiKanan - 1
		    det = det - 1
		}
		Indeks = Indeks + 1
	}
}

//Mengacak indeks mana saja yang akan saling bertukar tempat
func acakIndeks(Arandomed, Brandomed *int) {
	/*	I.S Terdefinisi Arandomed dan Brandomed bernilai bilangan bulat
		F.S Mengacak Arandomed dan Brandomed menggunakan rand
	*/
	*Arandomed = rand.Intn(28)
	*Brandomed = rand.Intn(28)
}

func acakIndeksSolo(randomed *int) {
	/*	I.S Terdefinisi randomed bernilai bilangan bulat
		F.S Mengacak randomed menggunakan rand
	*/
	*randomed = rand.Intn(28)
}

//Fungsi yang akan mengacak posisi 28 domino
func acakDomino(dominoSaatIni [28]Domino) [28]Domino {
	/*	Mengembalikan dominoSaatIni dalam posisi yang sudah teracak
	*/
	var indeksyangDitukarA, indeksyangDitukarB, jmlDominoTeracak, indeksPengecekan, det1x int
	var dominoRandomed [28]Domino
	var dominoTemp Domino
	var RandomedIndex [28]int
	jmlDominoTeracak = -1
	dominoRandomed = dominoSaatIni
	for jmlDominoTeracak < 27 {
		acakIndeks(&indeksyangDitukarA, &indeksyangDitukarB)
		for indeksyangDitukarA == indeksyangDitukarB {
			acakIndeks(&indeksyangDitukarA, &indeksyangDitukarB)
		}
		if det1x == 1 {
			indeksPengecekan = 0
			for indeksPengecekan <= jmlDominoTeracak {
				if RandomedIndex[indeksPengecekan] == indeksyangDitukarA {
					indeksPengecekan = -1
					acakIndeksSolo(&indeksyangDitukarA)
					for indeksyangDitukarA == indeksyangDitukarB {
						acakIndeksSolo(&indeksyangDitukarA)
					}
				}
				indeksPengecekan = indeksPengecekan + 1
			}
			indeksPengecekan = 0
			for indeksPengecekan <= jmlDominoTeracak {
				if RandomedIndex[indeksPengecekan] == indeksyangDitukarB {
					indeksPengecekan = -1
					acakIndeksSolo(&indeksyangDitukarB)
					for indeksyangDitukarA == indeksyangDitukarB {
						acakIndeksSolo(&indeksyangDitukarB)
					}
				}
				indeksPengecekan = indeksPengecekan + 1
			}
		}
		dominoTemp = dominoRandomed[indeksyangDitukarA]
		dominoRandomed[indeksyangDitukarA] = dominoRandomed[indeksyangDitukarB]
		dominoRandomed[indeksyangDitukarB] = dominoTemp
		jmlDominoTeracak = jmlDominoTeracak + 1
		RandomedIndex[jmlDominoTeracak] = indeksyangDitukarA
		jmlDominoTeracak = jmlDominoTeracak + 1
		RandomedIndex[jmlDominoTeracak] = indeksyangDitukarB
		det1x = 1
	}
	return dominoRandomed
}

//Fungsi yang dipanggil ketika pemain ingin mengambil domino dari boneyard berdasar urutan.
func milihDomino(pilihanPlayer *int) {
	/*	I.S Terdefinisi pilihanPlayer int
		F.S memilih pilihan domino antara 1 - 28
	*/
	var pilihanPlayerTemp int
	fmt.Println("\t\t\t\t\tDomino urutan berapa yang ingin diambil?")
	fmt.Print("\t\t\t\t\tDomino yang ingin saya ambil adalah ")
	fmt.Scan(&pilihanPlayerTemp)
	fmt.Println("")
	*pilihanPlayer = pilihanPlayerTemp
}

//Fungsi yang dipanggil untuk mengurutkan domino balak berdasar
func nyortirDominoBalak(dominoBalak [4]Domino) [4]Domino {
	/*	Mengembalikan dominoBalak secara ASCENDING dengan menggunakan metode Insertion Sort
	*/
	var balakTemp [4]Domino
	var domiTemp Domino
	var pass, indeksChecking int
	balakTemp = dominoBalak
	pass = 0
	for pass < 3 {
		indeksChecking = pass+1
		domiTemp = balakTemp[indeksChecking]
		for indeksChecking >= 1 && domiTemp.sisiKanan < balakTemp[indeksChecking-1].sisiKanan {
			balakTemp[indeksChecking] = balakTemp[indeksChecking-1]
			indeksChecking = indeksChecking - 1
		}
		balakTemp[indeksChecking] = domiTemp
		pass = pass + 1
	}
	return balakTemp
}

//Fungsi yang digunakan dan dipanggil untuk menentukan pemenang
func menentukanMenang(dominoPemain, dominoDealer [4]Domino, sedangBermain *Pemain) {
	/*	I.S Terdefinisi dominoPemain,dominoDealer,dan Player yang sedang bermain
		F.S menentukan Menang antara Pemain atau Dealer
	*/
	var jumlahTilesPemain, jumlahTilesDealer, indeksChecking, indeksBalakPemain, indeksBalakDealer int
	var detBalakPemain, detBalakDealer string
	var balakPemain, balakDealer [4]Domino
	detBalakPemain, detBalakDealer = "undecided", "undecided"
	jumlahTilesPemain = dominoPemain[0].sisiKanan+dominoPemain[0].sisiKiri+dominoPemain[1].sisiKanan+dominoPemain[1].sisiKiri+dominoPemain[2].sisiKanan+dominoPemain[2].sisiKiri+dominoPemain[3].sisiKanan+dominoPemain[3].sisiKiri
	jumlahTilesDealer = dominoDealer[0].sisiKanan+dominoDealer[0].sisiKiri+dominoDealer[1].sisiKanan+dominoDealer[1].sisiKiri+dominoDealer[2].sisiKanan+dominoDealer[2].sisiKiri+dominoDealer[3].sisiKanan+dominoDealer[3].sisiKiri
	for indeksChecking < 4 {
		if dominoPemain[indeksChecking].sisiKanan == dominoPemain[indeksChecking].sisiKiri {
			balakPemain[indeksBalakPemain] = dominoPemain[indeksChecking]
			indeksBalakPemain = indeksBalakPemain + 1
		}
		indeksChecking = indeksChecking + 1
	}
	if indeksBalakPemain >= 2 {
		detBalakPemain = "balak/double"
	}
	indeksChecking = 0
	for indeksChecking < 4 {
		if dominoDealer[indeksChecking].sisiKanan == dominoDealer[indeksChecking].sisiKiri {
			balakDealer[indeksBalakDealer] = dominoDealer[indeksChecking]
			indeksBalakDealer = indeksBalakDealer + 1
		}
		indeksChecking = indeksChecking + 1
	}
	if indeksBalakDealer >= 2 {
		detBalakDealer = "balak/double"
	}
	if detBalakPemain == "balak/double" && detBalakDealer == "balak/double" {
		balakPemain =  nyortirDominoBalak(balakPemain)
		balakDealer = nyortirDominoBalak(balakDealer)
		fmt.Println(" ")
		fmt.Println("\t\t\t\t\tDealer & Pemain memiliki 2 double/balak !")
		fmt.Println("\t\t\t\t\t[Mengadu jumlah 2 balak tertinggi...]")
		if balakPemain[3].sisiKanan+balakPemain[2].sisiKanan > balakDealer[3].sisiKiri+balakDealer[2].sisiKiri {
			fmt.Println("\t\t\t\t\tWoo! You've won.")
			sedangBermain.gameYangDimenangkan = sedangBermain.gameYangDimenangkan + 1
			sedangBermain.matchmakingPoints = sedangBermain.matchmakingPoints + 25
		} else {
			fmt.Println("\t\t\t\t\tYou've lost.")
			sedangBermain.matchmakingPoints = sedangBermain.matchmakingPoints - 25
		}
	} else if detBalakPemain == "balak/double" && detBalakDealer != "balak/double" {
		fmt.Println(" ")
		fmt.Println("\t\t\t\t\tAnda memiliki 2 balak/double!")
		fmt.Println("\t\t\t\t\tWoo! You've won.")
		sedangBermain.matchmakingPoints = sedangBermain.matchmakingPoints + 25
		sedangBermain.gameYangDimenangkan = sedangBermain.gameYangDimenangkan + 1
	} else if detBalakPemain != "balak/double" && detBalakDealer == "balak/double" {
		fmt.Println(" ")
		fmt.Println("\t\t\t\t\tDealer memiliki 2 balak/double!")
		fmt.Println("\t\t\t\t\tYou've lost.")
		sedangBermain.matchmakingPoints = sedangBermain.matchmakingPoints - 25
	} else if detBalakPemain != "balak/double" && detBalakDealer != "balak/double" {
		fmt.Println(" ")
		fmt.Println("\t\t\t\t\tPemain & Dealer tidak memiliki 2 double/balak")
		fmt.Println("\t\t\t\t\t[Mengadu jumlah skor tiles...]")
		fmt.Println("\t\t\t\t\tJumlah skor tiles anda   :", jumlahTilesPemain)
		fmt.Println("\t\t\t\t\tJumlah skor tiles dealer :", jumlahTilesDealer)
		if jumlahTilesPemain > jumlahTilesDealer {
			fmt.Println("\t\t\t\t\tWoo! You've won.")
			sedangBermain.gameYangDimenangkan = sedangBermain.gameYangDimenangkan + 1
			sedangBermain.matchmakingPoints = sedangBermain.matchmakingPoints + 25
		} else {
			fmt.Println("\t\t\t\t\tYou've lost.")
			sedangBermain.matchmakingPoints = sedangBermain.matchmakingPoints - 25
		}
	}
	sedangBermain.gameYangSudahDimainkan = sedangBermain.gameYangSudahDimainkan + 1
}

//Fungsi yang akan dipanggil saat pemain memilih option nomor 1 saat di main menu.
func optionPlayerBermain(sedangBermain *Pemain, dominoSaatIni *[28]Domino) {
	/*	I.S Terdefinisi sedangBermain yaitu orang yg Sedang Bermain dan dominoSaatIni yang berisi array Domino
		F.S dominoSaatIni teracak dan field-field pada tipe bentukan Pemain dari sedangBermain akan berubah, sesuai dengan apakah menang atau tidak.
	*/
	var dominoPemain, dominoDealer [4]Domino
	var dominoCurrent [28]Domino
	var dominoPilihanPemain, dominoPilihanDealer [4]int
	var pilihanPlayer, batasDomino, indeksPengecekan, indeksMenampilkan, indeksTilesDealer, pilihanDealer, indeksDealer int
	var decisionPemain string
	fmt.Println(" ")
	fmt.Println(" ")
	fmt.Println("\t\t\t\t\tShuffling the dominos...")
	dominoCurrent = acakDomino(*dominoSaatIni)
	dominoCurrent = acakDomino(dominoCurrent)
	fmt.Println("\t\t\t\t\tShuffled!")
	batasDomino = 0
	for batasDomino < 4 {
		milihDomino(&pilihanPlayer)
		for pilihanPlayer != 1 && pilihanPlayer != 2 && pilihanPlayer != 3 && pilihanPlayer != 4 && pilihanPlayer != 5 && pilihanPlayer != 6 && pilihanPlayer != 7 && pilihanPlayer != 8 && pilihanPlayer != 9 && pilihanPlayer != 10 && pilihanPlayer != 11 && pilihanPlayer != 12 && pilihanPlayer != 13 && pilihanPlayer != 14 && pilihanPlayer != 15 && pilihanPlayer != 16 && pilihanPlayer != 17 && pilihanPlayer != 18 && pilihanPlayer != 19 && pilihanPlayer != 20 && pilihanPlayer != 21 && pilihanPlayer != 22 && pilihanPlayer != 23 && pilihanPlayer != 24 && pilihanPlayer != 25 && pilihanPlayer != 26 && pilihanPlayer != 27 && pilihanPlayer != 28 {
			fmt.Println("\t\t\t\t\tInputan anda salah!")
			fmt.Println("\t\t\t\t\t========================================")
			milihDomino(&pilihanPlayer)
		}
		indeksPengecekan = 0
		for indeksPengecekan < batasDomino {
			if dominoPilihanPemain[indeksPengecekan] == pilihanPlayer {
				fmt.Println("\t\t\t\t\tDomino itu sudah anda ambil!")
				fmt.Println("\t\t\t\t\t========================================")
				milihDomino(&pilihanPlayer)
				for pilihanPlayer != 1 && pilihanPlayer != 2 && pilihanPlayer != 3 && pilihanPlayer != 4 && pilihanPlayer != 5 && pilihanPlayer != 6 && pilihanPlayer != 7 && pilihanPlayer != 8 && pilihanPlayer != 9 && pilihanPlayer != 10 && pilihanPlayer != 11 && pilihanPlayer != 12 && pilihanPlayer != 13 && pilihanPlayer != 14 && pilihanPlayer != 15 && pilihanPlayer != 16 && pilihanPlayer != 17 && pilihanPlayer != 18 && pilihanPlayer != 19 && pilihanPlayer != 20 && pilihanPlayer != 21 && pilihanPlayer != 22 && pilihanPlayer != 23 && pilihanPlayer != 24 && pilihanPlayer != 25 && pilihanPlayer != 26 && pilihanPlayer != 27 && pilihanPlayer != 28 {
					fmt.Println("\t\t\t\t\tInputan anda salah!")
					fmt.Println("\t\t\t\t\t========================================")
					milihDomino(&pilihanPlayer)
				}
				indeksPengecekan = -1
			}
			indeksPengecekan = indeksPengecekan + 1
		}
		clearScreen()
		fmt.Println(" ")
		fmt.Println("\t\t\t\t\tTiles berhasil ditambahkan!")
		dominoPemain[batasDomino] = dominoCurrent[pilihanPlayer-1]
		dominoPilihanPemain[batasDomino] = pilihanPlayer
		batasDomino = batasDomino + 1
		if batasDomino != 4 {
			fmt.Print("\t\t\t\t\tTiles anda saat ini : ")
			for indeksMenampilkan < batasDomino {
				fmt.Print("(", dominoPemain[indeksMenampilkan].sisiKiri, ",", dominoPemain[indeksMenampilkan].sisiKanan,") ")
				indeksMenampilkan = indeksMenampilkan + 1
			}
			fmt.Println(" ")
		}
		indeksMenampilkan = 0
		fmt.Println(" ")
	}
	fmt.Println("\t\t\t\t\tSettled!")
	fmt.Print("\t\t\t\t\tTiles anda : (", dominoPemain[0].sisiKiri, ",", dominoPemain[0].sisiKanan,") (", dominoPemain[1].sisiKiri, ",", dominoPemain[1].sisiKanan,") (", dominoPemain[2].sisiKiri, ",", dominoPemain[2].sisiKanan,") (", dominoPemain[3].sisiKiri, ",", dominoPemain[3].sisiKanan,")")
	fmt.Println("")
	fmt.Print("\t\t\t\t\tDecision : ")
	fmt.Scan(&decisionPemain)
	if sedangBermain.gameYangSudahDimainkan > 0 {
		for decisionPemain != "1" && decisionPemain != "2" && decisionPemain != "3" && decisionPemain != "4" && decisionPemain != "0" && decisionPemain != "9" {
			fmt.Println("\t\t\t\t\tDecision yang tersedia hanya 1, 2, 3, 4, 9, dan 0! (Selengkapnya baca 'Instruction' yang terdapat di mainmenu.)")
			fmt.Print("\t\t\t\t\tDecision : ")
			fmt.Scan(&decisionPemain)
		}
	} else {
		for decisionPemain != "1" && decisionPemain != "2" && decisionPemain != "3" && decisionPemain != "4" && decisionPemain != "0" {
			if decisionPemain == "9" {
				fmt.Println("\t\t\t\t\tDecision 9 hanya berlaku ketika anda setidaknya telah memainkan games 1x!")
			}
			fmt.Println("\t\t\t\t\tDecision yang tersedia untuk anda hanya 1, 2, 3, 4 dan 0! (Selengkapnya baca 'Instruction' yang terdapat di mainmenu.)")
			fmt.Print("\t\t\t\t\tDecision : ")
			fmt.Scan(&decisionPemain)
		}
	}
	if decisionPemain != "0" && decisionPemain != "9" {
		clearScreen()
		fmt.Println("")
		fmt.Println("\t\t\t\t\tTaking random tiles from the boneyard and replacing with your selected tiles...")
		opsiMenggantiTiles(decisionPemain, &dominoPilihanPemain, &dominoPemain, dominoCurrent)
		fmt.Print("\t\t\t\t\tTiles anda : (", dominoPemain[0].sisiKiri, ",", dominoPemain[0].sisiKanan,") (", dominoPemain[1].sisiKiri, ",", dominoPemain[1].sisiKanan,") (", dominoPemain[2].sisiKiri, ",", dominoPemain[2].sisiKanan,") (", dominoPemain[3].sisiKiri, ",", dominoPemain[3].sisiKanan,")")
		fmt.Println("")
		fmt.Print("\t\t\t\t\tDecision : ")
		fmt.Scan(&decisionPemain)
		if sedangBermain.gameYangSudahDimainkan > 0 {
			for decisionPemain != "1" && decisionPemain != "2" && decisionPemain != "3" && decisionPemain != "4" && decisionPemain != "0" && decisionPemain != "9" {
				fmt.Println("\t\t\t\t\tDecision yang tersedia hanya 1, 2, 3, 4, 9, dan 0! (Selengkapnya baca 'Instruction' yang terdapat di mainmenu.)")
				fmt.Print("\t\t\t\t\tDecision : ")
				fmt.Scan(&decisionPemain)
			}
		} else {
			for decisionPemain != "1" && decisionPemain != "2" && decisionPemain != "3" && decisionPemain != "4" && decisionPemain != "0" {
				if decisionPemain == "9" {
					fmt.Println("\t\t\t\t\tDecision 9 hanya berlaku ketika anda setidaknya telah memainkan games 1x!")
				}
				fmt.Println("\t\t\t\t\tDecision yang tersedia untuk anda hanya 1, 2, 3, 4 dan 0! (Selengkapnya baca 'Instruction' yang terdapat di mainmenu.)")
				fmt.Print("\t\t\t\t\tDecision : ")
				fmt.Scan(&decisionPemain)
			}
		}
		if decisionPemain != "0" && decisionPemain != "9" {
			clearScreen()
			fmt.Println("")
			fmt.Println("\t\t\t\t\tTaking random tiles from the boneyard and replacing with your selected tiles...")
			opsiMenggantiTiles(decisionPemain, &dominoPilihanPemain, &dominoPemain, dominoCurrent)
			fmt.Print("\t\t\t\t\tTiles anda : (", dominoPemain[0].sisiKiri, ",", dominoPemain[0].sisiKanan,") (", dominoPemain[1].sisiKiri, ",", dominoPemain[1].sisiKanan,") (", dominoPemain[2].sisiKiri, ",", dominoPemain[2].sisiKanan,") (", dominoPemain[3].sisiKiri, ",", dominoPemain[3].sisiKanan,")")
			fmt.Println("")
		}else if decisionPemain == "9"{
			fmt.Println("\t\t\t\t\tHello", sedangBermain.namaPemain, "!")
			fmt.Println("\t\t\t\t\tSejauh ini, ada", sedangBermain.gameYangDimenangkan, "games yang dimenangkan dari", sedangBermain.gameYangSudahDimainkan, "games.")
			if sedangBermain.gameYangSudahDimainkan != 0 {
				fmt.Print("\t\t\t\t\tWinrate kamu adalah ", 100*float64(sedangBermain.gameYangDimenangkan)/float64(sedangBermain.gameYangSudahDimainkan), "%\n")
			}
			fmt.Println("\t\t\t\t\tMMR", sedangBermain.namaPemain, "saat ini adalah", sedangBermain.matchmakingPoints)
		}
	} else if decisionPemain == "9"{
		fmt.Println("\t\t\t\t\tHello", sedangBermain.namaPemain, "!")
		fmt.Println("\t\t\t\t\tSejauh ini, ada", sedangBermain.gameYangDimenangkan, "games yang dimenangkan dari", sedangBermain.gameYangSudahDimainkan, "games.")
		if sedangBermain.gameYangSudahDimainkan != 0 {
			fmt.Print("\t\t\t\t\tWinrate kamu adalah ", 100*float64(sedangBermain.gameYangDimenangkan)/float64(sedangBermain.gameYangSudahDimainkan), "%\n")
			fmt.Println("\t\t\t\t\tMMR", sedangBermain.namaPemain, "saat ini adalah", sedangBermain.matchmakingPoints)
		}
	}
	if decisionPemain != "9" {
		for pilihanDealer < 4 {
			indeksTilesDealer = rand.Intn(28)
			indeksDealer = 0
			for indeksDealer < pilihanDealer {
				if dominoPilihanDealer[indeksDealer] == indeksTilesDealer+1 {
					indeksTilesDealer = rand.Intn(28)
					indeksDealer = -1
				}
				indeksDealer = indeksDealer + 1
			}
			if indeksTilesDealer+1 != dominoPilihanPemain[0] && indeksTilesDealer+1 != dominoPilihanPemain[1] && indeksTilesDealer+1 != dominoPilihanPemain[2] && indeksTilesDealer+1 != dominoPilihanPemain[3] {
				dominoDealer[pilihanDealer] = dominoCurrent[indeksTilesDealer]
				dominoPilihanDealer[pilihanDealer] = indeksTilesDealer+1
				pilihanDealer = pilihanDealer + 1
			}
		}
		fmt.Print("\t\t\t\t\tTiles dealer : (", dominoDealer[0].sisiKiri, ",", dominoDealer[0].sisiKanan,") (", dominoDealer[1].sisiKiri, ",", dominoDealer[1].sisiKanan,") (", dominoDealer[2].sisiKiri, ",", dominoDealer[2].sisiKanan,") (", dominoDealer[3].sisiKiri, ",", dominoDealer[3].sisiKanan,")")
		fmt.Println("")
		var playingParams Pemain
		playingParams = *sedangBermain
		menentukanMenang(dominoPemain, dominoDealer, &playingParams)
		*sedangBermain = playingParams
		fmt.Println("\t\t\t\t\tHello", sedangBermain.namaPemain, "!")
		fmt.Println("\t\t\t\t\tSejauh ini, ada", sedangBermain.gameYangDimenangkan, "games yang dimenangkan dari", sedangBermain.gameYangSudahDimainkan, "games.")
		fmt.Println("\t\t\t\t\tMMR", sedangBermain.namaPemain, "saat ini adalah", sedangBermain.matchmakingPoints)
		fmt.Print("\t\t\t\t\tWinrate kamu adalah ", 100*float64(sedangBermain.gameYangDimenangkan)/float64(sedangBermain.gameYangSudahDimainkan), "%\n")
	}
}

func opsiMenggantiTiles(tilesyangInginDiganti string, urutatilesPemain *[4]int, tilesPemain *[4]Domino, tilesRandom [28]Domino) {
	/*	I.S Terdefinisi array tilesyangInginDiganti,urutanTilesPemain,tilesPemain,dan array TilesRandom
		F.S Mengganti tilesPemain dengan tilesRandom yang belum diambil berdasarkan tilesyangInginDiganti
	*/
	var indeksPengganti, tilesyangInginDigantiTemp int
	var detPengganti string
	var tilesTemp [4]Domino
	var urutanTilesTemp [4]int
	tilesTemp = *tilesPemain
	urutanTilesTemp = *urutatilesPemain
	if tilesyangInginDiganti == "1" {
		tilesyangInginDigantiTemp = 1
	} else if tilesyangInginDiganti == "2" {
		tilesyangInginDigantiTemp = 2
	} else if tilesyangInginDiganti == "3" {
		tilesyangInginDigantiTemp = 3
	} else if tilesyangInginDiganti == "4" {
		tilesyangInginDigantiTemp = 4
	}
	indeksPengganti = rand.Intn(28)
	for indeksPengganti < 28 && detPengganti != "SuksesTerganti" {
		if urutanTilesTemp[0] != indeksPengganti+1 && urutanTilesTemp[1] != indeksPengganti+1 && urutanTilesTemp[2] != indeksPengganti+1 && urutanTilesTemp[3] != indeksPengganti+1 {
			tilesTemp[tilesyangInginDigantiTemp-1] = tilesRandom[indeksPengganti]
			urutanTilesTemp[tilesyangInginDigantiTemp-1] = indeksPengganti+1
			detPengganti = "SuksesTerganti"
		}
		indeksPengganti = rand.Intn(28)
	}
	*tilesPemain = tilesTemp
	*urutatilesPemain = urutanTilesTemp
}

//Fungsi untuk menampilkan papan skor.
func scoreBoard(listPemain himpunanPemain, totalPemain int, pilPemain int) {
	clearScreen()
	fmt.Println("\t\t\t\t\t|----------------------------------------------|")
	fmt.Println("\t\t\t\t\t[\tScoreboard\t\t Pemain\t       |")
	nyortirRankingPemain(listPemain, totalPemain, pilPemain)
	fmt.Println("")
}

func nyortirRankingPemain(listPemain himpunanPemain, totalPemain int, pilPemain int) {
	/*	I.S Terdefinisi array listPemain yang memiliki totalPemain bernilai bilangan bulat,dan pilPemain
		F.S array listPemain winrate akan terurut DESCENDING dengan menggunakan algoritma Insertion Sort apabila pilPemain 1
			array listPemain MMR akan terurut DESCENDING dengan menggunakan algoritma Selection Sort apabila pilPemain 2
	*/
	var IndeksPemainn, indeksChecking, pass, indeksHighest int
	var pemainTampungan Pemain
	var listPemainTemp himpunanPemain
	listPemainTemp = listPemain
	//Descending via Insertion Sort
	pass = 0
	if totalPemain != 0 {
		if pilPemain == 1 {
			for pass < totalPemain-1 {
				indeksChecking = pass+1
				pemainTampungan = listPemainTemp[indeksChecking]
				for indeksChecking >= 1 && pemainTampungan.gameYangSudahDimainkan != 0 &&float64(pemainTampungan.gameYangDimenangkan)/float64(pemainTampungan.gameYangSudahDimainkan) > float64(listPemainTemp[indeksChecking-1].gameYangDimenangkan)/float64(listPemainTemp[indeksChecking-1].gameYangSudahDimainkan) {
					listPemainTemp[indeksChecking] = listPemainTemp[indeksChecking-1]
					indeksChecking = indeksChecking - 1
				}
				listPemainTemp[indeksChecking] = pemainTampungan
				pass = pass+1
			}
		} else if pilPemain == 2 {
			for pass < totalPemain-1 {
				indeksHighest = pass
				indeksChecking = pass+1
				for indeksChecking < totalPemain {
					if listPemainTemp[indeksChecking].matchmakingPoints > listPemainTemp[indeksHighest].matchmakingPoints {
						indeksHighest = indeksChecking
					}
					indeksChecking = indeksChecking + 1
				}
				pemainTampungan = listPemainTemp[pass]
				listPemainTemp[pass] = listPemainTemp[indeksHighest]
				listPemainTemp[indeksHighest] = pemainTampungan
				pass = pass+1
			}
		}
		IndeksPemainn = 0
		fmt.Println("\t\t\t\t\t|----------------------------------------------|")
		fmt.Println("\t\t\t\t\t|No\tName\t\tWinrate\t\tMMR    |")
		fmt.Println("\t\t\t\t\t|----------------------------------------------|")
		for IndeksPemainn < totalPemain {
			if listPemainTemp[IndeksPemainn].gameYangSudahDimainkan != 0 {
				fmt.Print("\t\t\t\t\t|", IndeksPemainn+1, "| \t", listPemainTemp[IndeksPemainn].namaPemain, "\t    ")
				fmt.Printf("%.2f", 100*float64(listPemainTemp[IndeksPemainn].gameYangDimenangkan)/float64(listPemainTemp[IndeksPemainn].gameYangSudahDimainkan))
				fmt.Print("%\t\t", listPemainTemp[IndeksPemainn].matchmakingPoints,"\n")
			} else {
				fmt.Print("\t\t\t\t\t|", IndeksPemainn+1, "| \t", listPemainTemp[IndeksPemainn].namaPemain, "\t\t", "-\t\t", listPemainTemp[IndeksPemainn].matchmakingPoints,"\n")
			}
			IndeksPemainn = IndeksPemainn + 1
		}
		fmt.Println("\t\t\t\t\t|----------------------------------------------|")
	} else {
		fmt.Println("\t\t\t\t\t Scoreboard kosong ")
		fmt.Println("\t\t\t\t\t Belum ada yang bermain ")
	}
}


func back(keputusanPemain *string) {
	var decisionKey string
	fmt.Println("\t\t\t\t\tBack? [B]")
	fmt.Print("\t\t\t\t\tDecision: ")
	fmt.Scan(&decisionKey)
	for decisionKey != "b" && decisionKey != "B" {
		fmt.Println("\t\t\t\t\tInputan hanya menerima B!")
		fmt.Print("\t\t\t\t\tDecision: ")
		fmt.Scan(&decisionKey)
	}
	*keputusanPemain = decisionKey
}

func main() {
	clearScreen()
	rand.Seed(int64(time.Now().Nanosecond()+time.Now().Second()+time.Now().Minute()*60+time.Now().Hour()*3600))
	// var buat domino
	var dominoCurrent [28]Domino
	// variabel-variabel buat pemain
	var sedangBermain, mendaftar Pemain
	var keberhasilanLogin, statusNamaSama bool
	var listPemain himpunanPemain
	var jumlahygTerdaftar, nomerPemainCurrent, opsiScoreboard int
	var opsimainLagi, keputusanPemain string
	// var untuk pilihan di beberapa section di main menu
	var option, optionLoginRegistScreen int
	createDominoSets(&dominoCurrent)

	jumlahygTerdaftar = 0
	for option != 5 {
		keberhasilanLogin = false
		mainScreen()
		fmt.Scan(&option)
		for option < 1 && option > 5 {
			fmt.Println("\t\t\t\t\tInputan hanya ada 1,2,3,dan 4 !")
			fmt.Print("\t\t\t\t\tPilihan :")
			fmt.Scan(&option)
		}
		clearScreen()
		if option == 1 {
			loginRegistScreen(&optionLoginRegistScreen)
			if optionLoginRegistScreen == 1 {
				clearScreen()
				login(&sedangBermain, &nomerPemainCurrent, jumlahygTerdaftar, listPemain, &keberhasilanLogin)
			} else if optionLoginRegistScreen == 2 {
				clearScreen()
				register(&jumlahygTerdaftar, listPemain, &statusNamaSama, &mendaftar)
				if !statusNamaSama {
					listPemain[jumlahygTerdaftar-1] = mendaftar
				}
				sedangBermain = mendaftar
				nomerPemainCurrent = sedangBermain.nomerPemain
			}
			if optionLoginRegistScreen == 2 || keberhasilanLogin {
				if optionLoginRegistScreen == 2 && !statusNamaSama {
					fmt.Print("\t\t\t\t\t", sedangBermain.namaPemain,"? nama yang keren!\n")
					fmt.Println("\t\t\t\t\tSelamat bermain, goodluck havefun!")
				}
				if keberhasilanLogin || statusNamaSama {
					fmt.Println("\t\t\t\t\tWelcome back", sedangBermain.namaPemain, "!")
				}
				fmt.Println("\t\t\t\t\tLanjut? [Y]")
				fmt.Print("\t\t\t\t\tDecision : ")
				fmt.Scan(&opsimainLagi)
				for opsimainLagi != "Y" && opsimainLagi != "y" {
					fmt.Println("\t\t\t\t\tInputan hanya menerima 'Y' atau 'y'!")
					fmt.Print("\t\t\t\t\tDecision : ")
					fmt.Scan(&opsimainLagi)
				}
				for opsimainLagi == "Y" || opsimainLagi == "y" {
					loadingScreen()
					optionPlayerBermain(&sedangBermain, &dominoCurrent)
					listPemain[nomerPemainCurrent-1] = sedangBermain
					mainLagi(&opsimainLagi)
					clearScreen()
				}
			}
		} else if option == 2 {
			for keputusanPemain == "" {
				clearScreen()
				fmt.Println("\t\t\t\t\tSort by...")
				fmt.Println("\t\t\t\t\t[1] WinRate")
				fmt.Println("\t\t\t\t\t[2] Matchmaking Ratio")
				fmt.Print("\t\t\t\t\tDecision : ")
				fmt.Scan(&opsiScoreboard)
				for opsiScoreboard != 1 && opsiScoreboard != 2 {
					fmt.Println("\t\t\t\t\tInputan yang tersedia hanya 1 dan 2!")
					fmt.Print("\t\t\t\t\tDecision : ")
					fmt.Scan(&opsiScoreboard)
				}
				scoreBoard(listPemain, jumlahygTerdaftar, opsiScoreboard)
				back(&keputusanPemain)
			}
		} else if option == 3 {
			for keputusanPemain == "" {
				clearScreen()
				instructionScreen(&keputusanPemain)
			}
		} else if option == 4 {
			clearScreen()
			userStat(listPemain, jumlahygTerdaftar)
			back(&keputusanPemain)
		}
		clearScreen()
		keputusanPemain = ""
		opsiScoreboard = 0
	}
}

func mainLagi(opsiMainLagi *string) {
	/*	I.S Terdefinisi opsiMainLagi string
		F.S mengembalikan opsiMainlagi
	*/
	var opsiPlayingMore string
	fmt.Print("\t\t\t\t\tApakah anda ingin bermain lagi? [Y/N] ")
	fmt.Scan(&opsiPlayingMore)
	for opsiPlayingMore != "Y" && opsiPlayingMore != "y" && opsiPlayingMore != "N" && opsiPlayingMore != "n" {
		fmt.Println("\t\t\t\t\tInputan hanya menerima Y atau N!")
		fmt.Print("\t\t\t\t\tApakah anda ingin bermain lagi? [Y/N] ")
		fmt.Scan(&opsiPlayingMore)
	}
	*opsiMainLagi = opsiPlayingMore
}

func instructionScreen(keputusanPemain *string) {
	/*	I.S Terdefinisi keputusanPemain string
		F.S Menampilkan instruksi
	*/
	var decisionKey string
	fmt.Println("\t\t\t\t\t[Tata Cara Bermain]")
	fmt.Println("\t\t\t\t\t[1] Registrasi menggunakan nama & tidak berspasi (perhatikan besar kecilnya huruf)")
	fmt.Println("\t\t\t\t\t[2] Log-in dengan menggunakan nama yang anda daftarkan.")
	fmt.Println("\t\t\t\t\t[3] Registrasi dengan menggunakan username lama anda akan berakibat anda dialihkan")
	fmt.Println("\t\t\t\t\t    ke akun lama anda.")
	fmt.Println("\t\t\t\t\t\t\tHalaman [1]")
	fmt.Println("\t\t\t\t\tBack/Continue? [B/C]")
	fmt.Print("\t\t\t\t\tDecision: ")
	fmt.Scan(&decisionKey)
	for decisionKey != "b" && decisionKey != "B" && decisionKey != "c" && decisionKey != "C" {
    	fmt.Println("\t\t\t\t\tInputan hanya menerima B atau C!")
		fmt.Print("\t\t\t\t\tDecision: ")
		fmt.Scan(&decisionKey)
	}
	if decisionKey == "c" || decisionKey == "C" {
		clearScreen()
		fmt.Println("\t\t\t\t\t[Ketentuan Menang]")
		fmt.Println("\t\t\t\t\t[1] Memiliki dua tile dengan nilai kedua sisi sama (balak/double) dan lawan tidak memiliki dua tile dengan nilai kedua sisi yang sama.")
		fmt.Println("\t\t\t\t\t[2] Jika lawan memiliki balak juga, akan diadu total nilai tiles 2 balak tertinggi.")
		fmt.Println("\t\t\t\t\t[3] Memiliki nilai tile lebih tinggi dari lawan, jika memiliki nilai sama atau kurang dari, maka player akan kalah.")
		fmt.Println("\t\t\t\t\t\t\tHalaman [2]")
		fmt.Println("\t\t\t\t\tBack/Continue? [B/C]")
		fmt.Print("\t\t\t\t\tDecision : ")
		fmt.Scan(&decisionKey)
		for decisionKey != "b" && decisionKey != "B" && decisionKey != "c" && decisionKey != "C" {
	    	fmt.Println("\t\t\t\t\tInputan hanya menerima B dan C!")
			fmt.Print("\t\t\t\t\tDecision: ")
			fmt.Scan(&decisionKey)
		}
		if decisionKey == "C" || decisionKey == "c" {
			clearScreen()
			fmt.Println("\t\t\t\t\t[Matchmaking Ratio/MMR]")
			fmt.Println("\t\t\t\t\t[1] Untuk memperadil perankingan player, maka dibuat Matchmaking Ratio.")
			fmt.Println("\t\t\t\t\t[2] 1x menang berarti MMR player bertambah sebanyak 25, dan sebaliknya -25 jika kalah.")
			fmt.Println("\t\t\t\t\t[3] Dengan begitu, player lama akan lebih diuntungkan dibanding player dengan jumlah games 1x namun")
			fmt.Println("\t\t\t\t\t    memiliki WinRate sebesar 100%")
			fmt.Println("\t\t\t\t\t\t\tHalaman [3]")
			fmt.Println("\t\t\t\t\tBack/Continue? [B/C]")
			fmt.Print("\t\t\t\t\tDecision : ")
			fmt.Scan(&decisionKey)
			for decisionKey != "b" && decisionKey != "B" && decisionKey != "c" && decisionKey != "C" {
		    	fmt.Println("\t\t\t\t\tInputan hanya menerima B dan C!")
				fmt.Print("\t\t\t\t\tDecision: ")
				fmt.Scan(&decisionKey)
			}
			if decisionKey == "C" || decisionKey == "c" {
				clearScreen()
				fmt.Println("\t\t\t\t\t[Decision]")
				fmt.Println("\t\t\t\t\t[1] Saat anda telah mengambil 4 domino, anda diperkenankan untuk membuat keputusan.")
				fmt.Println("\t\t\t\t\t[2] Decision '1' berarti anda mengganti tiles ke-1 anda dengan tiles yang lain yang ada")
				fmt.Println("\t\t\t\t\t    di boneyard.")
				fmt.Println("\t\t\t\t\t[3] Begitu juga dengan decision '2', '3', dan '4'.")
				fmt.Println("\t\t\t\t\t[4] Decision '9' berarti anda ingin memberhentikan game.")
				fmt.Println("\t\t\t\t\t    dan games yang sedang anda mainkan tidak akan dihitung.")
				fmt.Println("\t\t\t\t\t    decision '9' hanya bisa digunakan saat anda setidaknya sudah bermain 1x.")
				fmt.Println("\t\t\t\t\t[5] Decision '0' berarti anda sudah puas dengan domino anda, dan langsung akan diadu.")
				fmt.Println("\t\t\t\t\t[6] Pemain hanya bisa melakukan decision mengganti tiles sebanyak 2x / 1 game.")
				fmt.Println("\t\t\t\t\t\t\tHalaman [4]")
				fmt.Println("\t\t\t\t\tBack? [B]")
				fmt.Print("\t\t\t\t\tDecision : ")
				fmt.Scan(&decisionKey)
				for decisionKey != "b" && decisionKey != "B" {
		    		fmt.Println("\t\t\t\t\tInputan hanya menerima B!")
					fmt.Print("\t\t\t\t\tDecision: ")
					fmt.Scan(&decisionKey)
				}
			}
		}
	}
	*keputusanPemain = decisionKey
}

func userStat(listPemain himpunanPemain, totalPemain int) {
	/*	I.S Terdefinisi array listPemain yang memiliki totalPemain bernilai bilangan bulat
		F.S Menampilkan listPemain yang terurut secara ASCENDING
	*/
	var idxPemain,idxUser int
	var key string
	idxPemain = 0
	clearScreen()
	fmt.Println("\t\t\t\t\t|-------- List Pemain ----------|")
	fmt.Println("\t\t\t\t\t|No\tName\t\t\t|")
	fmt.Println("\t\t\t\t\t|-------------------------------|")
	sortUser(&listPemain,totalPemain)
	if totalPemain != 0 {
		for idxPemain < totalPemain {
			fmt.Print("\t\t\t\t\t|",idxPemain+1,"\t",listPemain[idxPemain].namaPemain,"\n")
			idxPemain+=1
		}
		fmt.Println("\t\t\t\t\t|-------------------------------|")
		fmt.Println("")
		fmt.Print("\t\t\t\t\tMasukkan User yang ingin dilihat statusnya : ")
		fmt.Scan(&key)
		idxUser=searchUser(listPemain,totalPemain,key)
		if idxUser != -1 {
			fmt.Println("\t\t\t\t\t----- Stat ------------------------------")
			fmt.Println("\t\t\t\t\tNo User    :",listPemain[idxUser].nomerPemain)
			fmt.Println("\t\t\t\t\tName       :",listPemain[idxUser].namaPemain)
			fmt.Println("\t\t\t\t\tWin/Lose   :",listPemain[idxUser].gameYangDimenangkan,"/",listPemain[idxUser].gameYangSudahDimainkan-listPemain[idxUser].gameYangDimenangkan)
			fmt.Println("\t\t\t\t\tWin Rate   :",(float64(listPemain[idxUser].gameYangDimenangkan)/float64(listPemain[idxUser].gameYangSudahDimainkan))*100,"%")
			fmt.Println("\t\t\t\t\tMMR        :",listPemain[idxUser].matchmakingPoints)
			fmt.Println("\t\t\t\t\tTotal Game :",listPemain[idxUser].gameYangSudahDimainkan)
			fmt.Println("\t\t\t\t\t-----------------------------------------")
			}else{
				fmt.Println("\t\t\t\t\tUser tidak ditemukan.")
			}
	}else{
		fmt.Println("\t\t\t\t\t|       List Pemain Kosong      |")
		fmt.Println("\t\t\t\t\t|     Belum ada yang bermain    |")
		fmt.Println("\t\t\t\t\t|-------------------------------|")
	}
}

//Searching user w/ binary search
func searchUser(listPemain himpunanPemain,totalPemain int,key string) int{
	/*	Mengembalikan indeks hasil pencarian key di dalam array listPemain atau -1 apabila tidak ditemukan
		listPemain terurut Ascending
	*/
	var topData,botData,midData,found int
	found = -1
	topData = totalPemain
	midData = 0
	botData = 0
	for botData <= topData && found == -1 {
		midData = (botData + topData)/2
		if key < listPemain[midData].namaPemain{
			topData=midData-1
		}else if key > listPemain[midData].namaPemain{
			botData=midData+1
		}else{
			found = midData
		}
	}
	return found
}

//Sorting User w/ Selection Sort Ascending
func sortUser(listPemain *himpunanPemain,totalPemain int){
	/*	I.S Terdefinisi array listPemain yang memiliki totalPemain bernilai bilangan bulat
		F.S array listPemain namaPemain terurut ASCENDING dengan menggunakan algoritma Selection Sort
	*/
	var pass,idxMax,i int
	var temp Pemain
	pass = 0
	for pass < totalPemain {
		idxMax=pass
		i=pass+1
		for i < totalPemain {
			if listPemain[idxMax].namaPemain > listPemain[i].namaPemain{
				idxMax=i
			}
			i+=1
		}
		temp=listPemain[pass]
		listPemain[pass]=listPemain[idxMax]
		listPemain[idxMax]=temp
		pass+=1
	}
}
