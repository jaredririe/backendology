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
		FeaturedImage: "/public/images/profile-2019.png",
		Date:          date("July 2, 2018"),
		Category:      category.Personal,
		Tags:          []tagType{tag.Qualtrics, tag.BYU},
	},
	{
		Filename:      "../finished/what-this-blog-is-all-about.md",
		FeaturedImage: "/public/images/topics-outline.png",
		Date:          date("July 14, 2018"),
		Category:      category.General,
		Tags:          []tagType{tag.Microservices, tag.Golang, tag.Architecture, tag.Databases, tag.NoSQL},
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
		Tags:          []tagType{tag.Microservices, tag.Architecture, tag.Golang, tag.Javascript, tag.CodingInterview, tag.Databases, tag.NoSQL, tag.MachineLearning},
	},
	{
		Filename:      "../finished/lessons-from-adopting-go-qualtrics.md",
		FeaturedImage: "/public/images/utgo-qualtrics.png",
		Date:          date("August 6, 2018"),
		Category:      category.Golang,
		Tags:          []tagType{tag.Golang, tag.Qualtrics},
	},
	{
		Filename:      "../finished/experience-using-hugo.md",
		FeaturedImage: "/public/images/hugo.png",
		Date:          date("August 13, 2018"),
		Category:      category.Golang,
		Tags:          []tagType{tag.Golang, tag.Blogging, tag.Hugo},
	},
	{
		Filename:      "../finished/microservices-huge-mistake.md",
		FeaturedImage: "/public/images/microservices-entanglement.png",
		Date:          date("August 21, 2018"),
		Category:      category.Architecture,
		Tags:          []tagType{tag.Architecture, tag.Microservices},
	},
	{
		Filename:      "../finished/multiple-layers-caching.md",
		FeaturedImage: "/public/images/caching-worker.png",
		Date:          date("August 27, 2018"),
		Category:      category.Architecture,
		Tags:          []tagType{tag.Architecture, tag.Caching, tag.Microservices, tag.NoSQL},
	},
	{
		Filename:      "../finished/hugo-newsletter.md",
		FeaturedImage: "/public/images/newsletter-12.png",
		Date:          date("August 31, 2018"),
		Category:      category.General,
		Tags:          []tagType{tag.Hugo, tag.Blogging},
	},
	{
		Filename:      "../finished/distributed-systems-course-reading-list.md",
		FeaturedImage: "/public/images/distributed-systems-paper.png",
		Date:          date("September 10, 2018"),
		Category:      category.DistributedSystems,
		Tags:          []tagType{tag.ResearchPapers, tag.Architecture, tag.Microservices, tag.NoSQL},
	},
	{
		Filename:      "../finished/distributed-systems-course-introduction.md",
		FeaturedImage: "/public/images/distributed-systems-in-the-wild.png",
		Date:          date("September 19, 2018"),
		Category:      category.DistributedSystems,
		Tags:          []tagType{tag.ResearchPapers, tag.Architecture, tag.Microservices, tag.NoSQL, tag.Databases},
	},
	{
		Filename:      "../finished/distributed-systems-building-blocks.md",
		FeaturedImage: "/public/images/building-blocks-architecture-open-source.png",
		Date:          date("October 8, 2018"),
		Category:      category.Architecture,
		Tags:          []tagType{tag.Architecture, tag.Caching, tag.Microservices, tag.Golang},
	},
	{
		Filename:      "../finished/alexa-skill-lambda-golang.md",
		FeaturedImage: "/public/images/alexa-logo.png",
		Date:          date("October 27, 2018"),
		Category:      category.Golang,
		Tags:          []tagType{tag.Golang, tag.Microservices},
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
		DistributedSystems   categoryType
		Architecture         categoryType
	}{
		"Personal",
		"General",
		"Breaking Abstractions",
		"Books",
		"Golang",
		"Distributed Systems",
		"Architecture",
	}

	tag = struct {
		Databases       tagType
		Architecture    tagType
		Microservices   tagType
		ResearchPapers  tagType
		CodingInterview tagType
		DataStructures  tagType
		MachineLearning tagType
		Caching         tagType
		NoSQL           tagType
		SQL             tagType
		Golang          tagType
		Javascript      tagType
		Qualtrics       tagType
		BYU             tagType
		Blogging        tagType
		Hugo            tagType
	}{
		"databases",
		"architecture",
		"microservices",
		"research-papers",
		"coding-interview",
		"data-structures",
		"machine-learning",
		"caching",
		"nosql",
		"sql",
		"golang",
		"javascript",
		"qualtrics",
		"byu",
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
