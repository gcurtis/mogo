package mogo

type R []interface{}

type Returnable struct {
	returns []interface{}
}

func (this *Returnable) AndReturns(r R) {
	this.returns = r
}

func (this *Returnable) AndReturnsOne(r interface{}) {
	this.returns = R{r}
}
