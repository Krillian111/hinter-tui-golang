package common

import "fmt"

type Tab int

const (
	add = iota
	search
)

var tabLabels = map[Tab]string{
	add:    "Add",
	search: "Search",
}
var tabByValue = map[int]Tab{
	0: add,
	1: search,
}

func (t Tab) String() string {
	label, ok := tabLabels[t]
	if !ok {
		panic(fmt.Sprintf("Unknown tab: %d", int(t)))
	}
	return label
}

func FromIndex(index int) Tab {
	tab, ok := tabByValue[index]
	if !ok {
		panic(fmt.Sprintf("Unknown tab: %d", index))
	}
	return tab
}

func AddTab() Tab {
	return add
}
func SearchTab() Tab {
	return search
}

func AllTabs() []Tab {
	return []Tab{add, search}
}
