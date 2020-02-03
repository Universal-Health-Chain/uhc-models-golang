package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	math "math"
	"os"
	"reflect"
	"strconv"
)

type Geometry struct {
	Type        string                 `json:"type" bson:"type"`
	BoundingBox []float64              `json:"bbox,omitempty" bson:"bbox,omitempty"`
	Coordinates interface{}        `json:"coordinates,omitempty" bson:"coordinates,omitempty"`
	CRS         map[string]interface{} `json:"crs,omitempty" bson:"crs,omitempty"` // Coordinate Reference System Objects are not currently supported
}

type Feature struct {
	Type       string                 `json:"type" bson:"type"`
	Geometry   Geometry               `json:"geometry" bson:"geometry"`             // Coordinate Reference System Objects are not currently supported
	Properties map[string]interface{} `json:"properties,omitempty" bson:"properties,omitempty"` // Coordinate Reference System Objects are not currently supported
}

type Geojson struct {
	Type     string                 `json:"type" bson:"type"`
	CRS      map[string]interface{} `json:"crs,omitempty" bson:"crs,omitempty"` // Coordinate Reference System Objects are not currently supported
	Name     string                 `json:"name" bson:"name"`
	Features []Feature              `json:"features" bson:"features"`
}


var EmptyFeatures []Feature

func init() {
	EmptyFeatures = getEmptyFeatures()
}

func GeojsonFromFile(jsonFile *os.File) (geometry Geojson) {

	byteValue, _ := ioutil.ReadAll(jsonFile)

	geometry = Geojson{}
	json.Unmarshal(byteValue, &geometry)

	return geometry
}

func getEmptyFeatures() []Feature {
	coordinates := make([]interface{}, 2)
	coordinates[0] = 1
	coordinates[1] = 0
	feature:=Feature{Type:"Point", Geometry:Geometry{Type:"Point", Coordinates:coordinates}}
	features:=  make([]Feature, 1)
	features[0] = feature
	return features
}

func (feature *Feature) Validate() error {

	if feature.Geometry.Coordinates == nil {
		return errors.New("incorrect data")
	}

	if feature.Geometry.Type =="Point" {
		switch t:=feature.Geometry.Coordinates.(type) {
		case [][]interface{}:
			return errors.New("incorrect data, two levels for a point")
		case []interface{}:
			if len(t)==0 || t[0]==nil {
				return errors.New("incorrect data, coordintes can not be empty")
			}

			var err error
			feature.Geometry.Coordinates, err = CleanUpCoordinates(t)
			if err != nil {
				return err
			}

		default:
			return errors.New("error incorrect type for " + feature.Geometry.Type)
		}
	} else if feature.Geometry.Type =="Line" || feature.Geometry.Type =="LineString" || feature.Geometry.Type =="MultiPoint" || feature.Geometry.Type =="Polygon" {
		switch t:=feature.Geometry.Coordinates.(type) {
		case [][][]interface{}:
			return errors.New("incorrect data, three levels")
		case [][]interface{}:
			if len(t) == 0 {
				return errors.New("incorrect data, coordintes can not be empty")
			}
			var err error
			t[0], err = CleanUpCoordinates(t[0])
			if err != nil {
				return err
			}
		default:
			return errors.New("error incorrect type for " + feature.Geometry.Type)
		}
	} else if feature.Type =="MultiPolygon"  {
		switch t:=feature.Geometry.Coordinates.(type) {
		case [][][][]interface{}:
			return errors.New("incorrect data, four levels")
		case [][][]interface{}:
			if len(t) == 0 {
				return errors.New("incorrect data, coordintes can not be empty")
			}
		default:
			return errors.New("error incorrect type for " + feature.Geometry.Type)
		}
	}

	return nil
}

func CleanUpCoordinates(coordInterface interface{}) ([]interface{}, error){
	switch coordinates:=coordInterface.(type) {
	case []interface{}:
		for index, i := range coordinates {
			if i==nil || i=="" {
				return nil, errors.New("incorrect data, empty coordinates or incorrect values")
			} else if fmt.Sprintf("%T", i)=="string" {
				str := fmt.Sprintf("%v", i)

				s, err := strconv.ParseFloat(str, 64)
				if err != nil {
					return nil, errors.New("unable to parse string to float")
				}
				coordinates[index]=s
			}
		}
		x,_ := convertToFloat(coordinates[0])
		y,_ := convertToFloat(coordinates[1])

		if math.Abs(y)>90{
			return nil,errors.New("incorrect data, coordinate y too big")

		}

		if math.Abs(x)>180{
			return nil,errors.New("incorrect data, coordinate x too big")
		}

		return coordinates, nil
	default:
		return nil, errors.New("incorrect data found cleaning coordinates")
	}
}

func convertToFloat(unk interface{}) (float64, error){
	switch i := unk.(type) {
	case float64:
		return i, nil
	case float32:
		return float64(i), nil
	case int64:
		return float64(i), nil
	case int32:
		return float64(i), nil
	case int16:
		return float64(i), nil
	case int:
		return float64(i), nil
	// ...other cases...
	default:
		return math.NaN(), errors.New("getFloat: unknown value is of incompatible type")
	}
}

func (geojson *Geojson) CleanUpFeatures(){
	log.Println("cleaning up geojson")
	for index, feature := range(geojson.Features){
		err := feature.Validate()
		if (err!= nil){
			log.Println("error found in geojson: " + err.Error())
			log.Printf("cleaning up feature %v", index)
			geojson.Features[index] = EmptyFeatures[0]
		}
	}
}

func (geojson *Geojson) PropertyNames() map[string]string {
	propertiesMap := geojson.Features[0].Properties

	var propertiesWithTypes map[string]string
	propertiesWithTypes = make(map[string]string)
	for key, value := range propertiesMap {
		if key != "" && value != nil {
			typeStr := reflect.TypeOf(value).String()
			propertiesWithTypes[key] = typeStr
		}
	}
	return propertiesWithTypes

}
