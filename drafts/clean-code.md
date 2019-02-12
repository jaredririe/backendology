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
- Using a comment when a function or variable name works instead
- Position markers: // Section /////
- Attributions (quickly gets out of date, use source control instead)
- Commented-out code
- Describes far-away code
- Javadocs in non-public code

## Chapter 5: Formatting

When people look under the hood, you want them to be struck by the orderliness. You want them to think “professionals are at work” not “this code looks like it was written by a bevy of drunken sailors.” Formatting causes readers to conclude that the same in/attention to detail prevades other other aspect of the project.

Decide on formatting as a team and have an automatic tool enforce it.

Getting code working is not the first order of business. The functionality changes but the readability of your code will have a profound effect on all future changes. Style and discipline survive even when the code hardly resembles the original.

-> Example: Data Mapper, Marauder’s Map. Disciplined codebases encourage disciplined practices moving forward.

### Vertical Formatting

Small files are easier to understand

Newspaper Metaphor: headline, synopsis of the whole story, details (dates, quotes, names, etc.)

Name gives context of which module you’re in. Topmost parts provide high-level concepts and algorithms. Details increase as we move downward. Lowest-level functions at the end.

Vertical openness: Each group of lines represents a complete thought. Thoughts should be separated by blank lines.

Vertical density: lines of code tightly related should appear vertically close.

Vertical distance: keep related concepts in the same file and close within that file. We want to avoid forcing readers to hop around through the code.

- Variables should be declared as close to their usage as possible
- Instance variables should be declared at the top of the class (usually used by most methods)
- Dependent functions should be close and the caller should be above the callee

### Horizontal Formatting

Keep lines short. Anything beyond 100-120 characters is careless.

Spaces around operators and arguments keeps them separate. No spaces between a function name and its parentheses keeps them connected.

Breaking indentation (shortening if statements to one line) is not worth the space savings.

Every member of a team should use the same agreed-upon formatting.

## Chapter 6: Objects and Data Structures

- Objects: hide data behind abstractions and expose functions to operate on that data. Prefer functions over public variables.
- Data Structures: expose their data and have no meaningful functions. Prefer public variables over functions
- Hybrids: a combination of the two. Should be carefully avoided!

A class exposes abstract interfaces that allow users to manipulate the essence of the data without having to know its implementation.

It’s a mistake to blindly add getters and setters for all members.

Objects have functions that allow us to tell it to do something, not reveal their internals.

Procedural vs. OO

- Procedural makes it hard to add new data structures because all functions must change. Yet it’s easy to add new functions without changing existing data structures
- OO makes it hard to add new functions because all classes must change. Yet it’s easy to add new classes without changing existing functions

Law of Demeter: a method f of class C should only call methods of:

- C
- objects created by f
- an object passed as an argument to f
- an object held as an instance variable of C

A line like the following is a train wreck that should be split up:

```
final String outputDir = ctxt.getOptions().getScratchDir().getAbsolutePath();
```

Data Transfer Object (DTO) is a class with public variables and no functions. It is a data structure. These are very useful when communicating with databases and are the first in a series of translation stages.

## Chapter 7: Error Handling

Error handling that obscures logic is wrong

### Use exceptions rather than return codes

Allows you to look at the logic and error handling separately in the caller code

Try blocks are like transactions

### Use unchecked exceptions

Checked exceptions violate the open/closed principle as it’s required to import the exception types until they are caught

### Provide context

Each exception that you throw should provide enough context to determine the source and location of an error. In Java, you can get a stack trace from any exception; however, a stack trace can’t tell you the intent of the operation that failed.

Create informative error messages and pass them along with your exceptions. Mention the operation that failed and the type of failure. If you are logging in your application, pass along enough information to be able to log the error in your catch.

### Wrap third party APIs and translate their errors into yours

When you wrap a third-party API, you minimize your dependencies upon it: You can choose to move to a different library in the future without much penalty. Wrapping also makes it easier to mock out third-party calls when you are testing your own code.

Often a single exception class is ﬁne for a particular area of code. The information sent with the exception can distinguish the errors. Use different classes only if there are times when you want to catch one exception and allow the other one to pass through.

### Special Case Pattern

You create a class or conﬁgure an object so that it handles a special case for you. When you do, the client code doesn’t have to deal with exceptional behavior. That behavior is encapsulated in the special case object.

### Don’t return null

When we return null, we are essentially creating work for ourselves and foisting problems upon our callers. All it takes is one missing null check to send an application spinning out of control.

If you are tempted to return null from a method, consider throwing an exception or returning a SPECIAL CASE object instead. If you are calling a null-returning method from a third-party API, consider wrapping that method with a method that either throws an exception or returns a special case object.*

Ex: return an empty list instead of null from getEmployees

### Don’t pass null

Passing null is worse than returning it.

