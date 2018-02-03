package main

import "fmt"

var pitchesMap = map[string]int{
	"C":  C,
	"Cs": Cs,
	"Db": Db,
	"Df": Db,
	"D":  D,
	"Ds": Ds,
	"Eb": Eb,
	"Ef": Ef,
	"E":  E,
	"F":  F,
	"Fs": Fs,
	"Gb": Gb,
	"Gf": Gf,
	"G":  G,
	"Gs": Gs,
	"Ab": Ab,
	"Af": Af,
	"A":  A,
	"As": As,
	"Bb": Bb,
	"Bf": Bf,
	"B":  B,
}

type note struct {
	name   string
	number int
}

var notes = []note{
	note{"C", C},
	note{"Cs", Cs},
	note{"Db", Db},
	note{"Df", Db},
	note{"D", D},
	note{"Ds", Ds},
	note{"Eb", Eb},
	note{"Ef", Ef},
	note{"E", E},
	note{"F", F},
	note{"Fs", Fs},
	note{"Gb", Gb},
	note{"Gf", Gf},
	note{"G", G},
	note{"Gs", Gs},
	note{"Ab", Ab},
	note{"Af", Af},
	note{"A", A},
	note{"As", As},
	note{"Bb", Bb},
	note{"Bf", Bf},
	note{"B", B},
}

var octaves = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

func main() {
	for _, n := range notes {
		for _, i := range octaves {
			midiNoteNumber := n.number + ((i + 1) * 12)
			if midiNoteNumber < 128 {
				fmt.Printf("const %s%d = %d\n", n.name, i, midiNoteNumber)
			}
		}
	}
}

// C  Pitch Class
const C = 0

// Cs represents the C Sharp Pitch Class
const Cs = 1

// Df represents the D Flat Pitch Class
const Df = 1

// Db represents the Flat Pitch Class
const Db = 1

// D Pitch Class
const D = 2

// Ds represents the D Sharp Pitch Class
const Ds = 3

// Ef represents E Flat Pitch Class
const Ef = 3

// Eb represents the E Flat Pitch Class
const Eb = 3

// E  Pitch Class
const E = 4

// F  Pitch Class
const F = 5

// Fs represents the F Sharp Pitch Class
const Fs = 6

// Gf represents the G Flat Pitch Class
const Gf = 6

// Gb represents the G Flat Pitch Class
const Gb = 6

// G represetns the G Pitch Class
const G = 7

// Gs Represents the G Sharp Pitch Class
const Gs = 8

// Af represents the A Flat Pitch Class
const Af = 8

// Ab represents the A Flat Pitch Class
const Ab = 8

// A  Pitch Class
const A = 9

// As represents the A  Sharp Pitch Class
const As = 10

// Bf represents the B Flat Pitch Class
const Bf = 10

// Bb represents the Flat Pitch Class
const Bb = 10

// B  Pitch Class
const B = 11
