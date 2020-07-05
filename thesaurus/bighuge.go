package thesaurus

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type BigHuge struct {
	APIKey string
}

type words struct {
	Syn []string `json:"syn"`
}

type synonyms struct {
	Noun *words `json:"noun"`
	Verb *words `json:"verb"`
}

func (b *BigHuge) Synonyms(term string) ([]string, error) {
	var syns []string
	url := "http://words.bighugelabs.com/api/2/" + b.APIKey + "/" + term + "/json"
	response, err := http.Get(url)
	if err != nil {
		println(url)
		return syns, fmt.Errorf("bighuge: %qの類語検索に失敗しました: %v", term, err)
	}
	var data synonyms
	defer response.Body.Close()
	var r io.Reader = response.Body
	//r = io.TeeReader(r, os.Stderr)
	if err := json.NewDecoder(r).Decode(&data); err != nil {
		return syns, err
	}
	syns = append(syns, data.Noun.Syn...)
	syns = append(syns, data.Verb.Syn...)
	return syns, nil
}
