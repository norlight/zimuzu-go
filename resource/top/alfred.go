package top

import "github.com/norlight/zimuzu-go/alfred"
import "fmt"

func (o *TopOut) AlfResult() (r alfred.Result) {
	for i, v := range o.Data {
		var st string
		switch v.Channel {
		case "tv":
			switch v.Area {
			case "美国":
				st = "美剧"
			case "英国":
				st = "英剧"
			case "日本":
				st = "日剧"
			case "韩国":
				st = "韩剧"
			default:
				st = "影视剧"
			}

		case "movie":
			st = fmt.Sprintf("电影 [%s]", v.Area)
		case "openclass":
			st = "公开课"
		case "documentary":
			st = "纪录片"
		default:
			st = "影视剧"
		}

		item := alfred.Item{
			Title:     fmt.Sprintf("TOP%d《%s》%s", i+1, v.Cnname, v.PublishYear),
			Subtitle:  fmt.Sprintf("%s [%s] [%s]", st, v.PlayStatus, v.Category),
			Arg:       fmt.Sprintf("%s:%s", v.Channel, v.ID),
			Variables: map[string]string{"cnname": v.Cnname},
		}
		r.Append(item)
	}

	return
}
