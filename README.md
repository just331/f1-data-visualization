# F1 Data Visualization
This project provides a visualization to more easily determine which drivers have dominated Formula One (F1) racing circuits from the years 2009 to 2018. The data is from ergast.com/mrd, a site that provides an API with data for every circuit around the world starting with that of the year 1950. Example data includes race results, driver standings, driver information, and circuit information.


Click the image below to watch a short video demo.

[https://www.youtube.com/watch?v=gUjRPfxCDnA&feature=youtu.be]

Alternatively,


The following link is our live site:
http://ec2-18-191-95-221.us-east-2.compute.amazonaws.com:3005


## Data Processing
<!--The application uses # of datasets. From XYZ. -->
On initial page load, a wordle with the last names of the top drivers from a circuit from 2009 will be loaded on the right hand side. The circuit may change from time to time on refresh depending on what the API returns first in the array of all the circuits. On the left hand side, a map with points of locations with every circuit will be displayed and the currently selected circuit indicated with a different colored point.

<!--
[it converts the rawdata into the format of key value pairs where key is the term and values contains the properties of category, total frequency and monthly wise object.]
^^^ Replace that with how we convert the csv into like an array or json or aws stuff?

[each monthly wise object contains the blog numbers in which the term occurred and monthly frequency]
^^^replace that

Image of code array

The size of the word in the wordle
The application of calculating the frequency of related words by [this might not be necessary]
-->


## Functionality
The year of the selected race will be centered near the top of the page, and the name of the race will be directly below. To switch between circuits, the user shall click on the dots on the map. The dot shall change color and the wordle will display the corresponding circuit. To switch between years for a circuit, the user shall use the up and down arrow keys. Then, the wordle for the selected circuit will display the top drivers for that year.


***Wordle and Map*** <br />
In default, the application visualizes the top drivers with its word cloud and its corresponding location on the map.

When the user selects a different year to be displayed, then its word cloud and respective map is displayed. When the user selects a different circuit to be displayed, then its location on the map will be indicated and the word cloud will be updated.
<!--When the user hovers on a particular word, {does it do anything on the map?

[image of gif basically]

***Interesting Findings***

**Summary of Hamilton**
Hamilton consistently ...
[image]

**Summary of Haas**
Haas used to be...
[image]
-->

## Built With

* [D3](https://d3js.org/) - For data visualization
* [Go](https://golang.org/) - For web framework
* [Ergast Motorsport API](http://ergast.com/mrd/) - For F1 stats
* [MongoDB](https://www.mongodb.com/) - Used for data storage

### Prerequisites
- MongoDB
- Go
- mongo-tools

## Authors

* **Joshua Johnson**
* **Adriana Holley**
* **Justin Rodriguez**

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details
