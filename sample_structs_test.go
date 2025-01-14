package commando

import (
	"strconv"
	"time"
)

type Sample struct {
	Foo  string  `csv:"foo"`
	Bar  int     `csv:"BAR"`
	Baz  string  `csv:"Baz"`
	Frop float64 `csv:"Quux"`
	Blah *int    `csv:"Blah"`
	SPtr *string `csv:"SPtr"`
	Omit *string `csv:"Omit,omitempty"`
}

type SliceSample struct {
	Slice []int `csv:"Slice"`
}

type EmbedSample struct {
	Qux string `csv:"first"`
	Sample
	Ignore string  `csv:"-"`
	Grault float64 `csv:"garply"`
	Quux   string  `csv:"last"`
}

type MarshalSample struct {
	Dummy string
}

func (m MarshalSample) MarshalText() ([]byte, error) {
	return []byte(m.Dummy), nil
}
func (m *MarshalSample) UnmarshalText(text []byte) error {
	m.Dummy = string(text)
	return nil
}

type EmbedMarshal struct {
	Foo *MarshalSample `csv:"foo"`
}

type EmbedPtrSample struct {
	Qux string `csv:"first"`
	*Sample
	Ignore string  `csv:"-"`
	Grault float64 `csv:"garply"`
	Quux   string  `csv:"last"`
}

type SkipFieldSample struct {
	EmbedSample
	MoreIgnore string `csv:"-"`
	Corge      string `csv:"abc"`
}

// Testtype for unmarshal/marshal functions on renamed basic types
type RenamedFloat64Unmarshaler float64
type RenamedFloat64Default float64

type RenamedSample struct {
	RenamedFloatUnmarshaler RenamedFloat64Unmarshaler `csv:"foo"`
	RenamedFloatDefault     RenamedFloat64Default     `csv:"bar"`
}

type MultiTagSample struct {
	Foo string `csv:"Baz,foo"`
	Bar int    `csv:"BAR"`
}

type CustomTagSample struct {
	Foo string `custom:"foo"`
	Bar string `csv:"BAR"`
}

type DateTime struct {
	Foo time.Time `csv:"Foo"`
}

type Level0Struct struct {
	Level0Field level1Struct `csv:"-"`
}

type level1Struct struct {
	Level1Field level2Struct `csv:"-"`
}

type level2Struct struct {
	InnerStruct
}

type InnerStruct struct {
	BoolIgnoreField0 bool   `csv:"-"`
	BoolField1       bool   `csv:"boolField1"`
	StringField2     string `csv:"stringField2"`
}

type UnmarshalCSVWithFieldsSample struct {
	Foo  string  `csv:"foo"`
	Bar  int     `csv:"bar"`
	Baz  string  `csv:"baz"`
	Frop float64 `csv:"frop"`
}

func (u *UnmarshalCSVWithFieldsSample) UnmarshalCSVWithFields(key, value string) error {
	switch key {
	case "foo":
		u.Foo = value
	case "bar":
		i, err := strconv.Atoi(value)
		if err != nil {
			return err
		}
		u.Bar = i
	case "baz":
		u.Baz = value
	case "frop":
		f, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return err
		}
		u.Frop = f * 100
	}
	return nil
}

var _ TypeUnmarshalCSVWithFields = (*UnmarshalCSVWithFieldsSample)(nil)
