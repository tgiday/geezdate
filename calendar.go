// Package geezdate implements functions for converting Gregorian calendar to Geez calander.
package geezdate

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/tgiday/mgn2"
)

const (
	delta     = -2810
	leapday   = 11
	leapmonth = 9
)

type Gdate struct {
	d, m, y int
}

// Geezday return a geez date (eg "2024-01-09" to "፴ ታኅሣሥ ፳፻፲፮") taking Gregorian calander date of format yyyy-mm-dd
func Geezday(date string) Gdate {
	s := Convert(date)
	lst := strings.Split(s, "-")
	g := Gdate{}
	g.d, _ = strconv.Atoi(lst[0])
	g.m, _ = strconv.Atoi(lst[1])
	g.y, _ = strconv.Atoi(lst[2])
	return g
}

// Today return a todays date according to Geez calander
func Today() Gdate {
	t := time.Now()
	ls := strings.Split(t.String(), " ")
	td := ls[0]
	s := Convert(td)
	lst := strings.Split(s, "-")
	g := Gdate{}
	g.d, _ = strconv.Atoi(lst[0])
	g.m, _ = strconv.Atoi(lst[1])
	g.y, _ = strconv.Atoi(lst[2])

	return g
}

func (g Gdate) String() string {
	month := []string{"መስከረም", "ጥቅምት", "ኅዳር", "ታኅሣሥ", "ጥር", "የካቲት", "መጋቢት", "ሚያዝያ", "ግንቦት", "ሰኔ", "ሐምሌ", "ነሐሴ", "ጳጉሜ"}
	d := mgn2.Fmtint(g.d)
	m := month[g.m-1]
	y := mgn2.Fmtint(g.y)
	str := fmt.Sprintf("%s %s %s", d, m, y)
	return str
}

// daysBefore[m] counts the number of days in a non-leap year
// before month m begins. There is an entry for m=13, counting
// the number of days before Meskerem of next year (365).
var daysBefore = [...]int32{
	0,
	30,
	30 + 30,
	30 + 30 + 30,
	30 + 30 + 30 + 30,
	30 + 30 + 30 + 30 + 30,
	30 + 30 + 30 + 30 + 30 + 30,
	30 + 30 + 30 + 30 + 30 + 30 + 30,
	30 + 30 + 30 + 30 + 30 + 30 + 30 + 30,
	30 + 30 + 30 + 30 + 30 + 30 + 30 + 30 + 30,
	30 + 30 + 30 + 30 + 30 + 30 + 30 + 30 + 30 + 30,
	30 + 30 + 30 + 30 + 30 + 30 + 30 + 30 + 30 + 30 + 30,
	30 + 30 + 30 + 30 + 30 + 30 + 30 + 30 + 30 + 30 + 30 + 30,
	30 + 30 + 30 + 30 + 30 + 30 + 30 + 30 + 30 + 30 + 30 + 30 + 5,
}

// Convert return a Geez calander date ,take string Gregorian calendar date  ("1991-05-24" to 16-9-1983)
func Convert(date string) string {
	d, _ := time.Parse("2006-01-02", date)
	g := d.AddDate(0, 0, delta)
	// leap day
	if isLeap(d.Year()+1) && d.Day() == leapday && d.Month() == leapmonth {
		g = g.AddDate(0, 0, -1)
		_, mm, yy := convert(g)
		str := fmt.Sprintf("%v-%v-%v", 6, mm, yy)
		return str
	}
	//  >leap day
	if isLeap(g.Year()) {
		g = g.AddDate(0, 0, -1)
		dd, mm, yy := convert(g)
		str := fmt.Sprintf("%v-%v-%v", dd, mm, yy)
		return str
	}
	// <leap day
	dd, mm, yy := convert(g)
	str := fmt.Sprintf("%v-%v-%v", dd, mm, yy)
	return str
}

func isLeap(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}

func convert(x time.Time) (int, int, int) {
	yr := x.Year()
	day := x.YearDay() - 1
	month := int(day / 30)
	end := int(daysBefore[month+1])
	var begin int
	if day >= end {
		month++
		begin = end
	} else {
		begin = int(daysBefore[month])
	}
	month++
	day = day - begin + 1
	if month > 13 {
		month = month - 13
		return day, month, yr
	}
	return day, month, yr
}
