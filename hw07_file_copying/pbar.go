package main

import (
	"errors"
	"fmt"
)

var ErrInvalidProgress = errors.New("non-positive progress value")

type progressBar interface {
	AddProgress(pos int64) error
	SetMin(v int64)
	SetMax(v int64)
}

type pBar struct {
	min, max int64
}

func ProgressBar() progressBar { return new(pBar) }
func (p *pBar) SetMax(v int64) { p.max = v }
func (p *pBar) SetMin(v int64) { p.min = v }

func (p *pBar) AddProgress(pos int64) error {
	if pos < 0 || pos > p.max {
		return ErrInvalidProgress
	}
	p.min += pos
	percents := p.min * 100 / p.max
	fmt.Printf("\r[")
	for i := 0; i < 20; i++ {
		if i < int(percents/5) {
			fmt.Printf("#")
		} else {
			fmt.Printf("_")
		}
	}
	fmt.Printf("]%v%s", percents, "%")
	if percents == 100 {
		fmt.Printf("\n")
	}
	return nil
}
