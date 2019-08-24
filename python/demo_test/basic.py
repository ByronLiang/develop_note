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