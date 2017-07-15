package taninari

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"regexp"
	"time"
)

const blogPostEndpoint = "https://api.amebaowndme.com/v2/public/blogPosts?siteId=18381&searchType=recent&limit=15"

type BlogPost struct {
	Meta struct {
		Code       int `json:"code"`
		Pagination struct {
			Total   int `json:"total"`
			Offset  int `json:"offset"`
			Limit   int `json:"limit"`
			Cursors struct {
				After  string `json:"after"`
				Before string `json:"before"`
			} `json:"cursors"`
		} `json:"pagination"`
	} `json:"meta"`
	Body []struct {
		ID       string `json:"id"`
		Contents []struct {
			Type   string `json:"type"`
			Format string `json:"format"`
			Value  string `json:"value"`
			Url    string `json:"url"`
		} `json:"contents"`
		PublishedURL string `json:"publishedUrl"`
		PublishedAt  string `json:"publishedAt"`
	} `json:"body"`
}

type Goroku struct {
	Text         string
	ImageURL     string
	PublishedURL string
	PublishedAt  string
}

func getBlogPosts(api string) string {
	req, err := http.NewRequest("GET", api, nil)
	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Fatal(res)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	return string(body)
}

func parseJson(jsonStr string) (*BlogPost, error) {
	jsonBytes := []byte(jsonStr)
	blogPost := new(BlogPost)

	if err := json.Unmarshal(jsonBytes, blogPost); err != nil {
		return nil, err
	}

	return blogPost, nil
}

func GetAllGorokus() []Goroku {
	gorokus := []Goroku{}

	url := blogPostEndpoint
	for {
		blogPostsStr := getBlogPosts(url)
		blogPost, err := parseJson(blogPostsStr)
		if err != nil {
			log.Fatal(err)
		}

		tagRegexp, _ := regexp.Compile("\\<[\\S\\s]+?\\>")

		for _, post := range blogPost.Body {
			goroku := Goroku{
				PublishedURL: post.PublishedURL,
				PublishedAt:  post.PublishedAt,
			}

			for _, content := range post.Contents {
				if content.Type == "text" {
					text := tagRegexp.ReplaceAllString(content.Value, "")
					goroku.Text = text
				} else if content.Type == "image" {
					goroku.ImageURL = content.Url
				}
			}

			gorokus = append(gorokus, goroku)
		}

		if len(gorokus) >= blogPost.Meta.Pagination.Total {
			break
		}

		url = blogPostEndpoint + "&cursor=" + blogPost.Meta.Pagination.Cursors.After
		time.Sleep(5 * time.Millisecond)
	}

	return gorokus
}

func GetGoroku() Goroku {
	gorokus := GetAllGorokus()

	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(gorokus))

	return gorokus[index]
}
