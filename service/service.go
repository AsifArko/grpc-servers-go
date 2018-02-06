package service

func (p Person)FullName()string  {
	return p.FirstName+" "+p.LastName
}
