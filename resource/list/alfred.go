package list

import (
	"fmt"

	"strconv"

	"github.com/norlight/zimuzu-go/alfred"
)

func (o *ItemListOut) AlfResult() (r alfred.Result) {
	for _, v := range o.Data {
		item := alfred.Item{
			Title:    v.Name,
			Subtitle: v.Size,
		}
		r.Append(item)
	}
	return
}

func (v *Item) AlfItem() (alf alfred.Item) {
	var title string
	var size string
	if s, _ := strconv.Atoi(v.Season); s == 0 {
		title = fmt.Sprintf("%s %s", v.Format, v.Name)
	} else {
		title = fmt.Sprintf("S%sE%s %s %s", v.Season, v.Episode, v.Format, v.Name)
	}
	if s, err := strconv.Atoi(v.Size); err == nil && s == 0 {
		size = "不提供下载"
	} else {
		size = v.Size
	}

	item := alfred.Item{
		Title:    title,
		Subtitle: size,
		Mods:     make(map[string]*alfred.Mod),
	}
	for _, str := range []string{"cmd", "alt", "ctrl", "shift"} {
		item.Mods[str] = &alfred.Mod{
			Valid:    false,
			Subtitle: v.Size,
			Arg:      "",
		}
	}

	//临时借用字段存储下格式和季度信息，传出去后再修改
	//等Alfred 3.4.1出来Item支持变量就不用这么麻烦了
	item.Mods["ctrl"] = &alfred.Mod{
		Valid:    false,
		Subtitle: v.Size,
		Arg:      v.Format,
	}
	item.Mods["shift"] = &alfred.Mod{
		Valid:    false,
		Subtitle: v.Episode,
		Arg:      v.Season,
	}
	for _, link := range v.Link {
		switch link.Way {
		case "1":
			item.Arg = link.Address
			item.Mods["cmd"] = &alfred.Mod{
				Valid:    true,
				Subtitle: "复制[电驴]链接到剪贴板",
				Arg:      link.Address,
			}
		case "2":
			item.Mods["alt"] = &alfred.Mod{
				Valid:    true,
				Subtitle: "复制[磁力]链接到剪贴板",
				Arg:      link.Address,
			}
		}
	}

	return item
}
