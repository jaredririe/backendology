---
title: Lessons From Adopting Go at Qualtrics
author: Jared Ririe
categories: Golang
tags:
- golang
- qualtrics
- software
date: "2018-08-06"
slug: lessons-from-adopting-go-qualtrics
---

![UTGO + Qualtrics](/public/images/utgo-qualtrics.png)

Back in September 2017, I teamed up with a coworker to give a presentation for the [Utah Golang User Group](http://utahgolang.com/). We chose to share the lessons we had learned as we adopted and scaled Go as one of the core programming languages at Qualtrics. Our intended audience was companies and developers who were interested in trying out Go and increasing its use within an organization. Many companies are considering Go because its popularity has only continued to rise (for good reason), so I'm excited to repost this content via my blog.

## Summary of the content

Here is a summary of the content of our presentation which will serve as a teaser of the slides and video below:

### `context.Context` Day in the life of a Qualtrics engineer

There were key software development life cycle (SDLC) problems at Qualtrics that started pushing us towards significant changes in our architecture. We adopted Docker and Consul and wrote our first microservices in Node.js.

### `reflect.Value` How/why Go works for us

Adopting microservices solved many problems, such as decreasing the size of our codebases and enabling clear assignment of ownership. Some problems remained, however, which is what encouraged us to try Go, a language we knew was proven for microservices.

### `func main()` Getting the initial buy-in

Convincing an organization to adopt a new programming language is challenging. This is what worked for us:

* Learn Go well enough to sell it (perhaps through side projects outside of work)
* Start with something small
* Stand up a Go service that compares favorably to a service it replaces
* Attain manager buy-in
* Deploy a real service into production

### `go build && ./scale` Process of driving Go adoption among engineers

Once you managed to get your initial service in production, how do you continue the momentum? We focused on lowering the learning barrier for the base case, helping developers get started and be productive in Go as quickly as possible. We did this by providing getting-started documentation, educating developers through trainings, and building tools and libraries which addressed common needs.

### `if err != nil` Things we wish we knew / things we’d do differently

We wish we would known how long it can take to garner interest around a new language--it took nearly a year to expand the number of Go developers to more than one. Also, in building shared libraries, we focused too much on libraries that would be useful to our work rather than core functionality needed by engineering at large.

### `wg.Wait()` Ideas for future work

To continue to scale Go at Qualtrics, we'd like to augment our documentation with tutorial-style guides and provide more specific training. It would also be great to offer some form of tailored service template that other developers could use as a starting point for their Go projects.

## Presentation slides and recording

[Here are the slides](https://docs.google.com/presentation/d/18JiufQTTm8GxFRb2uyg2C8RHmm5NbtWU9Se7vBtwkUY/edit?usp=sharing) we created for the presentation. We retroactively fleshed out the notes section to help the slides stand alone without the presentation. If you prefer the original video, however, the user group recorded our presentation and put together [this Youtube video](https://www.youtube.com/watch?v=8wmEL0JwHQA&feature=youtu.be).

<iframe width="560" height="315" src="https://www.youtube.com/embed/8wmEL0JwHQA" frameborder="0" allow="autoplay; encrypted-media" allowfullscreen></iframe>

