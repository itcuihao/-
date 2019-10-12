package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

var gene = map[string]string{
	"GGGCCC":   "Apal",
	"GGATCC":   "BamHI",
	"AGATCT":   "BGIII",
	"GAATTC":   "EcoRI",
	"AAGCTT":   "HindIII",
	"GGTACC":   "Kpnl",
	"CCATGG":   "Ncol",
	"CATATG":   "Ndel",
	"GCGGCCGC": "Notl",
	"GAGCTC":   "Sacl",
	"GTCGAC":   "Sall",
	"GCATGC":   "Sphl",
	"TCTAGA":   "Xbal",
	"CTCGAG":   "Xhol",
}

func getGene() map[string]string {
	data := make(map[string]string)
	for k, v := range gene {
		data[k] = v
	}
	return data
}

func main() {
	path, err := getDirFile("./data")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("目录文件：", path)

	notGene := getGene()
	for _, f := range path {
		fbyte, err := ioutil.ReadFile(f)
		if err != nil {
			log.Println(err)
			continue
		}
		fstr := string(fbyte)
		for seq := range gene {
			index := SearchString(fstr, seq)
			if index >= 0 {
				log.Printf("%s 存在序列：%s，位置：%d；\n", f, seq, index)
				delete(notGene, seq)
			}
		}
	}

	log.Println("文件中都不存在的为：")
	enzyme := make([]string, 0, len(notGene))
	for _, e := range notGene {
		enzyme = append(enzyme, e)
	}
	log.Println("酶：", enzyme)
}

func getDirFile(d string) ([]string, error) {
	paths := make([]string, 0, 1)
	dir, err := ioutil.ReadDir(d)
	if err != nil {
		return paths, err
	}

	for _, f := range dir {
		if strings.HasSuffix(f.Name(), ".fasta") {
			paths = append(paths, fmt.Sprintf("%s/%s", d, f.Name()))
		}
	}
	return paths, nil
}

const (
	PatternSize int = 100
)

func SearchNext(haystack string, needle string) int {
	retSlice := KMP(haystack, needle)
	if len(retSlice) > 0 {
		return retSlice[len(retSlice)-1]
	}

	return -1
}

func SearchString(haystack string, needle string) int {
	retSlice := KMP(haystack, needle)
	if len(retSlice) > 0 {
		return retSlice[0]
	}

	return -1
}

func KMP(haystack string, needle string) []int {
	next := preKMP(needle)
	i := 0
	j := 0
	m := len(needle)
	n := len(haystack)

	x := []byte(needle)
	y := []byte(haystack)
	var ret []int

	//got zero target or want, just return empty result
	if m == 0 || n == 0 {
		return ret
	}

	//want string bigger than target string
	if n < m {
		return ret
	}

	for j < n {
		for i > -1 && x[i] != y[j] {
			i = next[i]
		}
		i++
		j++

		//fmt.Println(i, j)
		if i >= m {
			ret = append(ret, j-i)
			//fmt.Println("find:", j, i)
			i = next[i]
		}
	}

	return ret
}

func preMP(x string) [PatternSize]int {
	var i, j int
	length := len(x) - 1
	var mpNext [PatternSize]int
	i = 0
	j = -1
	mpNext[0] = -1

	for i < length {
		for j > -1 && x[i] != x[j] {
			j = mpNext[j]
		}
		i++
		j++
		mpNext[i] = j
	}
	return mpNext
}

func preKMP(x string) [PatternSize]int {
	var i, j int
	length := len(x) - 1
	var kmpNext [PatternSize]int
	i = 0
	j = -1
	kmpNext[0] = -1

	for i < length {
		for j > -1 && x[i] != x[j] {
			j = kmpNext[j]
		}

		i++
		j++

		if x[i] == x[j] {
			kmpNext[i] = kmpNext[j]
		} else {
			kmpNext[i] = j
		}
	}
	return kmpNext
}
