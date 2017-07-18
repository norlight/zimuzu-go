package resource

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"sync"

	"github.com/norlight/zimuzu-go/alfred"
	"github.com/norlight/zimuzu-go/resource/list"
	"github.com/norlight/zimuzu-go/resource/top"
)

func (l *List) AlfResult() (r alfred.Result) {
	for _, v := range l.Data.List {
		item := alfred.Item{
			Title:    v.CNName,
			Subtitle: fmt.Sprintf("%s%s", v.Itemupdate, v.Remark),
		}
		r.Append(item)
	}
	return
}

func (r *Resource) AlfFetchList() {
	resp, _ := r.FetchList("", "", "", "", "", "", "")
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	var info List
	json.Unmarshal(body, &info)
	//log.Println(string(body))
	//log.Println(info)
	as := info.AlfResult()
	b, _ := json.Marshal(as)
	fmt.Print(string(b))
	os.Exit(0)

}

func (r *Resource) AlfTop() (result []byte, err error) {
	resp, err := r.Top("", "11")
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var out top.TopOut
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

func (r *Resource) AlfList(chID string) (result []byte, err error) {
	channelRe, _ := regexp.Compile(`[a-zA-Z]+`)
	idRe, _ := regexp.Compile(`\d+`)
	channel := channelRe.FindString(chID)
	id := idRe.FindString(chID)
	switch channel {
	case "tv":
		return r.AlfTV(id)
	case "movie":
		return r.AlfMovie(id)
	default:
		return r.AlfTV(id)
	}
}
func (r *Resource) AlfTV(id string) (result []byte, err error) {
	resp, err := r.List(id)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var out list.SEOut
	if err := json.Unmarshal(body, &out); err != nil {
		return nil, err
	}

	eps := make(chan []alfred.Item)
	var wg sync.WaitGroup
	for i := 0; i < len(out.Data); i++ {
		se := out.Data[i]
		var count int
		fmt.Sscanf(se.Episode, "%d", &count)

		for j := 0; j < count; j++ {
			wg.Add(1)
			go func(i, j int) {
				defer wg.Done()
				sStr := strconv.Itoa(i + 1)
				eStr := strconv.Itoa(j + 1)
				data, err := r.AlfItemList(id, sStr, eStr)
				if err != nil {
					log.Println(err)
					return
				}
				eps <- data
			}(i, j)
		}
	}

	go func() {
		wg.Wait()
		close(eps)
	}()

	var alfresult alfred.Result

	//信道传过来的为包含各集所有格式解析成AltItem的slice
	//取出来放到一块方便遍历处理
	allitem := make([]alfred.Item, 0)
	for items := range eps {
		allitem = append(allitem, items...)
	}
	sort.Sort(byEpisode(allitem))

	//提取Item对应季度对应格式的所有链接
	for _, item := range allitem {
		format := item.Mods["ctrl"].Arg
		season := item.Mods["shift"].Arg

		var ed2ks string
		var magnets string

		for _, itemagain := range allitem {
			formatagain := itemagain.Mods["ctrl"].Arg
			seasonagain := itemagain.Mods["shift"].Arg
			if formatagain == format && seasonagain == season {
				ed2k := itemagain.Mods["cmd"].Arg
				magnet := itemagain.Mods["alt"].Arg
				ed2ks += fmt.Sprintln(ed2k)
				magnets += fmt.Sprintln(magnet)
			}
		}

		//更新AltItem，使按下相应键可以复制所有链接
		item.Mods["ctrl"] = &alfred.Mod{
			Valid:    true,
			Subtitle: fmt.Sprintf("复制全季[%s][电驴]链接", format),
			Arg:      ed2ks,
		}
		item.Mods["shift"] = &alfred.Mod{
			Valid:    true,
			Subtitle: fmt.Sprintf("复制全季[%s][磁力]链接", format),
			Arg:      magnets,
		}

		alfresult.Append(item)
	}

	result, err = json.Marshal(alfresult)
	return
}

func (r *Resource) AlfMovie(id string) (result []byte, err error) {
	resp, err := r.ItemList(id, "", "", "1")
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var out list.ItemListOut
	if err := json.Unmarshal(body, &out); err != nil {
		return nil, err
	}

	var alfresult alfred.Result
	for _, v := range out.Data {
		item := v.AlfItem()
		//将被用于临时存储的字段改回来
		item.Mods["shift"].Subtitle = item.Mods["ctrl"].Subtitle
		alfresult.Append(item)
	}
	result, err = json.Marshal(alfresult)
	return
}

func (r *Resource) AlfItemList(id, season, episode string) (slice []alfred.Item, err error) {
	resp, err := r.ItemList(id, season, episode, "1")
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var out list.ItemListOut
	if err := json.Unmarshal(body, &out); err != nil {
		return nil, err
	}
	for _, v := range out.Data {

		item := v.AlfItem()

		slice = append(slice, item)
	}
	return
}

//排序
type byEpisode []alfred.Item

func (x byEpisode) Len() int {
	return len(x)
}
func (x byEpisode) Less(i, j int) bool {
	itemi := x[i]
	itemj := x[j]
	eiStr := itemi.Mods["shift"].Subtitle
	ejStr := itemj.Mods["shift"].Subtitle
	ei, _ := strconv.Atoi(eiStr)
	ej, _ := strconv.Atoi(ejStr)
	return ei < ej
}
func (x byEpisode) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}
