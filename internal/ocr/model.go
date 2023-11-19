// Copyright 2023 Christoph Fichtmüller. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package ocr

import (
	"github.com/cfichtmueller/jug"
	"github.com/otiai10/gosseract/v2"
)

var iteratorLevels = map[string]gosseract.PageIteratorLevel{
	"word":     gosseract.RIL_WORD,
	"textline": gosseract.RIL_TEXTLINE,
	"para":     gosseract.RIL_PARA,
	"symbol":   gosseract.RIL_SYMBOL,
	"block":    gosseract.RIL_BLOCK,
}

type Request struct {
	Lang          string
	IteratorLevel string
	Data          []byte
}

func (r Request) Validate() error {
	return jug.NewValidator().
		RequireEnum(r.IteratorLevel, "iteratorLevel_invalid", "word", "textline", "para", "symbol", "block").
		Require(len(r.Data) > 0, "data_required").
		Validate()
}

type Result struct {
	Boxes []*Box `json:"boxes"`
}

func newResult(boxes []gosseract.BoundingBox) *Result {
	return &Result{
		Boxes: jug.MapMany(boxes, newBox),
	}
}

type Box struct {
	MinX       int     `json:"minX"`
	MinY       int     `json:"minY"`
	MaxX       int     `json:"maxX"`
	MaxY       int     `json:"maxY"`
	Word       string  `json:"word"`
	Confidence float64 `json:"confidence"`
}

func newBox(box gosseract.BoundingBox) *Box {
	return &Box{
		MinX:       box.Box.Min.X,
		MinY:       box.Box.Min.Y,
		MaxX:       box.Box.Max.X,
		MaxY:       box.Box.Max.Y,
		Word:       box.Word,
		Confidence: box.Confidence,
	}
}
