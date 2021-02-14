package common

type Entry struct {
	Key   string
	Value string
}

type Repository struct {
	entries map[string]string
}

var instance = Repository{
	entries: map[string]string{},
}

func Repo() Repository {
	return instance
}

func (r Repository) Add(entry Entry) {
	r.entries[entry.Key] = entry.Value
}

func (r Repository) Search(query string) (result []Entry) {
	exactMatch := r.entries[query]
	return append(result, Entry{query, exactMatch})
}
