# freeCodeCamp Challenges w/Golang
***

A couple of FreeCodeCamp challenges using Golang. Will update occasionally, but mostly a project for fun and playing around with Golang.

## challenges
***

### 1. Reverse a String
This was one surprisingly tricky for me in Go - I initially started using strings and strings.Split() until I revised a few Stack Overflow questions that pointed me in the right direction of using runes instead. Effectively though, for single strings breaking up the character into a rune makes more logical sense.

### 2. Factorial
In this case, I created a channel, then started a go routine where I looped over the values from n to 1 and added them to the final value. I then pushed that value into a channel.

### 3. Palindrome
Get the string from the user/input string. Compare to the reversed form of the string. If the two are equal, return that this is a palindrome.

### 4. Longest Word in a String
Using strings.Fields() which returns a slice of string. Looping over that string and then returning the longest version of that string.

### 5. Title Case
