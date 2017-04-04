package main

func main() {

}
/*

cards from 1 to n

four cards for example
if two cards are paired together they together make a single element reducing n by 1
if three cards

1234  2, 3, 4
1243  2
1324
1342  2
1423  2
1432

2134  2
2143
2314  2
2341  2 3
2413
2431

3124 2
3142
3214
3241
3412 2
3421 2

4123 2 3
4132
4213
4231 2
4312
4321
[1]
1 = 1/1
[2]
1 = 2/2
2 = 1/2
[3]
1 = 6/6
2 = 2/6 =
3 = 1/6
[4]
1 24/24
2 8/24
3 3/24
4 1/24
[5]
1 = 120/120
2
3
4
5 = 1/120

123  3
132
213
231 2
312 2
321



4 = 1/24
3 = 3/24
2 = 8/24
1 = 24/24

perm(4,1) = 4!/3! = 4
perm(4,2) = 4!/2! = 12
perm(4,3) = 4!/1! = 24
perm(4, 4) = 4!/0! = 24

5
4
3
2
1


*/

n iteration number
k cards remaining
l sequence length
c choices
s subsets


5 seq as exp

12xxx (16, 6, 2)
x12xx
xx12x
xxx12

23xxx
x23xx
xx23x
xxx23

34xxx
x34xx
xx34x
xxx34

45xxx
x45xx
xx45x
xxx45

123xx (9, 2, 3)
x123x
xx123

234xx
x234x
xx234

345xx
xx345
x345x

1234x (4, 1, 4)
x1234

2345x
x2345

12345 (1, 1, 5)

every time there is a one in k! chance of instant exit when l=k
the odds then branch off into the smaller sequences

n iteration number
k cards remaining
l sequence length
c choices
s subsets

