const proxy = window.location.href;
const svgWidth = 900;
const svgHeight = 600;
const margin = { top: 20, right: 20, bottom: 30, left: 50 };
const width = svgWidth - margin.left - margin.right;
const height = svgHeight - margin.top - margin.bottom;
const colorPallete = [
  "#e6194B",
  "#3cb44b",
  "#ffe119",
  "#4363d8",
  "#911eb4",
  "#42d4f4",
  "#f032e6",
  "#bfef45",
  "#fabebe",
  "#469990",
  "#e6beff",
  "#9A6324",
  "#fffac8",
  "#800000",
  "#aaffc3",
  "#808000",
  "#ffd8b1",
  "#000075",
]

var delayTimer;
var g;
var line;
var availableColors = colorPallete;
var circuits;
var currCircuit = 0;

document.addEventListener("DOMContentLoaded", async (e) => {
  await getCircuits();
  await getRace();
})

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
    await getCircuits();
    await getRace();
  } else if (e.key === "ArrowLeft" || e.key === "ArrowRight") {
    await changeCircuit(e);
    await getRace();
  }
});

const getLapTimes = async (raceid, driverid) => {
  let url = proxy + "/LapTimes/" + raceid + "/" + driverid;
  let result = await ajaxRequest("GET", url);
  // Sort laptimes by lap number
  result.sort((a, b) => (a.lap > b.lap ? 1 : -1));
  return result
}

const getCircuits = async (e) => {
  let yearEle = document.getElementById('year');
  let url = proxy + "Circuits/" + yearEle.dataset.year
  let result = await ajaxRequest("GET", url);
  circuits = result.MRData.CircuitTable.Circuits;
  console.log(circuits);

  // Set first circuit for the year and reset currCircuit
  currCircuit = 0;
  let ele = document.getElementById('circuit');
  let circuit = circuits[currCircuit];
  ele.innerHTML = circuit.circuitName;
  ele.dataset.circuitid = circuit.circuitId;
}

const changeCircuit = async (e) => {
  let ele = document.getElementById('circuit');
  if (e.key === "ArrowLeft" && currCircuit != 0) {
    let arg = circuits[currCircuit--];
    ele.innerHTML = arg.circuitName;
    ele.dataset.circuitid = arg.circuitId;
  } else if (e.key === "ArrowRight" && currCircuit != circuits.length - 1) {
    let arg = circuits[currCircuit++];
    ele.innerHTML = arg.circuitName;
    ele.dataset.circuitid = arg.circuitId;
  }
}

const getRace = async () => {
  clearTimeout(delayTimer);
  delayTimer = setTimeout( async () => {
    // document.getElementById('chart').innerHTML = '';
    // document.getElementById('drivers').innerHTML = '';
    availableColors = colorPallete.slice(0)
    let circuitid = document.getElementById('circuit').dataset.circuitid;
    let year = document.getElementById('year').dataset.year;
    let url = proxy + "Results/" + year + "/" + circuitid;
    let result = await ajaxRequest("GET", url);
    if (result == 0) {
      console.log("No results for that race.")
    } else {
      await drawWordle(result);
    }
  }, 1000);
}

const ajaxRequest = async (protocol, url) => {
  const res = await fetch(url);
  return await res.json();
}
