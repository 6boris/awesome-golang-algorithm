package Solution

type Robot struct {
	moved bool
	idx   int
	pos   [][2]int
	dir   []int
	toDir map[int]string
}

func Constructor(width, height int) Robot {
	robot := Robot{
		toDir: map[int]string{
			0: "East",
			1: "North",
			2: "West",
			3: "South",
		},
	}

	for i := 0; i < width; i++ {
		robot.pos = append(robot.pos, [2]int{i, 0})
		robot.dir = append(robot.dir, 0)
	}
	for i := 1; i < height; i++ {
		robot.pos = append(robot.pos, [2]int{width - 1, i})
		robot.dir = append(robot.dir, 1)
	}
	for i := width - 2; i >= 0; i-- {
		robot.pos = append(robot.pos, [2]int{i, height - 1})
		robot.dir = append(robot.dir, 2)
	}
	for i := height - 2; i > 0; i-- {
		robot.pos = append(robot.pos, [2]int{0, i})
		robot.dir = append(robot.dir, 3)
	}
	robot.dir[0] = 3

	return robot
}

func (this *Robot) Step(num int) {
	this.moved = true
	this.idx = (this.idx + num) % len(this.pos)
}

func (this *Robot) GetPos() []int {
	return []int{this.pos[this.idx][0], this.pos[this.idx][1]}
}

func (this *Robot) GetDir() string {
	if !this.moved {
		return "East"
	}
	return this.toDir[this.dir[this.idx]]
}

type input struct {
	name string
	step int
}

func Solution(width, height int, inputs []input) []any {
	ret := make([]any, 0)
	c := Constructor(width, height)
	for i := range inputs {
		if inputs[i].name == "step" {
			c.Step(inputs[i].step)
			continue
		}
		if inputs[i].name == "pos" {
			ret = append(ret, c.GetPos())
			continue
		}
		ret = append(ret, c.GetDir())
	}
	return ret
}
