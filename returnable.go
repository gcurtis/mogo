package mogo

type R []interface{}

type Returnable struct {
	returns []interface{}
}

func (this *Returnable) AndReturn(r R) {
	this.returns = r
}

func (this *Returnable) AndReturnOne(r interface{}) {
	this.returns = R{r}
}
