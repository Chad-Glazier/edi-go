/*
This package contains a set of re-implemented versions of existing functions
(e.g., alpha-beta search) but with some added overhead to provide diagnostics
(e.g., tracking search depth and nodes pruned). These functions are meant
strictly for collecting diagnostic data and should not be used in place of the
normal implementations.

Note that this package is meant specifically for diagnosing search functions,
not for diagnosing VI's. VI's typically use search functions to make decisions,
but they are not the same thing.
*/
package diag
