package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	rep := 0


	for name, list := range cb {
        if name == file {
            for _, val := range list {
                if val {
                    rep++
            	}
        	}
        
    	}
    }

    return rep
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	rep := 0
    
	if (rank < 1 || rank > 8) {
        return 0
    }
	
	for _, list := range cb {
    	if list[rank-1] {
            rep++
        }    
    }
    


    return rep
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	rep := 0

	for _, list := range cb {
    	for range list {
            rep++
        }   
    }

    return rep
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	rep := 0


	for name, _ := range cb {
        rep += CountInFile(cb, name)
    }



    return rep
}
