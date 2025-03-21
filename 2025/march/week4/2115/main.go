package main
import (
	"fmt"
)

func findAllRecipes(recipes []string, ingredients [][]string, supplies []string) []string {
	availableSupplies := make(map[string]bool)
	recipeToIndex := make(map[string]int)
	dependencyGraph := make(map[string][]string)

	for _, s := range supplies {
		availableSupplies[s] = true
	}

	for i, r := range recipes {
		recipeToIndex[r] = i
	}

	inDegree := make([]int, len(recipes))

	for i := 0; i < len(recipes); i++ {
		for _, ingredient := range ingredients[i] {
			if !availableSupplies[ingredient] {
				if dependencyGraph[ingredient] == nil {
					dependencyGraph[ingredient] = make([]string, 0)
				}

				dependencyGraph[ingredient] = append(dependencyGraph[ingredient], recipes[i])
				inDegree[i]++
			}
		}
	}

	queue := make([]int, 0)
	for i := 0; i < len(recipes); i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	result := make([]string, 0, len(recipes))
	for len(queue) > 0 {
		i := queue[0]
		queue = queue[1:]
		recipe := recipes[i]
		result = append(result, recipe)

		if dependencyGraph[recipe] == nil {
			continue
		}

		for _, dependentRecipe := range dependencyGraph[recipe] {
			inDegree[recipeToIndex[dependentRecipe]]--
			if inDegree[recipeToIndex[dependentRecipe]] == 0 {
				queue = append(queue, recipeToIndex[dependentRecipe])
			}
		}
	}

	return result
}

func main() {
	// result: ["bread"]
	// recipes := []string{"bread"}
	// ingredients := [][]string{{"yeast","flour"}}
	// supplies := []string{"yeast","flour","corn"}

	// result: ["bread","sandwich"]
	// recipes := []string{"bread","sandwich"}
	// ingredients := [][]string{{"yeast","flour"},{"bread","meat"}}
	// supplies := []string{"yeast","flour","meat"}

	// result: ["bread","sandwich","burger"]
	recipes := []string{"bread","sandwich","burger"}
	ingredients := [][]string{{"yeast","flour"},{"bread","meat"},{"sandwich","meat","bread"}}
	supplies := []string{"yeast","flour","meat"}

	// result: []
	// recipes := []string{}
	// ingredients := [][]string{}
	// supplies := []string{}

	result := findAllRecipes(recipes, ingredients, supplies)
	fmt.Printf("result = %v\n", result)
}

