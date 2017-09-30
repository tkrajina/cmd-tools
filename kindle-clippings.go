package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"regexp"
	"sort"
	"strings"
)

var metadataReg = regexp.MustCompile(`\-\s+Your\s(.*?)\s+on\s+(.*?)\s*\|.*`)

func panicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	flag.Parse()

	for _, filename := range flag.Args() {
		bytes, err := ioutil.ReadFile(filename)
		panicIfErr(err)

		parseKindleNotes(string(bytes))
	}
}

func parseKindleNotes(str string) {
	var keys []string
	var texts = map[string][]string{}

	var unknown bytes.Buffer
	parts := strings.Split(str, "==========")
	for _, part := range parts {
		part = strings.TrimLeft(part, " \n\r\t")
		lines := strings.Split(part, "\n")
		if len(lines) < 3 {
			unknown.WriteString(part)
		} else {
			title := lines[0]
			metadata := lines[1]
			text := strings.TrimSpace(strings.Join(lines[2:], "\n"))
			//fmt.Println("title:", title)
			//fmt.Println("text:", text)

			if metadataReg.MatchString(metadata) {
				all := metadataReg.FindAllStringSubmatch(metadata, -1)
				typ := all[0][1]
				loc := all[0][2]
				_ = loc
				//fmt.Printf("type=%s location=%s\n", typ, loc)

				key := fmt.Sprintf("%s :: %s", strings.TrimSpace(title), strings.TrimSpace(typ))
				if _, found := texts[key]; !found {
					keys = append(keys, key)
					texts[key] = []string{}
				}
				texts[key] = append(texts[key], text)
			} else {
				unknown.WriteString(part)
			}
		}
	}

	sort.Strings(keys)

	for _, key := range keys {
		fmt.Println("#", key, "\n")
		for _, text := range texts[key] {
			fmt.Println(text)
		}
		fmt.Println("\n\n")
	}

	if unknown.Len() > 0 {
		fmt.Println("# Unknown")
		fmt.Println()
		fmt.Println(unknown.String())
	}
}

