# Use official PHP 7.3 image as base image
FROM php:7.3-apache
 
# Copy index.php file to the docker image
COPY index.php /var/www/html/
 
# Replace port 80 with the value from PORT environment variable in apache2 configuration files
RUN sed -i 's/80/${PORT}/g' /etc/apache2/sites-available/000-default.conf /etc/apache2/ports.conf
 
# Use the default production configuration file
RUN mv "$PHP_INI_DIR/php.ini-production" "$PHP_INI_DIR/php.ini"

