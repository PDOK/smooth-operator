// Package model
// +kubebuilder:object:generate=true
// +groupName=pdok.nl
package model

import (
	"fmt"
	"strconv"
	"strings"
)

// BBox defines a bounding box with coordinates
type BBox struct {
	// Linksboven X coördinaat
	// +kubebuilder:validation:Pattern="^[+-]?([0-9]+([.][0-9]*)?|[.][0-9]+)$"
	MinX string `json:"minx"`
	// Rechtsonder X coördinaat
	// +kubebuilder:validation:Pattern="^[+-]?([0-9]+([.][0-9]*)?|[.][0-9]+)$"
	MaxX string `json:"maxx"`
	// Linksboven Y coördinaat
	// +kubebuilder:validation:Pattern="^[+-]?([0-9]+([.][0-9]*)?|[.][0-9]+)$"
	MinY string `json:"miny"`
	// Rechtsonder Y coördinaat
	// +kubebuilder:validation:Pattern="^[+-]?([0-9]+([.][0-9]*)?|[.][0-9]+)$"
	MaxY string `json:"maxy"`
}

func (b BBox) ToExtent() string {
	return strings.Trim(fmt.Sprintf("%s %s %s %s", b.MinX, b.MinY, b.MaxX, b.MaxY), " ")
}

func ExtentToBBox(extent string) BBox {
	coords := strings.Split(extent, " ")
	fmt.Println("Coords", coords)
	if len(coords) != 4 {
		panic(fmt.Errorf("Extent has %d coordinates, need 4.", len(coords)))
	}

	return BBox{
		MinX: coords[0],
		MaxX: coords[2],
		MinY: coords[1],
		MaxY: coords[3],
	}
}

func (box *BBox) Combine(other BBox) {
	minXA, err := strconv.ParseFloat(box.MinX, 64)
	if err != nil {
		panic(fmt.Errorf("Error while parsing minX to float: %w", err))
	}

	minYA, err := strconv.ParseFloat(box.MinY, 64)
	if err != nil {
		panic(fmt.Errorf("Error while parsing minY to float: %w", err))
	}

	maxXA, err := strconv.ParseFloat(box.MaxX, 64)
	if err != nil {
		panic(fmt.Errorf("Error while parsing maxX to float: %w", err))
	}

	maxYA, err := strconv.ParseFloat(box.MaxY, 64)
	if err != nil {
		panic(fmt.Errorf("Error while parsing maxY to float: %w", err))
	}

	minXB, err := strconv.ParseFloat(other.MinX, 64)
	if err != nil {
		panic(fmt.Errorf("Error while parsing minX to float: %w", err))
	}

	minYB, err := strconv.ParseFloat(other.MinY, 64)
	if err != nil {
		panic(fmt.Errorf("Error while parsing minY to float: %w", err))
	}

	maxXB, err := strconv.ParseFloat(other.MaxX, 64)
	if err != nil {
		panic(fmt.Errorf("Error while parsing maxX to float: %w", err))
	}

	maxYB, err := strconv.ParseFloat(other.MaxY, 64)
	if err != nil {
		panic(fmt.Errorf("Error while parsing maxY to float: %w", err))
	}

	box.MinX = strconv.FormatFloat(min(minXA, minXB), 'f', -1, 64)
	box.MinY = strconv.FormatFloat(min(minYA, minYB), 'f', -1, 64)
	box.MaxX = strconv.FormatFloat(max(maxXA, maxXB), 'f', -1, 64)
	box.MaxY = strconv.FormatFloat(max(maxYA, maxYB), 'f', -1, 64)
}
