package main

import "github.com/golddranks/TiraLabra/src/langinspect/builder"
import "github.com/golddranks/TiraLabra/src/langinspect/viterbi"
import "github.com/golddranks/TiraLabra/src/orderedmap/naivetrie"
import "testing"
import "fmt"

func printLangTable(db *naivetrie.Node, ngram string) {
	lTable := db.TryAndGet([]byte(ngram)).(builder.LangTable)
	lTable2 := db.Value.(builder.LangTable)
	for i, l := range lTable {
		fmt.Println(ngram, "langtable: ", i, builder.LangIndexToTag(builder.LangIndex(i)), l, float32(int(10000*float32(l)/float32(lTable2[i])))/100)
	}
	fmt.Println()
}

func hasErrors(path []int, expected_path []int, text []byte) bool {
	for i, p := range path {
		if expected_path[i] != p {
			fmt.Println("Got:")
			printPath(text, path)
			fmt.Println("Should be:")
			printPath(text, expected_path)
			fmt.Println()
			return true
		}
	}
	return false
}

func TestInspect(t *testing.T) {
	db := builder.Build("testdata")
	states := builder.AmountOfLangs // TODO Ugly, ugly global value. Need to refactor.
	start_prob := []float64{0.33333, 0.33333, 0.33333}
	trans_prob := [][]float64{
		[]float64{0.96, 0.02, 0.02},
		[]float64{0.02, 0.96, 0.02},
		[]float64{0.02, 0.02, 0.96},
	}
	emit_prob := viterbi.GetEmitProbFunction(db)
	var text []byte
	var path []int
	var expected_path []int
	text = []byte("joopa joo")
	path = viterbi.Run(slice(text), states, trans_prob, start_prob, emit_prob)
	expected_path = []int{1, 1, 1, 1, 1, 1, 1, 1, 1}
	if hasErrors(path, expected_path, text) {
		t.Fail()
	}

	text = []byte("dumb")
	path = viterbi.Run(slice(text), states, trans_prob, start_prob, emit_prob)
	expected_path = []int{0, 0, 0, 0}
	if hasErrors(path, expected_path, text) {
		t.Fail()
	}

	text = []byte("jack")
	path = viterbi.Run(slice(text), states, trans_prob, start_prob, emit_prob)
	expected_path = []int{1, 1, 1, 1}
	if hasErrors(path, expected_path, text) {
		t.Fail()
	}

	text = []byte("jac")
	path = viterbi.Run(slice(text), states, trans_prob, start_prob, emit_prob)
	expected_path = []int{0, 0, 0, 0}
	if hasErrors(path, expected_path, text) {
		t.Fail()
	}

}
