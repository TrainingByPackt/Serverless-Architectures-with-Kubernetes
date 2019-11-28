const fetch = require('node-fetch');
const Slack = require('slack-node');

module.exports.weather = (event, context, callback) => {

    const webhookUri = process.env.SLACK_WEBHOOK_URL;
    const location = process.env.CITY;

    const slack = new Slack();
    slack.setWebhook(webhookUri);

    weatherURL = "http://wttr.in/" + encodeURIComponent(location) + "?m&&format=1"

    console.log(weatherURL)

    fetch(weatherURL)
        .then(response => response.text())
        .then(data => {

            console.log("======== WEATHER TEXT ========")
            console.error(data);
            console.log("======== WEATHER TEXT ========")

            slack.webhook({
                text: "Current weather status is " + data
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
