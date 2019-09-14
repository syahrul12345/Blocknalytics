from matplotlib import pyplot as plt
class Plotter:
	def plot(self,values):
		sortedValues = sorted(values.items())
		x,y = zip(*sortedValues)
		plt.plot(x,y)
		plt.savefig('transactionHistoryGraph')
		