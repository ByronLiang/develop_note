def f(a, L=[]):
    L.append(a)
    return L

def f1(a, L=None):
    if L is None:
        L = []
    L.append(a)
    return L

def f2():
    words = ['monkey', 'cat', 'dog', 'birds']
    for word in words:
        print(word)

def f3(a):
	''' if sample '''
	if a > 5:
		print('aa')
	elif a > 1 and a < 5:
		print('xx')
	else:
		print('df')

f3(-2)