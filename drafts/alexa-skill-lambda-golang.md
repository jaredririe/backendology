# Implementing an Amazon Alexa Skill with an AWS Lambda Function in Go

Amazon is currently [running a promotion](https://build.amazonalexadev.com/echodot.html) in which developers who publish a new skill to the Alexa Skills Store receive the new Echo Dot. While I do not own any Amazon devices, I decided to participate to get something for free, learn more about developing voice-based applications, and write my first AWS Lambda function in Go.

Before I dive into the implementation, I think Amazon deserves some praise for this promotion. It's a win-win that more companies should learn from. Getting to the point of publishing a skill in the Alexa Skills Store requires a developer to go through the entire process: set up an Amazon Developer Account, learn Alexa terminology like invocations and intents, set up an Amazon Web Services account, write an AWS Lambda function that understands how to parse Alexa's requests and generate valid responses, and submit the skill for review (with a customer-ready description and icon). While the expectation for the skills for the competition is not high, with a little incentive, Amazon will get an impressive number of novel skills submitted to their store, making their product better and more comprehensive. Amazon is also managing to get at least a few hours of work for the cost it takes them to make an Echo Dot (which is likely less than the $50 price tag). The developer gets an interesting side project, diverse development experience, and a consumer product that can run the code he or she writes. Very cool!

## Alexa skill: Picking Apples

The simple skill I wrote for the purpose of the promotion is called the **Picking Apples**. It offers a convenient way to check whether it's a good time to buy a new Apple product. Through a real-time look at the [MacRumors Buyer's Guide](https://buyersguide.macrumors.com/), this skill allows Alexa to tell you which of four states an Apple product is in: {updated, neutral, caution, and outdated}. A product in the _caution_ state, for example, has not been updated for quite some time, so it may be wise to be patient and wait for a new update. The status _updated_, on the other hand, means that the Apple product was just updated and you're safe to go ahead with the purchase.

To create my skill, I followed two tutorials. The first, ["How To Build A Custom Amazon Alexa Skill, Step-By-Step: My Favorite Chess Player"](https://medium.com/crowdbotics/how-to-build-a-custom-amazon-alexa-skill-step-by-step-my-favorite-chess-player-dcc0edae53fb), shows you how how to create a skill from beginning to end. I highly recommend following this kind of tutorial to get a skill working as it's crucial to have a high-level overview to know what you're signing yourself up to create, what configuration is possible, etc.

The first tutorial gives a great high-level overview, but the AWS Lambda function is written in Python and is not explained in much detail. As I wanted to write my Lambda function in Go, I supplemented with this second tutorial, ["Alexa Skills with Go"](https://medium.com/@amalec/alexa-skills-with-go-54db0c21e758) which covers the following topics:

* Automating deployment of Go-based AWS Lambda functions
* Handling a variety of Alexa Skill request attributes
* Creating and returning Alexa Skill responses
* Deploying code to AWS Lambda without manually zipping the binary created by Go

## Configuring the Alexa skill interface

An Alexa skill is not entirely written in code. Configuration, tests, and distribution are done in the [Alexa Developer Console](https://developer.amazon.com/alexa/console/ask). The console is where you create the **skill interface**, the code is where you create the **skill service**:

> The Alexa skill consists of two main components: the skill interface and the skill service.
>
> The skill interface processes the user’s speech requests and then maps them to intents within the interaction model. ...
>
> The skill service determines what actions to take in response to the JSON encoded event received from the skill interface. Upon reaching a decision the skill service returns a JSON encoded response to the skill interface for further processing. After processing, the speech response is sent back to the user through the Echo.[^1]

### Invocation

The invocation name is the phrase users speak to trigger a particular skill. Keeping it simple and understandable by Alexa is critical--a skill that's hard to launch will never be used.

My initial invocation name was "Apple Buyer's Guide" and turned out to have two major problems. First, it refers to the brand Apple too directly, implying this was an official skill or sponsored by Apple. Second, "Buyer's" proved difficult for Alexa to understand which led to a frustrating experience interacting with the skill. "Picking Apples" resolved both of these problems.

[Here](https://developer.amazon.com/docs/custom-skills/understanding-how-users-invoke-custom-skills.html) is Amazon's documentation on invocation.

### Intents

![Alexa skill intents](../static/public/images/alexa-skill-intents.png)

Intents capture what the user "intends" to do, such as ask for help, interact with the skill, or exit the skill. Amazon takes care of the defaults (help, cancel, etc.) but requires configuration for the unique aspects of your skill.

In my case, I needed to define an intent to ask for a product recommendation. Each intent has at least one utterance (word or phrase) the user speaks to invoke the intent. My utterances included:

* should I buy the {product}
* is now a good time to buy the {product}
* Apple {product}

Where `{product}` is a defined slot for the part of the phrase that is variable. In the skill service, I extract the contents of this slot to know which product the user is asking about:

```go
product := request.Body.Intent.Slots["product"].Value
```

### Endpoint

The endpoint is where you define the web location of the skill service. There are two options: an AWS Lambda function or an HTTPS service.

## Writing the AWS Lambda Function (the skill service)

![AWS Lambda](../static/public/images/alexa-skill-aws-lambda.png)

(Basics of Lambda functions and the serverless movement)

(Lambda configuration, implementation in Go)

(Alexa request/response)

## Demonstration

![Alexa Skill Demonstration](../static/public/images/alexa-skill-demonstration.png)

[^1]: https://medium.com/crowdbotics/how-to-build-a-custom-amazon-alexa-skill-step-by-step-my-favorite-chess-player-dcc0edae53fb

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