var str = `Allocation Efficiency in High-Performance Go Services (Achille Roussel)
- Your Highlight on Location 151-152 | Added on Thursday, September 28, 2017 12:03:35 PM

method calls on interface types are more expensive than those on struct types.
==========
A Short History of the World (Lascelles, Christopher)
- Your Highlight on page 43 | Location 461-462 | Added on Thursday, September 28, 2017 12:53:29 PM

Jesus was born sometime between 6 BC and 4 BC.
==========
A Short History of the World (Lascelles, Christopher)
- Your Highlight on page 43 | Location 463-464 | Added on Thursday, September 28, 2017 12:53:54 PM

He challenged and angered the established Pharisee leaders, who successfully called for the Roman occupiers to crucify him
==========
A Short History of the World (Lascelles, Christopher)
- Your Highlight on page 43 | Location 467-468 | Added on Thursday, September 28, 2017 12:55:11 PM

crucifixion in circa AD 28-29
==========
A Short History of the World (Lascelles, Christopher)
- Your Highlight on page 45 | Location 470-471 | Added on Thursday, September 28, 2017 12:55:36 PM

AD 380, Christianity had become the state religion of Rome.
==========
A Short History of the World (Lascelles, Christopher)
- Your Highlight on page 45 | Location 475-475 | Added on Thursday, September 28, 2017 12:56:12 PM

Emperor Claudius launched a major invasion of England in AD 43
==========
A Short History of the World (Lascelles, Christopher)
- Your Highlight on page 45 | Location 476-477 | Added on Thursday, September 28, 2017 12:56:37 PM

Emperor Nero had his mother and wife murdered and blamed the Great Fire of Rome in AD 64 on Christians, whom he had promptly thrown to the lions before he eventually committed suicide.
==========
A Short History of the World (Lascelles, Christopher)
- Your Highlight on page 45 | Location 478-478 | Added on Thursday, September 28, 2017 12:56:53 PM

the eruption of Mount Vesuvius in AD 79,
==========
A Short History of the World (Lascelles, Christopher)
- Your Highlight on page 45 | Location 479-480 | Added on Thursday, September 28, 2017 12:57:05 PM

After Titus’ death in AD 81, until the end of the 2nd century, emperors adopted their successors, as opposed to passing the crown down through family lines.
==========
A Short History of the World (Lascelles, Christopher)
- Your Highlight on page 45 | Location 480-480 | Added on Thursday, September 28, 2017 12:57:20 PM

This led to a succession of capable emperors,
==========
A Short History of the World (Lascelles, Christopher)
- Your Highlight on page 45 | Location 481-482 | Added on Thursday, September 28, 2017 12:57:39 PM

The appointment to emperor in AD 180 of Lucius Commodus, after the death of his father, Marcus Aurelius, was the first time a son had succeeded his father since AD 79.
==========
A Short History of the World (Lascelles, Christopher)
- Your Highlight on page 46 | Location 484-484 | Added on Thursday, September 28, 2017 12:58:06 PM

During a period of 50 years in the middle of the 3rd century AD there were more than 20 emperors,
==========
A Short History of the World (Lascelles, Christopher)
- Your Highlight on page 46 | Location 488-489 | Added on Thursday, September 28, 2017 12:58:45 PM

Rome was also increasingly threatened by the rise of the Persian Sassanids, who sensed weakness in their neighbour.
==========
Allocation Efficiency in High-Performance Go Services (Achille Roussel)
- Your Highlight on Location 216-216 | Added on Thursday, September 28, 2017 6:03:27 PM

Stack allocation is cheap, heap allocation is expensive.
==========
Love and Ruin (magazine.atavist.com)
- Your Highlight on Location 382-383 | Added on Thursday, September 28, 2017 10:12:03 PM

through an interpreter. “Is Kabul listening to you?” Murmurs.
==========
Hrvatski jezik - povijest (postjugo.filg.uj.edu.pl)
- Your Highlight on Location 11-12 | Added on Thursday, September 28, 2017 10:13:22 PM

Baltoslavenska jezična zajednica raspala se oko 1500-1300. g. pr. n. e.
==========
Hrvatski jezik - povijest (postjugo.filg.uj.edu.pl)
- Your Highlight on Location 11-12 | Added on Thursday, September 28, 2017 10:13:26 PM

Baltoslavenska jezična zajednica raspala se oko 1500-1300. g. pr. n. e.
==========
Hrvatski jezik - povijest (postjugo.filg.uj.edu.pl)
- Your Highlight on Location 18-21 | Added on Thursday, September 28, 2017 10:14:35 PM

Istočnoslavenski je danas mrtav jezik starocrkvenoslavenski (ili staroslavenski). To je jezik prve slavenske pismenosti i književnosti. Utemeljen je na makedonskom (to jest Istočno-južnoslavenskom) dijalektu iz okolice Soluna što su ga sveta braća Ćiril i Metod u drugoj polovici 9. stoljeća uzela za osnovu jezika slavenskoga bogoslužja
==========
Hrvatski jezik - povijest (postjugo.filg.uj.edu.pl)
- Your Highlight on Location 24-25 | Added on Thursday, September 28, 2017 10:15:58 PM

crkvenoslavenske redakcije: ruska, bugarska, makedonska, srpska, hrvatska. Hrvatskom se redakcijom služila hrvatska glagoljička liturgijska književnost,
==========
Hrvatski jezik - povijest (postjugo.filg.uj.edu.pl)
- Your Highlight on Location 29-30 | Added on Friday, September 29, 2017 7:17:31 AM

čakavskoj podlozi za liturgijske knjige razvijen crkvenoslavenski jezik hrvatske redakcije
==========
Hrvatski jezik - povijest (postjugo.filg.uj.edu.pl)
- Your Highlight on Location 18-21 | Added on Friday, September 29, 2017 7:18:46 AM

Istočnoslavenski je danas mrtav jezik starocrkvenoslavenski (ili staroslavenski). To je jezik prve slavenske pismenosti i književnosti. Utemeljen je na makedonskom (to jest Istočno-južnoslavenskom) dijalektu iz okolice Soluna što su ga sveta braća Ćiril i Metod u drugoj polovici 9. stoljeća uzela za osnovu jezika slavenskoga bogoslužja
==========
Hrvatski jezik - povijest (postjugo.filg.uj.edu.pl)
- Your Highlight on Location 24-24 | Added on Friday, September 29, 2017 7:19:18 AM

crkvenoslavenske redakcije: ruska, bugarska, makedonska, srpska, hrvatska.
==========
Hrvatski jezik - povijest (postjugo.filg.uj.edu.pl)
- Your Highlight on Location 27-28 | Added on Friday, September 29, 2017 7:19:35 AM

Pod kraj 9. stoljeća Hrvati su sa slavenskim bogoslužjem dobili i književni jezik, starocrkvenoslavenski i pismo, glagoljicu.
==========
Hrvatski jezik - povijest (postjugo.filg.uj.edu.pl)
- Your Highlight on Location 36-37 | Added on Friday, September 29, 2017 7:22:17 AM

Najstariji su spomenici hrvatskoga narodnoga jezika čakavskoga narječja "Istarski razvod" iz 1275, te "Vinodolski zakonik", 1288.
==========
Hrvatski jezik - povijest (postjugo.filg.uj.edu.pl)
- Your Highlight on Location 36-37 | Added on Friday, September 29, 2017 7:24:12 AM

Najstariji su spomenici hrvatskoga narodnoga jezika čakavskoga narječja "Istarski razvod" iz 1275, te "Vinodolski zakonik", 1288.
==========
How to: The Art of Memorisation and the Power of Spaced Repetition - WhyWhatHow.xyz (whywhathow.xyz)
- Your Highlight on Location 145-145 | Added on Friday, September 29, 2017 3:35:55 PM

Leitner system
==========
How to: The Art of Memorisation and the Power of Spaced Repetition - WhyWhatHow.xyz (whywhathow.xyz)
- Your Highlight on Location 183-183 | Added on Friday, September 29, 2017 3:40:11 PM

We often greatly overestimate what we can achieve in day and greatly underestimate what we can achieve in a year.
==========
How to: The Art of Memorisation and the Power of Spaced Repetition - WhyWhatHow.xyz (whywhathow.xyz)
- Your Highlight on Location 191-191 | Added on Friday, September 29, 2017 3:41:00 PM

interference:
==========
The International New York Times (The International New York Times)
- Your Highlight on Location 1531-1531 | Added on Friday, September 29, 2017 5:13:53 PM

Recurring
==========
The International New York Times (The International New York Times)
- Your Highlight on Location 1531-1531 | Added on Friday, September 29, 2017 5:14:10 PM

pickup trucks
==========
153-206-unenc  
- Your Highlight on page 159-159 | Added on Friday, September 29, 2017 8:42:44 PM

U međuvremenu, godine 568. Langobardi su ušli u Italiju i zauzeli veći dio sjeveroistočnih teritorija Apeninskog poluotoka
==========
153-206-unenc  
- Your Highlight on page 159-159 | Added on Friday, September 29, 2017 8:49:48 PM

prvi langobardski upad tek površno dotaknuo. Zbog tih je događaja akvilejski patrijarh, zajedno s većinom svojih vjernika i s crkvenim blagom, prebjegao u Gradež. Najvjerojatnije su u to vrijeme, nakon pljač- ke Trsta i bijega stanovništva s tog područja, nastali Justinopolis (današnji Kopar) i Novigrad
==========
153-206-unenc  
- Your Highlight on page 160-160 | Added on Friday, September 29, 2017 8:50:46 PM

panonski prostor, a potom i područje u blizini primorskih gradova. Od 599. do 611. godine ta, tada još poganska, plemena često upadaju u Istru, pljačkajući je i ubijajući tamošnje stanovništvo
==========
153-206-unenc  
- Your Highlight on page 160-160 | Added on Friday, September 29, 2017 8:51:52 PM

akti Rižanskog placita (sabora) iz 804.godine
==========
153-206-unenc  
- Your Highlight on page 161-161 | Added on Friday, September 29, 2017 8:53:46 PM

Slobodno se stanovništvo gradova i kaštela dijelilo u tri klase: svećenstvo, kao najistaknutiji dio stanovništva, posjednici te puk organiziran u razna cehovska udruženja
==========
153-206-unenc  
- Your Highlight on page 161-161 | Added on Friday, September 29, 2017 8:54:30 PM

Pula 66 solida mancosa, Poreč 66, Trst 60, Rovinj 40, Labin 30, Motovun 30, Buzet 20, Pićan 20 i Novigrad 12. Ta podjela djelomično odražava odnos veličine i važnosti pojedinih istarskih mjesta
==========
153-206-unenc  
- Your Highlight on page 163-163 | Added on Friday, September 29, 2017 8:57:48 PM

. Maura (III.st.
==========
153-206-unenc  
- Your Highlight on page 163-163 | Added on Friday, September 29, 2017 8:58:02 PM

spomenuti sv. Maura (III.st
==========
153-206-unenc  
- Your Highlight on page 163-163 | Added on Friday, September 29, 2017 8:58:19 PM

spomenuti sv. Maura (III.st.), prvoga porečkog biskupa
==========
A Short History of the World (Lascelles, Christopher)
- Your Highlight on page 46 | Location 489-491 | Added on Friday, September 29, 2017 9:08:33 PM

commanders in the more remote provinces increasingly began to behave as independent rulers, paying scant attention to Rome. It was in response to such problems that the Emperor Valerian split the empire into two zones of responsibility, one in the East and one in the West.
==========
jkljkl`
