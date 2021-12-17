package main

import (
	"fmt"
)

type vec struct {
	x, y int
}

type ball struct {
	pos, v vec
}

func newBall(v vec) ball {
	return ball{
		vec{0, 0},
		v,
	}
}

func (b *ball) move() vec {
	b.pos.x += b.v.x
	b.pos.y += b.v.y
	if b.v.x > 0 {
		b.v.x -= 1
	} else if b.v.x < 0 {
		b.v.x += 1
	}
	b.v.y -= 1
	return b.pos
}

type rect struct {
	xmin, xmax, ymin, ymax int
}

func (r rect) contains(v vec) bool {
	return v.x >= r.xmin && v.x <= r.xmax && v.y >= r.ymin && v.y <= r.ymax
}

func shoot(v vec, t rect) (bool, int) {
	ball := newBall(v)
	ym := 0
	for {
		pos := ball.move()
		if pos.y > ym {
			ym = pos.y
		}
		if t.contains(pos) {
			return true, ym
		}
		if pos.x > t.xmax || pos.y < t.ymin {
			return false, ym
		}
	}
}

func main() {
	target := rect{150, 193, -136, -86}

	my := 0
	cc := 0
	for y := -1000; y < 1000; y++ {
		for x := -1000; x < 1000; x++ {
			result, p := shoot(vec{x, y}, target)
			if result {
				cc++
				if p > my {
					my = p
				}
			}
		}
	}
	fmt.Println(my)
	fmt.Println(cc)
}
