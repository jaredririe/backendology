# Implementing an Amazon Alexa Skill with an AWS Lambda Function in Go

Amazon is currently [running a promotion](https://build.amazonalexadev.com/echodot.html) in which developers who publish a new skill to the Alexa Skills Store receive the new Echo Dot. While I do not own any Amazon devices, I decided to participate to get something for free, learn more about developing voice-based applications, and write my first AWS Lambda function in Go.

Before I dive into the implementation, I think Amazon deserves some praise for this promotion. It's a win-win and I think more companies should learn from this. Getting to the point of publishing a skill in the Alexa Skills Store requires a developer to go through the entire process: set up an Amazon Developer Account, learn Alexa terminology like invocations and intents, set up an Amazon Web Services account, write an AWS Lambda function that understands how to parse Alexa's requests and generate valid responses, and submit the skill for review (with a customer-ready description and icon). While the expectation for the skills for the competition is not high, with a little incentive, Amazon will get an impressive number of novel skills submitted to their store, making their product better and more comprehensive. Amazon is also managing to get at least a few hours of work for the cost it takes them to make an Echo Dot (which is assuredly less than the $50 price tag). The developer gets an interesting side project, diverse development experience, and a $50 consumer product that can run the code he or she writes. Very cool!

## Alexa skill: Apple Buyer's Guide

The simple skill I wrote for the purpose of the promotion is called the **Apple Buyer's Guide**. It offers a convenient way to check whether it's a good time to buy a new Apple product. Through a real-time look at the [MacRumors Buyer's Guide](https://buyersguide.macrumors.com/), this skill allows Alexa to tell you which of four states an Apple product is in: {updated, neutral, caution, and outdated}. A product in the caution state, for example, has not been updated for quite some time, so it may be wise to be patient and wait for a new update. The status updated, on the other hand, means that the Apple product was just updated and you're safe to go ahead with the purchase.

To create my skill, I followed two tutorials. The first, ["How To Build A Custom Amazon Alexa Skill, Step-By-Step: My Favorite Chess Player"](https://medium.com/crowdbotics/how-to-build-a-custom-amazon-alexa-skill-step-by-step-my-favorite-chess-player-dcc0edae53fb), shows you how how to create a skill from beginning to end. I highly recommend following this kind of tutorial to get a skill working as it's crucial to have a high-level overview to know what you're signing yourself up to create, what configuration is possible, etc.

The first tutorial gives a great high-level overview, but the AWS Lambda function is written in Python and is not explained in much detail. As I wanted to write my Lambda function in Go, I supplemented with this second tutorial, ["Alexa Skills with Go"](https://medium.com/@amalec/alexa-skills-with-go-54db0c21e758) which covers the following topics:

* Automating deployment of Go-based AWS Lambda functions
* Handling a variety of Alexa Skill request attributes
* Creating and returning Alexa Skill responses
* Deploying code to AWS Lambda without manually zipping the binary created by Go

## Configuring the Alexa Skill

![Alexa skill intents](../static/public/images/alexa-skill-intents.png)

(Overview of configuring the skill, including intents)

## Writing the AWS Lambda Function

![AWS Lambda](../static/public/images/alexa-skill-aws-lambda.png)

(Basics of Lambda functions and the serverless movement)

(Lambda configuration, implementation in Go)

(Alexa request/response)

## Demonstration

![Alexa Skill Demonstration](../static/public/images/alexa-skill-demonstration.png)

---

# Notes (supplementary to blog post)

https://medium.com/crowdbotics/how-to-build-a-custom-amazon-alexa-skill-step-by-step-my-favorite-chess-player-dcc0edae53fb

* Create a simple Amazon Alexa Skill step-by-step

https://medium.com/@amalec/alexa-skills-with-go-54db0c21e758

* Automate deployment of Go-based AWS Lambdas
* Create and return Alexa Skill Responses
* Handle a variety of Alexa Skill Request attributes (Locale and Attributes)

https://docs.aws.amazon.com/lambda/latest/dg/go-programming-model-handler-types.html

> A Lambda function written in Go is authored as a Go executable. In your Lambda function code, you need to include the github.com/aws/aws-lambda-go/lambda package, which implements the Lambda programming model for Go. In addition, you need to implement handler function code and a main() function.
