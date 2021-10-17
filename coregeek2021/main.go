package main

type (
	Ctx struct {
		Map MapDto
		Player Player
	}
	Player struct {
		Pos Pos
		Comps []string
	}

	MapDto struct {
		Width,Height int
		Zones []Zone
	}
	Zone struct {
		Pos Pos
		Units []EnumRsrc
	}
	Pos struct {
		X,Y int
	}
)
type EnumRsrc string
const (
	RsrcWall EnumRsrc = "WALL"
	RsrcBlank EnumRsrc = "BLANK"
	RsrcMec EnumRsrc = "MEC"
	RsrcPower EnumRsrc = "POWER"
	RsrcDevice EnumRsrc = "DEVICE"
	RsrcRack EnumRsrc = "RACK"
	RsrcChallenger EnumRsrc = "CHALLENGER"
	RsrcChallenged EnumRsrc = "CHALLENGED"
)

func main() {
	
}
