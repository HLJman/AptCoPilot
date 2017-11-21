var Properties = (function (PropertiesMap) {
  var properties = [];

  function searchProperties(e) {
    var val = e.value
    var resultsView = document.querySelector(".properties-search-results");

    if (val === "") {
      resultsView.style.display = 'none'
      return
    }

    var options = { extract: function (el) { return el.name } }

    var filteredProperties = fuzzy.filter(val, properties, options)

    _clearSearchResults(resultsView)
    _buildSearchResults(resultsView, filteredProperties)
  }

  function _buildSearchResults(view, data) {
    view.style.display = 'block';

    data.forEach(function(d) {
      var d = d.original
      view.innerHTML += `<div class="properties-search-result">` + d.name + `</div>`
    });
  }

  function _clearSearchResults(view) {
    view.innerHTML = "";
  }

  function loadAllProperties(data) {
    if (data == null) {
      _makeRequest('http://localhost:8081/properties/all', loadAllProperties)
      return 
    }

    properties = data;
    _buildSideList(data)
    PropertiesMap.populateMap(data)
  }

  function loadMarketRateProperties(data) {
    if (data == null) {
      _makeRequest('http://localhost:8081/properties/marketrate', loadAffordableProperties)
      return 
    }

    _buildSideList(data)
    PropertiesMap.populateMap(data)
  }

  function loadAffordableProperties(data) {
    if (data == null) {
      _makeRequest('http://localhost:8081/properties/affordable', loadAffordableProperties)
      return 
    }

    _buildSideList(data)
    PropertiesMap.populateMap(data)
  }

  function _makeRequest(url, callback) {
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

  function mapFilterProperties(f) {
    mapFilteredProperites = properties.filter(f);

    _buildSideList(mapFilteredProperites)
  }

  function _buildSideList(data) {
    console.log(data)
    var scrollView = document.querySelector('.properties-side-list')
    scrollView.innerHTML = ""

    data.forEach(function (apt, i) {
      var infowincontent = new MarkerInfo(apt.id, apt.name, apt.type,
        apt.units, apt.address, apt.city,
        apt.mainpic)

      scrollView.appendChild(infowincontent)
    });
  }

  return {
    loadAllProperties: loadAllProperties,
    mapFilterProperties: mapFilterProperties,
    searchProperties: searchProperties,
    loadAffordableProperties: loadAffordableProperties,
    loadMarketRateProperties: loadMarketRateProperties
  }
})(PropertiesMap);