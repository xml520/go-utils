package layuiUtil

import "sort"

type LayuiMenuModel struct {
	ID     int              `json:"id" xorm:"pk autoincr"`
	PID    int              `json:"pid"`
	IsMenu bool             `json:"is_menu"`
	Weight int              `json:"weight" xorm:"default(100)"`
	Name   string           `json:"name"`
	Icon   string           `json:"icon,omitempty"`
	Title  string           `json:"title"`
	Jump   string           `json:"jump,omitempty"`
	List   []LayuiMenuModel `json:"list,omitempty" xorm:"-"`
}

func (l *LayuiMenuModel) Sort() {
	if len(l.List) == 0 {
		return
	}
	sort.Slice(l.List, func(i, j int) bool {
		if l.List[i].Weight == l.List[j].Weight {
			return l.List[i].ID < l.List[j].ID
		} else {
			return l.List[i].Weight < l.List[j].Weight
		}
	})
	for _, model := range l.List {
		model.Sort()
	}

}
