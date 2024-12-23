package day23

import (
	"testing"
)

var input = []byte(`kh-tc
qp-kh
de-cg
ka-co
yn-aq
qp-ub
cg-tb
vc-aq
tb-ka
wh-tc
yn-cg
kh-ub
ta-co
de-co
tc-td
tb-wq
wh-td
ta-ka
td-qp
aq-cg
wq-ub
ub-vc
de-ta
wq-aq
wq-vc
wh-yn
ka-de
kh-ta
co-tc
wh-qp
tb-vc
td-yn`)

func TestPart1(t *testing.T) {
	want := 7
	graph, err := parseData(input)
	if err != nil {
		t.Fatal(err)
	}
	got := part1(graph)
	if got != want {
		t.Errorf("part1(%s) = %d, want %d", input, got, want)
	}
}

func TestPart2(t *testing.T) {
	want := "co,de,ka,ta"
	graph, err := parseData(input)
	if err != nil {
		t.Fatal(err)
	}
	got := part2(graph)
	if got != want {
		t.Errorf("part2(%s) = %s, want %s", input, got, want)
	}
}
