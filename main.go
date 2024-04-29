package main

import (
	autopsy "BruteForce/Autopsy"
	"errors"
	"fmt"
	"os"
)

const SCREEN_HEIGHT = 900
const SCREEN_WIDTH = 900
const MAX_VERTICES = 5
const OFFSET = 0

func print_int_slice(slc []int) string {
	out := "{"
	for i := 0; i < len(slc); i++ {
		out += fmt.Sprintf("%d", slc[i])
		if i != len(slc)-1 {
			out += ","
		}
	}
	out += "}"
	return out
}
func reached_contains(reached *[]Quiver, quiv Quiver) bool {
	for i := 0; i < len(*reached); i++ {
		if QuiverStrictlyEqual((*reached)[i], quiv) {
			return true
		}
	}
	return false
}
func GreedyReduction(quiv Quiver, reached *[]Quiver) (Quiver, error) {
	*reached = append(*reached, quiv)
	mutation_targets := make([]int, 0)
	min := 2147483647
	quivm := quiv.ToMatrix()
	autopsy.Store(fmt.Sprintf("staring matrix:\n%s\n", quivm.ToString()))
	autopsy.Store(fmt.Sprintf("starting cost: %d\n", Cost(quiv)))
	for i := 0; i < 4; i++ {
		m := Cost(quiv.MutateAt(i))
		autopsy.Store(fmt.Sprintf("index: %d, cost: %d", i, m))
		if m == min {
			mutation_targets = append(mutation_targets, i)
		}
		if m < min {
			min = m
			mutation_targets = make([]int, 0)
			mutation_targets = append(mutation_targets, i)
		}
	}
	if min > Cost(quiv) {
		s := quiv.ToMatrix()
		autopsy.Store(s.ToString())
		return quiv, nil
	}
	var old Quiver
	hit := false
	for i := 0; i < len(mutation_targets); i++ {
		q := quiv.MutateAt(mutation_targets[i])
		if reached_contains(reached, q) {
			continue
		}
		tmp, err := GreedyReduction(q, reached)
		if err != nil {
			return tmp, err
		}
		if hit {
			l := old.ToMatrix()
			v := tmp.ToMatrix()
			if !QuiverIsEqual(old, tmp) {
				autopsy.Store(fmt.Sprintf("not equal\n%s\n%s\n", l.ToString(), v.ToString()))
				autopsy.Store(fmt.Sprint("l mutations: ", print_int_slice(tmp.mutations)))
				autopsy.Store(fmt.Sprint("v mutations: ", print_int_slice(old.mutations)))
				return old, errors.New("not equal quivers")
			}
		}
		old = tmp
		hit = true
	}
	if !hit {
		s := quiv.ToMatrix()
		autopsy.Store(s.ToString())
		return quiv, nil
	}
	return old, nil
}
func main() {
	autopsy.Init()
	for i := 1; i < 1000; i++ {
		for j := 0; j < i*100; j++ {
			reached := make([]Quiver, 0)
			//recursed = false
			autopsy.Reset()
			q := RandomQuiver(4, i+2)
			m := q.ToMatrix()
			_, err := GreedyReduction(q, &reached)
			if err != nil {
				println(err.Error())
				println("counter example:\n")
				println(m.ToString())
				autopsy.Dump()
				os.Exit(1)
			}

		}
		println(i)
	}

}

/*
tan(9) =

*/
