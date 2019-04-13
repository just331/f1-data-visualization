const drawWordle = async (results) => {
  let race = results.MRData.RaceTable.Races[0];
  let words = race.Results
    .map(function(raceResult) {

      console.log(raceResult);
      // TODO: Refine this equation
      return {
        text: raceResult.Driver.familyName,
        size: 10 + Math.abs(raceResult.position - race.Results.length) * 2
      };
  });

  var layout = d3.layout.cloud()

  // TODO: Fix dimensions
  layout.size([960, 500])
    .words(words)
    .font("Impact")
    .fontSize(function(d, i) { return d.size; })
    .on("end", draw)
    .start();
}

function getRandColor() {
  let color = '#'+(0x1000000+(Math.random())*0xffffff).toString(16).substr(1,6);
  return color;
}

function draw(words) {
  d3.select("#wordle").append("svg")
    .attr("width", 850)
    .attr("height", 350)
    .attr("class", "wordcloud")
    .append("g")
    // without the transform, words words would get cutoff to the left and top, they would
    // appear outside of the SVG area
    .attr("transform", "translate(320,200)")
    .selectAll("text")
    .data(words)
    .enter().append("text")
    .style("font-size", function(d) { return d.size + "px"; })
    .style("fill", function() { return getRandColor(); })
    .attr("transform", function(d) {
        return "translate(" + [d.x, d.y] + ")rotate(" + d.rotate + ")";
    })
    .text(function(d) { return d.text; });
}
