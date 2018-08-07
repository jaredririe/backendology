package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"gopkg.in/yaml.v2"
)

type frontMatter struct {
	Title         string       `yaml:"title"`
	FeaturedImage string       `yaml:"featured_image"`
	Author        string       `yaml:"author"`
	Categories    categoryType `yaml:"categories"`
	Tags          []tagType    `yaml:"tags,omitempty"`
	Date          string       `yaml:"date"`
	Slug          string       `yaml:"slug"`
}

const (
	postDirectory = "../content/post/"
)

var (
	notesRegex = regexp.MustCompile(`[\n]---[\W]+# Notes`)
	titleRegex = regexp.MustCompile(`^#[\s]+(.+)\n`)
)

func main() {
	for _, post := range blogPosts {
		post.PublishedFilename = publishedFilename(post.Filename, post.Date)
		post.Slug = slug(post.Filename)
		processPost(post)
	}
}

// publishedFilename transforms the filename into
// the pulished form (YEAR-MONTH-DAY-name-of-post.md).
func publishedFilename(filename string, date time.Time) string {
	// get just the file's name (without path)
	filename = filepath.Base(filename)

	// prepend YEAR-MONTH-DAY- prefix
	filenameParts := strings.Split(filename, "-")
	filenameParts = append([]string{
		fmt.Sprintf("%d", date.Year()),
		fmt.Sprintf("%02d", int(date.Month())),
		fmt.Sprintf("%02d", date.Day()),
	}, filenameParts...)

	return strings.Join(filenameParts, "-")
}

func slug(filename string) string {
	// get just the file's name (without path)
	filename = filepath.Base(filename)

	return strings.TrimSuffix(filename, filepath.Ext(filename))
}

func processPost(post Post) {
	fmt.Println("Processing", post.PublishedFilename)

	// extract contents of original file
	bytes, err := ioutil.ReadFile(post.Filename)
	if err != nil {
		panic(err)
	}

	contents := string(bytes)

	// apply several transformations to the content
	post.Title = extractTitle(&contents)
	removeNotes(&contents)
	correctRelativeImages(&contents)
	addFrontMatter(&contents, post)

	// write the post under the _posts directory, as it's now
	// ready to be published.
	os.MkdirAll(postDirectory, 0755)
	publishedFile, err := os.Create(postDirectory + post.PublishedFilename)
	if err != nil {
		panic(err)
	}
	defer publishedFile.Close()

	publishedFile.Write([]byte(contents))
}

func extractTitle(contents *string) string {
	var title string

	matches := titleRegex.FindAllStringSubmatchIndex(*contents, -1)
	if matches != nil && len(matches) > 0 && len(matches[0]) >= 4 {
		// extract title
		titleStart, titleEnd := matches[0][2], matches[0][3]
		title = (*contents)[titleStart:titleEnd]

		// remove heading (which contained the title)
		headingStart, headingEnd := matches[0][0], matches[0][1]
		*contents = (*contents)[0:headingStart] + (*contents)[headingEnd+1:]
	}

	return title
}

// removeNotes removes any notes at the end of the blog post.
// The expected notes format is as follows:
//
// ---
//
// # Notes
func removeNotes(contents *string) {
	matches := notesRegex.FindAllStringIndex(*contents, -1)
	if matches != nil {
		notesSectionStart := matches[len(matches)-1][0]
		*contents = (*contents)[:notesSectionStart]
	}
}

func correctRelativeImages(contents *string) {
	*contents = strings.Replace(*contents, "../static/public/images", "/public/images", -1)
}

func addFrontMatter(contents *string, post Post) {
	fm := frontMatter{
		Title:         post.Title,
		FeaturedImage: post.FeaturedImage,
		Author:        "Jared Ririe",
		Categories:    post.Category,
		Tags:          post.Tags,
		Date:          fmt.Sprintf("%d-%02d-%02d", post.Date.Year(), int(post.Date.Month()), post.Date.Day()),
		Slug:          post.Slug,
	}

	yaml, err := yaml.Marshal(&fm)
	if err != nil {
		panic(err)
	}

	yamlStr := fmt.Sprintf("---\n%s---\n\n", string(yaml))
	*contents = yamlStr + *contents
}
