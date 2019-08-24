from Animal import Pet

# 继承抽象类, 多态处理
class MyPet(Pet.Pet):
    
    def __init__(self, name, age):
        super().__init__(name)
        self._age = age
    
    def habit(self):
        print('like eating fish')

    def make_vocie(self):
        print('miao miao~~')
    
p = MyPet('Bob', 1)
p.habit()
p.make_vocie()