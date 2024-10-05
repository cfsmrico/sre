import http.client

connection = http.client.HTTPSConnection("www.fitcalcs.life")
connection.request("GET", "/")
response = connection.getresponse()
statusCode = response.status
data = response.read().decode("utf-8")
print(f' * * * Status Code: {statusCode} * * *')
print(' * * * Data: --v-- * * * ')
print(data)
