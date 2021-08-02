package domain

type Holder struct {
	Name, Cpf, Job string
}

func CreateHolder(Name string, Cpf string, Job string) Holder {
	newHolder := Holder{}
	newHolder.Name = Name
	newHolder.Cpf = Cpf
	newHolder.Job = Job
	return newHolder
}
