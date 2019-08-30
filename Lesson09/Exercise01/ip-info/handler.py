import requests
import json
 
def handle(req):
    api_response = requests.get('https://httpbin.org/ip')
    json_object = api_response.json()
    origin_ip = json_object["origin"]
 
    return "Origin IP is " + origin_ip

