class Book:
	"""doc string for Book"""
	def __init__(self, name):
		self.name = name
		self.version = []

	def add_version(self, version):
		self.version.append(version)

	def get_name(self):
		return self.name


		