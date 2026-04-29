# Notes

This document contains explanations and justifications for certain concepts that can't be sufficiently explicated in doc comments.

## Terminology

A *position index* is an integer $n \in \set{0, 1, \ldots, 99}$ that represents a position on the $10\times10$ Amazons board. The board state is stored in row-major order, visualized below:
$$
\begin{bmatrix}
	00 & 01 & 02 & \ldots & 09 \\
	10 & 11 & 12 & \ldots & 19 \\
	\vdots & & \ddots	& & \vdots \\
	80 & 81 & 82 & \ldots & 89 \\
	90 & 91 & 92 & \ldots & 99
\end{bmatrix}
$$

