from enum import Enum
from types import MethodType
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

# the use of Enum
def f3():
	Month = Enum('Month', ('Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun', 'Jul', 'Aug', 'Sep', 'Oct', 'Nov', 'Dec'))
	# 遍历枚举类型
	for name, member in Month.__members__.items():
		print(name, '---------', member, '----------', member.value)	
	# print(Month.__members__.items())
	# 直接引用一个常量
	print('\n', Month.Jan)

class Gender(Enum):
	Male='1'
	Female='2'
	wait_send = 'wait_send'

# for x in Gender:
# 	print(x.name, '----', x.value, '---', x)
# 	pass		
# print(Gender['wait_send'].value)

f = lambda x,y: x + y
print(list(range(1, 20)))
# print(f(, range(1, 20)))