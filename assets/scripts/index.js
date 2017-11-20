var Main = (function(Properties, PropertiesMap) {

    function init() {
        PropertiesMap.initMap(Properties.mapFilterProperties);

        Properties.loadAllProperties(_propertiesLoaded);
    }

    function _propertiesLoaded(data) {
        Properties.setProperties(data)
        PropertiesMap.populateMap(data);
    }

    return {
        init: init
    }

})(Properties, PropertiesMap);