const proxy = "http://localhost:3005"
const svgWidth = 900;
const svgHeight = 600;
const margin = { top: 20, right: 20, bottom: 30, left: 50 };
const width = svgWidth - margin.left - margin.right;
const height = svgHeight - margin.top - margin.bottom;
const colorPallete = [
  "#4dEEEA",
  "#74EE15",
  "#FFE700",
  "#F000FF",
  "#001EFF"
]

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

const drawAxes = (laptimes) => {
  let svg = d3.select('svg')
    .attr("width", svgWidth)
    .attr("height", svgHeight);

  let g = svg.append("g")
    .attr("transform",
      "translate(" + margin.left + "," + margin.top + ")"
    );

  let x = d3.scaleLinear().rangeRound([0, width]);
  let y = d3.scaleLinear().rangeRound([height, 0]);

  let line = d3.line()
   .x(function(d) { return x(d.lap)})
   .y(function(d) { return y(d.milliseconds)})
   x.domain(d3.extent(laptimes, function(d) { return d.lap }));
   y.domain(d3.extent(laptimes, function(d) { return d.milliseconds }));

   g.append("g")
     .attr("transform", "translate(0," + height + ")")
     .call(d3.axisBottom(x))
     .select(".domain")
     .remove();

   g.append("g")
     .call(d3.axisLeft(y))
     .append("text")
     .attr("fill", "white")
     .attr("transform", "rotate(-90)")
     .attr("y", 6)
     .attr("dy", "-6em")
     .attr("text-anchor", "end")
   .text("Time (ms)");

  return {line, g};
}

const drawChart = async (results) => {
  // Sort results by final position
  results.sort((a, b) => (a.position > b.position ? 1 : -1));
  var availableColors = colorPallete;
  // Get the top three drivers by default
  for (let i = 0; i < 3; i++) {
    var g;
    var line;
    result = results[i];
    console.log(result);
    let laptimes = await getLapTimes(result.raceId, result.driverId);
    // Sort laptimes by lap number
    laptimes.sort((a, b) => (a.lap > b.lap ? 1 : -1));
    if (i == 0) {
      let axes = drawAxes(laptimes);
      g = axes.g;
      line = axes.line;
    }

    let lineColor = getRandomColor(availableColors);


    g.append("path")
      .datum(laptimes)
      .attr("data-driverid", result.driverId)
      .attr("fill", "none")
      .attr("stroke", lineColor)
      .attr("stroke-linejoin", "round")
      .attr("stroke-linecap", "round")
      .attr("stroke-width", 1.0)
      .attr("d", line);
   }
}

const getRandomColor = (availableColors) => {
  let colorIndex = Math.floor(Math.random() * availableColors.length);
  console.log(colorIndex);
  return availableColors[colorIndex];
}

const getLapTimes = async (raceid, driverid) => {
  let url = proxy + "/LapTimes/" + raceid + "/" + driverid;
  let result = await ajaxRequest("GET", url);
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
    let circuitid = document.getElementById('circuit').dataset.circuitid;
    let year = document.getElementById('year').dataset.year;
    let url = proxy + "/Races/" + year + "/" + circuitid;
    let result = await ajaxRequest("GET", url);
    console.log(result);
    getResults(result);
  }, 3000);
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
  console.log(protocol, url);
  const res = await fetch(url);
  return await res.json();
}
