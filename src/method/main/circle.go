package main 

import (

)
type circle struct{
	redius float64
}

func (c circle) area() float64 {
	return c.redius*c.redius
}

