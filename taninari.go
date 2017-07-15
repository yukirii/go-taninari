package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const limit = 12
const blogPostEndPoint = "https://api.amebaowndme.com/v2/public/blogPosts?siteId=18381&searchType=recent"

//&page=

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

func getBlogPosts() string {
	req, err := http.NewRequest("GET", blogPostEndPoint, nil)
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

func GetGorokus() []Goroku {
	blogPostsStr := getBlogPosts()
	blogPost, err := parseJson(blogPostsStr)
	if err != nil {
		log.Fatal(err)
	}

	gorokus := []Goroku{}

	for _, post := range blogPost.Body {
		goroku := Goroku{
			PublishedURL: post.PublishedURL,
			PublishedAt:  post.PublishedAt,
		}

		for _, content := range post.Contents {
			if content.Type == "text" {
				goroku.Text = content.Value
			} else if content.Type == "image" {
				goroku.ImageURL = content.Url
			}
		}

		gorokus = append(gorokus, goroku)
	}

	return gorokus
}

func main() {
	gorokus := GetGorokus()

	for _, goroku := range gorokus {
		fmt.Println(goroku.Text)
		if goroku.ImageURL != "" {
			fmt.Println(goroku.ImageURL)
		}
		fmt.Println(goroku.PublishedURL)
		fmt.Println(goroku.PublishedAt)
	}
}
