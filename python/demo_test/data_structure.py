def f1():
    print('In the %d year, the growth rate is %.2f %%' % (2019, 4.245))

def f2():
    a_tuple = (2, 3)
    # 逗号默认建立元组数据结构
    b_tuple = 1,2,3,4,5
    # 元组无法变更元素数据
    # a_tuple[0] = 1
    a_list = [2, 3]
    # list可对指定下标数据进行变更数据操作
    a_list[0] = 1
    # 对list进行追加数据
    a_list.append(4)

    print(b_tuple)
    print(a_list)

f2()