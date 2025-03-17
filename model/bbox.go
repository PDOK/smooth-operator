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
	return fmt.Sprintf("%s %s %s %s", b.MinX, b.MinY, b.MaxX, b.MaxY)
}

func (b BBox) ToPolygon() string {
	var sb strings.Builder
	sb.WriteString(b.MinX + " " + b.MinY + " ") // lower left
	sb.WriteString(b.MinX + " " + b.MaxY + " ") // upper left
	sb.WriteString(b.MaxX + " " + b.MaxY + " ") // upper right
	sb.WriteString(b.MaxX + " " + b.MinY + " ") // lower right
	sb.WriteString(b.MinX + " " + b.MinY)       // lower left, final vertice is equal to the first one
	return sb.String()
}

func ExtentToBBox(extent string) BBox {
	coords := strings.Split(extent, " ")
	if len(coords) != 4 {
		panic(fmt.Errorf("extent has %d coordinates, needs 4", len(coords)))
	}

	return BBox{
		MinX: coords[0],
		MaxX: coords[2],
		MinY: coords[1],
		MaxY: coords[3],
	}
}
