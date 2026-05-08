# To-Do

- Refactor search functions to be stateless.
- Make `board.Apply` return an error on invalid moves; don't let it mutate in place.
- Maybe make a `move.Infer` function to avoid having to store move data in the game state.
- Make a board state `interface` that the search functions take.
- Figure out the best way to implement analytics.