---
title: My Experience Using Hugo
featured_image: /public/images/hugo.png
author: Jared Ririe
categories: Golang
tags:
- golang
- blogging
- hugo
date: "2018-08-13"
slug: experience-using-hugo
---

As my blog is still in its infancy, I have been exploring various blogging platforms and tools to find the ones that I would like to use moving forward. This post introduces Hugo, considers its alternatives, and walks through an example of setting up a Hugo blog and publishing a post.

## What is Hugo?

![Hugo](/public/images/hugo.png)

[Hugo](https://gohugo.io/) is an open-source static site generator known for its performance and flexibility. Hugo was initially released in 2013 and has quickly grown in popularity. It is written in Go which is what pushed me to give it a shot.

Static sites take the content (usually markdown in the case of blogs) and apply it to layouts or templates, ultimately generating static HTML files that can be served to the user. They make a lot of sense for websites like blogs that don't offer a dynamic experience that requires custom JavaScript to run on the user's web browser. As the HTML is served directly, static sites are highly performant and more secure.

### Alternatives to Hugo

Let's take a look at some alternatives to Hugo which will help us learn more about it and help us make a better decision.

#### Wordpress

Wordpress is a fully-featured and proven website creation platform. It has everything you need to write a blog, including plenty of free themes, plugins to share your new posts via social media providers (and automatically send an email version of your post to those who subscribe), and usable (but not great) markdown support. The websites it produces are *not* static sites which makes Wordpress a common target for malicious attacks.

It's worth noting here that [wordpress.com](https://wordpress.com/) is different from [wordpress.org](https://wordpress.org/). This confused me at first! wordpress.com is a solution built on top of the open source software which is Wordpress (wordpress.org). wordpress.com offers things like free and paid hosting options and a simpler GUI for writing posts.

When I first started writing this blog, I chose to host my site under a free wordpress.com subdomain--this removed the barrier to start blogging in terms of both effort and cost. I was still settling on a name for my blog anyway, so I didn't want to pay for a domain quite yet. If I choose Wordpress over Hugo going forward, I will use the open source version directly and buy my own domain. The conveniences offered by wordpress.com are not valuable enough to me for their relatively high cost.

Wordpress has some clear disadvantages in my mind. Foremost, its target audience is not a software engineer who thinks some of the challenges of blogging might be fun to solve as a form of side project. Because it's so fully featured, there aren't that many problems left to solve. I think someone who doesn't write software might find it strange I'm listing this as a negative! Perhaps fellow developers will agree that it would be fun to learn how to integrate with a commenting system or learn the Twitter API well enough to update followers of new posts.

Even as I write this, though, I recognize that needing to solve these problems will slow me down from getting content to users. There are arguments on both sides which make it a tough call!

#### Jekyll

[Jekyll](https://jekyllrb.com/) was started by GitHub’s founder Tom Preston-Werner in 2008. It is more directly comparable to Hugo as it's primarily a static site generator, not a full blogging solution. In fact, Jekyll is the software that started the current excitement around static site generation.

As Jekyll installs as a RubyGem, it does require that you install and configure a Ruby environment on your computer. I did not already have everything set up for Ruby, so it was a bit of a hassle to get the `jekyll` command to work. The set up for Hugo was a breeze in comparison as Brew just downloaded the Go binary that was immediately runnable. While I have Go installed for other reasons, this is unnecessary.

Jekyll has been around longer than Hugo (2008 vs. 2013) and therefore is known to be more extensible through plugins. Hugo is still under heavy development, so new features are being added all the time. When I added Google Analytics to my site, for example, some guides showed how to do it manually but then newer ones pointed out built-in support.

Hugo is the clear winner in terms of the speed of static site generation.[^2] This is noticeable when you're actively changing the site's contents and Hugo is able to easily keep up with your changes in real time. Jekyll also supports live reload like this, but some files require a full rebuild (such as the changes to the config). Hugo does not have this limitation.

## Setting up a simple blog with Hugo

Without further ado, let's jump into a real world example of using Hugo (my blog!). You could follow these steps to learn how to make a simple blog generated with Hugo. Most of this content is adapted from [this tutorial](https://gohugo.io/hosting-and-deployment/hosting-on-github/).

### Install Hugo

Install Hugo by following the steps outlined [here](https://gohugo.io/getting-started/installing/). On a Mac, it is as simple as `brew install hugo`.

### Create GitHub repositories and submodule

GitHub is a great way to store the contents of your blog to keep it under version control. Several of the guides I followed recommended that you make two repositories, one for the site and another for the generated site that will actually be hosted.

Make the first repository for the content of the blog with any name you prefer. I chose `backendology`. ([github.com/jaredririe/backendology](https://github.com/jaredririe/backendology))

Now, make the second repository named `<username>.github.io` for the generated site: ([github.com/jaredririe/jaredririe.github.io](https://github.com/jaredririe/jaredririe.github.io)).

With both repositories in place, run this command to make the `public` directory a submodule for the second repository.

> Submodules allow you to keep a Git repository as a subdirectory of another Git repository. This lets you clone another repository into your project and keep your commits separate.[^1]

```bash
git submodule add -b master https://github.com/<username>/<username>.github.io.git public
```

To make it easy to generate your site and push changes to the second repository, write a script like the following:

```bash
#!/bin/bash

echo -e "\033[0;32mDeploying updates to GitHub...\033[0m"

# Generate the static site in the default /public folder.
hugo

# Add changes to the git submodule.
cd public
git add .

# Commit changes.
msg="Rebuilding site on `date`"
if [ $# -eq 1 ]
  then msg="$1"
fi
git commit -m "$msg"

# Push changes to origin.
git push origin master

# Come back up to the project's root.
cd ..
```

The script runs the `hugo` command which by default overwrites the public directory. It pushes the changes to GitHub with a default message or the one specified as an argument to the script.

### Copy someone else's repository using Hugo

Rather than starting from an empty repo, I recommend copying someone else's who is using Hugo, getting it to work, and then modifying and customizing it as you see fit. Feel free to fork [my GitHub repo](https://github.com/jaredririe/backendology) for this purpose.

### Enable GitHub Pages

This option is found in the GitHub settings under "GitHub Pages" on the repository for the generated site.

## Deploying my Hugo site to "production"

To illustrate my workflow for deploying my site, let's follow a post from draft to finished product.

### 1. Write a draft

Begin a post as a draft in markdown. I store these in a separate directory (`drafts/`). Once I'm happy with what I have written, I move the file into a finished directory (`finished`/).

### 2. Publish the post

I consider publishing to be the step which turns a post you are finished writing into one that is ready for Hugo. One could choose to write posts directly in this Hugo format, but I chose not to for a few reasons:

* I include a notes section at the bottom of my posts that's intended to be used for research while writing the post, not for the users reading my blog. I could just manually delete it, but I might find the raw content or quotes to be useful at some later point.
* Hugo expects the post to be named according to its publish date (e.g. `2018-07-23-database-indexes.md`) but I don't know that information until I'm done writing.
* Each post includes "front matter" (a few lines of JSON/YAML/TOML at the top of the markdown) that I think is better to automatically generate:

```yaml
---
title: 'Breaking Down Abstractions: Database Indexes'
featured_image: /public/images/b+-tree.png
author: Jared Ririe
categories: Breaking Abstractions
tags:
- databases
- sql
- nosql
- data-structures
date: "2018-07-23"
slug: database-indexes
---
```

If I wrote the front matter directly, I could accidentally make a typo when assigning the category which would have the effect of making a new category. With my categories and tags defined in code, it's easy to find ones I've previously used through auto-completion. I keep track of my categories in a Go struct:

```go
type categoryType string

var category = struct {
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
```

[Here](https://github.com/jaredririe/backendology/tree/master/publish) is the Go code I use to publish a post. It makes a copy of the post in the `finished/` directory, makes the changes described (strips off the notes, renames the file, and adds front matter) and writes to `content/post`.

### 3. Generate the site locally

Run the command `hugo server` to start up a server running at localhost:1313. Take a look at the new post: Did the markdown render as you expected? Are the images displaying correctly?

### 4. Run the deploy script

Once you're happy, run the deploy script mentioned in the initial setup to push your changes to production. Within a few minutes, your site should be deployed to `https://<username>.github.io/` or a custom domain (if you configured it this way).

[^1]: https://git-scm.com/book/en/v2/Git-Tools-Submodules
[^2]: https://forestry.io/blog/hugo-vs-jekyll-benchmark/