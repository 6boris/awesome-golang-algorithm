package Solution

type idWithVal struct {
	id, val int
}
type SnapshotArray struct {
	snpaId int
	data   [][]idWithVal
}

func Constructor(length int) SnapshotArray {
	data := make([][]idWithVal, length)
	for i := 0; i < length; i++ {
		data[i] = []idWithVal{{0, 0}}
	}
	return SnapshotArray{snpaId: 0, data: data}
}

func (this *SnapshotArray) Set(index int, val int) {
	length := len(this.data[index])
	if this.data[index][length-1].id != this.snpaId {
		this.data[index] = append(this.data[index], idWithVal{id: this.snpaId, val: val})
		return
	}
	this.data[index][length-1].val = val
}

func (this *SnapshotArray) Snap() int {
	id := this.snpaId
	this.snpaId++
	return id
}

func (this *SnapshotArray) Get(index int, snap_id int) int {
	// 该用二分
	length := len(this.data[index])
	for i := length - 1; i >= 0; i-- {
		if this.data[index][i].id <= snap_id {
			return this.data[index][i].val
		}
	}
	return 0
}

type op struct {
	name           string
	index, val, id int
}

func Solution(length int, ops []op) []int {
	c := Constructor(length)
	ans := make([]int, 0)
	for _, _op := range ops {
		if _op.name == "set" {
			c.Set(_op.index, _op.val)
			continue
		}
		if _op.name == "snap" {
			ans = append(ans, c.Snap())
			continue
		}
		ans = append(ans, c.Get(_op.index, _op.id))
	}
	return ans
}
