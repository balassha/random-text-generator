# random-text-generator
## Description

It is a Go program that generates random texts based on
previously seen text corpus.

To accomplish this task, we want to parse the given text into a list of words.
These words are then structured into "trigrams" which are simply groups of three words.
Now for each such trigram we calculate the frequency of how often a word appears
after the previous two words. Using these frequencies we can generate text based
on the last two words that we have generated.

### Example

Lets assume our input text corpus is the sentence:

> To be or not to be that is the question

From this we may extract the words:

```
["To", "be", "or", "not", "to", "be", "that", "is", "the", "question"]
```

The eight trigrams are:

```
["To", "be", "or"]
["be", "or", "not"]
["or", "not", "to"]
["not", "to", "be"]
["to", "be", "that"]
["be", "that", "is"]
["that", "is", "the"]
["is", "the", "question"]
```

The frequencies count how often we have seen a word, given the previous two words.
For the input above we would end up with the following frequencies:

| Prefix  | word     | count |
|---------|----------|-------|
| be or   | not      | 1     |
| be that | is       | 1     |
| is the  | question | 1     |
| not to  | be       | 1     |
| or not  | to       | 1     |
| that is | the      | 1     |
| to be   | or       | 1     |
| to be   | that     | 1     |

Now when we want to generate a new text we first pick one of the trigrams at
random and then, word by word, we create the text using the calculated frequencies.

For instance we may start with "[To be or]" as the first trigram. From this we
take the first two words to start our text:

> To be

We now pick the next word from the two options we have for this prefix which are
"or" (1) and "that" (1). Both words have the same count so the chances to pick
either is a 50:50 (or 1:1). Lets say we pick "that":

> To be that

The new prefix now is created from the last two words of our generated text:
"be that". Looking at our frequency table again we see there is only a single
option we can take from here which is the word "is":

> To be that is

You can see with the limited corpus we provided there is not a lot of room to
generate an interesting text. However if we train our model on a large enough
corpus, we will have more options, often with different counts for the next
possible word. These counts should be used to make a _weighted_  random
selection of the next word. For instance if the options for the next word are
"a" with a weight of 2 and "b" with a weight of 1, the chance that we pick "a"
should be twice as high as the chance of picking "b".