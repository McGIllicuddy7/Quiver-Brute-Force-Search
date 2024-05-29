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
func GreedyReduction(quiv Quiver, reached *[]Quiver, is_initial bool) (Quiver, int, error) {
	*reached = append(*reached, quiv)
	mutation_targets := make([]int, 0)
	min := 2147483647
	quivm := quiv.ToMatrix()
	if is_initial {
		autopsy.Store(fmt.Sprintf("staring matrix:\n%s\n", quivm.ToString()))
		autopsy.Store(fmt.Sprintf("starting cost: %d\n", Cost(quiv)))
	}

	for i := 0; i < 4; i++ {
		m := Cost(quiv.MutateAt(i))

		//autopsy.Store(fmt.Sprintf("index: %d, cost: %d", i, m))
		if m == min {
			mutation_targets = append(mutation_targets, i)
		}
		if m < min {
			min = m
			mutation_targets = make([]int, 0)
			mutation_targets = append(mutation_targets, i)
		}
	}
	if min >= Cost(quiv) {
		//s := quiv.ToMatrix()
		//autopsy.Store(s.ToString())
		return quiv, 1, nil
	}
	var old Quiver
	hit := false
	for i := 0; i < len(mutation_targets); i++ {
		q := quiv.MutateAt(mutation_targets[i])
		if reached_contains(reached, q) {
			continue
		}
		tmp, cst, err := GreedyReduction(q, reached, false)
		if err != nil {
			return tmp, cst + 1, err
		}
		if hit {
			l := old.ToMatrix()
			v := tmp.ToMatrix()
			if !QuiverIsEqual(old, tmp) {
				autopsy.Store(fmt.Sprintf("not equal\n%s\n%s\n", l.ToString(), v.ToString()))
				autopsy.Store(fmt.Sprint("l mutations: ", print_int_slice(tmp.mutations)))
				autopsy.Store(fmt.Sprint("v mutations: ", print_int_slice(old.mutations)))
				return old, cst + 1, errors.New("not equal quivers")
			}
		}
		old = tmp
		hit = true
	}
	if !hit {
		s := quiv.ToMatrix()
		autopsy.Store(s.ToString())
		return quiv, 1, nil
	}
	return old, 1, nil
}
func GreedyReduction2(quiv Quiver, reached *[]Quiver, is_initial bool) (Quiver, int, error) {
	*reached = append(*reached, quiv)
	mutation_targets := make([]int, 0)
	min := 2147483647
	quivm := quiv.ToMatrix()
	if is_initial {
		autopsy.Store(fmt.Sprintf("staring matrix:\n%s\n", quivm.ToString()))
		autopsy.Store(fmt.Sprintf("starting cost: %d\n", Cost(quiv)))
	}

	for i := 0; i < 4; i++ {
		m := Cost(quiv.MutateAt(i))

		//autopsy.Store(fmt.Sprintf("index: %d, cost: %d", i, m))
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
		//s := quiv.ToMatrix()
		//autopsy.Store(s.ToString())
		return quiv, 1, nil
	}
	var old Quiver
	hit := false
	for i := 0; i < len(mutation_targets); i++ {
		q := quiv.MutateAt(mutation_targets[i])
		if reached_contains(reached, q) {
			continue
		}
		tmp, cst, err := GreedyReduction2(q, reached, false)
		if err != nil {
			return tmp, cst + 1, err
		}
		if hit {
			l := old.ToMatrix()
			v := tmp.ToMatrix()
			if !QuiverIsEqual(old, tmp) {
				autopsy.Store(fmt.Sprintf("not equal\n%s\n%s\n", l.ToString(), v.ToString()))
				autopsy.Store(fmt.Sprint("l mutations: ", print_int_slice(tmp.mutations)))
				autopsy.Store(fmt.Sprint("v mutations: ", print_int_slice(old.mutations)))
				return old, cst + 1, errors.New("not equal quivers")
			}
		}
		old = tmp
		hit = true
	}
	if !hit {
		s := quiv.ToMatrix()
		autopsy.Store(s.ToString())
		return quiv, 1, nil
	}
	return old, 1, nil
}
func old_func() {
	autopsy.Init()
	//min := 0
	for i := 1; i < 100; i++ {
		for j := 0; j < (i)*100000; j++ {
			reached := make([]Quiver, 0)
			//recursed = false
			autopsy.Reset()
			q := RandomQuiver(4, i+1)
			_, _, err := GreedyReduction2(q, &reached, true)
			if err != nil {
				//min = cst
				//fmt.Fprintln(os.Stdout, os.Stdout, err.Error())
				//fmt.Fprintln(os.Stdout, "counter example:\n")
				//fmt.Fprintln(os.Stdout, m.ToString())
				autopsy.Dump()
				//fmt.Fprintln(os.Stdout, "done\n")
				os.Exit(1)
			}
		}
		println(i)
	}
}
func main() {
	old_func()
	return
	autopsy.Init()
	max := 10
	for a0 := 0; a0 < max; a0++ {
		for a1 := 0; a1 < max; a1++ {
			for a2 := 0; a2 < max; a2++ {
				for a3 := 0; a3 < max; a3++ {
					for a4 := 0; a4 < max; a4++ {
						for a5 := 0; a5 < max; a5++ {

						}
					}
				}
			}
		}
	}
}

/*
tan(9) =

*/
