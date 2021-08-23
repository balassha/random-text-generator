package generator

import (
	"random-text-generator/trigram"
)

type Solution struct {
	dictionary map[string][]string
	weightage  map[string]int
}

func New() *Solution {
	return &Solution{
		// TODO: initialize your store here
		dictionary: make(map[string][]string),
		weightage:  make(map[string]int),
	}
}

func (s *Solution) Add(t trigram.Trigram) {
	s.dictionary[t.Word1+" "+t.Word2] = append(s.dictionary[t.Word1+" "+t.Word2], t.Word3)
	s.weightage[t.Word1+" "+t.Word2+" "+t.Word3]++
}

func (s *Solution) Generate(numWords int) (string, error) {
	var response string
	var count int
	for k, v := range s.dictionary {
		if count < 1 {
			response = response + k
			count++
		}
		var max int
		var word string
		for _, word3 := range v {
			value := s.weightage[k+" "+word3]
			if value > max {
				max = value
				word = word3
			}
		}
		//response = fmt.Sprintf(response, word)
		response = response + " " + word
	}

	return response, nil
}
