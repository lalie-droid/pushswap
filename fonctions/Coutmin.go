package Fonctions

func QuiACoutMin (A []int) int {

	minCout =cout(A, 0) 			//appel cout A[0]
		for i:=0; i<len(A); i++{
 			cout(A, i)
			if cout(A[i]) < minCout {
				minCout = cout(A[i])
			}
		}
		return i
}


func cout (A []int, i int) {
cout:=0
chiffre := A[i]
cible:= trouverCible(chiffre , A)
	if chiffre =< len(a)/2
	cout += i  // pas i +1 car [1,5,2,7,9,8,3] 2 prend 2 couts a etre 1er
	else if chiffre > len(a)/2
	cout = (len(A)-i)+1

	if cible =< len(a)/2
	cout += i  
	else if cible > len(a)/2
	cout = (len(A)-i)+1

	
return cout +1
}




func trouverCible (chiffre int, A []int) 
	cible :=max[A]
	for i range len(A); i++
		if A[i] < chiffre 
			continue
		if A[i]-chiffre < cible-chiffre
			cible = A[i]
	return cible
