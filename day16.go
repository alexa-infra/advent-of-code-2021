package main

import (
	"bufio"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"
)

func hexToBin(hex string) string {
	m := map[rune]string{
		'0': "0000",
		'1': "0001",
		'2': "0010",
		'3': "0011",
		'4': "0100",
		'5': "0101",
		'6': "0110",
		'7': "0111",
		'8': "1000",
		'9': "1001",
		'A': "1010",
		'B': "1011",
		'C': "1100",
		'D': "1101",
		'E': "1110",
		'F': "1111",
	}
	parts := make([]string, 0, len(hex))
	for _, ch := range []rune(hex) {
		part, ok := m[ch]
		if !ok {
			log.Fatalf("wrong hex")
		}
		parts = append(parts, part)
	}
	return strings.Join(parts, "")
}

type packet struct {
	ver, typ, val int
	packets       []packet
}

func binToInt(bin string) int {
	b := new(big.Int)
	_, ok := b.SetString(bin, 2)
	if !ok {
		log.Fatalf("wrong bin")
	}
	return int(b.Int64())
}

func parseOne(bin string) (int, packet) {
	i := 0
	ver := binToInt(bin[i : i+3])
	i += 3
	typ := binToInt(bin[i : i+3])
	i += 3
	if typ == 4 {
		buf := ""
		for {
			five := bin[i : i+5]
			i += 5
			buf += five[1:]
			if five[:1] == "0" {
				break
			}
		}
		val := binToInt(buf)
		pack := packet{ver, typ, val, []packet{}}
		return i, pack
	}
	typLen := bin[i : i+1]
	i += 1
	if typLen == "0" {
		bitLen := binToInt(bin[i : i+15])
		i += 15
		packets := parseMany(bin[i : i+bitLen])
		i += bitLen
		pack := packet{ver, typ, 0, packets}
		return i, pack
	} else {
		num := binToInt(bin[i : i+11])
		i += 11
		packets := []packet{}
		for j := 0; j < num; j++ {
			ix, p := parseOne(bin[i:])
			i += ix
			packets = append(packets, p)
		}
		pack := packet{ver, typ, 0, packets}
		return i, pack
	}
}

func parseMany(bin string) []packet {
	packets := []packet{}
	i := 0
	for i < len(bin) {
		if len(bin)-i < 10 {
			break
		}
		ix, p := parseOne(bin[i:])
		i += ix
		packets = append(packets, p)
	}
	return packets
}

func (p *packet) versionValue() int {
	val := p.ver
	for _, pp := range p.packets {
		val += pp.versionValue()
	}
	return val
}

func (p *packet) value() int {
	if p.typ == 0 {
		s := 0
		for _, pp := range p.packets {
			s += pp.value()
		}
		return s
	}
	if p.typ == 1 {
		s := 1
		for _, pp := range p.packets {
			s *= pp.value()
		}
		return s
	}
	if p.typ == 2 {
		s := 0
		first := true
		for _, pp := range p.packets {
			v := pp.value()
			if first {
				s = v
				first = false
			} else if v < s {
				s = v
			}
		}
		return s
	}
	if p.typ == 3 {
		s := 0
		first := true
		for _, pp := range p.packets {
			v := pp.value()
			if first {
				s = v
				first = false
			} else if v > s {
				s = v
			}
		}
		return s
	}
	if p.typ == 4 {
		return p.val
	}
	if p.typ == 5 {
		v1 := p.packets[0].value()
		v2 := p.packets[1].value()
		if v1 > v2 {
			return 1
		}
		return 0
	}
	if p.typ == 6 {
		v1 := p.packets[0].value()
		v2 := p.packets[1].value()
		if v1 < v2 {
			return 1
		}
		return 0
	}
	if p.typ == 7 {
		v1 := p.packets[0].value()
		v2 := p.packets[1].value()
		if v1 == v2 {
			return 1
		}
		return 0
	}
	log.Fatalf("wrong typ")
	return 0
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	bin := hexToBin(line)
	packets := parseMany(bin)
	r1 := 0
	for _, p := range packets {
		r1 += p.versionValue()
	}
	fmt.Println("Part 1:", r1)
	r2 := 0
	for _, p := range packets {
		r2 = p.value()
	}
	fmt.Println("Part 2:", r2)
}
