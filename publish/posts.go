package main

import "time"

type Post struct {
	Filename          string
	FeaturedImage     string
	PublishedFilename string
	Slug              string
	Title             string
	Date              time.Time
	Category          categoryType
	Tags              []tagType
}

var blogPosts = []Post{
	{
		Filename:      "../finished/yet-another-software-blog.md",
		FeaturedImage: "/public/images/DSC_0186.jpg",
		Date:          date("July 2, 2018"),
		Category:      category.Personal,
		Tags:          []tagType{tag.Software, tag.Qualtrics, tag.BYU},
	},
	{
		Filename:      "../finished/what-this-blog-is-all-about.md",
		FeaturedImage: "/public/images/topics-outline.png",
		Date:          date("July 14, 2018"),
		Category:      category.General,
		Tags:          []tagType{tag.Microservices, tag.Architecture, tag.DistributedSystems, tag.Databases, tag.NoSQL, tag.Golang, tag.Software},
	},
	{
		Filename:      "../finished/database-indexes.md",
		FeaturedImage: "/public/images/b+-tree.png",
		Date:          date("July 23, 2018"),
		Category:      category.BreakingAbstractions,
		Tags:          []tagType{tag.Databases, tag.SQL, tag.NoSQL, tag.DataStructures},
	},
	{
		Filename:      "../finished/top-software-books.md",
		FeaturedImage: "/public/images/top-software-books.jpg",
		Date:          date("July 30, 2018"),
		Category:      category.Books,
		Tags:          []tagType{tag.Microservices, tag.Architecture, tag.CodingInterview, tag.Golang, tag.Javascript, tag.MachineLearning, tag.Databases, tag.NoSQL, tag.Software},
	},
	{
		Filename:      "../finished/lessons-from-adopting-go-qualtrics.md",
		FeaturedImage: "/public/images/utgo-qualtrics.png",
		Date:          date("August 6, 2018"),
		Category:      category.Golang,
		Tags:          []tagType{tag.Golang, tag.Qualtrics, tag.Software},
	},
	{
		Filename:      "../finished/experience-using-hugo.md",
		FeaturedImage: "/public/images/hugo.png",
		Date:          date("August 13, 2018"),
		Category:      category.Golang,
		Tags:          []tagType{tag.Golang, tag.Blogging, tag.Hugo},
	},
}

type (
	categoryType string
	tagType      string
)

var (
	category = struct {
		Personal             categoryType
		General              categoryType
		BreakingAbstractions categoryType
		Books                categoryType
		Golang               categoryType
		Research             categoryType
	}{
		"Personal",
		"General",
		"Breaking Abstractions",
		"Books",
		"Golang",
		"Research",
	}

	tag = struct {
		Databases          tagType
		Architecture       tagType
		Microservices      tagType
		DistributedSystems tagType
		CodingInterview    tagType
		DataStructures     tagType
		MachineLearning    tagType
		NoSQL              tagType
		SQL                tagType
		Golang             tagType
		Javascript         tagType
		Qualtrics          tagType
		BYU                tagType
		Software           tagType
		Blogging           tagType
		Hugo               tagType
	}{
		"databases",
		"architecture",
		"microservices",
		"distributed-systems",
		"coding-interview",
		"data-structures",
		"machine-learning",
		"nosql",
		"sql",
		"golang",
		"javascript",
		"qualtrics",
		"byu",
		"software",
		"blogging",
		"hugo",
	}
)

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
