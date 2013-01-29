package main

import (
	"math/rand"
	"time"
)

type LifeGame struct {
	board    [][]bool
	row, col int
	time     int
	repeated uint
}

func new_board(row, col int) *LifeGame {
	p := new(LifeGame)
	p.row = row
	p.col = col

	p.board = make([][]bool, row)
	for r := 0; r < row; r++ {
		p.board[r] = make([]bool, col)
	}

	return p
}

func (p *LifeGame) clone() *LifeGame {
	cloned := new_board(p.row, p.col)

	for r := 0; r < p.row; r++ {
		for c := 0; c < p.col; c++ {
			cloned.board[r][c] = p.board[r][c]

		}
	}

	return cloned
}

func (p *LifeGame) init_rand() *LifeGame {
	rand.Seed(time.Now().Unix())

	for r := 0; r < p.row; r++ {
		for c := 0; c < p.col; c++ {
			if rand.Float32() < 0.3 {
				p.board[r][c] = true
			}
		}
	}

	return p
}

func (p *LifeGame) init_pattern_block() *LifeGame {
	p.init_rand()

	if p.row >= 4 || p.col >= 4 {
		for i := 0; i < 4; i++ {
			for j := 0; j < 4; j++ {
				p.board[i][j] = false
			}
		}

		p.board[1][1] = true
		p.board[1][2] = true
		p.board[2][1] = true
		p.board[2][2] = true
	}

	return p
}

func (p *LifeGame) generate() *LifeGame {
	gened := new_board(p.row, p.col)

	for r := 0; r < p.row; r++ {
		for c := 0; c < p.col; c++ {
			gened.board[r][c] = p.is_dead_or_alive(r, c)
		}
	}
	gened.time = p.time + 1
	return gened
}

func (p *LifeGame) show() {
	if p.repeated != 0 {
		println("time =", p.time, " repeated! period:", p.repeated)
	} else {
		println("time =", p.time)
	}
	f := func(b bool) (s string) {
		if b {
			s = "[*]"
		} else {
			s = "[ ]"
		}
		return s
	}
	for r := 0; r < p.row; r++ {
		for c := 0; c < p.col; c++ {
			print(f(p.board[r][c]))
		}
		println()
	}
}

func (p *LifeGame) count_now_alive_roll(r, c int) (i int) {
	if r < 0 {
		r += p.row
	}
	if p.row <= r {
		r -= p.row
	}
	if c < 0 {
		c += p.col
	}
	if p.col <= c {
		c -= p.col
	}
	if p.board[r][c] {
		i = 1
	} else {
		i = 0
	}
	return i
}

func (p *LifeGame) count_now_alive_normal(r, c int) int {
	if r < 0 || r >= p.row || c < 0 || c >= p.col || !p.board[r][c] {
		return 0
	}

	return 1
}

func (p *LifeGame) count_now_alive(r, c int) int {
	return p.count_now_alive_roll(r, c)
}

func (p *LifeGame) is_dead_or_alive(r, c int) (b bool) {
	count := p.count_now_alive(r-1, c-1) +
		p.count_now_alive(r-1, c) +
		p.count_now_alive(r-1, c+1) +
		p.count_now_alive(r, c-1) +
		p.count_now_alive(r, c) +
		p.count_now_alive(r, c+1) +
		p.count_now_alive(r+1, c-1) +
		p.count_now_alive(r+1, c) +
		p.count_now_alive(r+1, c+1)
	switch count {
	case 3:
		b = true
	case 4:
		b = p.board[r][c]
	default:
		b = false
	}
	return b
}

func (p *LifeGame) is_same(d *LifeGame) (b bool) {
	if p.row != d.row || p.col != d.col {
		return false
	}

	for r := 0; r < p.row; r++ {
		for c := 0; c < p.col; c++ {
			if p.board[r][c] != d.board[r][c] {
				return false
			}
		}
	}

	return true
}

func Run(row, col int) {
	game := new_board(row, col).init_pattern_block()

	var prev *LifeGame = nil
	for game.repeated == 0 {
		game.show()

		next := game.generate()

		if next.is_same(game) {
			next.repeated = 1
		} else if prev != nil && next.is_same(prev) {
			next.repeated = 2
		}

		prev = game
		game = next

		time.Sleep(time.Millisecond * 500)
	}
	game.show()
}
