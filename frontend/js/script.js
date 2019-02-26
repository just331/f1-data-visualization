const proxy = "http://localhost:3005"

document.addEventListener("keydown", function(e) {
  if (e.key === "ArrowUp" || e.key === "ArrowDown") {
    var ele = document.getElementById('year');
    e.key === "ArrowUp" ? ele.innerHTML++ : ele.innerHTML--;
  } else if (e.key === "ArrowLeft" || e.key === "ArrowRight") {
    var ele = document.getElementById('circuit');
    if (e.key === "ArrowLeft") {
      ele.innerHTML = "prev"; // Remove this after adding db logic
      var arg = ele.dataset.circuitid++;
      var url = proxy + "/GetRace" + arg.toString();
      ajaxRequest("GET", url)
    } else {
      ele.innerHTML = "next"; // Remove this after adding db logic
      var arg = ele.dataset.circuitid--;
      var url = proxy + "/GetRace" + arg.toString();
      console.log(ajaxRequest("GET", url))
    }
  }
});

var ajaxRequest = function (protocol, url) {
  console.log(protocol, url)
  var httpRequest = new XMLHttpRequest();
  if (!httpRequest) {
    alert('Can not create XMLHTTP instance');
    return false;
  }
  httpRequest.open(protocol, url);
  httpRequest.send();
  return httpRequest.responseXML;
}
