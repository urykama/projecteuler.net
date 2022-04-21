# aviso.bz


def pr(q):
    for aa in range (2,q):
        if q/aa == q//aa:
            return False
    return True


w = 1
ini = 8
q=w


print ( q, " / ", w, " = ", 60*60/300*0.262)
print ( q, " / ", w, " = ", 60*60/10*0.029)
print ( q, " / ", w, " = ", 60*60/30*0.044)
print ( q, " / ", w, " = ", 60*60/180*0.159)
print (1.276 + 0.262*4+0.159)