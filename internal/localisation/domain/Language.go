package domain

import (
	"errors"

	"golang.org/x/text/language"
)

type Language struct {
	Code       language.Tag
	Name       string
	NativeName string
	IsRtl      bool
}

func (l *Language) Validate() error {
	if l.Code == (language.Tag{}) || l.Code == language.Und {
		return errors.New("code is empty or undetermined")
	}

	if l.Code.String() == "und" {
		return errors.New("entity locale must be a valid ISO code")
	}

	if l.Name == "" {
		return errors.New("name is empty")
	}

	if l.NativeName == "" {
		return errors.New("native name is empty")
	}

	return nil
}
