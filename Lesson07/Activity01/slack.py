import json
import requests

def main(event, context):

    webhook_url = 'YOUR_INCOMMING_WEBHOOK_URL'

    response = requests.post(
        webhook_url, data=json.dumps(event['data']),
        headers={'Content-Type': 'application/json'}
    )

    if response.status_code == 200:
        return "Your message successfully sent to Slack"
    else:
        return "Error while sending your message to Slack: " + response.get('error')