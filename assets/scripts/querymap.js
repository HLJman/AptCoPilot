/*jslint devel: true */
function isChecked() {
    "use strict";
    var allPropsChecked = document.getElementById("allprop").checked;
    var mrkPropsChecked = document.getElementById("mrkprop").checked;
    var affPropsChecked = document.getElementById("affprop").checked;
    var landPropsChecked = document.getElementById("landprop").checked;
    if (allPropsChecked === false && mrkPropsChecked === false && affPropsChecked === false && landPropsChecked === false) {
        alert("No Query Selected!");
        return false;
    } else {
        var propertySelect = document.getElementsByName('proptype');
        var propertySelectValue;
        for(var i = 0; i < controlForm.length; i++){
        if(controlForm[i].checked){
        propertySelectValue = controlForm[i].value;
            
            
        var selectAllProperties = "SELECT * FROM properties";
        var theQuery = document.getElementById("hiddenQuery");    
        theQuery.value = selectAllProperties;       
                }
            }
        
        }
    
    }




