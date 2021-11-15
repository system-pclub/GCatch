package toy

type MyInter1 interface {
	Equals(other MyInter2)
}

type MyInter2 interface {

}

type MyType1 struct {
	Index int
}

func (m *MyType1) Equals(other MyInter2) {
	cast := other.(*MyType1)
	_ = cast
	m.Index++
}