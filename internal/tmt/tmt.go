package tmt

import (
	"sort"
	"strings"
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
	repl := strings.NewReplacer(
		"Javascript", "JavaScript",
		"Typescript", "TypeScript",
		"ReactJS", "React",
		"htmx", "HTMX",
		"Golang", "Go",
		".NET", "DotNET",
		"c++", "CPP",
		"C++", "CPP",
		// "/", "/",
		"+", "/",
	)
	for _, v := range migrations {
		splits := strings.Split(v.Description, " ")
		from := splits[1]
		from = repl.Replace(from)
		fromSplitted := strings.Split(from, "/")
		to := splits[3]
		to = repl.Replace(to)
		toSplitted := strings.Split(to, "/")
		for _, v := range fromSplitted {
			if v == "?" {
				continue
			}
			f := frommap[v]
			f++
			frommap[v] = f
		}
		for _, v := range toSplitted {
			if v == "?" {
				continue
			}
			t := tomap[v]
			t++
			tomap[v] = t
		}
		for _, v := range fromSplitted {
			for _, w := range toSplitted {
				ft := v + " " + w
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
