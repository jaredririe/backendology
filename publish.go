package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

type Post struct {
	Filename          string
	PublishedFilename string
	Slug              string
	Title             string
	Date              time.Time
	Categories        []string
	Tags              []string
}

const (
	postDirectory = "content/post/"
)

var (
	notesRegex = regexp.MustCompile(`[\n]---[\W]+# Notes`)
	titleRegex = regexp.MustCompile(`^#[\s]+(.+)\n`)
)

var blogPosts = []Post{
	{
		Filename:   "finished/yet-another-software-blog.md",
		Date:       date("July 2, 2018"),
		Categories: []string{"Personal"},
	},
	{
		Filename:   "finished/what-this-blog-is-all-about.md",
		Date:       date("July 14, 2018"),
		Categories: []string{"General"},
	},
	{
		Filename:   "finished/database-indexes.md",
		Date:       date("July 23, 2018"),
		Categories: []string{"Breaking Abstractions"},
		Tags:       []string{"databases"},
	},
	{
		Filename:   "finished/top-software-books.md",
		Date:       date("July 30, 2018"),
		Categories: []string{"Technical Books"},
	},
}

func main() {
	for _, post := range blogPosts {
		post.PublishedFilename = publishedFilename(post.Filename, post.Date)
		post.Slug = slug(post.Filename)
		processPost(post)
	}
}

// date expects a date string formatted like 'January 2, 2006'
// and parses this format into a time.Time struct.
func date(str string) time.Time {
	layout := "January 2, 2006"
	date, err := time.Parse(layout, str)
	if err != nil {
		panic(err)
	}

	return date
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
	addHeader(&contents, post)

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

func addHeader(contents *string, post Post) {
	header := fmt.Sprintf(`---
title: "%s"
author: "Jared Ririe"
categories: "%s"
date: %d-%02d-%02d
slug: %s`,
		post.Title,
		strings.Join(post.Categories, " "),
		post.Date.Year(), int(post.Date.Month()), post.Date.Day(),
		post.Slug,
	)

	if len(post.Tags) > 0 {
		tags := "\ntags:"
		for _, tag := range post.Tags {
			tags += "\n- " + tag
		}
		header += tags
	}

	header += "\n---\n\n"

	*contents = header + *contents
}
