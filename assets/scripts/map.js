var PropertiesMap = (function () {
    var map;
    var markers = [];
    var markerClusterer;

    function initMap(mapChanged) {
        map = new google.maps.Map(document.getElementById('map'), {
            center: new google.maps.LatLng(40.354528, -111.757338),
            zoom: 7
        });

        map.addListener('bounds_changed', function () {
            var bounds = this.getBounds()
            mapChanged(function (prop) {
                return bounds.contains(new google.maps.LatLng(parseFloat(prop.lat), parseFloat(prop.lng)))
            })
        });
    }

    function populateMap(properties) {
        _clearMarkers();
        _buildMarkers(properties);
    }

    function _buildMarkers(properties) {
        var customIcon = {
            url: "../images/mapdot.png"
        };

        var infoWindow = new google.maps.InfoWindow;

        markers = properties.map(function (apt, i) {
            var infowincontent = new MarkerInfo(apt.id, apt.name, apt.type,
                apt.units, apt.address, apt.city,
                apt.mainpic)

            var marker = new google.maps.Marker({
                map: map,
                position: new google.maps.LatLng(parseFloat(apt.lat), parseFloat(apt.lng)),
                icon: customIcon,
            });

            marker.addListener('click', function () {
                infoWindow.setContent(infowincontent);
                infoWindow.open(map, marker);
            });

            return marker
        })

        // markerClusterer = new MarkerClusterer(map, markers,
        //     { imagePath: 'https://cdn.rawgit.com/googlemaps/js-marker-clusterer/gh-pages/images/m' });
    };

    function _setMapOnAll(map) {
        for (var i = 0; i < markers.length; i++) {
            markers[i].setMap(map);
          }
    }

    function _clearMarkers() {
        if (markerClusterer) {
            markerClusterer.removeMarkers(markers)
        }
        _setMapOnAll(null)
    }

    return {
        initMap: initMap,
        populateMap: populateMap
    }
})();