package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
	"time"
)

type Post struct {
	Filename          string
	PublishedFilename string
	Title             string
	Date              time.Time
	Categories        []string
}

var (
	notesRegex = regexp.MustCompile(`[\n]---[\W]+# Notes`)
)

var blogPosts = []Post{
	{
		Filename: "001-yet-another-software-blog.md",
		Title:    "Yet Another Software Blog",
		Date:     date("July 2, 2018"),
		// Categories: []string{"Personal"},
	},
	{
		Filename: "003-what-this-blog-is-all-about.md",
		Title:    "What This Blog is All About",
		Date:     date("July 14, 2018"),
	},
	{
		Filename: "006-database-indexes.md",
		Title:    "Breaking Down Abstractions: Database Indexes",
		Date:     date("July 23, 2018"),
		// Categories: []string{"Breaking Abstractions"},
	},
}

func main() {
	for _, post := range blogPosts {
		post.PublishedFilename = publishedFilename(post.Filename, post.Date)
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
	filenameParts := strings.Split(filename, "-")
	filenameParts = filenameParts[1:] // remove the ###- prefix

	// prepend YEAR-MONTH-DAY- prefix
	filenameParts = append([]string{
		fmt.Sprintf("%d", date.Year()),
		fmt.Sprintf("%02d", int(date.Month())),
		fmt.Sprintf("%02d", date.Day()),
	}, filenameParts...)

	return strings.Join(filenameParts, "-")
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
	removeNotes(&contents)
	addHeader(&contents, post)
	correctRelativeLinks(&contents)

	// write the post under the _posts directory, as it's now
	// ready to be published.
	os.MkdirAll("_posts", 0755)
	publishedFile, err := os.Create("_posts/" + post.PublishedFilename)
	if err != nil {
		panic(err)
	}
	defer publishedFile.Close()

	publishedFile.Write([]byte(contents))
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

func addHeader(contents *string, post Post) {
	header := fmt.Sprintf(`---
layout: post
title:  "%s"
date:   %d-%02d-%02d
categories: %s
---

`, post.Title, post.Date.Year(), int(post.Date.Month()), post.Date.Day(), strings.Join(post.Categories, " "))

	*contents = header + *contents
}

func correctRelativeLinks(contents *string) {
	*contents = strings.Replace(*contents, "./images/", "../images/", -1)
}
