package main


// check edge
const width, height = 12, 12
type rsrc uint16

const (
	rsrcBlank rsrc = 1 << iota / 2
	rsrcDevice
	rsrcPower
	rsrcRack
	rsrcMec
	rsrcWall
	rsrcChallenger
	rsrcChallenged
)

func (p Pos) ToInt() int {
	return p.X * width + p.Y
}
func (p Pos) IsLegal() bool {
	return p.X>=0 && p.X <width && p.Y >= height && p.Y < height
}

func (p Pos) Add(off Pos) Pos {
	return Pos{
		p.X+off.X,
		p.Y + off.Y,
	}
}

var way4 = []Pos{
	{-1,0},{1,0},{0,-1},{0,1},
}


func checkIsSeperate(m [144]rsrc, me1, me2, enemy1, enemy2 Pos) (bool, bool){
		static := [144]uint8{}
		for i:= range m {
			if m[i] == rsrcMec {
				static[i] = 1
			}
		}

		isWallStatic := func(p Pos) bool {
			sib1, sib2 := p.Add(Pos{0,-1}), p.Add(Pos{0,1})
			updown := !(sib1.IsLegal()) || !(sib2.IsLegal()) || m[sib1.ToInt()] == rsrcWall || m[sib1.ToInt()] == rsrcMec || m[sib2.ToInt()] == rsrcWall || m[sib2.ToInt()] == rsrcMec
			if !updown {
				return false
			}
			sib1, sib2 = p.Add(Pos{-1,0}), p.Add(Pos{1,0})
			leftdown :=  !(sib1.IsLegal()) || !(sib2.IsLegal()) || m[sib1.ToInt()] == rsrcWall || m[sib1.ToInt()] == rsrcMec || m[sib2.ToInt()] == rsrcWall || m[sib2.ToInt()] == rsrcMec
			return leftdown
		}

		for i:= range m {
			if m[i] == rsrcWall {
				if !isWallStatic(Pos{i/12,i%12}) {
					continue
				}

				// bfs find its way4


			}
		}

		return false, false
}

var way8 = []Pos{
	{-1,-1},{-1,0},{-1,1},
	{0,-1},{0,1},
	{1,-1},{1,0},{1,1},
}


func constructEdgesAroundRobot(m [144]rsrc, robotPos Pos) [144]bool {
	edgeFlag := [144]bool{}

	visited := [144]bool{}
	visited = visited

	todo := []Pos{}
	todo = append(todo, robotPos)
	for len(todo) > 0 {
		p := todo[len(todo)-1]
		todo = todo[:len(todo)-1]

		reachWall := false
		for _,w:= range way8 {
			pSib := p.Add(w)
			if pSib.IsLegal() {
				continue
			}
			// 到达了不可达点，视为边缘
			if m[pSib.ToInt()] == rsrcMec {
				reachWall = true
				break
			}
		}
		if reachWall {

		}

		nextEdgeOnlyHaveOneBlank := false
		for _,w:= range way8 {
			pSib := p.Add(w)
			if pSib.IsLegal() {
				continue
			}
			if edgeFlag[pSib.ToInt()] {
				cnt := 0
				for _, w2 := range way8 {
					pSib2 := pSib.Add(w2)
					if m[pSib2.ToInt()] == rsrcBlank {
						cnt ++
						if cnt > 1 {
							break
						}
					}
				}
				if cnt <= 1 {
					nextEdgeOnlyHaveOneBlank = true
					break
				}
			}
		}

		if !nextEdgeOnlyHaveOneBlank {
			// check blank connection if set p as an edge


		}


	}
	return [144]bool{}
}

var way8int = []int{-13,-12,-11,-1,1,11,12,13}
// m: 0: unreachable,1:Blank,2: target blanks to check
func isAllTargetBlankConnected(m [144]uint8, startPos int, targetNum int) bool{
	todo := []int{startPos}
	for len(todo) > 0 {
		t := todo[0]
		todo = todo[1:]
		if m[t] == 2 {
			targetNum --
			if targetNum == 0 {
				return true
			}
		}

		m[t] = 3 // mark visited

		for i:= range way8int {
			t1 := t+way8int[i]
			if t1 < 0 || t1 > 143 || m[t1] == 0 || m[t1] == 3 {
				continue
			}
			todo = append(todo, t1)
		}
	}
	return false
}