In most programming languages there is no good way to deal with a null that is passed by a caller accidentally. Because this is the case, the rational approach is to forbid passing null by default. When you do, you can code with the knowledge that a null in an argument list is an indication of a problem, and end up with far fewer careless mistakes.

## Chapter 8: Boundaries

There is a natural tension between the provider of an interface and the user of an interface. Providers of third-party packages and frameworks strive for broad applicability so they can work in many environments and appeal to a wide audience. Users, on the other hand, want an interface that is focused on their particular needs. This tension can cause problems at the boundaries of our systems.

Ex: giving access to variables of type map violates boundaries, especially if we don’t want the caller to modify that map.

public class Sensors { private Map sensors = new HashMap(); public Sensor getById(String id) { return (Sensor) sensors.get(id); } //snip }

The interface at the boundary (Map) is hidden. It is able to evolve with very little impact on

the rest of the application. The use of generics is no longer a big issue because the casting and type management is handled inside the Sensors class.

### Learning tests

Instead of experimenting and trying out the new stuff in our production code, we could write some tests to explore our understanding of the third-party code. Jim Newkirk calls such tests learning tests. 1

In learning tests we call the third-party API, as we expect to use it in our application. We’re essentially doing controlled experiments that check our understanding of that API. The tests focus on what we want out of the API.

Learning tests verify that the third-party packages we are using work the way we expect them to. Once integrated, there are no guarantees that the third-party code will stay compatible with our needs. The original authors will have pressures to change their code to meet new needs of their own. They will ﬁx bugs and add new capabilities. With each release comes new risk. If the third-party package changes in some way incompatible with our tests, we will ﬁnd out right away.

## Unknown code

To keep from being blocked, we deﬁned our own interface. We called it something catchy, like Transmitter. We gave it a method called transmit that took a frequency and a data stream. This was the interface we wished we had.

One good thing about writing the interface we wish we had is that it’s under our control. This helps keep client code more readable and focused on what it is trying to accomplish.

Code at the boundaries needs clear separation and tests that deﬁne expectations. We should avoid letting too much of our code know about the third-party particulars. It’s better to depend on something you control than on something you don’t control, lest it end up controlling you.

## Chapter 9: Unit Tests

### Example

We've come a long way in our profession in terms of testing

> Nowadays I would write a test that made sure that every nook and cranny of that code worked as I expected it to. I would isolate my code from the operating system rather than just calling the standard timing functions. I would mock out those timing functions so that I had absolute control over the time. I would schedule commands that set boolean flags, and then I would step the time forward, watching those flags and ensuring that they went from false to true just as I changed the time to the right value.

### The Three Lwas of TDD

> First Law: You may not write production code until you have written a failing unit test.
>
> Second Law: You may not write more of a unit test than is sufficient to fail, and not compiling is failing.
>
> Third Law: You may not write more production code than is sufficient to pass the cur- rently failing test.

The tests and the production code are written together, with the tests just a few seconds ahead of the production code.

### Keeping tests clean

> Having dirty tests is equivalent to, if not worse than, having no tests. The problem is that tests must change as the production code evolves. The dirtier the tests, the harder they are to change. The more tangled the test code, the more likely it is that you will spend more time cramming new tests into the suite than it takes to write the new production code. As you modify the production code, old tests start to fail, and the mess in the test code makes it hard to get those tests to pass again. So the tests become viewed as an ever-increasing liability.

> From release to release the cost of maintaining my team’s test suite rose. ... In the end they were forced to discard the test suite entirely. But, without a test suite they lost the ability to make sure that changes to their code base worked as expected. Without a test suite they could not ensure that changes to one part of their system did not break other parts of their system. So their defect rate began to rise. As the number of unintended defects rose, they started to fear making changes. They stopped cleaning their production code because they feared the changes would do more harm than good. Their production code began to rot. In the end they were left with no tests, tangled and bug-riddled production code, frustrated customers, and the feeling that their testing effort had failed them

"Test code is just as important as production code."

Tests allow you to improve the design, introduce new features, and maintain the system without breaking anything along the way. High code coverage from automated tests dispels fear.

#### What makes a clean test?

> Three things. Readability, readability, and readability. Read- ability is perhaps even more important in unit tests than it is in production code. What makes tests readable? The same thing that makes all code readable: clarity, simplicity, and density of expression. In a test you want to say a lot with as few expressions as possible.

#### BUILD-OPERATE-CHECK

> The BUILD-OPERATE-CHECK pattern is made obvious by the structure of these tests. Each of the tests is clearly split into three parts. The first part builds up the test data, the second part operates on that test data, and the third part checks that the operation yielded the expected results.

Rather than using the APIs that programmers use to manipulate the system, we build up a set of functions and utilities that make use of those APIs and that make the tests more convenient to write and easier to read.

#### Dual Standard

While tests should be clean, they should favor readability over efficiency. There are some things you would never do in production code that are fine in tests.

#### Single Concept Per Test

Perhaps a better rule is that we want to test a single concept in each test function. We don’t want long test functions that go testing one miscellaneous thing after another.

#### FIRST

Clean tests follow ﬁve other rules that form the above acronym:

Fast: Tests should be fast. They should run quickly. When tests run slow, you won’t want to run them frequently. If you don’t run them frequently, you won’t ﬁnd problems early enough to ﬁx them easily. You won’t feel as free to clean up the code. Eventually the code will begin to rot.

Independent: Tests should not depend on each other. One test should not set up the conditions for the next test. You should be able to run each test independently and run the tests in any order you like. When tests depend on each other, then the ﬁrst one to fail causes a cascade of downstream failures, making diagnosis difﬁcult and hiding downstream defects.

Repeatable: Tests should be repeatable in any environment. You should be able to run the tests in the production environment, in the QA environment, and on your laptop while riding home on the train without a network. If your tests aren’t repeatable in any environment, then you’ll always have an excuse for why they fail.

Self-Validating: The tests should have a boolean output. Either they pass or fail. You should not have to read through a log ﬁle to tell whether the tests pass. If the tests aren’t self-validating, then failure can become subjective and running the tests can require a long manual evaluation.

Timely: The tests need to be written in a timely fashion. Unit tests should be written just before the production code that makes them pass. If you write tests after the production code, then you may ﬁnd the production code to be hard to test. You may decide that some production code is too hard to test. You may not design the production code to be testable.

## Chapter 10: Classes

### Class Organization

Begin with a list of variables, starting with public static constants. Then private static variables, then private instance variables. Public functions come after variables.

### Classes should be small!

Classes should be small as measured by the number of responsibilities they have. If it's hard to come up with a short, concrete name for a class, it's likely too large. Words like "Processor" "Manager" "Super" are a hint at unfortunate aggregation of responsibility.

Small vs. large class analogy: organizing tools into small and well-labeled drawers or throwing them into large bins.

Too many think they are done when their code works. They fail to spend time on the other concern, which is organization/cleanliness.

#### Cohesion

If a subset of variables are used by a group of methods, that could mean that a smaller class is trying to get out of the larger class.

### Abstract vs. Concrete classes

Abstract: represent concepts only

Concrete: contain implementation details

A client class depending upon concrete details is at risk when those details change. We can introduce interfaces and abstract classes to help isolate the impact of those details.

When a system is decoupled for testability, it is more flexible and promotes better reuse. Classes should depend on abstractions, not concrete details.

## Chapter 11: Systems

consider that construction is a very different process from use

Software systems should separate the startup process, when the application objects are
constructed and the dependencies are “wired” together, from the runtime logic that takes
over after startup.

The startup process of object con-
struction and wiring is no exception. We should modularize this process separately from
the normal runtime logic and we should make sure that we have a global, consistent strat-
egy for resolving our major dependencies.

One way to separate construction from use is simply to move all aspects of construction to
main, or modules called by main, and to design the rest of the system assuming that all
objects have been constructed and wired up appropriately. (See Figure 11-1.)
The flow of control is easy to follow. The main function builds the objects necessary
for the system, then passes them to the application, which simply uses them. Notice the
direction of the dependency arrows crossing the barrier between main and the application.
They all go one direction, pointing away from main. This means that the application has no
knowledge of main or of the construction process. It simply expects that everything has
been built properly.

It is a myth that we can get systems “right the ﬁrst time.” Instead, we should implement only today’s stories, then refactor and expand the system to implement new stories tomorrow. This is the essence of iterative and incremental agility. Test-driven development, refactoring, and the clean code they produce make this work at the code level.

## Chapter 13: Emergence

a design is “simple” if it follows these rules:

• Runs all the tests

A system that is comprehensively  tested and passes all of its tests all of the time is a testable system. 

Fortunately,  making our systems testable pushes us toward a design where our classes are small and single purpose. It’s just easier to test classes that conform to the SRP.  The more tests we  write, the more we’ll continue to push toward things that are simpler to test. So making sure our system is fully testable helps us create better designs.

Tight coupling makes it difficult to write tests. So, similarly,  the more tests we  write, the more we  use principles like DIP and tools like dependency injection, interfaces, and abstraction to minimize coupling. Our designs improve  even more.

• Contains no duplication 

Duplication is the primary enemy of a well-designed system. It represents additional work, additional risk, and additional unnecessary complexity. Duplication manifests itself in many forms. Lines of code that look exactly alike are, of course, duplication. Lines of code that are similar can often be massaged to look even more alike so that they can be more easily refactored.  And duplication can exist in other forms such as duplication of implementation4

• Expresses the intent of the programmer 
• Minimizes the number of classes and methods The rules are given in order of importance. 