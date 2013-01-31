package mogo

type R []interface{}

type returnable struct {
	returns []interface{}
}

func (this *returnable) AndReturn(r ...interface{}) {
	this.returns = r
}
