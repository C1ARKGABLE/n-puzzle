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
For a class assignment, we are working with simple search algorithms. This is an implementation of [A*](https://en.wikipedia.org/wiki/A*_search_algorithm) with [manhattan distance](https://en.wikipedia.org/wiki/Manhattan_distance) and [Nilsson Sequence of Moves score](https://www.cse.iitk.ac.in/users/cs365/2009/ppt/13jan_Aman.pdf#page=7) as [scoring heuristics](https://en.wikipedia.org/wiki/Heuristic_(computer_science)). After learning how A* works, we've be asked to implement it in a coding language of our own choice.

### What're A* and heuristics?
A* is a simple and light pathfinding algorithm used for soling problems. A* is unique because it uses two different scores in its decision making:
* Score one is a sum of the past moves, and while this allows the algorithm to minimize its moves, it doesn't easily find the goal state. 
* Score two is a heuristic, or a function that ranks the favorability of a move based on some known information. 

In the case of heuristics, we are using manhattan distance from the solution state and the Nilsson Sequence of Moves score. These the two most accurate heuristics for the n-puzzle, as described in [this lecture](https://www.cse.iitk.ac.in/users/cs365/2009/ppt/13jan_Aman.pdf).

### Why Go?
Honestly, we could have picked [any programming language](https://en.wikipedia.org/wiki/List_of_programming_languages) (Like, all the way down to [assembly](https://en.wikipedia.org/wiki/Assembly_language)... It'd be fast 🖥💨). We were curious about how Go works, and this problem is the perfect test case for learning. So, why not Go?

#### Breakdowns:

.| **S** |.||.| **G** |.
-|-|-|-|-|-|-
1 | 2 | 3 |->| 1 | 2 | 3
7 | 4 | 5 |->| 8 | 6 | 4
6 | 8 | 0 |->| 7 | 5 | 0

* **Nilsson:** true 
**Sequence of Moves:** [down left down left up right down right] 
**Nodes expanded:** 19 
**Nodes visited:** 8
* **Nilsson:** false 
**Sequence of Moves:** [down left down left up right down right] 
**Nodes expanded:** 19 
**Nodes visited:** 8


.| **S** |.||.| **G** |.
-|-|-|-|-|-|-
 0 | 1 | 3 |->| 1 | 2 | 3
 8 | 2 | 4 |->| 8 | 0 | 4
 7 | 6 | 5 |->| 7 | 6 | 5
* **Nilsson:** true 
**Sequence of Moves:** [left up] 
**Nodes expanded:** 2 
**Nodes visited:** 2
* **Nilsson:** false 
**Sequence of Moves:** [left up] 
**Nodes expanded:** 2 
**Nodes visited:** 2


.| **S** |.||.| **G** |.
-|-|-|-|-|-|-
 2 | 8 | 1 |->| 3 | 2 | 1
 3 | 4 | 6 |->| 8 | 0 | 4
 7 | 5 | 0 |->| 7 | 5 | 6
* **Nilsson:** true 
**Sequence of Moves:** [down left up left down right] 
**Nodes expanded:** 6 
**Nodes visited:** 6
* **Nilsson:** false 
**Sequence of Moves:** [down left up left down right] 
**Nodes expanded:** 6   


.| **S** |.||.| **G** |.
-|-|-|-|-|-|-
 2 | 8 | 3 |->| 1 | 2 | 3
 1 | 6 | 4 |->| 8 | 0 | 4
 7 | 0 | 5 |->| 7 | 6 | 5
* **Nilsson:** true 
**Sequence of Moves:** [down down right up left] 
**Nodes expanded:** 6 
**Nodes visited:** 5
* **Nilsson:** false 
**Sequence of Moves:** [down down right up left] 
**Nodes expanded:** 6 
**Nodes visited:** 5


.| **S** |.||.| **G** |.
-|-|-|-|-|-|-
8 | 6 | 7 |->| 1 | 2 | 3
2 | 5 | 4 |->| 4 | 5 | 6
3 | 0 | 1 |->| 7 | 8 | 0
* **Nilsson:** true 
**Sequence of Moves:** [right down left up left down right down left up right up right down down left up up right down down left up up left down right up right down left left up]
**Nodes expanded:** 5473 
**Nodes visited:** 33
* **Nilsson:** false 
**Sequence of Moves:** [down right up left left down down right right up left down right up up left left down right right up left left down down right right up up left left]
**Nodes expanded:** 28090 
**Nodes visited:** 31


.| **S** |.||.| **G** |.
-|-|-|-|-|-|-
 1 | 2 | 5 |->| 0 | 1 | 2
 6 | 3 | 4 |->| 3 | 4 | 5
 7 | 8 | 0 |->| 6 | 7 | 8
* **Nilsson:** true 
**Sequence of Moves:** [right right down left left down right right] 
**Nodes expanded:** 8 
**Nodes visited:** 8
* **Nilsson:** false 
**Sequence of Moves:** [right right down left left down right right] 
**Nodes expanded:** 8 
**Nodes visited:** 8
