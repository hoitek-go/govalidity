package govalidityconv

import (
	"testing"
)

func TestToNumberWhenValueIsInt(t *testing.T) {
	t.Run("When value is int", func(t *testing.T) {
		_, err := ToNumber(1)
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("When value is int8", func(t *testing.T) {
		_, err := ToNumber(int8(1))
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("When value is int16", func(t *testing.T) {
		_, err := ToNumber(int16(1))
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("When value is int64", func(t *testing.T) {
		_, err := ToNumber(int64(1))
		if err != nil {
			t.Error(err)
		}
	})
}

func TestToNumberWhenValueIsUInt(t *testing.T) {
	t.Run("When value is uint", func(t *testing.T) {
		_, err := ToNumber(uint(1))
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("When value is uint8", func(t *testing.T) {
		_, err := ToNumber(uint8(1))
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("When value is uint16", func(t *testing.T) {
		_, err := ToNumber(uint16(1))
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("When value is uint64", func(t *testing.T) {
		_, err := ToNumber(uint64(1))
		if err != nil {
			t.Error(err)
		}
	})
}

func TestToNumberWhenValueIsFloat(t *testing.T) {
	t.Run("When value is float32", func(t *testing.T) {
		_, err := ToNumber(float32(1))
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("When value is float64", func(t *testing.T) {
		_, err := ToNumber(float64(1))
		if err != nil {
			t.Error(err)
		}
	})
}

func TestToNumberWhenValueIsString(t *testing.T) {
	t.Run("When string is not number", func(t *testing.T) {
		value, _ := ToNumber("test")
		if value != nil {
			t.Error("value should be nil when there is and error")
		}
	})

	t.Run("When string is number", func(t *testing.T) {
		_, err := ToNumber("1")
		if err != nil {
			t.Error(err)
		}
	})
}

func TestToNumberWhenUnknownError(t *testing.T) {
	t.Run("When unknown error", func(t *testing.T) {
		x := 1
		_, err := ToNumber(&x)
		if err == nil {
			t.Error("should return error when type is not detected")
		}
	})
}
