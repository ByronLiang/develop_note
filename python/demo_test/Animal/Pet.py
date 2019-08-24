from abc import ABCMeta, abstractmethod

# 抽象类
class Pet(object, metaclass=ABCMeta):
    
    def __init__(self, name):
        self._name = name
    
    # 抽象方法
    @abstractmethod
    def habit(self):
        pass