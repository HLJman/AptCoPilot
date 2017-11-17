var Properties = (function () {
  var properties = [];
  var filteredProperites = [];
  var map;

  function searchProperties(e) {
    var val = e.value

    var options = { extract: function (el) { return el.name } }

    var filteredProperties = fuzzy.filter(val, filteredProperites, options)

    var scrollView = document.querySelector('.properties-search-view')
    scrollView.innerHTML = ""

    filteredProperties.map(function (apt, i) {
      apt = apt.original
      var infowincontent = new MarkerInfo(apt.id, apt.name, apt.type,
        apt.units, apt.address, apt.city,
        apt.mainpic)

      scrollView.appendChild(infowincontent)
    });

  }

  var customIcon = {
    url: "../images/mapdot.png"
  };

  function buildMarkers(data) {
    properties = data
    filteredProperites = data

    var infoWindow = new google.maps.InfoWindow;

    var scrollView =  document.querySelector('.properties-search-view');

    var markers = data.map(function (apt, i) {
      var infowincontent = new MarkerInfo(apt.id, apt.name, apt.type,
        apt.units, apt.address, apt.city,
        apt.mainpic)

      scrollView.appendChild(infowincontent)

      var marker = new google.maps.Marker({
        position: new google.maps.LatLng(parseFloat(apt.lat), parseFloat(apt.lng)),
        icon: customIcon,
      });

      marker.addListener('click', function () {
        infoWindow.setContent(infowincontent);
        infoWindow.open(map, marker);
      });

      return marker
    })

    var markerCluster = new MarkerClusterer(map, markers,
      { imagePath: 'https://cdn.rawgit.com/googlemaps/js-marker-clusterer/gh-pages/images/m' });

  };

  function loadAllProperties() {
    makeRequest('http://localhost:8081/properties/all', buildMarkers)
  }

  function loadMarketRateProperties() {
    makeRequest('http://localhost:8081/properties/marketrate', buildMarkers)
  }

  function loadAffordableProperties() {
    makeRequest('http://localhost:8081/properties/affordable', buildMarkers)
  }

  //Add border to map
  // document.getElementById("map").style.border = "1px solid #018ED0";

  // document.getElementById("map").style.marginRight = "10px";


  function makeRequest(url, callback) {
    var request = new XMLHttpRequest();
    request.open("GET", url, true);
    request.setRequestHeader("Accept", "application/json")
    request.setRequestHeader("x-pingpong", "pingpong")

    request.onload = function (e) {
      if (this.readyState === 4 && this.status === 200) {
        callback(JSON.parse(this.responseText));
      } else {
        console.error(this.statusText);
      }
    };

    request.onerror = function (e) {
      console.error(this.statusText);
    }

    request.send();
  }

  function filterProperties(f) {
    filteredProperites = filteredProperites.filter(f);

    console.log(filteredProperites)
  }

  function setMap(m) {
    map = m
  }

  return {
    loadAllProperties: loadAllProperties,
    filterProperties: filterProperties,
    setMap: setMap,
    searchProperties: searchProperties
  }
})();