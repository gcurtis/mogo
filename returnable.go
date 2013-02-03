package mogo

type R []interface{}

var DefaultR R = make(R, 32)

type returnable struct {
	returns []interface{}
	set     bool
}

func (this *returnable) AndReturn(r ...interface{}) {
	this.set = true
	this.returns = r
}

func (this *returnable) ret() R {
	if !this.set {
		return DefaultR
	}

	return this.returns
}
