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

The objective of this code is to find the optimal (least total movements) from the initial state to the goal state.

### Why n-puzzle problem?
For a class assignment, we are working with simple search algorithms. This is an implementation of [A*](https://en.wikipedia.org/wiki/A*_search_algorithm) with [manhattan distance](https://en.wikipedia.org/wiki/Manhattan_distance) and [Nilsson Sequence score](https://www.cse.iitk.ac.in/users/cs365/2009/ppt/13jan_Aman.pdf#page=7) as [scoring heuristics](https://en.wikipedia.org/wiki/Heuristic_(computer_science)). After learning how A* works, we've be asked to implement it in a coding language of our own choice.

### What're A* and heuristics?
A* is a simple and light pathfinding algorithm used for soling problems. A* is unique because it uses two different scores in its decision making:
* Score one is a sum of the past moves, and while this allows the algorithm to minimize its moves, it doesn't easily find the goal state. 
* Score two is a heuristic, or a function that ranks the favorability of a move based on some known information. 

In the case of heuristics, we are using manhattan distance from the solution state and the Nilsson Sequence score. These the two most accurate heuristics for the n-puzzle, as described in [this lecture](https://www.cse.iitk.ac.in/users/cs365/2009/ppt/13jan_Aman.pdf).

### Why Go?
Honestly, we could have picked [any programming language](https://en.wikipedia.org/wiki/List_of_programming_languages) (Like, all the way down to [assembly](https://en.wikipedia.org/wiki/Assembly_language)... It'd be fast 🖥💨). We were curious about how Go works, and this problem is the perfect test case for learning. So, why not Go?

#### Breakdowns:

.| **S** |.||.| **G** |.
-|-|-|-|-|-|-
1 | 2 | 3 |->| 1 | 2 | 3
7 | 4 | 5 |->| 8 | 6 | 4
6 | 8 | 0 |->| 7 | 5 | 0

1) **Nilsson:** true 
**Sequence:** [down left down left up right down right] 
**Nodes expanded:** 19 
**Nodes visited:** 8
1) **Nilsson:** false 
**Sequence:** [down left down left up right down right] 
**Nodes expanded:** 19 
**Nodes visited:** 8


.| **S** |.||.| **G** |.
-|-|-|-|-|-|-
 0 | 1 | 3 |->| 1 | 2 | 3
 8 | 2 | 4 |->| 8 | 0 | 4
 7 | 6 | 5 |->| 7 | 6 | 5
1) **Nilsson:** true 
**Sequence:** [left up] 
**Nodes expanded:** 2 
**Nodes visited:** 2
1) **Nilsson:** false 
**Sequence:** [left up] 
**Nodes expanded:** 2 
**Nodes visited:** 2


.| **S** |.||.| **G** |.
-|-|-|-|-|-|-
 2 | 8 | 1 |->| 3 | 2 | 1
 3 | 4 | 6 |->| 8 | 0 | 4
 7 | 5 | 0 |->| 7 | 5 | 6
1) **Nilsson:** true 
**Sequence:** [down left up left down right] 
**Nodes expanded:** 6 
**Nodes visited:** 6
1) **Nilsson:** false 
**Sequence:** [down left up left down right] 
**Nodes expanded:** 6   


.| **S** |.||.| **G** |.
-|-|-|-|-|-|-
 2 | 8 | 3 |->| 1 | 2 | 3
 1 | 6 | 4 |->| 8 | 0 | 4
 7 | 0 | 5 |->| 7 | 6 | 5
1) **Nilsson:** true 
**Sequence:** [down down right up left] 
**Nodes expanded:** 6 
**Nodes visited:** 5
1) **Nilsson:** false 
**Sequence:** [down down right up left] 
**Nodes expanded:** 6 
**Nodes visited:** 5


.| **S** |.||.| **G** |.
-|-|-|-|-|-|-
 8 | 0 | 6 |->| 0 | 1 | 2
 5 | 4 | 7 |->| 3 | 4 | 5
 2 | 3 | 1 |->| 6 | 7 | 8
1) **Nilsson:** true 
**Sequence:** [left down right up left down up right up right down left up right up left down left up right down right up left down left down right up down left down right down left up down] 
**Nodes expanded:** 1208 
**Nodes visited:** 37
1) **Nilsson:** false 
**Sequence:** [left down right up left down up right up right down left up right up left down left up right down right up left down left down right up down left down right down left up down] 
**Nodes expanded:** 1208 
**Nodes visited:** 37


.| **S** |.||.| **G** |.
-|-|-|-|-|-|-
 1 | 2 | 5 |->| 0 | 1 | 2
 6 | 3 | 4 |->| 3 | 4 | 5
 7 | 8 | 0 |->| 6 | 7 | 8
1) **Nilsson:** true 
**Sequence:** [right right down left left down right right] 
**Nodes expanded:** 8 
**Nodes visited:** 8
1) **Nilsson:** false 
**Sequence:** [right right down left left down right right] 
**Nodes expanded:** 8 
**Nodes visited:** 8
