<?php
 
namespace App;
 
class Handler
{
    public function handle($data) {
        // Retrieve page name from path params
	$path_params = getenv('Http_Path');
	$path_params_array = explode('/',$path_params);
	$last_index = count($path_params_array);
	$page_name = $path_params_array[$last_index-1];
		
	// Set the page name
	$current_dir = __DIR__;
	$html_file_path = $current_dir . "/html/" . $page_name . ".html";
		
	// Read the file
	$html_file = fopen($html_file_path, "r") or die("Unable to open HTML file!");
	$html_output = fread($html_file,filesize($html_file_path));
	fclose($html_file);
		
	// Return file content
	return $html_output;	
    }
}

