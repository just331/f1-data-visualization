const drawChart = async (results) => {
  // Sort results by final position
  results.sort((a, b) => (a.position > b.position ? 1 : -1));

  // Get the top three drivers by default
  for (let i = 0; i < 3; i++) {
    result = results[i];
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
      .attr("class", "driverC" + result.driverid)
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
