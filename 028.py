#Number spiral diagonals
#Problem 28
spiral = 1001
rr = int(spiral/2+1)
print (rr)
r = 1
s = 1
n4 = 1
while r < rr:	
	n1 = n4 + 2 * r
	n2 = n1 + 2 * r
	n3 = n2 + 2 * r
	n4 = n3 + 2 * r
	s = s + n1 + n2 + n3 + n4
	print (r, n1, n2, n3, n4, s)
	r = r + 1	
	pass
