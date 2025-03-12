// Package model
// +kubebuilder:object:generate=true
// +groupName=pdok.nl
package model

import (
	"fmt"
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
