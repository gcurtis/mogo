package mogo

type R []interface{}

var DefaultR R = make(R, 32)

type returnable struct {
	returns []interface{}
	set     bool
}

func (this *returnable) AndReturn(r ...interface{}) {
	this.returns = r

	for i, r := range this.returns {
		if a, ok := r.(any); ok {
			this.returns[i] = a.zero()
		}
	}

	this.set = true
}

func (this *returnable) ret() R {
	if !this.set {
		return DefaultR
	}

	return this.returns
}
