

var widthMap = 1100,
    heightMap = 600;

var projection = d3.geo.mercator()
    .scale(200)
    .translate([widthMap / 2, heightMap / 2]);


var path = d3.geo.path()
    .projection(projection);

var graticule = d3.geo.graticule();


var svg = d3.select("#map").append("svg")
    .attr("width", widthMap)
    .attr("height", heightMap);


svg.append("path")
    .datum(graticule)
    .attr("class", "graticule")
    .attr("d", path)
    .style("fill", "white");



queue()
    .defer(d3.json, "js/world-110m.json")
    .defer(d3.json, "js/circuits.json")
    .await(ready);

function ready(error, world, data) {
    if(error) return console.error(error);

    // TODO Implement map for circuits

    /*
    let race =
    let circuit = race.Circuits
        .map(function(raceResult) {

            console.log(raceResult);
            return {
                circuitName: raceResult.properties.circuitName,
                circuitID: raceResult.properties.circuitID
            };
        });
    */
    svg.insert("path", "graticule")
        .datum(topojson.feature(world, world.objects.land))
        .attr("class", "land")
        .attr("d", path)
        .style("fill", "gray");
    svg.insert("path", "graticule")
        .datum(topojson.mesh(world, world.objects.countries, function(a, b) { return a !== b; }))
        .attr("class", "boundary")
        .attr("d", path)
        .style("fill", "gray");

    svg.append("path")
        .datum(data)
        .classed("geopath",true)
        .style("fill", "blue")
        .attr("d", path)
        .call(d3.helper.tooltip(
            function(d){
                return "<b>"+ d.features.circuitName + "</b><br/>"+ d.circuitID;

            }

        ));

    /*var zoom = d3.behavior.zoom()
        .on("zoom",function() {
            svg.attr("transform","translate("+
                d3.event.translate.join(",")+")scale("+d3.event.scale+")");
            svg.selectAll("circle")
                .attr("d", path.projection(projection));
            svg.selectAll("path")
                .attr("d", path.projection(projection));

        });

    svg.call(zoom)*/
}

