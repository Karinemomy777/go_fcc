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
Using strings.Fields() to break the word into a slice of string. Iterating over the slice, and if the word is equal to the first item in the array, then I capitalize it via strings.Title(), leave the rest, and join the string back together.


### 6. Longest Number in a Slice

Note - the freeCodeCamp says to find the number in an array, but what kind of Gopher would I be if I didn't use slices? 

In this case, I created a function that takes in a slice of int64 (though you can change this to int, int32, float, etc.) and iterated over the values in the slice. I then used an if statement to compare which one wqas the biggest and return the largest value.

### 7. Confirm the Ending

For this challange, I took advantage of the features in the Go strings package. I created a function that would take in a string (the word) and a rune (the letter we're comparing). I created a variable to slice of the last letter in the string, and compare it to the rune. If correct, then we confirm the ending.

### 8. Repeat Strings

This one is still a work in progress :) For now, I'm taking advatage of Go's strings package, but will work on a more elaborate solution later.