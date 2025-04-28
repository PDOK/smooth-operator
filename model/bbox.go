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

func (b BBox) ToPolygon() string {
	var sb strings.Builder
	sb.WriteString(b.MinY + " " + b.MinX + " ")
	sb.WriteString(b.MinY + " " + b.MaxX + " ")
	sb.WriteString(b.MaxY + " " + b.MaxX + " ")
	sb.WriteString(b.MaxY + " " + b.MinX + " ")
	sb.WriteString(b.MinY + " " + b.MinX)
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

func (b *BBox) Combine(other BBox) {
	minXA, err := strconv.ParseFloat(b.MinX, 64)
	if err != nil {
		panic(fmt.Errorf("Error while parsing minX to float: %w", err))
	}

	minYA, err := strconv.ParseFloat(b.MinY, 64)
	if err != nil {
		panic(fmt.Errorf("Error while parsing minY to float: %w", err))
	}

	maxXA, err := strconv.ParseFloat(b.MaxX, 64)
	if err != nil {
		panic(fmt.Errorf("Error while parsing maxX to float: %w", err))
	}

	maxYA, err := strconv.ParseFloat(b.MaxY, 64)
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

	b.MinX = strconv.FormatFloat(min(minXA, minXB), 'f', -1, 64)
	b.MinY = strconv.FormatFloat(min(minYA, minYB), 'f', -1, 64)
	b.MaxX = strconv.FormatFloat(max(maxXA, maxXB), 'f', -1, 64)
	b.MaxY = strconv.FormatFloat(max(maxYA, maxYB), 'f', -1, 64)
}
