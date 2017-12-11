var Properties = (function () {
  var properties = [];

  function loadAllProperties(callback) {
    callback(JSON.parse('[{"id":1,"name":"Name","address":"Address","lat":0,"lng":0,"type":"Type","city":"City","units":0,"county":"County","mainpic":"PictureName"},{"id":2,"name":"Heritage Magna","address":"3544 South Kingsburg Cove","lat":40.69519,"lng":-112.086815,"type":"Affordable","city":"Magna","units":80,"county":"","mainpic":"Heritage Magna sm.jpg"},{"id":3,"name":"Solara Renovated","address":"780 North 900 West","lat":0,"lng":0,"type":"","city":"Salt Lake City","units":418,"county":"","mainpic":""},{"id":4,"name":"Deer Creek Village","address":"358 North Gateway Drive","lat":41.715027,"lng":-111.8286,"type":"Market Rate","city":"Providence","units":96,"county":"","mainpic":"Deer Creek.jpg"},{"id":5,"name":"Harris  (Tremonton)","address":"350 West 550 North","lat":41.71566,"lng":-112.17346,"type":"","city":"Tremonton","units":60,"county":"BOX ELDER","mainpic":"Kelly \u0026 Julie 4plexes.jpg"},{"id":6,"name":"Brittany Green","address":"460 West Westland Drive","lat":41.487434,"lng":-112.02489,"type":"Affordable","city":"Brigham City","units":90,"county":"BOX ELDER","mainpic":"Brittany Green.jpg"},{"id":7,"name":"Canyon Cove (Brigham)","address":"848 West 1075 South","lat":41.48779,"lng":-112.03002,"type":"Market Rate","city":"Brigham City","units":60,"county":"BOX ELDER","mainpic":"canyon cove.jpg"},{"id":8,"name":"Continental","address":"729 East 900 North","lat":41.74835,"lng":-111.81571,"type":"Student","city":"Logan","units":35,"county":"CACHE","mainpic":"Continental, Logan.jpg"},{"id":9,"name":"Fairfield Townhomes","address":"1121 North 360 West","lat":41.724903,"lng":-111.8516,"type":"","city":"Logan","units":50,"county":"CACHE","mainpic":"Fairfield Townhomes.jpg"},{"id":10,"name":"Pine View","address":"780 East 1000 North","lat":41.750015,"lng":-111.81453,"type":"Student","city":"Logan","units":50,"county":"CACHE","mainpic":"Pine View Logan.jpg"},{"id":11,"name":"Spring Hollow","address":"1300 North 200 East","lat":41.75567,"lng":-111.82834,"type":"Affordable","city":"Logan","units":50,"county":"CACHE","mainpic":"Spring Hollow.jpeg"},{"id":12,"name":"Austin Ridge","address":"760 East 900 North","lat":41.747997,"lng":-111.81437,"type":"Student","city":"Logan","units":40,"county":"CACHE","mainpic":""},{"id":13,"name":"Cambridge Court","address":"590 Canyon Road","lat":41.736073,"lng":-111.81973,"type":"Student","city":"Logan","units":35,"county":"CACHE","mainpic":"Cambridge Court sm.jpg"},{"id":14,"name":"Forest Gate","address":"454 North 400 East","lat":41.740005,"lng":-111.82442,"type":"Student","city":"Logan","units":56,"county":"CACHE","mainpic":"Forest Gate.jpg"},{"id":15,"name":"North Pointe","address":"1580 North 200 East","lat":41.734818,"lng":-111.812706,"type":"Affordable","city":"Logan","units":80,"county":"CACHE","mainpic":"North Pointe Logan.jpg"},{"id":16,"name":"Oakridge","address":"1355 North 800 East","lat":41.756428,"lng":-111.81513,"type":"Student","city":"Logan","units":150,"county":"CACHE","mainpic":"Oakridge in logan.jpg"},{"id":17,"name":"Old Farm","address":"777 East  1000 North","lat":41.750813,"lng":-111.816,"type":"Student","city":"Logan","units":97,"county":"CACHE","mainpic":"Old Farm.jpg"},{"id":18,"name":"Pineview","address":"780 East 1000 North","lat":41.749714,"lng":-111.814674,"type":"Student","city":"Logan","units":50,"county":"CACHE","mainpic":"Pine View.jpg"},{"id":19,"name":"Riverwalk","address":"781 South Riverwalk Parkway","lat":41.717976,"lng":-111.83952,"type":"Affordable; Senior","city":"Logan","units":88,"county":"CACHE","mainpic":"Riverwalk logan.jpg"},{"id":20,"name":"Spectrum","address":"760 East 900 north","lat":41.747833,"lng":-111.815,"type":"Student","city":"Logan","units":36,"county":"CACHE","mainpic":"Spectrum.jpg"}]'))
    // _makeRequest(config.apibaseurl + '/properties', callback)
  }


  function loadMarketRateProperties(callback) {
    _makeRequest(config.apibaseurl + '/properties?type=marketrate', callback)
  }

  function loadAffordableProperties(callback) {
    _makeRequest(config.apibaseurl + '/properties?type=affordable', callback)
  }


  function _makeRequest(url, callback) {
    var request = new XMLHttpRequest();
    request.open("GET", url, true);
    request.setRequestHeader("Accept", "application/json")
    request.setRequestHeader("x-pingpong", "pingpong")

    var that = this;
    request.onload = function (e) {
      if (this.readyState === 4 && this.status === 200) {
        var data = JSON.parse(this.responseText);
        properties = data
        callback(data);
      } else {
        console.error(this.statusText);
      }
    };

    request.onerror = function (e) {
      console.error(this.statusText);
    }

    request.send();
  }


  return {
    loadAllProperties: loadAllProperties,
    // mapFilterProperties: mapFilterProperties,
    // searchProperties: searchProperties,
    loadAffordableProperties: loadAffordableProperties,
    loadMarketRateProperties: loadMarketRateProperties
  }
})();