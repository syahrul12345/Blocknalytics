from get import InfoGetter
from plot import Plotter
import aiohttp
import asyncio
import json


async def main():
	transactionCounts99 = {}
	plotter = Plotter()
	infoGetter = InfoGetter("https://adoring-snyder:humped-muster-device-mousy-bauble-appear@nd-806-802-183.p2pify.com")
	latestBlock = infoGetter.getLatestTransactions()
	tasks =[]
	async with aiohttp.ClientSession() as session:
		for selectedBlock in range(int(latestBlock,16)-100,int(latestBlock,16)):
			task = asyncio.ensure_future(infoGetter.getTransactions(session,hex(selectedBlock)))
			tasks.append(task)
		
		responses = await asyncio.gather(*tasks)
		for response in responses:
			valuesAndKey = next(iter(response.items()))
			transactionCounts99[valuesAndKey[0]] = valuesAndKey[1]			
		#we've completed the request, so now we can plot
		plotter.plot(transactionCounts99)



loop = asyncio.get_event_loop()
loop.run_until_complete(main())

# transactionCounts99[selectedBlock] = len(json.loads(response)["result"]["transactions"])
