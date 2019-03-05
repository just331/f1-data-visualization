# F1 Data Visualization

##Video

https://www.youtube.com/watch?v=VfHqvTstWRk&feature=youtu.be

## Live Site
http://ec2-18-191-95-221.us-east-2.compute.amazonaws.com:3005


## About the project and data
This projects goal is to better determine what drivers and teams can improve to decrease laptime in F1 racing. The data is very vast
and we can uncover even more hidden secrets using the visualization. So far I have used the data to graph individual laptimes for drivers
in a given race. The data also contains info about, accidents, pitstops, constructors, and season results. I found it interesting that the lap
before a pitstop is usually the same length time wise of the first lap.

### Prerequisites
- MongoDB
- Go
- mongo-tools
- gorilla/mux
- mongo driver for Go

```
Use up and down arrows to change year, use left and right to change circuit, select drivers using dropdown.

```
## Built With

* [D3](https://d3js.org/) - For data visualization
* [Go](https://golang.org/) - For web framework
* [MongoDB](https://www.mongodb.com/) - Used for data storage

## Authors

* **Joshua Johnson**

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details
