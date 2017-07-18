package search

import (
	"encoding/json"
	"io/ioutil"

	"fmt"
	"strconv"

	"github.com/norlight/zimuzu-go/alfred"
	"github.com/norlight/zimuzu-go/utils"
)

func (o *SearchOut) AlfResult() (r alfred.Result) {
	for _, v := range o.Data.List {
		if v.Type == "resource" {
			unix, _ := strconv.Atoi(v.Uptime)
			s := fmt.Sprintf("更新：%s", utils.FormatUnix(int64(unix)))
			item := alfred.Item{
				Title:    v.Title,
				Subtitle: s,
				Arg:      v.Itemid,
			}
			r.Append(item)
		}
	}
	return
}

func (s *Search) AlfSearch(k string) (result []byte, err error) {
	resp, err := s.Search(k)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var out SearchOut
	if err := json.Unmarshal(body, &out); err != nil {
		return nil, err
	}
	as := out.AlfResult()
	result, err = json.Marshal(as)
	if err != nil {
		return nil, err
	}
	return
}
