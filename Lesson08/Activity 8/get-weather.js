const request = require('request');
 
function main(params) {
 
    const cityName = params.cityName
    const openWeatherApiKey = '<OPEN_WEATHER_API_KEY>';
    const openWeatherUrl = 'https://api.openweathermap.org/data/2.5/weather?q=' + cityName + '&mode=json&units=metric&appid=' + openWeatherApiKey ;
 
    return new Promise(function(resolve, reject) {
 
        request(openWeatherUrl, function(error, response, body) {   
 
            if (error) {
                reject('Requesting weather data from provider failed ' 
                       + 'with status code ' 
                       + response.statusCode + '.\n' 
                       + 'Please check the provided cityName argument.');
            } else {
                try {
                    var weatherData = JSON.parse(body);
                    resolve({weatherData:weatherData});
                } catch (ex) {
                    reject('Error occurred while parsing weather data.');
                }
            }
            
        });
    });
}
