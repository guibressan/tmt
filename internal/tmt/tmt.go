package tmt

import (
	"sort"
	"sync"
)

var (
	processedData ProcessedData
	processOnce   sync.Once
)

type ProcessedData struct {
	Year   []migrationYear
	From   []StrCount
	To     []StrCount
	FromTo []StrCount
}

type migrationYear struct {
	Year       int16
	Migrations []migration
}

type StrCount struct {
	Name  string
	Count int16
}

func Process() ProcessedData {
	processOnce.Do(processLazy)
	return processedData
}

func processLazy() {
	frommap := map[string]int16{}
	tomap := map[string]int16{}
	fromtomap := map[string]int16{}
	yearmap := map[int16][]migration{}
	for _, v := range migrations {
		yearmig := yearmap[int16(v.Year)]
		yearmig = append(yearmig, v)
		yearmap[int16(v.Year)] = yearmig
	}
	for _, v := range migrations {
		for _, w := range v.From {
			if w == "unknown" {
				continue
			}
			f := frommap[Tech[w]]
			f++
			frommap[Tech[w]] = f
		}
		for _, w := range v.To {
			if w == "unknown" {
				continue
			}
			t := tomap[Tech[w]]
			t++
			tomap[Tech[w]] = t
		}
		for _, w := range v.From {
			for _, x := range v.To {
				ft := Tech[w] + " " + Tech[x]
				ftCount := fromtomap[ft]
				ftCount++
				fromtomap[ft] = ftCount
			}
		}
	}
	from := make([]StrCount, 0, len(frommap))
	to := make([]StrCount, 0, len(tomap))
	fromto := make([]StrCount, 0, len(fromtomap))
	years := []migrationYear{}
	for k, v := range frommap {
		from = append(from, StrCount{k, v})
	}
	for k, v := range tomap {
		to = append(to, StrCount{k, v})
	}
	for k, v := range fromtomap {
		fromto = append(fromto, StrCount{k, v})
	}
	for k, v := range yearmap {
		years = append(years, migrationYear{
			Year:       k,
			Migrations: v,
		})
	}
	s := func(slice []StrCount) {
		sort.Slice(slice, func(i, j int) bool {
			return slice[i].Count > slice[j].Count
		})
	}
	s(from)
	s(to)
	s(fromto)
	sort.Slice(years, func(i, j int) bool {
		return years[i].Year < years[j].Year
	})
	processedData = ProcessedData{
		Year:   years,
		From:   from,
		To:     to,
		FromTo: fromto,
	}
}
