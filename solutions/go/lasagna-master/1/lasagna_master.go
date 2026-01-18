package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, time int) int {
    
    if time == 0 {
    	time = 2
    }
    
	return len(layers) * time
    
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
    var noodles int
    var sauce float64

	for _, layer := range layers {
        if layer == "sauce" {
            sauce += 0.2
        }
        if layer == "noodles" {
            noodles += 50
        }       
    }  
    return noodles, sauce
}


// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList []string, myList []string) {
    secret := friendsList[len(friendsList)-1]
    myList[len(myList)-1] = secret  
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portions int) []float64 {
    factor := float64(portions) / 2.0
    scaled := make([]float64, len(quantities))
	for i := 0; i < len(quantities); i++ {
        scaled[i] = quantities[i] * factor
    }

    return scaled
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
