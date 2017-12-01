class MarkerInfo extends HTMLElement {

    constructor(id, name, type, units, address, city, pic) {
        super();

        this.innerHTML = `
            <div>
                <h1><a href='/views/property.html?id=`+id+`'>` + name + `</a></h1>
                <h3> Units: ` + units + `</h3>
                <h3> Address: ` + address + `</h3>
                <h3> City: ` + city + `</h3>
                <h3> Type: ` + type + `</h3>
                <img src=../proppics/` + encodeURI(pic) + ` width='400' height='228' id='pic'>
            </div>`
    }
}

customElements.define('marker-info', MarkerInfo);