/*
This package and its sub-packages implement game tree search algorithms, which
are functions that look at a given board state and try to predict the best
move. Search functions in this package are entirely stateless, which means that
certain optimizations like transposition tables have to be passed as arguments
and maintained outside of the function.
*/
package search
