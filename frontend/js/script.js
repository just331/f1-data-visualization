const proxy = "http://localhost:3005"

document.addEventListener("keydown", async (e) => {
  if (e.key === "ArrowUp" || e.key === "ArrowDown") {
    let ele = document.getElementById('year');
    e.key === "ArrowUp" ? ele.innerHTML++ : ele.innerHTML--;
  } else if (e.key === "ArrowLeft" || e.key === "ArrowRight") {
    let ele = document.getElementById('circuit');
    if (e.key === "ArrowLeft") {
      ele.innerHTML = "prev"; // Remove this after adding db logic
      let arg = ele.dataset.circuitid--;
      let url = proxy + "/Circuit/" + arg.toString();
      ajaxRequest("GET", url)
    } else {
      ele.innerHTML = "next"; // Remove this after adding db logic
      let arg = ele.dataset.circuitid++;
      let url = proxy + "/Circuit/" + arg.toString();
      let result = await ajaxRequest("GET", url);
      console.log(result);
    }
  }
});

const ajaxRequest = async (protocol, url) => {
  console.log(protocol, url);
  let httpRequest = new XMLHttpRequest();
  if (!httpRequest) {
    alert('Can not create XMLHTTP instance');
    return false;
  }
  httpRequest.onreadystatechange = () => {

     if (httpRequest.readyState === 4) {
        // Javascript function JSON.parse to parse JSON data
        return JSON.parse(httpRequest.responseText);
     }
  }
  httpRequest.open(protocol, url, true);
  httpRequest.send();
}
