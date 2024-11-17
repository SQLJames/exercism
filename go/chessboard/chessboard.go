package chessboard

// Declare a type named Rank which stores if a square is occupied by a piece - this will be a slice of bools
type Rank []bool

// Declare a type named Chessboard which contains a map of eight Ranks, accessed with keys from "A" to "H"
type Chessboard map[string]Rank

func countPieces(rank Rank) (pieces int) {
	for _, v := range rank {
		if v {
			pieces++
		}
	}
	return pieces
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank string) int {
	r := cb[rank]
	return countPieces(r)
}

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file int) (pieces int) {
	if file < 1 || file > 8 {
		return 0
	}
	for _, v := range cb {
		if v[file-1] {
			pieces++
		}
	}
	return pieces
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) (totalSquares int) {
	for _, v := range cb {
		totalSquares += len(v)
	}
	return totalSquares
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) (OccupiedSquares int) {
	for _, rank := range cb {
		OccupiedSquares += countPieces(rank)
	}
	return OccupiedSquares
}
