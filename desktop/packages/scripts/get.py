import requests
import asyncio
from aiohttp import ClientSession
import json

class InfoGetter:
	def __init__(self,ethRPC):
		self.ethRPC = ethRPC
	
	def getRPC(self):
		return self.ethRPC

	def getLatestTransactions(self):
		payload = {"jsonrpc":"2.0","method":"eth_blockNumber","params":[],"id":1}
		r = requests.post(self.ethRPC,json=payload)
		return json.loads(r.text)["result"]


	async def getTransactions(self,session,blockNumber):
		payload = {"jsonrpc":"2.0","method":"eth_getBlockByNumber","params":[blockNumber, True],"id":1}
		async with session.post(self.ethRPC,json=payload) as r:
			response = await r.text()
			return {int(blockNumber,16):len(json.loads(response)["result"]["transactions"])}
	