from random import randint

def roll_dice(n=2):
    total = 0
    lists = []
    for index in range(n):
        num = randint(1, 6)
        total += num
        lists.append(num)
    return total, lists


def add(a=0, b=0, c=0):
    return a + b + c

# Python查找一个变量时会按照“局部作用域”、“嵌套作用域”、“全局作用域”和“内置作用域”的顺序进行搜索
# 无法从下一级去获取变量
def foo():
    # 嵌套作用域
    b = 'hello'

    def bar():
        # 局部作用域
        c = True
        print(a)
        print(b)
        print(c)

    bar()
    # print(c)  # NameError: name 'c' is not defined

def f3(a):
	''' if sample '''
	if a > 5:
		print('aa')
	elif a > 1 and a < 5:
		print('xx')
	else:
		print('df')


if __name__ == '__main__':
    a = 100
    # print(b)  # NameError: name 'b' is not defined
    foo()
    f3(-2)
# res = roll_dice(2)
# print(res)