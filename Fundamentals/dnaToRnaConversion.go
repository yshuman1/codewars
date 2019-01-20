/*
Deoxyribonucleic acid, DNA is the primary information storage molecule in biological systems. It is composed of four nucleic acid bases Guanine ('G'), Cytosine ('C'), Adenine ('A'), and Thymine ('T').

Ribonucleic acid, RNA, is the primary messenger molecule in cells. RNA differs slightly from DNA its chemical structure and contains no Thymine. In RNA Thymine is replaced by another nucleic acid Uracil ('U').

Create a funciton which translates a given DNA string into RNA.

For example:

DNAtoRNA("GCAT") // returns "GCAU"
The input string can be of arbitrary length - in particular, it may be empty. All input is guaranteed to be valid, i.e. each input string will only ever consist of 'G', 'C', 'A' and/or 'T'.

*/

package main

import (
	"fmt"
	"strings"
)

func DNAtoRNA(dna string) string {
	// your code here
	return strings.Replace(dna, "T", "U", -1)
}

func main() {
	dna := "GCAT"
	// expected output "GCAU"
	fmt.Println(DNAtoRNA(dna))
}
