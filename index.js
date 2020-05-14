function login() {
    var xmlHttp = new XMLHttpRequest();
    xmlHttp.open( "GET", "http://asymptoter-practice.nctu.me:8080/ping", false ); // false for synchronous request
    xmlHttp.send( null );
    console.log(xmlHttp.responseText);
}
