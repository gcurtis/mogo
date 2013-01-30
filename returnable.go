package mogo

type R []interface{}

type Returnable struct {
	returns []interface{}
}

func (this *Returnable) AndReturn(r ...interface{}) {
	this.returns = r
}
