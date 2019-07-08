package enigma

import (
	"strings"
)

type Transformation interface {
	Transform(from Alphabet) Alphabet
	Clone() Transformation
	Inverse() Transformation
}

func CombineTransformations(tfs ...Transformation) Transformation {
	combined := TransformationList{}
	for _, tf := range tfs {
		combined = append(combined, tf)
	}
	return combined
}

type TransformationList []Transformation

func (tl TransformationList) Transform(from Alphabet) Alphabet {
	to := from
	for _, tf := range tl {
		to = tf.Transform(to)
	}
	return to
}

func (tl TransformationList) Clone() Transformation {
	clone := []Transformation{}
	for _, tf := range tl {
		clone = append(clone, tf.Clone())
	}
	return TransformationList(clone)
}

func (tl TransformationList) Inverse() Transformation {
	inverse := []Transformation{}
	for i := len(tl) - 1; i >= 0; i-- {
		inverse = append(inverse, tl[i].Inverse())
	}
	return TransformationList(inverse)
}

func TransformString(tf Transformation, from string) string {
	to := []byte{}
	for _, alph := range []byte(from) {
		to = append(to, byte(tf.Transform(Alphabet(alph))))
	}
	return string(to)
}

func TransformAsString(tf Transformation) string {
	str := []byte{}
	for i := 0; i < 26; i++ {
		str = append(str, byte(tf.Transform(Alphabet('a'+byte(i)))))
	}
	return strings.ToUpper(string(str))
}
