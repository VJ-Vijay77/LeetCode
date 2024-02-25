package main

import "fmt"

func main() {
	words := []string{"gin","zen","gig","msg"}
	// Output: 2
	// Explanation: The transformation of each word is:
	// "gin" -> "--...-."
	// "zen" -> "--...-."
	// "gig" -> "--...--."
	// "msg" -> "--...--."
	// There are 2 different transformations: "--...-." and "--...--.".
	fmt.Println("Result := ",uniqueMorseRepresentations(words))
}


func uniqueMorseRepresentations(words []string) int {
	morseMap := map[rune]string{
		'a': ".-", 'b': "-...", 'c': "-.-.", 'd': "-..", 'e': ".", 'f': "..-.",
		'g': "--.", 'h': "....", 'i': "..", 'j': ".---", 'k': "-.-", 'l': ".-..",
		'm': "--", 'n': "-.", 'o': "---", 'p': ".--.", 'q': "--.-", 'r': ".-.",
		's': "...", 't': "-", 'u': "..-", 'v': "...-", 'w': ".--", 'x': "-..-",
		'y': "-.--", 'z': "--..",
	}
	
	newArr := make(map[string]bool) 
	for _,v := range words{
		str := ""
		for _,k := range v{
			str += morseMap[k]
		}
		if _,ok := newArr[str]; !ok{
			newArr[str] = true
		}
	}
	return len(newArr)
}