// Copyright 2023 Christoph Fichtmüller. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package ocr

import (
	"github.com/cfichtmueller/jug"
	"github.com/otiai10/gosseract/v2"
)

func Recognize(r Request) (*Result, error) {
	client := gosseract.NewClient()
	defer client.Close()

	if len(r.Lang) > 0 {
		if err := client.SetLanguage(r.Lang); err != nil {
			//TODO: wrap error
			return nil, err
		}
	}

	if err := client.SetImageFromBytes(r.Data); err != nil {
		//TODO: wrap error
		return nil, err
	}

	iteratorLevel, ok := iteratorLevels[r.IteratorLevel]
	if !ok {
		return nil, jug.NewBadRequestError("iteratorLevel_invalid")
	}

	boxes, err := client.GetBoundingBoxes(iteratorLevel)
	if err != nil {
		//TODO: wrap error
		return nil, err
	}

	return newResult(boxes), nil
}
