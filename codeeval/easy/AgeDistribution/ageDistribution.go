package ageDistribution

/*
AGE DISTRIBUTION

Challenge Description:

You"re responsible for providing a demographic report for your local school district based on
age. To do this,
you"re going determine which "category" each person fits into based on their age.
The person"s age will determine which category they should be in:

If they"re from 0 to 2 the child should be with parents print : "Still in Mama"s arms"
If they"re 3 or 4 and should be in preschool print : "Preschool Maniac"
If they"re from 5 to 11 and should be in elementary school print : "Elementary school"
From 12 to 14: "Middle school"
From 15 to 18: "High school"
From 19 to 22: "College"
From 23 to 65: "Working for the man"
From 66 to 100: "The Golden Years"
If the age of the person less than 0 or more than 100 - it might be an alien - print: "This
program is for humans"

Input sample:
Your program should accept as its first argument a path to a filename.
Each line of input contains one integer - age of the person:

0
19

Output sample:
For each line of input print out where the person is:

Still in Mama's arms
College
*/

func Process(age int) string {
	switch {
	case age < 0 || age > 100:
		return "This program is for humans"
	case age > 65:
		return "The Golden Years"
	case age > 22:
		return "Working for the man"
	case age > 18:
		return "College"
	case age > 14:
		return "High school"
	case age > 11:
		return "Middle school"
	case age > 4:
		return "Elementary school"
	case age > 2:
		return "Preschool Maniac"
	default:
		return "Still in Mama's arms"
	}
}
