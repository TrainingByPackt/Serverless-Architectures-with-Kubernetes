import json
from sendgrid import SendGridAPIClient
from sendgrid.helpers.mail import Mail
 
def handle(req):
    
    SENDGRID_API_KEY = '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'
    TO_EMAIL = 'sathsara89@gmail.com'
    EMAIL_SUBJECT = 'New Message from OpenFaaS Contact Form'
    
    json_req = json.loads(req)
    email = json_req["email"]
    name = json_req["name"]
    message = json_req["message"]
    email_body = '<strong>Name: </strong>' + name + '<br><br> <strong>Email: </strong>' + email + '<br><br> <strong>Message: </strong>' + message
    
    email_object = Mail(
        from_email= email,
        to_emails=TO_EMAIL,
        subject=EMAIL_SUBJECT,
        html_content=email_body)
    
    try:
        sg = SendGridAPIClient(SENDGRID_API_KEY)
        response = sg.send(email_object)
        sendingStatus = "Message sent successfully"
    except Exception as e:
        sendingStatus = "Message sending failed"
    
    return sendingStatus

