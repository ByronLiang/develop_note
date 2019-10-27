import random
import string

# @property and @property.setter
class Screen(object):
	@property	
	def width():
		return self._width;
	@width.setter
	def width(self, value):
		if not isinstance(value, int):
			raise ValueError('score must be an integer!')
		self._width = value
	@property
	def height(self):
		return self._height
	@height.setter
	def height(self, value):
		if not isinstance(value, int):
			raise ValueError('score must be an integer!')
		self._height = value
	@property
	def resolution(self):
		return self._width * self._height
		pass

s = Screen()
# 赋值时调用属性的校验方法 @width.setter
s.width = 1024
# @height.setter
s.height = 768
print('resolution =', s.resolution)

# create random string in define length (ex: 32)
random_txt = ''.join(random.sample(string.ascii_letters + string.digits, 32))
print(random_txt)
