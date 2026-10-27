package Fonctions

import (
	"testing"
)

func TestTri3OuMoins(t *testing.T) {
	// 3 elements non tries
	reponse1 := Tri3OuMoins([]int{3, 1, 2})
	if reponse1[0] != 1 || reponse1[1] != 2 || reponse1[2] != 3 {
		t.Errorf("reponse1 = %v, attendu %v", reponse1, []int{1, 2, 3})
	}

	// elements deja tries
	reponse2 := Tri3OuMoins([]int{1, 2, 3})
	if reponse2[0] != 1 || reponse2[1] != 2 || reponse2[2] != 3 {
		t.Errorf("reponse2 = %v, attendu %v", reponse2, []int{1, 2, 3})
	}

	// 2 elements
	reponse3 := Tri3OuMoins([]int{2, 1})
	if reponse3[0] != 1 || reponse3[1] != 2 {
		t.Errorf("reponse3 = %v, attendu %v", reponse3, []int{1, 2})
	}

	// 2 elements tries
	reponse4 := Tri3OuMoins([]int{1, 2})
	if reponse4[0] != 1 || reponse4[1] != 2 {
		t.Errorf("reponse4 = %v, attendu %v", reponse4, []int{1, 2})
	}

	//tableau vide
	reponse5 := Tri3OuMoins([]int{})
	if len(reponse5) != 0 {
		t.Errorf("reponse5 = %v, attendu %v", reponse5, []int{})
	}

	// tab 1 element
	reponse6 := Tri3OuMoins([]int{5})
	if reponse6[0] != 5 {
		t.Errorf("reponse6 = %v, attendu %v", reponse6, []int{5})
	}
}
