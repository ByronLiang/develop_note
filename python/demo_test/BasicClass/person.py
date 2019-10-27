class person(object):
    # 定义__slots__变量来进行限定。需要注意的是__slots__的限定只对当前类的对象生效，对子类并不起任何作用。
    # 限定Person对象只能绑定_name, _age和_gender属性
    __slots__ = ('_name', '_age', '_gender')

    def __init__(self, name):
        self._gender = ''
        self._name = name
    
p = person('jack')
p._age = 19
p._gender = 'male'
print(p._name, p._age)