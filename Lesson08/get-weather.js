const request = require('request');
 
function main(params) {
 
    const cityName = params.cityName
    const openWeatherApiKey = '3855a05d3fa3171bbb852a6b44087108';
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
                    resolve(weatherData);
                } catch (ex) {
                    reject('Error occurred while parsing weather data.');
                }
            }
            
        });
    });
}
