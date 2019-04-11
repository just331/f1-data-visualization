# F1 Data Visualization

##Video


## Live Site
http://ec2-18-191-95-221.us-east-2.compute.amazonaws.com:3005


## About the project and data
This projects goal is to better determine what drivers and teams can improve to decrease laptime in F1 racing. The data is very vast
and we can uncover even more hidden secrets using the visualization. So far we have added the abilities to see a text visualization, via a wordcloud,
of drivers who have won the most races in a given year and season. You can change the year to see who has been winning the most (visualized by larger text size).
We have also added a map of every circuit around the world, and you can hover over the marker to see the circuit name and other info on the races that have 
been held there in a given year. 

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
* **Adriana Holley**
* **Justin Rodriguez**

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details
