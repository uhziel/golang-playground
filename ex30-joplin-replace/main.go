package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/zeromicro/go-zero/rest/httpc"
)

const (
	BaseURL = "http://localhost:41184"
	Token   = "dc9dd18b2ea101b5aee131441bf7f6e71a7ddf8886ad0d111ddaed2212f91d28a828d546885a05dc6f17f860133d1d05e6ccb1bfa8b5b17ad73c32c5d022e106"
)

var (
	Limit         int
	MaxNum        int
	Text, Replace string
)

func init() {
	flag.IntVar(
		&Limit,
		"limit",
		100,
		"the number of items to be returned(the maximum being 100 items)",
	)
	flag.IntVar(&MaxNum, "maxnum", 0, "the max num of items")
	flag.StringVar(&Text, "text", " ", "the searched text")
	flag.StringVar(&Replace, "replace", "", "")
	flag.Parse()
}

func main() {
	ctx := context.Background()

	items, err := GetNotes(ctx, Limit, MaxNum)
	if err != nil {
		panic(err)
	}

	fmt.Println("len:", len(items))
	replacedNum := 0
	for _, item := range items {
		replaced := strings.ReplaceAll(item.Body, Text, Replace)
		if replaced == item.Body {
			continue
		}

		replacedNum++
		if len(Replace) == 0 {
			//fmt.Printf("original: %#v\n", item)
			fmt.Printf("original id=%s title=%#v\n", item.ID, item.Title)
		} else {
			//fmt.Printf("replaced: %#v", replaced)
			err := APIModifyNote(ctx, item.ID, replaced, item.UserUpdatedTime)
			if err != nil {
				panic(err)
			} else {
				fmt.Printf("replace success. id=%s title=%#v\n", item.ID, item.Title)
			}
		}
	}

	if len(Replace) == 0 {
		fmt.Println("willReplacedNum:", replacedNum)
	} else {
		fmt.Println("replacedNum:", replacedNum)
	}
}

func Search(ctx context.Context, query string, pageLimit int, maxItemNum int) ([]Item, error) {
	items := []Item{}
	page := 1
Outter:
	for {
		resp, err := APISearch(ctx, query, pageLimit, page)
		if err != nil {
			return nil, fmt.Errorf("Search fail: %w", err)
		}

		for _, item := range resp.Items {
			items = append(items, item)
			if maxItemNum > 0 && len(items) >= maxItemNum {
				break Outter
			}
		}

		if !resp.HasMore {
			break
		}
		page += 1
	}

	return items, nil
}

func GetNotes(ctx context.Context, pageLimit int, maxItemNum int) ([]Item, error) {
	items := []Item{}
	page := 1
Outter:
	for {
		resp, err := APIGetNotes(ctx, pageLimit, page)
		if err != nil {
			return nil, fmt.Errorf("GetNotes fail: %w", err)
		}

		for _, item := range resp.Items {
			items = append(items, item)
			if maxItemNum > 0 && len(items) >= maxItemNum {
				break Outter
			}
		}

		if !resp.HasMore {
			break
		}
		page += 1
	}

	return items, nil
}

// API
type SearchReq struct {
	Fields   string `form:"fields"`
	Limit    int    `form:"limit"`
	Page     int    `form:"page"`
	OrderBy  string `form:"order_by"`
	OrderDir string `form:"order_dir"`
	Query    string `form:"query"`
	Token    string `form:"token"`
	Type     string `form:"type"`
}

type SearchResp struct {
	Items   []Item `json:"items"`
	HasMore bool   `json:"has_more,optional"`
}

type Item struct {
	ID              string `json:"id"`
	Title           string `json:"title"`
	Body            string `json:"body"`
	UserUpdatedTime int64  `json:"user_updated_time"`
}

func APISearch(ctx context.Context, query string, limit int, page int) (*SearchResp, error) {
	req := SearchReq{
		Fields:   "id,title,body,user_updated_time",
		Limit:    limit,
		Page:     page,
		OrderBy:  "user_updated_time",
		OrderDir: "DESC",
		Query:    query,
		Token:    Token,
		Type:     "note",
	}

	resp, err := httpc.Do(ctx, http.MethodGet, BaseURL+"/search", req)
	if err != nil {
		return nil, fmt.Errorf("APISearch request fail: %w", err)
	}

	if resp.StatusCode >= 400 {
		var builder strings.Builder
		io.Copy(&builder, resp.Body)
		return nil, fmt.Errorf(
			"APISearch status fail StatusCode: %s Body: %s",
			resp.Status,
			builder.String(),
		)
	}

	searchResp := SearchResp{}
	if err := httpc.Parse(resp, &searchResp); err != nil {
		return nil, fmt.Errorf("APISearch parse fail: %w", err)
	}

	return &searchResp, nil
}

type ModifyNoteReq struct {
	ID              string `path:"id"`
	Token           string `          form:"token"`
	Body            string `                       json:"body"`
	UserUpdatedTime int64  `                       json:"user_updated_time"`
}

func APIModifyNote(ctx context.Context, id, body string, userUpdatedTime int64) error {
	req := ModifyNoteReq{
		ID:              id,
		Token:           Token,
		Body:            body,
		UserUpdatedTime: userUpdatedTime,
	}

	resp, err := httpc.Do(ctx, http.MethodPut, BaseURL+"/notes/:id", req)
	if err != nil {
		return fmt.Errorf("APIModifyNote request fail: %w", err)
	}

	if resp.StatusCode >= 400 {
		var builder strings.Builder
		io.Copy(&builder, resp.Body)
		return fmt.Errorf(
			"APIModifyNote status fail StatusCode: %s Body: %s",
			resp.Status,
			builder.String(),
		)
	}
	return nil
}

type GetNotesReq struct {
	Fields   string `form:"fields"`
	Limit    int    `form:"limit"`
	Page     int    `form:"page"`
	OrderBy  string `form:"order_by"`
	OrderDir string `form:"order_dir"`
	Token    string `form:"token"`
}

type GetNotesResp struct {
	Items   []Item `json:"items"`
	HasMore bool   `json:"has_more,optional"`
}

func APIGetNotes(ctx context.Context, limit, page int) (*GetNotesResp, error) {
	req := GetNotesReq{
		Fields:   "id,title,body,user_updated_time",
		Limit:    limit,
		Page:     page,
		OrderBy:  "user_updated_time",
		OrderDir: "DESC",
		Token:    Token,
	}

	resp, err := httpc.Do(ctx, http.MethodGet, BaseURL+"/notes", req)
	if err != nil {
		return nil, fmt.Errorf("APIGetNotes request fail: %w", err)
	}

	if resp.StatusCode >= 400 {
		var builder strings.Builder
		io.Copy(&builder, resp.Body)
		return nil, fmt.Errorf(
			"APIGetNotes status fail StatusCode: %s Body: %s",
			resp.Status,
			builder.String(),
		)
	}

	getNotesResp := GetNotesResp{}
	if err := httpc.Parse(resp, &getNotesResp); err != nil {
		return nil, fmt.Errorf("APIGetNotes parse fail: %w", err)
	}

	return &getNotesResp, nil
}
