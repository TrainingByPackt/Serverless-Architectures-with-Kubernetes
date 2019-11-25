function main(params) {
 
    return new Promise(function(resolve, reject) {
 
        if (!params.weatherData) {
            reject("Weather data not provided");
        }
 
        const weatherData = params.weatherData;
        const cityName = weatherData.name;
        const currentTemperature = weatherData.main.temp;
 
        weatherMessage = "It's " + currentTemperature
                                 + " degrees celsius in " + cityName;
 
        resolve({message: weatherMessage});
 
   });
}
