const proxy = "http://localhost:3005"
var delayTimer;

document.addEventListener("keydown", async (e) => {
  if (e.key === "ArrowUp" || e.key === "ArrowDown") {
    let ele = document.getElementById('year');
    if (e.key === "ArrowUp") {
      ++ele.innerHTML;
      ++ele.dataset.year;
    } else {
      --ele.innerHTML;
      --ele.dataset.year;
    }
    await getRace();
  } else if (e.key === "ArrowLeft" || e.key === "ArrowRight") {
      await changeCircuit(e);
      await getRace();
  }
});

const changeCircuit = async (e) => {
  let ele = document.getElementById('circuit');
  if (e.key === "ArrowLeft" && ele.dataset.circuitid != 1) {
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

const getRace = async () => {
  clearTimeout(delayTimer);
  delayTimer = setTimeout( async () => {
    let circuitid = document.getElementById('circuit').dataset.circuitid;
    let year = document.getElementById('year').dataset.year;
    let url = proxy + "/Races/" + year + "/" + circuitid;
    let result = await ajaxRequest("GET", url);
    console.log(result);
    getResults(result);
  }, 500);
}

const getResults = async (race) => {
  let url = proxy + "/Results/" + race.raceId;
  let result = await ajaxRequest("GET", url);
  if (result == 0) {
    console.log("No results for that race.")
  } else {
    console.log(result);
  }
}

const ajaxRequest = async (protocol, url) => {
  console.log(protocol, url);
  const res = await fetch(url);
  return await res.json();
}
