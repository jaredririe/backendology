# Clean Code

## Introduction

The only valid code quality metric: WTFs/minute

Learning to write clean code is hard work. You must practice and fail. You must sweat over your code.

Never think you’ll clean up code later. *Later equals never*.

## Chapter 1

It’s your responsibility as the developer to defend the code. The manager defends the schedule and requirements, but he really wants clean code as well.

The only way to truly go fast is keep the code as clean as possible at all times.

Clean code

- Elegant
- Pleasing to read
- Hides bugs
- Has minimal dependencies
- Avoids broken windows
- Focused: does one thing well
- Contains crisp abstractions
- Tested
- Minimal API
- Easy for others to understand and change
- Written by someone who cares about the craft
- Avoids duplication

Most of our time writing code is spent referencing other parts of the code. Code is read, whether we think it is or not.

Leave code a little cleaner than how you found it.

## Chapter 2: Naming

Use intention-revealing names

Don’t reveal implementation details in the name.

Avoid the practice of satisfying the compiler by changing a variable name to a close form when it’s already used. yaml -> yamlStr.

Use pronounceable names as programming is a social activity.

The length of a name should correspond to the size of its scope.

Don’t embed types in the name, such as String, interface (IShapeFactory, ShapeFsctoryImp, StorageInterface).

Class names should be a noun or noun phrase. Avoid words like Manager, Processor, Data, Info.

Method names should be a verb or verb phrase.

Avoid cleverness in names. Say what you mean, mean what you say. Don’t put in jokes or culturally-dependent words.

What’s the difference between a manager, processor, and driver? There’ll all essentially equivalent. Only use one in a given codebase if at all.

Use technical words which assume a CS background.

Groups related words under a class/struct. State should be within an address struct or else it is ambiguous.

Shorter names are generally better than longer ones, so long as they are clear. Add no more context than is necessary.

## Chapter 3: Functions

The first rule of functions is that they should be small. The second rule of functions is that they should be smaller than that. Functions should hardly ever be 20 lines long

The indent level of a function should not be greater than one or two.

Functions should do one thing well.

“If all the steps of a function are one level of abstraction below the stated name, it is doing one thing well.”

If you find yourself dividing code into “sections” (perhaps with a comment and whitespace), this is a symptom of a function that does not do just one thing well. Functions that do one thing well cannot reasonably be divided into sections.

Don’t mix levels of abstraction in a single function. This is confusing and causes details to creep into high-level concepts

1. getHTML() - highest
2. String pagePathName = PathParser.render(pagePath)
3. .append("\n") - lowest

Code should be read top-down (step-down rule)

- Abstraction 1
	- Abstraction 1.a
	- Abstraction 1.b

Switch statements and if-else chains should be buried in a low-level class. Should appear only once and are used to create polymorphic objects behind an inheritance relationship

Use the same phrases, nouns, and verbs in the names you choose for consistency.

The ideal number of arguments for a function is zero, then one, then two. Three should be avoided where possible. More than three requires very special justification and then should be used anyway.

Prefer storing values in the object/struct to passing them around as arguments.

Avoid passing in values that are modified in place (output arguments) as it causes a double take.

Monadic forms are common (one argument)

- Operate on an argument
- Ask a question about an argument
- (Rare) Event
- Avoid output arguments and instead prefer returning the input so it’s still a transformation

Flag arguments (passing a Boolean into a function) is truly a terrible practice. It loudly proclaims this function does more than one thing. Instead, make two functions for each path and put the Boolean check in the calling function.

Dyadic functions (two arguments)

- Take longer to understand
- Not evil, will be written
- Order is easy to switch

Triadic functions (three arguments)

- Significantly harder to understand than dyads
- Ordering concerns and complexity increase

Writing good functions

- Write the function as you would write a draft for an article
- It may be long and complicated, but add unit tests
- Then, massage and refine, reduce duplication, think of better names all the while keeping the tests passing

## Chapter 4: Comments

Don’t comment bad code, rewrite it

Compensate for our failure to express ourselves in code. Always try to express yourself in code. Don’t pat yourself in the back when you write a comment.

Inaccurate comments are worse than no comments.

Truth can only be found in the code.

Often you can write a function that says the same thing as the comment you were going to write.

Decent uses

- Explanation of intent
- Informative
- Clarification
- Warning of consequences
- Todo comments
- Amplification
- Javadocs in Public APIs

Bad comments

- Mumbling: writing a comment because you think you should
- Redundant
- Misleading
- Mandated comments, like requiring one on every function and giving no new information
- Document changes (that’s what Git is for)
