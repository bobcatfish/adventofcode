package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func load() ([]Food, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		return nil, fmt.Errorf("couldn't open file: %v", err)
	}
	defer file.Close()

	foods := []Food{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()

		f := Food{Ingredients: map[string]struct{}{}}
		ss := strings.Split(s, " (contains ")

		for _, ff := range strings.Fields(ss[0]) {
			f.Ingredients[ff] = struct{}{}
		}

		a := strings.Trim(ss[1], "()")
		f.Allergens = strings.Split(a, ", ")

		foods = append(foods, f)
	}
	return foods, err
}

type Food struct {
	Ingredients map[string]struct{}
	Allergens   []string
}

func remove(k string, poss map[string]map[string]struct{}) {
	for _, p := range poss {
		if len(p) > 1 {
			delete(p, k)
		}
	}
}

func firstMapKey(p map[string]struct{}) string {
	var v string
	for f, _ := range p {
		v = f
	}
	return v
}

func removeKnown(known map[string]string, poss map[string]map[string]struct{}) {
	still := false
	for a, p := range poss {
		if len(p) == 1 {
			v := firstMapKey(p)
			if _, ok := known[v]; !ok {
				still = true
				known[v] = a
				remove(v, poss)
			}
		}
	}
	if still {
		removeKnown(known, poss)
	}
}

func allKnown(all map[string]struct{}, known map[string]string) bool {
	for a, _ := range all {
		found := false
		for _, aa := range known {
			if a == aa {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func main() {
	foods, err := load()
	if err != nil {
		log.Fatalf("Couldn't load from file: %v", err)
	}

	all := map[string]struct{}{}
	ingredients := []string{}
	for _, food := range foods {
		for _, a := range food.Allergens {
			all[a] = struct{}{}
		}
		for f, _ := range food.Ingredients {
			ingredients = append(ingredients, f)
		}
	}

	// a map of an ingredient to its allergen
	known := map[string]string{}
	// a map of allergens to ingredients that might contian it
	poss := map[string]map[string]struct{}{}

	for _, food := range foods {
		for _, a := range food.Allergens {
			// if we've already seen this alergen
			if p, ok := poss[a]; ok {
				// we need to update the poss list to only include these ingredients
				for pp, _ := range p {
					if _, ok := food.Ingredients[pp]; !ok {
						delete(poss[a], pp)
					}
				}
			} else {
				// This is the first time seeing the allegen, so all the ingredients are candidates
				poss[a] = map[string]struct{}{}
				for f, _ := range food.Ingredients {
					if _, ok := known[f]; !ok {
						poss[a][f] = struct{}{}
					}
				}
			}
			removeKnown(known, poss)
		}
		if allKnown(all, known) {
			break
		}
	}

	safe := []string{}
	for _, f := range ingredients {
		if _, ok := known[f]; !ok {
			safe = append(safe, f)
		}
	}
	fmt.Println(len(safe))

	allergens := []string{}
	for a, _ := range all {
		allergens = append(allergens, a)
	}
	sort.Strings(allergens)

	badFoods := []string{}
	for _, a := range allergens {
		for i, aa := range known {
			if aa == a {
				badFoods = append(badFoods, i)
			}
		}
	}
	fmt.Println(strings.Join(badFoods, ","))
}
