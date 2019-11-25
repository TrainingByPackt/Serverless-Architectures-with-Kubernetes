const sendGridMailer = require('@sendgrid/mail');
 
function main(params) {
 
    const sendGridApiKey = '<SEND_GRID_API_KEY>';
    const toMail = '<TO_EMAIL>';
    const fromMail = '<FROM_EMAIL>';
    const mailSubject = 'Weather Information for Today';
    const mailContent = params.message;
    
    return new Promise(function(resolve, reject) {
 
        sendGridMailer.setApiKey(sendGridApiKey);
 
        const msg = {
            to: toMail,
            from: fromMail,
            subject: mailSubject,
            text: mailContent,
        };
 
        sendGridMailer.send(msg, (error, result) => {
            if (error) {
                reject({msg: "Message sending failed."});
            } else {
                resolve({msg: "Message sent!"});
            }
        });       
       
   });
}

exports.main = main;