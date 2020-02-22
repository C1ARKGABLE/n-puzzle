# n-puzzle in Go
This is a simple solution to the [n-puzzle](https://en.wikipedia.org/wiki/N-puzzle) problem in [Go](https://en.wikipedia.org/wiki/Go_(programming_language)).

### What is the n-puzzle problem?

A grid of tiles n+1 tiles are given in a pattern like below.

1|2|3
-|-|-
7|4|5
6|8|0

From the above state, the tiles may be slid up, down, left, and right into the free space (represented by 0). Below is a representation of the grid after one move (5 -> 0) in the downward direction.

1|2|3
-|-|-
7|4|0
6|8|5

The goal state of the n-puzzle is defined by the user as a corresponding grid. 

The objective of this program 



A report covering 8-puzzle problem formulation, program structure, global variables, functions and procedures, etc.

### Why n-puzzle problem?
For a class assignment, we are working with simple search algorithms. This is an implementation of [A*](https://en.wikipedia.org/wiki/A*_search_algorithm) with [manhattan distance](https://en.wikipedia.org/wiki/Manhattan_distance) and [Nilsson Sequence score](https://www.cse.iitk.ac.in/users/cs365/2009/ppt/13jan_Aman.pdf#page=7) as [scoring heuristics](https://en.wikipedia.org/wiki/Heuristic_(computer_science)). After learning how A* works, we've be asked to implement it in a coding language of our own choice.

### Why Go?
Honestly, we could have picked [any programming language](https://en.wikipedia.org/wiki/List_of_programming_languages) (Like, all the way down to [assembly](https://en.wikipedia.org/wiki/Assembly_language)... It'd be fast 🖥💨). We were curious about how Go works, and this problem is the perfect test case for learning. So, why not Go?

