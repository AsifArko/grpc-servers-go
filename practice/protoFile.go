package protofile

import "fmt"

func (p Person) Fullname()string{
	return fmt.Sprintf("%s %s",p.FirstName,p.LastName)
}
