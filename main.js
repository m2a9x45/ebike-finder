// const API_URL = "data.json" 
const API_URL = "http://localhost:8080/"
const mapId = document.querySelector("#mapid");

let mymap = L.map(mapId).setView([55.951040, -3.186479], 13);

window.onload = () => {

    L.tileLayer('https://api.mapbox.com/styles/v1/mapbox/{id}/tiles/{z}/{x}/{y}?access_token=pk.eyJ1IjoibTJhOXg0NSIsImEiOiJjazgwZ3AzczcwZXJuM25tdzRpMGc0ajBkIn0.h4nHqQ9zYSLyP5bRiynMGg', {
        attribution: 'Map data &copy; <a href="https://www.openstreetmap.org/">OpenStreetMap</a> contributors, <a href="https://creativecommons.org/licenses/by-sa/2.0/">CC-BY-SA</a>, Imagery © <a href="https://www.mapbox.com/">Mapbox</a>',
        maxZoom: 18,
        id: 'streets-v11',
        accessToken: 'pk.eyJ1IjoibTJhOXg0NSIsImEiOiJjamxqeWkxc2gwZm8xM3ByMDQzOXI5Z2x3In0.0AlI9PSaNnfUKy_bT9849A',
        tileSize: 512,
        zoomOffset: -1
        }).addTo(mymap);

    fetch(API_URL)
        .then(res => res.json())
        .then((data) => {
            console.log(data);

            for (let i = 0; i < data.data.dockGroups.length; i++) {
                if (data.data.dockGroups[i].availabilityInfo.availableVehicleCategories[1].count >= 1) {
                    // this should give the info for each dock that has at least one e-bike
                    console.log(data.data.dockGroups[i]);
                    addmarker(data.data.dockGroups[i].coord.lat,data.data.dockGroups[i].coord.lng,
                        data.data.dockGroups[i].title, data.data.dockGroups[i].availabilityInfo.availableVehicleCategories[1].count);
                }     
            }
        })
        .catch((err) => console.error(err))
}

function addmarker(lat, lon, name, ebikes) {
    marker = L.marker([lat, lon],{"id":name, "ebikes":ebikes,"type":"marker"}).addTo(mymap).on("click", getLivecap)
    .bindPopup("Loading...");
}

function getLivecap(e){
    onMapClick(e);

    let popup = e.target.getPopup();
    let ebikes = this.options.ebikes;
    let stationId = this.options.id;

    popup.setContent(stationId + "<br>" + "e-bikes: " + ebikes + "<br>");
    popup.update();
    
}

function onMapClick(e) {
    console.log(e.target);
    
    const myCustomColour = '#583470'

    const markerHtmlStyles = `
    background-color: ${myCustomColour};
    width: 2rem;
    height: 2rem;
    display: block;
    left: -1.5rem;
    top: -1.5rem;
    position: relative;
    border-radius: 3rem 3rem 0;
    transform: rotate(45deg);
    border: 1px solid #FFFFFF`

    const icon = L.divIcon({
    className: "my-custom-pin",
    iconAnchor: [0, 24],
    labelAnchor: [-6, 0],
    popupAnchor: [0, -36],
    html: `<span style="${markerHtmlStyles}" />`
    })
}

mymap.on('click', onMapClick);
