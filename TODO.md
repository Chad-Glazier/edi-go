# To-Do

- Sorting by history is slow and is effectively useless in the early turns. We can make it more efficient by updating the history scores but not actually sorting by them for the first $n$ turns. To find the optimal $n$, we need to do run some diagnostics.
	- Note: the number of cutoffs $c$ probably matters more than $n$. If It's cheap enough to keep track of the number of cutoffs (as a property of the table), then this might be better. It would also allow us to handle the sorting at the table level--once the threshold is reached, a switch flips and the table's `Sort` method turns on; before then it just returns the unaltered slice.
- Implement the diagnostics package.
