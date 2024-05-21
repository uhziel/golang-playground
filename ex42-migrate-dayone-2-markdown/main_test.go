package main

import (
	"reflect"
	"strings"
	"testing"
	"time"
)

const (
	input0 = `	Date:	September 23, 2014 at 13:34
	Tags:	皮鞋
	Starred

第一次使用 Day One

`
	input1 = `	Date:	September 23, 2014 at 13:38
	Location:	北京市市辖区, 北京市, 中国
`
	input2 = `	Date:	October 31, 2014 at 23:44
	Photo:	2014-10-31.jpg

`
	input3 = `	Date:	January 20, 2018 at 19:50
	Location:	建国路75号, 北京市, 中国
	Weather:	-1° Mostly Clear
	Photo:	2018-1-20.jpg

# 今天发年终奖

今年的年终奖比去年少一个月。我的总资产终于超过100万了。

2017年我个人没有任何进步，看到别人的进步，我确实没有努力了，得改过。
`

	input4 = `	Date:	September 23, 2014 at 13:34

第一次使用 Day One

	Date:	September 23, 2014 at 13:38
	Location:	北京市市辖区, 北京市, 中国



	Date:	October 31, 2014 at 23:44
	Photo:	2014-10-31.jpg



	Date:	January 20, 2018 at 19:50
	Location:	建国路75号, 北京市, 中国
	Weather:	-1° Mostly Clear
	Photo:	2018-1-20.jpg

# 今天发年终奖

今年的年终奖比去年少一个月。我的总资产终于超过100万了。

2017年我个人没有任何进步，看到别人的进步，我确实没有努力了，得改过。

`
)

func TestImportDayOne(t *testing.T) {
	tests := []struct {
		input string
		want  []DiaryEntry
	}{
		{
			input: input0,
			want: []DiaryEntry{
				{
					Date: time.Date(2014, time.September, 23, 13, 34, 0, 0, beijingTZ),
					Content: []string{
						"第一次使用 Day One",
					},
				},
			},
		},
		{
			input: input1,
			want: []DiaryEntry{
				{
					Date:     time.Date(2014, time.September, 23, 13, 38, 0, 0, beijingTZ),
					Location: "北京市市辖区, 北京市, 中国",
				},
			},
		},
		{
			input: input2,
			want: []DiaryEntry{
				{
					Date:  time.Date(2014, time.October, 31, 23, 44, 0, 0, beijingTZ),
					Photo: "2014-10-31.jpg",
				},
			},
		},
		{
			input: input3,
			want: []DiaryEntry{
				{
					Date:     time.Date(2018, time.January, 20, 19, 50, 0, 0, beijingTZ),
					Location: "建国路75号, 北京市, 中国",
					Weather:  "-1° Mostly Clear",
					Photo:    "2018-1-20.jpg",
					Content: []string{
						"# 今天发年终奖",
						"",
						"今年的年终奖比去年少一个月。我的总资产终于超过100万了。",
						"",
						"2017年我个人没有任何进步，看到别人的进步，我确实没有努力了，得改过。",
					},
				},
			},
		},
		{
			input: input4,
			want: []DiaryEntry{
				{
					Date: time.Date(2014, time.September, 23, 13, 34, 0, 0, beijingTZ),
					Content: []string{
						"第一次使用 Day One",
					},
				},
				{
					Date:     time.Date(2014, time.September, 23, 13, 38, 0, 0, beijingTZ),
					Location: "北京市市辖区, 北京市, 中国",
					Content: []string{
						"",
					},
				},
				{
					Date:  time.Date(2014, time.October, 31, 23, 44, 0, 0, beijingTZ),
					Photo: "2014-10-31.jpg",
					Content: []string{
						"",
					},
				},
				{
					Date:     time.Date(2018, time.January, 20, 19, 50, 0, 0, beijingTZ),
					Location: "建国路75号, 北京市, 中国",
					Weather:  "-1° Mostly Clear",
					Photo:    "2018-1-20.jpg",
					Content: []string{
						"# 今天发年终奖",
						"",
						"今年的年终奖比去年少一个月。我的总资产终于超过100万了。",
						"",
						"2017年我个人没有任何进步，看到别人的进步，我确实没有努力了，得改过。",
					},
				},
			},
		},
	}

	for i, test := range tests {
		inputStream := strings.NewReader(test.input)
		diaryEntries, err := importDayOne(inputStream)
		if err != nil {
			t.Errorf("#%d err=%#v", i, err)
			continue
		}
		if !reflect.DeepEqual(diaryEntries, test.want) {
			t.Errorf("#%d \nreal %#v\nwant %#v", i, diaryEntries, test.want)
		}
	}
}

const content0 = `---
title: 2014-09-23 13:34
updated: 2014-09-23T13:34:00+08:00
created: 2014-09-23T13:34:00+08:00
---

第一次使用 Day One`

const content1 = `---
title: 2014-09-23 13:34
updated: 2014-09-23T13:34:00+08:00
created: 2014-09-23T13:34:00+08:00
location: 建国路75号, 北京市, 中国
Weather: -1° Mostly Clear
---

![2014-10-26 (1).jpg](../../assets/2014-10-26%20%281%29.jpg)

第一次使用 Day One`

func TestExportDayOne(t *testing.T) {
	tests := []struct {
		input []DiaryEntry
		want  []FileEntry
	}{
		{
			input: []DiaryEntry{
				{
					Date: time.Date(2014, time.September, 23, 13, 34, 0, 0, beijingTZ),
					Content: []string{
						"第一次使用 Day One",
					},
				},
			},
			want: []FileEntry{
				{
					FileName: "2014-09-23 13:34.md",
					Content:  content0,
				},
			},
		},
		{
			input: []DiaryEntry{
				{
					Date:     time.Date(2014, time.September, 23, 13, 34, 0, 0, beijingTZ),
					Location: "建国路75号, 北京市, 中国",
					Weather:  "-1° Mostly Clear",
					Photo:    "2014-10-26 (1).jpg",
					Content: []string{
						"第一次使用 Day One",
					},
				},
			},
			want: []FileEntry{
				{
					FileName: "2014-09-23 13:34.md",
					Content:  content1,
				},
			},
		},
	}

	for i, test := range tests {
		result := exportDayOne(test.input)
		if !reflect.DeepEqual(result, test.want) {
			t.Errorf("#%d \nreal %#v\nwant %#v", i, result, test.want)
		}
	}
}
