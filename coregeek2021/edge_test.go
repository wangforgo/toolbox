package main

import (
	"testing"
	"time"
)

func TestIsAllTargetConnected(t *testing.T) {
	t1 := time.Now()
	m := [144]uint8{}
	targets := []Pos{
		{0,0},
		{2,0},
		{0,2},
		{3,3},
		{4,4},
	}

	blanks := []Pos{
		{1,1},
		{2,2},
	}

	for i:= range blanks{
		m[blanks[i].ToInt()] = 1
	}

	for i:= range targets{
		m[targets[i].ToInt()] = 2
	}

	ans := isAllTargetBlankConnected(m,0,len(targets))
	println(time.Now().Sub(t1).Milliseconds())
	println(ans)

}
