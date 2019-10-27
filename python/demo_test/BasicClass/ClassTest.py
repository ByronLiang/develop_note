class Test(object):
    bar = 1
    
    def __init__(self):
        self._name = 'test'

    def foo(self):
        print 'foo'
 
    @staticmethod
    def static_foo():
        print 'static_foo'
        print Test.bar
 
    @classmethod
    def class_foo(cls):
        print 'class_foo'
        print cls.bar
        cls().foo()
    @property
    def get_name(self):
        return 'VIP ' + self._name

if __name__ == "__main__":
    Test.static_foo()
    Test.class_foo()
    t = Test()
    print(t.get_name)