package main
import (
	"fmt"
	"sort"
	"strings"
)

func wordSubsets(A []string, B []string) []string {
	uniqueB := []string{}
	bMap := map[string]int{}
	sort.Slice(B, func (a, b int) bool { return len(B[a]) > len(B[b]) })
TOP:
	for _, b := range B {
		for key, _ := range bMap {
			// fmt.Printf("key = %v, b = %v\n", key, b)
			if strings.Contains(key, b) {
				continue TOP
			}
		}
		if bMap[b] == 0 {
			uniqueB = append(uniqueB, b)
			bMap[b] = 1
		}
	}

	for _, b := range uniqueB {
		i := int(0)
		for i < len(A) {
			a := A[i]
			// fmt.Printf("a = %v, b = %v\n", sortA, sortB)
			if !contains(b, a) {
				A = append(A[:i], A[i + 1:]...)
				// fmt.Printf("current A %v\n", A)
				continue
			}
			i++
		}
	}
	return A
}

func contains(target string, search string) bool {
	for _, r := range target {
		c := string(r)
		index := strings.Index(search, c)
		if index == - 1 {
			return false
		}
		search = search[:index] + search[index + 1:]
	}
	return true
}

func main() {
	// A := []string{"amazon", "apple", "facebook", "google", "leetcode"}
	// B := []string{"e", "o"}

	// A := []string{"amazon", "apple", "facebook", "google", "leetcode"}
	// B := []string{"e", "oo"}

	// A := []string{"amazon", "apple", "facebook", "google", "leetcode"}
	// B := []string{"lo", "eo"}

	// A := []string{"amazon", "apple", "facebook", "google", "leetcode"}
	// B := []string{"ec", "oc", "ceo"}

	// A := []string{"acaaa", "cbabb", "ababc", "aabbb", "cbbbc"}
	// B := []string{"ba", "bb", "cc"}

	A := []string{"rmiunsaykn","msnnukvysi","nwusmkinyu","myinnpsukg","unmknnykis","kjnyisnumy","yrnwkiusnm","usmpnkriny","mnkswuinly","nymaunkiis","inuymrhkns","dwpnsrsxbr","lfkwdffvwe","kncyismqnu","bagyjgaxjg","nnyoysmuki","nzudsynkmi","minspnuwky","mnunaytisk","yniknimfus","ymnkssyuni","skunyoikmn","mynumkpnsi","yusgmnunik","rwgnfiiddc","ndumskynri","gnkamsuyni","zismuyknmn","swfuhdmqay","nwpngrantp","krsnmiyunp","umyslykinn","lsuniknymi","nnjuymkisg","nmkeusnbyi","uysmknlion","unmzikyvns","kbmemvylrx","ismkunniiy","yuisnemnkw","iumsnykwnz","mukynnqids","rikmnuksyn","kninupkysm","fmjikyusnn","inykdmslnu","ujinynvksm","yukishninm","gmzlhrwllq","ikmymusnnk","konmyjuins","mhnyisukyn","isgncmkuyn","mimkunfysn","mugnnksyei","maykuswnin","utvnmykins","ynksduiman","sikytdnmnu","nkngiusmys","iutnzsnymk","mingunyisk","inkkinusmy","myuiksndnu","ymvgksunin","kndmuiycsn","nnmwiysufk","uybknnsimc","nnmiusyknm","ovlwaghfnl","nmniuktays","nliknlmusy","spnmyiuunk","mksnnyiwua","fkisunycmn","nnynmiukds","matkmeaniw","munfnxyiks","ypnknimeus","ukknysnmiw","skmsyninmu","ckypinmuns","cyjjdjrvib","dkimqysnnu","myeskfnuin","mgxwetvick","ukmyinnysu","kiksnnyumq","nyiubskmnz","zykiusxmnn","mnsiukyaon","mnkunsnuiy","nukaynkmis","nkncsyiumb","nmsuinyvnk","numksymniy","usmnoitkny","nsmyukijdn","jhjxynlodg","snkuknimye","umnsvgkyni","yukusmnnli","fukyiisnnm","mnnouyiksk","kvnnyifusm","szdepanvbf","yxqumnfxqy","aewcvndari","kxsyinnmvu","rsmnuknyif","nrsyimpkun","lkxobjkjcf","kdnniysumn","ynasnfiukm","ggpltusdpa","ytapccrlic","mhuyshinnk","lkmsnicuyn","vmiunyukns","yskminnmsu","nyitmksenu","ynimksnusi","hnunmskiye","bksfpcuhwj","sknrtnmiyu","suymnincmk","kynvsigumn","mdcrtdcqyq","noykmufnsi","uiskfnnmjy","kniruysnmp","iumskynnfn","zsiymnnknu","kmylnsniyu","umgionysnk","knnwmyiuns","inuuynmsnk","nikpooglay","kfcnbiiwaq","mulbsnkiyn","yusknnmrci","noymnszuik","yuimsnnukr","nwuybmskni","mungkisynn","ysmubqnink","ymziusmknn","nkduyimisn","mkyiausnxn","emyqnsniuk","umkkysitnn","txucqmqvkh","nnulsyuimk","rpkjbottnv","mkiuknpnys","ismnvuknye","ineksuimny","smsunykgni","msgknycuin","qyvalprhet","vimusnyfnk"}
	B := []string{"si","nnyi","nink","y","kui","nnuiykm","mynniku","u","sknyium","nni","mi","ukmniy","uims","uysnn","ny","m","nnmy","myuksin","ysin","nnym","kmi","iuks","sy","nmnus","ikuymn","nnmy","k","iks","mys","nykins","m","n","nk","knunm","smyun","suy","yinu","inmkny","nk","ninsuyk","sm","nsnm","inmnk","mu","numsn","ys","inysum","ki","msy","ms","n","inyuk","imuyskn","syknnm","s","nnikmyu","i","siuk","mksyinu","nsk","i","innsy","sim","knin","smn","mkn","n","ks","yusnn","ikym","nuskinm","i","niu","snyukim","ynksi","iknsynm","msnn","n","ni","i","iny","ns","y","knusiy","msinnk","unisymk","knsin","m","m","uikmnny","ysn","uksnnmy","kmniusy","mny","kinym","sk","i","n","nmsuk","nmnsky","nsiykum","nim","yknmisu","knu","smnui","synmn","nkimun","simunk","nskmni","sm","i","ninkus","kmi","mkuynin","i","imnn","nnumkiy","ys","n","n","kmysniu","uki","i","in","in","in","uninmys","ksinyum","kuinsmy","mky","yn","uynk","m","nuysk","u","synnkiu","ynmsn","uim","nink","nikyn","nnm","n","nnyms","nu","s","umnn","yknum","usimynk","y","unni","nusyi","y","ukm","ksni","nunys","nsyknim","nsk","unsny","yimnsk","nn","kmsnuyi","munin","nimkyn","kunmysn","kn","unkym","ksi","iknu","iknnus","ksnuiym","simu","u","yn","msui","nn","ikynums","n","yinsuk","kmsyn","kn","ns","n","iy","yinmnu","nyim","sik","n","unmskni","uinkys","uknsiym","m","uiknn","ninkmsy","ny","musyinn","mun","sn","suknn","n","n","nynm","nunmk","um","k","mknn","m","ny","imnyusn","mni","kym","sm","miknsu","imnsuny","umnkiys","nsyk","nsikym","nuymsi","k","ysinku","k","inymkn","u","iysum","ksm","ns","ynsiknu","smnu","kn","ysm","ninkuys","mnukys","sm","yus","mnkyi","n","minns","nsnmu","n","ymksn","isnnku","k","uminy","munk","nki","mnny","s","kn","suyminn","kmnyni","yums","nnm","nsmkiyn","mkys","ikm","uminy","i","nsmuny","k","ykui","unmnk","knsunmy","ysunkmi","ksyiunm","uyik","nsmy","mun","sm","uisn","m","nyknu","myinsk","ynki","nik","ksn","u","unym","nsi","mni","nmnkiuy","uy","mk","kims","mnkyus","sn","inynk","nuy","nkus","synu","iumnkn","n","nmniysu","s","n","s","iuskynn","ym","nynkium","yumink","sinmn","ksm","y","snymiuk","nnm","nusk","n","usinm","smn","kny","m","y","yn","syknui","nsnu","uksniyn","yniusnm","msnn","iym","nnsmky","smnnkiu","knmi","uinksn","ukns","nuiysn","suky","mksin","isn","kuymnis","kmisn","sy","niusmyk","nmsnu","nimyuks","nsm","ksumn","muksin","um","mnnikys","suk","knymin","kins","unn","nsmkiy","ysnn","nnkuyi","nkynsm","isn","nskinm","smnyu","m","ys","ukn","snikm","n","ykis","yk","m","isnkuym","knsimyu","m","ksinyn","yuski","yum","nsmi","nis","i","k","smunk","m","kyusm","n","iy","ink","msnknui","s","kn","mkni","n","ysnnuk","nkns","knyu","nmuyik","kysiunn","nkusni","skiynmn","myisknu","snkn","yism","knus","mknny","ymks","n","knyu","miknys","nuys","ynuknis","nynsmu","nns","nnsmky","imunyn","inuky","nsiym","nuinmk","mnu","umk","k","u","mknin","mk","s","skn","nkm","kiyms","ksm","nkuym","k","msiyukn","knuyim","nk","ik","k","isym","my","ymkisn","nui","ns","nkuns","snknmy","uy","nksyun","iymuskn","ki","iumny","isnmk","sm","niksuy","isnkm","nk","nky","yiu","umksnin","kis","imyk","inmynku","kmnsuyi","ns","kni","nnsumy","nki","m","ny","k","nkm","n","k","m","n","nsin","munykns","mukisy","k","nnuk","iyskn","yniksmn","kuy","mkiuny","nsnkumy","y","y","ink","nkyumi","insmk","nkuin","iskn","inu","ukyims","uynknms","nnu","kiunmny","mi","imnysun","mnsiny","isnmuy","nymsn","uyis","yisnn","nyksi","synku","nmksn","sniymu","ninmys","usnkim","nn","kunys","ykumni","siknumy","nnm","mnuiys","iynnus","yinm","ui","nnkuiym","yn","si","iymuksn","sunnkym","knimnys","nys","mynn","nins","y","ms","kmiyn","knsm","ynm","nsmny","n","yns","sinyn","sukimnn","mnyks","msknuyi","mukn","nisy","in","n","n","nmniu","snykim","sniumy","iynnm","mikynun","uyik","nyk","nkmiu","umyis","nys","knysiu","nmny","sn","nk","nkiumyn","umsn","sumk","us","iunm","sykm","n","s","skm","mnk","iy","nkyumn","nuims","niukn","s","y","mknynis","kyn","n","niknm","unsykn","ny","m","munkisy","nsyink","s","m","synikum","muiknsy","nk","y","mysiuk","nn","nk","niy","nu","munn","suinky","inyum","usyin","u","sy","ius","unsni","iyknsu","kyins","sn","nks","umsyin","unnm","yn","yninm","n","nsymunk","nunyi","nuniys","ski","misy","nmnsiu","skuni","mn","kinsnu","iymusn","minyn","iymkns","nmsuiny","suyn","inkyns","yinksum","uknim","yisk","iuskny","nunmky","nu","nmyukis","uim","nimskun","nsinky","yumk","kmni","nykni","inmnkys","ksiumny","mu","ynku","yknni","nyknm","msn","ims","nysm","syn","ysunn","snu","ukni","iskn","imnks","nknsiuy","m","smnu","unysmi","uysnmki","mu","syukmni","usnk","ks","knsuyni","im","nkisy","ykms","yks","ns","u","uysm","nskmyun","uniy","nyiu","ny","nm","u","mny","sm","inksnm","nusnkym","skniym","y","ynukin","nik","unk","mkin","inmukns","ns","m","nkyui","s","uknnysm","ninum","mnykun","kmnsn","usnikn","u","nsyk","ukmyi","nnsy","kimnu","usimn","usmkin","iumn","i","insnmky","km","nimun","kysunm","kmnynui","nusy","kyum","yumkisn","n","iy"}

	result := wordSubsets(A, B)
	fmt.Printf("result = %v\n", result)
}

