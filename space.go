package space

import (
	"container/list"
	"fmt"
	"math"
)

type Space struct {
	vectors *List
}

const g_dimention = 2
const g_width = 400
const g_hight = 400

func (s *Space) AddVector(v Vector) {

}

func (s *Space) Init() {
	vectors := list.New()
	s.vectors = vectors
	return s
}
func New() *Space {
	fmt.Println("creating a world")
	fmt.Println(math.Pi)
	return new(Space).Init()
}
