const proxy = "http://localhost:3005"

document.addEventListener("keydown", async (e) => {
  if (e.key === "ArrowUp" || e.key === "ArrowDown") {
    let ele = document.getElementById('year');
    e.key === "ArrowUp" ? ele.innerHTML++ : ele.innerHTML--;
  } else if (e.key === "ArrowLeft" || e.key === "ArrowRight"){
      await changeCircuit(e);
  }
});

const changeCircuit = async (e) => {
  let ele = document.getElementById('circuit');
  if (e.key === "ArrowLeft" && ele.dataset.circuitid != 1) {
    ele.innerHTML = "prev"; // Remove this after adding db logic
    let arg = --ele.dataset.circuitid;
    let url = proxy + "/Circuit/" + arg;
    let result = await ajaxRequest("GET", url);
    ele.innerHTML = result.name;
  } else if (e.key === "ArrowRight" && ele.dataset.circuitid != 73) {
    let arg = ++ele.dataset.circuitid;
    let url = proxy + "/Circuit/" + arg;
    let result = await ajaxRequest("GET", url);
    ele.innerHTML = result.name;
  }
}

const ajaxRequest = async (protocol, url) => {
  console.log(protocol, url);
  const res = await fetch(url);
  return await res.json();
}
