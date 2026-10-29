package Fonctions

import (
	"testing"
)

func TestQuiACoutMin(t *testing.T) {
	// Test 1 : Cas normal avec un tableau B de taille moyenne
	A := []int{7, 8}
	B := []int{2, 1, 3}
	result := QuiACoutMin(A, B)
	if result != 0 {
		t.Errorf("Test 1 échoué : QuiACoutMin(%v) = %d, attendu 0", B, result)
	}

	// Test 2 : Minimum au début de B
	A = []int{5, 6, 7}
	B = []int{1, 5, 3, 2, 8}
	result = QuiACoutMin(A, B)
	if result != 0 {
		t.Errorf("Test 2 échoué : QuiACoutMin(%v) = %d, attendu 0", B, result)
	}

	// Test 3 : Minimum à la fin de B

	A = []int{1, 3, 4, 8}
	B = []int{5, 1}
	result = QuiACoutMin(A, B)
	if result != 1 {
		t.Errorf("Test 3 échoué : QuiACoutMin(%v) = %d, attendu 1", B, result)
	}

}

func TestCout(t *testing.T) {
	// Initialisation de A une seule fois pour tous les tests
	A := []int{1, 2, 4}

	// Test 1 : Coût pour un élément au début de B
	B := []int{5, 3, 8, 42, 79, 63}
	result := Cout(A, B, 0)
	if result != 2 {
		t.Errorf("Test 1 échoué : Cout(B, %v, 0) = %d, attendu 2", B, result)
	}

	// Test 2 : Coût pour un élément milieu
	result = Cout(A, B, 2)
	if result != 4 {
		t.Errorf("Test 2 échoué : Cout(B, %v, 2) = %d, attendu 4", B, result)
	}

	// Test 3 : Coût pour un élément à la fin de B
	result = Cout(A, B, 5)
	if result != 4 {
		t.Errorf("Test 3 échoué : Cout(B, %v, 5) = %d, attendu 4", B, result)
	}
	result = Cout(A, B, 4)
	if result != 5 {
		t.Errorf("Test 3 échoué : Cout(B, %v, 4) = %d, attendu 5", B, result)
	}
}

func TestTrouverCible(t *testing.T) {
	// Test 1 : Cible pour un chiffre au début
	A := []int{1, 3, 7}
	B := []int{5, 2, 9, 8, 11}
	cible, index := TrouverCible(0, B, A)
	if cible != 7 || index != 2 {
		t.Errorf("Test 1 échoué : trouverCible(0, %v) = (%d, %d), attendu (7, 2)", A, cible, index)
	}

	// Test 2 : Cible pour un chiffre au milieu
	cible, index = TrouverCible(2, B, A)
	if cible != 1 || index != 0 {
		t.Errorf("Test 2 échoué : trouverCible(2, %v) = (%d, %d), attendu (1, 0)", A, cible, index)
	}

	// Test 3 : Cible pour un chiffre à la fin
	cible, index = TrouverCible(4, B, A)
	if cible != 1 || index != 0 {
		t.Errorf("Test 3 échoué : trouverCible(4, %v) = (%d, %d), attendu (1, 0)", A, cible, index)
	}
}
