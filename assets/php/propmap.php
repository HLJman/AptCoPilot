<?php
require("connect.php");

$dom = new DOMDocument("1.0");
$node = $dom->createElement("markers");
$parnode = $dom->appendChild($node);

// Opens a connection to a MySQL server
$conn = new mysqli($servername, $username, $password);

// Set the active MySQL database
mysqli_select_db($conn, $database);

  

// Select all the rows in the markers table
$result = mysqli_query($conn, "SELECT * FROM properties");


header("Content-type: text/xml");

while ($row = @mysqli_fetch_assoc($result)){
  // Add to XML document node
  $node = $dom->createElement("marker");
  $newnode = $parnode->appendChild($node);
  $newnode->setAttribute("id",$row['id']);
  $newnode->setAttribute("name",$row['name']);
  $newnode->setAttribute("address", $row['address']);
  $newnode->setAttribute("lat", $row['lat']);
  $newnode->setAttribute("lng", $row['lng']);
  $newnode->setAttribute("type", $row['type']);
  $newnode->setAttribute("city", $row['city']);
  $newnode->setAttribute("units", $row['units']); 
  $newnode->setAttribute("county", $row['county']);
  $newnode->setAttribute("mainpic", $row['mainpic']);
}

echo $dom->saveXML();

?>