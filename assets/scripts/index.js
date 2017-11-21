var Main = (function(Properties, PropertiesMap) {

    function init() {
        PropertiesMap.initMap(Properties.mapFilterProperties);

        Properties.loadAllProperties();
    }

    return {
        init: init
    }

})(Properties, PropertiesMap);