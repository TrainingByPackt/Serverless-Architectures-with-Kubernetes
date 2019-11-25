import os
 
def handle(req):
 
    current_directory = os.path.dirname(__file__)
    html_file_path = os.path.join(current_directory, 'html', 'contact-us.html')
 
    with(open(html_file_path, 'r')) as html_file:
        html = html_file.read()
		
    return html

