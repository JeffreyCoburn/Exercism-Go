package chessboard

const totalRanks = 8
const totalFiles = 8

// Rank stores if a square is occupied by a piece
type Rank []bool

// Chessboard contains a map of eight Ranks, accessed with values from 1 to 8
type Chessboard map[int]Rank

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank
func CountInRank(cb Chessboard, rank int) (ret int) {
	if len(cb) < rank || rank < 0 {
		return 0
	}
	for _, v := range cb[rank] {
		if v {
			ret++
		}
	}
	return
}

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file
func CountInFile(cb Chessboard, file int) (ret int) {
	if file > totalFiles || file < 0 {
		return 0
	}
	file--
	for x := 1; x <= totalFiles; x++ {
		if cb[x][file] {
			ret++
		}
	}
	return
}

// CountAll should count how many squares are present in the chessboard
func CountAll(cb Chessboard) (ret int) {
	return totalFiles * totalRanks
}

// CountOccupied returns how many squares are occupied in the chessboard
func CountOccupied(cb Chessboard) (ret int) {
	for _, rank := range cb {
		for _, v := range rank {
			if v {
				ret++
			}
		}
	}
	return
}
