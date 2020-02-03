package models

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestGeojsonValidation(t *testing.T) {
	features := getEmptyFeatures()
	feature := features[0]

 	err:= feature.Validate()
	log.Println(err)
	assert.Nil(t,err,"data loaded")

	featureLine := feature
	featureLine.Geometry.Type = "Line"
	err= featureLine.Validate()
	log.Println(err)
	assert.NotNil(t,err,"returns error")

	featureLine2 := feature
	featureLine2.Geometry.Type = "Line"
	coordinates := make([][]interface{}, 1)
	coordinates2 := make([]interface{}, 2)

	coordinates2[0] = 1
	coordinates2[1] = 0
	coordinates[0] = coordinates2

	featureLine.Geometry.Coordinates = coordinates
	err= featureLine.Validate()
	log.Println(err)
	assert.Nil(t,err,"data loaded for line")

	featureLine3 := feature
	featureLine3.Geometry.Type = "Line"
	var coordinatesLine [][]interface{} = make([][]interface{}, 1)
	coordinates3 := make([]interface{}, 2)

	coordinatesLine[0] = coordinates3

	featureLine.Geometry.Coordinates = coordinatesLine
	err= featureLine.Validate()
	log.Println(err)
	assert.NotNil(t,err,"returns error")

}

func TestGeojsonValidationIncorrectValues(t *testing.T) {
	features := getEmptyFeatures()

	feature1 := features[0]
	feature1.Geometry.Type = "Point"
	coordinates := make([]interface{}, 2)

	coordinates[0] = 0
	coordinates[1] = 100

	feature1.Geometry.Coordinates = coordinates
	err := feature1.Validate()

	log.Println(err)
	assert.NotNil(t, err, "returns error becuse out of rage")

	feature2 := features[0]
	feature2.Geometry.Type = "Line"

	var coordinates2 [][]interface{} = make([][]interface{}, 1)

	coordinates = make([]interface{}, 2)

	coordinates[0] = 190
	coordinates[1] = 90

	coordinates2[0] = coordinates


	feature2.Geometry.Coordinates = coordinates2
	err = feature2.Validate()

	log.Println(err)
	assert.NotNil(t, err, "returns error becuse out of rage")
}