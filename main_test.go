package main

import "testing"

func TestPredSucc(t *testing.T) {
	root := tn(8, tn(3, tn(1, nil, nil), tn(6, tn(4, nil, nil), tn(7, nil, nil))), tn(10, nil, tn(14, tn(13, nil, nil), nil)))
	cases := []struct {
		key        int
		pred, succ int // -1 means nil expected
	}{
		{8, 7, 10},
		{1, -1, 3},   // minimum: no predecessor
		{14, 13, -1}, // maximum: no successor
		{5, 4, 6},    // key absent: neighbors of where it would sit
		{0, -1, 1},   // below everything
		{99, 14, -1}, // above everything
		{6, 4, 7},
	}
	for _, c := range cases {
		p, s := PredSucc(root, c.key)
		if (c.pred == -1) != (p == nil) || (p != nil && p.Val != c.pred) {
			t.Errorf("PredSucc(%d) pred = %v, want %d", c.key, p, c.pred)
		}
		if (c.succ == -1) != (s == nil) || (s != nil && s.Val != c.succ) {
			t.Errorf("PredSucc(%d) succ = %v, want %d", c.key, s, c.succ)
		}
	}
	p, s := PredSucc(nil, 5)
	if p != nil || s != nil {
		t.Error("empty tree should give nil, nil")
	}
}
