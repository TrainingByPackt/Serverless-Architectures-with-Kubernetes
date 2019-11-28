import datetime as dt
def main(event, context):
    currentHour = dt.datetime.now().hour
    greetingMessage = ''
    if currentHour < 12:
        greetingMessage = 'Hello, Good morning!'
    elif currentHour < 18:
        greetingMessage = 'Hello, Good afternoon!'
    else:
        greetingMessage = 'Hello, Good evening!'
    return greetingMessage
	
	
$ kubectl create namespace lesson-7
$ kubeless function deploy greetings --runtime python3.7 \
--from-file greetings.py \
--handler greetings.main \
--namespace lesson-7