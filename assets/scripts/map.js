var Map = (function() {
    var map;

    function setMap(m) {
        map = m;
    }

    function getMap() {
        return map;
    }

    return {
        setMap: setMap,
        getMap: getMap
    }
})();