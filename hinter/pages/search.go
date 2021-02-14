package pages

import (
	"hinter/hinter/common"
)

type SearchModel struct {
}

func InitialSearch() SearchModel {
	return SearchModel{}
}

func (m SearchModel) View(entries []common.Entry) string {
	s := "\n"
	for i := 0; i < len(entries); i++ {
		s += entries[i].Key
		s += " | "
		s += entries[i].Value
		s += "\n"
	}
	return s
}
