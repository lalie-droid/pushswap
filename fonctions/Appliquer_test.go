package Fonctions

import "testing"

func TestAppliquer(t *testing.T) {

	//test basique
	A := []int{1, 2, 3}
	B := []int{4, 5}
	Appliquer(A, B, 0)
	if len(A) != 4 || A[1] != 1 || A[2] != 2 || A[3] != 3 || A[0] != 4 {
		t.Errorf("Test 1 échoué : A = %v, attendu [1 2 3 4]", A)
	}
	if len(B) != 1 || B[0] != 5 {
		t.Errorf("Test 1 échoué : B = %v, attendu [5]", B)
	}
	// test
	A = []int{1, 2, 3}
	B = []int{0}
	Appliquer(A, B, 0)
	if len(A) != 4 || A[0] != 0 || A[1] != 1 || A[2] != 2 || A[3] != 3 {
		t.Errorf("Test 2 échoué : A = %v, attendu [0 1 2 3]", A)
	}
	if len(B) != 0 {
		t.Errorf("Test 2 échoué : B = %v, attendu []", B)
	}

	A = []int{1, 2, 3}
	B = []int{4, 5, 12}
	Appliquer(A, B, 3)
	if len(B) != 4 || B[0] != 1 || B[1] != 2 || B[2] != 3 || B[3] != 12 {
		t.Errorf("Test 1 échoué : B = %v, attendu [1 2 3 12]", B)
	}
	if len(A) != 2 || A[0] != 4 || A[1] != 5 {
		t.Errorf("Test 1 échoué : A = %v, attendu [4 5]", A)
	}
}
