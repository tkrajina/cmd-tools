package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"sort"
	"time"
)

type Data struct {
	SubredditID           string        `json:"subreddit_id"`
	ApprovedAtUtc         interface{}   `json:"approved_at_utc"`
	Edited                interface{}   `json:"edited"`
	ModReasonBy           interface{}   `json:"mod_reason_by"`
	BannedBy              interface{}   `json:"banned_by"`
	AuthorFlairType       string        `json:"author_flair_type"`
	RemovalReason         interface{}   `json:"removal_reason"`
	LinkID                string        `json:"link_id"`
	AuthorFlairTemplateID interface{}   `json:"author_flair_template_id"`
	Likes                 interface{}   `json:"likes"`
	Replies               string        `json:"replies"`
	UserReports           []interface{} `json:"user_reports"`
	Saved                 bool          `json:"saved"`
	ID                    string        `json:"id"`
	BannedAtUtc           interface{}   `json:"banned_at_utc"`
	ModReasonTitle        interface{}   `json:"mod_reason_title"`
	Gilded                int           `json:"gilded"`
	Archived              bool          `json:"archived"`
	NoFollow              bool          `json:"no_follow"`
	Author                string        `json:"author"`
	NumComments           int           `json:"num_comments"`
	CanModPost            bool          `json:"can_mod_post"`
	SendReplies           bool          `json:"send_replies"`
	ParentID              string        `json:"parent_id"`
	Score                 int           `json:"score"`
	AuthorFullname        string        `json:"author_fullname"`
	Over18                bool          `json:"over_18"`
	ApprovedBy            interface{}   `json:"approved_by"`
	ModNote               interface{}   `json:"mod_note"`
	Collapsed             bool          `json:"collapsed"`
	Body                  string        `json:"body"`
	LinkTitle             string        `json:"link_title"`
	AuthorFlairCSSClass   string        `json:"author_flair_css_class"`
	Name                  string        `json:"name"`
	AuthorPatreonFlair    bool          `json:"author_patreon_flair"`
	Downs                 int           `json:"downs"`
	AuthorFlairRichtext   []interface{} `json:"author_flair_richtext"`
	IsSubmitter           bool          `json:"is_submitter"`
	BodyHTML              string        `json:"body_html"`
	Gildings              struct {
		Gid1 int `json:"gid_1"`
		Gid2 int `json:"gid_2"`
		Gid3 int `json:"gid_3"`
	} `json:"gildings"`
	CollapsedReason            interface{}   `json:"collapsed_reason"`
	Distinguished              interface{}   `json:"distinguished"`
	Stickied                   bool          `json:"stickied"`
	CanGild                    bool          `json:"can_gild"`
	Subreddit                  string        `json:"subreddit"`
	AuthorFlairTextColor       string        `json:"author_flair_text_color"`
	ScoreHidden                bool          `json:"score_hidden"`
	Permalink                  string        `json:"permalink"`
	NumReports                 interface{}   `json:"num_reports"`
	LinkPermalink              string        `json:"link_permalink"`
	ReportReasons              interface{}   `json:"report_reasons"`
	LinkAuthor                 string        `json:"link_author"`
	AuthorFlairText            string        `json:"author_flair_text"`
	LinkURL                    string        `json:"link_url"`
	Created                    float64       `json:"created"`
	CreatedUtc                 float64       `json:"created_utc"`
	SubredditNamePrefixed      string        `json:"subreddit_name_prefixed"`
	Controversiality           int           `json:"controversiality"`
	AuthorFlairBackgroundColor string        `json:"author_flair_background_color"`
	ModReports                 []interface{} `json:"mod_reports"`
	Quarantine                 bool          `json:"quarantine"`
	SubredditType              string        `json:"subreddit_type"`
	Ups                        int           `json:"ups"`
}

type Datas []Data

var _ sort.Interface = Datas([]Data{})

func (d Datas) Len() int           { return len(d) }
func (d Datas) Less(i, j int) bool { return d[i].Created < d[j].Created }
func (d Datas) Swap(i, j int)      { d[i], d[j] = d[j], d[i] }

type Posts struct {
	Kind string `json:"kind"`
	Data struct {
		Modhash  string `json:"modhash"`
		Dist     int    `json:"dist"`
		Children []struct {
			Kind string `json:"kind"`
			Data Data   `json:"data"`
		} `json:"children"`
		After  interface{} `json:"after"`
		Before interface{} `json:"before"`
	} `json:"data"`
}

func panicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	flag.Parse()
	byts, err := ioutil.ReadFile(flag.Arg(0))
	panicIfErr(err)

	var posts Posts
	panicIfErr(json.Unmarshal(byts, &posts))

	datas := []Data{}
	for _, ch := range posts.Data.Children {
		if ch.Kind == "t1" {
			datas = append(datas, ch.Data)
		}
	}

	sort.Sort(Datas(datas))

	var res bytes.Buffer
	res.WriteString("# Reddit doc\n\n")
	for _, d := range datas {
		title := d.LinkTitle
		desc := ""
		if len(title) > 50 {
			desc = title
			title = title[:50] + "..."
		}
		res.WriteString("## " + title)
		res.WriteString("\n\n")
		if title != "" {
			res.WriteString("> " + desc)
		}
		res.WriteString("\n\n")
		res.WriteString("* " + time.Unix(int64(d.CreatedUtc), 0).String() + "\n")
		res.WriteString("* " + d.LinkAuthor + "\n")
		res.WriteString("* " + d.LinkPermalink + "\n")
		res.WriteString("\n\n")
		res.WriteString(d.Body)
		res.WriteString("\n\n")
	}
	fmt.Println(res.String())
}
