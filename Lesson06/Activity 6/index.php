<?php
 
if ( !isset ( $_GET['timezone'] ) ) {
    // Returns error if the timezone parameter is not provided
    $output_message = "Error: Timezone not provided"; 
} else if ( empty ( $_GET['timezone'] ) ) {
    // Returns error if the timezone parameter value is empty
    $output_message = "Error: Timezone cannot be empty"; 
} else {
    // Save the timezone parameter value to a variable
    $timezone = $_GET['timezone'];
    
    try {
        // Generates the current time for the provided timezone
        $date = new DateTime("now", new DateTimeZone($timezone) );
        $formatted_date_time = $date->format('Y-m-d H:i:s');
        $output_message = "Current date and time for $timezone is $formatted_date_time";
    } catch(Exception $e) {
        // Returns error if the timezone is invalid
        $output_message = "Error: Invalid timezone value"; 
    }
 
}
 
// Return the output message
echo $output_message;

