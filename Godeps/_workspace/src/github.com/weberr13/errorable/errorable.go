package errorable

type IntErr struct {
	I int
	E error
}
type IntErrs []IntErr

func NewIntErr(f func(string) (int, error), s string) *IntErr {
	this := &IntErr{}
	this.I, this.E = f(s)
	return this
}
func NewIntErrs(f func(string) (int, error), s ...string) *IntErrs {
	interr := make(IntErrs, len(s), len(s))
	for i, v := range s {
		interr[i] = *NewIntErr(f, v)
	}
	return (&interr)
}
func (this *IntErrs) Get(i int) int {
	return (*this)[i].I
}
func (this *IntErrs) Len() int {
	return (len(*this))
}
func (this *IntErrs) GetFirstErr() (err error) {
	for _, v := range *this {
		if nil != v.E {
			return v.E
		}
	}
	return nil
}
