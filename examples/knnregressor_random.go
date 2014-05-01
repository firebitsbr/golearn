package main

import (
	"fmt"

	data "github.com/sjwhitworth/golearn/data"
	knnclass "github.com/sjwhitworth/golearn/knn"
	util "github.com/sjwhitworth/golearn/utilities"
	mat "github.com/skelterjohn/go.matrix"
)

func main() {
	//Parses the infamous Iris data.
	cols, rows, _, labels, data := data.ParseCsv("datasets/randomdata.csv", 2, []int{0, 1})
	newlabels := util.ConvertLabelsToFloat(labels)

	//Initialises a new KNN classifier
	knn := knnclass.KNNRegressor{}
	knn.New(newlabels, data, rows, cols, "euclidean")

	for {
		//Creates a random array of N float64s between 0 and Y
		randArray := util.RandomArray(2, 100)

		//Initialises a vector with this array
		random := mat.MakeDenseMatrix(randArray, 1, 2)

		//Calculates the Euclidean distance and returns the most popular label
		outcome, _ := knn.Predict(random, 3)
		fmt.Println(outcome)
	}
}
