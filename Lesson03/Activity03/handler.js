const fetch = require('node-fetch');
const Slack = require('slack-node');

module.exports.joker = (event, context, callback) => {

    const webhookUri = process.env.SLACK_WEBHOOK_URL;

    const slack = new Slack();
    slack.setWebhook(webhookUri);

    fetch('https://icanhazdadjoke.com/slack')
        .then(response => response.json())
        .then(data => {

            joke = data.attachments[0]

            console.log("======== JOKE TEXT ========")
            console.error(joke.text);
            console.log("======== JOKE TEXT ========")

            slack.webhook({
                text: joke.text
            }, function(err, response) {
                console.log("======== SLACK SEND STATUS ========")
                console.error(response.status);
                return callback(null, {statusCode: 200, body: "ok" });
                console.log("======== SLACK SEND STATUS ========")

                if (err) {
                    console.log("======== ERROR ========")
                    console.error(error);
                    console.log("======== ERROR ========")
                    return callback(null, {statusCode: 500, body: JSON.stringify({ error}) });
                }
            });

        }).catch((error) => {
            console.log("======== ERROR ========")
            console.error(error);
            console.log("======== ERROR ========")
             return callback(null, {statusCode: 500, body: JSON.stringify({ error}) });
        });
};
