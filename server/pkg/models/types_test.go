package models

import (
	"fmt"
	"testing"
)

func TestWin(t *testing.T) {
	i := Item{Name: "Miami", Score: 2, Rounds: 1}
	i.Win()
	expected := Item{Name: "Miami", Score: 4, Rounds: 2}
	if i != expected {
		t.Errorf("Got %+v, want %+v", i, expected)
	}
}

func ExampleRanking_Final() {
	r := Ranking{}
	r.Final("Timbuktu")
	fmt.Println(r.RankingsList)

	// Output:
	// [Timbuktu]
}

func TestAddHolder(t *testing.T) {
	r := Ranking{}
	item := &Item{Name: "Rome", Score: -2, Rounds: 3}
	r.AddHolder(item)

	if len(r.RankingsHolder) != 1 {
		t.Errorf("expected 1 item, got %d", len(r.RankingsHolder))
	}
}

func ExampleRanking_AddHolder() {
	r := &Ranking{}
	item := Item{Name: "Rome", Score: -2, Rounds: 3}
	r.AddHolder(&item)

	for _, i := range r.RankingsHolder {
		fmt.Println(i.Name, i.Score, i.Rounds)
	}
	// Output:
	// Rome -2 3
}
