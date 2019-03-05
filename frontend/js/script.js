const proxy = "http://localhost:3005"
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

document.addEventListener("DOMContentLoaded", async (e) => {
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
    await getRace();
  } else if (e.key === "ArrowLeft" || e.key === "ArrowRight") {
      await changeCircuit(e);
      await getRace();
  }
});

const drawAxes = (laptimes) => {
  let svg = d3.select("#chart")
    .attr("width", svgWidth)
    .attr("height", svgHeight);

  let g = svg.append("g")
    .attr("transform",
      "translate(" + margin.left + "," + margin.top + ")"
    )
    .attr("class", "axis");

  var x = d3.scaleLinear().rangeRound([0, width]);
  var y = d3.scaleLinear().rangeRound([height, 0]);
  x.domain(d3.extent(laptimes, function(d) { return d.lap }));
  y.domain(d3.extent(laptimes, function(d) { return d.milliseconds }));


  let line = d3.line()
    .x(function(d) { return x(d.lap)})
    .y(function(d) { return y(d.milliseconds)});

   g.append("g")
    .attr("transform", "translate(0," + height + ")")
    .call(d3.axisBottom(x))

   g.append("text")
    .attr("transform",
          "translate(" + (width/2) + " ," +
                         (height + margin.top) + ")")
    .style("text-anchor", "middle")
    .text("Lap");

   g.append("g")
    .call(d3.axisLeft(y));

   g.append("text")
    .attr("transform", "rotate(-90)")
    .attr("y", 6)
    .attr("dy", ".71em")
    .style("text-anchor", "end")
    .text("Time (ms)");

  return {line, g, x, y};
}

const drawChart = async (results) => {
  // Sort results by final position
  results.sort((a, b) => (a.position > b.position ? 1 : -1));

  // Get the top three drivers by default
  for (let i = 0; i < 3; i++) {
    result = results[i];
    console.log(result);
    let laptimes = await getLapTimes(result.raceId, result.driverId);
    if (i == 0) {
      let axes = drawAxes(laptimes);
      g = axes.g;
      line = axes.line;
      x = axes.x;
      y = axes.y;
    }

    let lineColor = availableColors.pop();

    g.append("path")
      .datum(laptimes)
      .attr("id", "driver" + result.driverId)
      .attr("data-driverid", result.driverId)
      .attr("fill", "none")
      .attr("stroke", lineColor)
      .attr("stroke-linejoin", "round")
      .attr("stroke-linecap", "round")
      .attr("stroke-width", 1.0)
      .attr("d", line);

    g.selectAll(".dot")
      .data(laptimes)
    .enter().append("circle")
      .attr("fill", lineColor)
      .attr("cx", function(d) { return x(d.lap) })
      .attr("cy", function(d) { return y(d.milliseconds) })
      .attr("r", 2);
  }

  // Populate driver dropdown
  let driverDropdown = document.getElementById('drivers');
  for (let result of results) {
    let ele = document.createElement("div");
    ele.dataset.driverid = result.driverId;
    ele.dataset.raceid = result.raceId;
    ele.className = "w3-bar-item w3-button";

    // Get driver name later
    ele.innerHTML = "Driver " + result.driverId;
    ele.addEventListener("click", toggleDriver);
    driverDropdown.appendChild(ele);
  }
}

const toggleDriver = async (e) => {
  let driverid = e.target.dataset.driverid;
  let raceid = e.target.dataset.raceid;
  if (document.getElementById("driver" + driverid)) {
    console.log("Already exists!")
    availableColors.push()
    d3.select("#driver" + driverid).remove();
  } else {
    let laptimes = await getLapTimes(raceid, driverid);
    let lineColor = availableColors.pop();
    e.target.style.borderColor = lineColor;

    g.append("path")
      .datum(laptimes)
      .attr("id", "driver" + driverid)
      .attr("data-driverid", driverid)
      .attr("fill", "none")
      .attr("stroke", lineColor)
      .attr("stroke-linejoin", "round")
      .attr("stroke-linecap", "round")
      .attr("stroke-width", 1.0)
      .attr("d", line);
  }
}

const getLapTimes = async (raceid, driverid) => {
  let url = proxy + "/LapTimes/" + raceid + "/" + driverid;
  let result = await ajaxRequest("GET", url);
  // Sort laptimes by lap number
  result.sort((a, b) => (a.lap > b.lap ? 1 : -1));
  return result
}

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
    document.getElementById('chart').innerHTML = '';
    document.getElementById('drivers').innerHTML = '';
    availableColors = colorPallete.slice(0)
    let circuitid = document.getElementById('circuit').dataset.circuitid;
    let year = document.getElementById('year').dataset.year;
    let url = proxy + "/Races/" + year + "/" + circuitid;
    let result = await ajaxRequest("GET", url);
    console.log(result);
    getResults(result);
  }, 1000);
}

const getResults = async (race) => {
  let url = proxy + "/Results/" + race.raceId;
  let result = await ajaxRequest("GET", url);
  if (result == 0) {
    console.log("No results for that race.")
  } else {
    console.log(result);
    await drawChart(result);
  }
}

const ajaxRequest = async (protocol, url) => {
  const res = await fetch(url);
  return await res.json();
}
