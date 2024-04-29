package main

import (
	LA "BruteForce/LA"
	"BruteForce/utils"
	"math/rand"
)

type Quiver struct {
	points     []vertex_t
	num_points int
	mutations  []int
}

func (q *Quiver) MutateAt(point int) Quiver {
	m := utils.Clone[int](q.mutations)
	return Quiver{mutate(q.points, q.num_points, point), q.num_points, append(m, point)}
}
func (q *Quiver) ToMatrix() LA.Matrix {
	return make_matrix_from_quiver(q.points, q.num_points)
}
func AddCost(quiv Quiver) int {
	cost := 0
	for i := 0; i < quiv.num_points; i++ {
		for j := 0; j < quiv.num_points; j++ {
			if quiv.points[i].edges[j] > 0 {
				cost += quiv.points[i].edges[j]
			}
		}
	}
	return cost
}
func MaxCost(quiv Quiver) int {
	cost := 0
	for i := 0; i < quiv.num_points; i++ {
		for j := 0; j < quiv.num_points; j++ {
			if quiv.points[i].edges[j] > cost {
				cost = quiv.points[i].edges[j]
			}
		}
	}
	return cost
}
func Cost(quiv Quiver) int {
	return AddCost(quiv)
}

func RandomQuiver(dim int, max int) Quiver {
	out := Quiver{make([]vertex_t, dim), dim, make([]int, 0)}
	for i := 0; i < dim; i++ {
		for j := 0; j < dim; j++ {
			if j == i {
				continue
			}
			r := rand.Int()%(max*2) - max
			out.points[i].edges[j] = r
			out.points[j].edges[i] = -r
		}
	}
	return out
}

type vertx struct {
	edges []int
	used  bool
}

func v_is_equal(a vertx, b vertx) bool {
	for i := 0; i < len(a.edges); i++ {
		if a.edges[i] != b.edges[i] {
			return false
		}
	}
	return true
}
func QuiverStrictlyEqual(aq Quiver, bq Quiver) bool {
	if aq.num_points != bq.num_points {
		return false
	}
	for i := 0; i < aq.num_points; i++ {
		for j := 0; j < aq.num_points; j++ {
			if aq.points[i].edges[j] != bq.points[i].edges[j] {
				return false
			}
		}
	}
	return true
}

func QuiverIsEqual(aq Quiver, bq Quiver) bool {
	if aq.num_points != bq.num_points {
		return false
	}
	/*mp := [][]int{{0, 1, 2, 3}, {0, 1, 3, 2}, {0, 2, 1, 3}, {0, 2, 3, 1},
	{0, 3, 2, 1}, {0, 3, 1, 2}, {1, 0, 2, 3}, {1, 0, 3, 2}, {1, 2, 0, 3},
	{1, 2, 3, 0}, {1, 3, 2, 0}, {1, 3, 0, 2}, {2, 1, 0, 3}, {2, 1, 3, 0},
	{2, 0, 1, 3}, {2, 0, 3, 1}, {2, 3, 0, 1}, {2, 3, 1, 0}, {3, 1, 2, 0},
	{3, 1, 0, 2}, {3, 2, 1, 0}, {3, 2, 0, 1}, {3, 0, 2, 1}, {3, 0, 1, 2}}
	*/
	//mp := [][]int{{0, 1, 2}, {0, 2, 1}, {2, 1, 0}, {2, 0, 1}, {1, 2, 0}, {1, 0, 2}}
	max := 0
	for i := 0; i < aq.num_points; i++ {
		for j := 0; j < aq.num_points; j++ {
			if aq.points[i].edges[j] > max {
				max = aq.points[i].edges[j]
			}
		}
	}
	for i := 0; i < bq.num_points; i++ {
		for j := 0; j < bq.num_points; j++ {
			if bq.points[i].edges[j] > max {
				max = bq.points[i].edges[j]
			}
		}
	}
	acounts := make([]int, max+1)
	bcounts := make([]int, max+1)
	for i := 0; i < aq.num_points; i++ {
		for j := 0; j < aq.num_points; j++ {
			if aq.points[i].edges[j] > 0 {
				acounts[aq.points[i].edges[j]] += 1
			}
		}
	}
	for i := 0; i < bq.num_points; i++ {
		for j := 0; j < bq.num_points; j++ {
			if bq.points[i].edges[j] > 0 {
				bcounts[bq.points[i].edges[j]] += 1
			}
		}
	}
	for i := 0; i < len(acounts); i++ {
		if acounts[i] != bcounts[i] {
			return false
		}
	}
	return true
}
