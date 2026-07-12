package lasagnamaster 


func PreparationTime(slice []string, avgPrepTime int) int {
    if avgPrepTime == 0 {
        avgPrepTime = 2
    }

    return avgPrepTime * len(slice)
}

// TODO: define the 'Quantities()' function
func Quantities(slice []string) (int, float64) {
    nbGram := 0
    litreSauce := 0.0

    for i:=0; i<len(slice); i++ {
        if slice[i] == "noodles" {
            nbGram += 50
        } else if slice[i] == "sauce" {
            litreSauce += 0.2
        }
    }

    return nbGram, litreSauce
}



// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList, myList []string) () {
    myList[len(myList) - 1] = friendsList[len(friendsList) - 1]
}




// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, nbPortion int) []float64 {

	var res []float64 = make([]float64, len(quantities))
    
    for i:=0; i<len(quantities); i++ {
    	res[i] = (quantities[i]/ 2.0) * float64(nbPortion)
    }

    return res
}




// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
