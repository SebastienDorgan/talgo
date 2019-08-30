package talgo

import "strings"

//IntSlice collection of int
type IntSlice []int

//Len len
func (s IntSlice) Len() int {
	return len(s)
}

//Greater check if s[i] is greater or greater or equal to value
func (s IntSlice) Greater(value int, orEqual bool) Predicate {
	return func(i int) bool {
		if orEqual {
			return s[i] >= value
		}
		return s[i] > value
	}
}

//Lesser check if s[i] is lesser or lesser or equal to value
func (s IntSlice) Lesser(value int, orEqual bool) Predicate {
	return func(i int) bool {
		if orEqual {
			return s[i] <= value
		}
		return s[i] < value
	}
}

//Equal check if s[i] is equal to value
func (s IntSlice) Equal(value int) Predicate {
	return func(i int) bool {
		return s[i] == value
	}
}

//Float64Slice collection of float64
type Float64Slice []float64

//Len len
func (s Float64Slice) Len() int {
	return len(s)
}

//Greater check if s[i] is greater or greater or equal to value
func (s Float64Slice) Greater(value float64, orEqual bool) Predicate {
	return func(i int) bool {
		if orEqual {
			return s[i] >= value
		}
		return s[i] > value
	}
}

//Lesser check if s[i] is lesser or lesser or equal to value
func (s Float64Slice) Lesser(value float64, orEqual bool) Predicate {
	return func(i int) bool {
		if orEqual {
			return s[i] <= value
		}
		return s[i] < value
	}
}

//Equal check if s[i] is equal to value
func (s Float64Slice) Equal(value float64) Predicate {
	return func(i int) bool {
		return s[i] == value
	}
}

//Float32Slice collection of float32
type Float32Slice []float32

//Len len
func (s Float32Slice) Len() int {
	return len(s)
}

//Greater check if s[i] is greater or greater or equal to value
func (s Float32Slice) Greater(value float32, orEqual bool) Predicate {
	return func(i int) bool {
		if orEqual {
			return s[i] >= value
		}
		return s[i] > value
	}
}

//Lesser check if s[i] is lesser or lesser or equal to value
func (s Float32Slice) Lesser(value float32, orEqual bool) Predicate {
	return func(i int) bool {
		if orEqual {
			return s[i] <= value
		}
		return s[i] < value
	}
}

//Equal check if s[i] is equal to value
func (s Float32Slice) Equal(value float32) Predicate {
	return func(i int) bool {
		return s[i] == value
	}
}

//Int8Slice collection of int16
type Int8Slice []int8

//Len len
func (s Int8Slice) Len() int {
	return len(s)
}

//Greater check if s[i] is greater or greater or equal to value
func (s Int8Slice) Greater(value int8, orEqual bool) Predicate {
	return func(i int) bool {
		if orEqual {
			return s[i] >= value
		}
		return s[i] > value
	}
}

//Lesser check if s[i] is lesser or lesser or equal to value
func (s Int8Slice) Lesser(value int8, orEqual bool) Predicate {
	return func(i int) bool {
		if orEqual {
			return s[i] <= value
		}
		return s[i] < value
	}
}

//Equal check if s[i] is equal to value
func (s Int8Slice) Equal(value int8) Predicate {
	return func(i int) bool {
		return s[i] == value
	}
}

//Int16Slice collection of int16
type Int16Slice []int16

//Len len
func (s Int16Slice) Len() int {
	return len(s)
}

//Greater check if s[i] is greater or greater or equal to value
func (s Int16Slice) Greater(value int16, orEqual bool) Predicate {
	return func(i int) bool {
		if orEqual {
			return s[i] >= value
		}
		return s[i] > value
	}
}

//Lesser check if s[i] is lesser or lesser or equal to value
func (s Int16Slice) Lesser(value int16, orEqual bool) Predicate {
	return func(i int) bool {
		if orEqual {
			return s[i] <= value
		}
		return s[i] < value
	}
}

//Equal check if s[i] is equal to value
func (s Int16Slice) Equal(value int16) Predicate {
	return func(i int) bool {
		return s[i] == value
	}
}

//Int32Slice collection of int32
type Int32Slice []int32

//Len len
func (s Int32Slice) Len() int {
	return len(s)
}

//Greater check if s[i] is greater or greater or equal to value
func (s Int32Slice) Greater(value int32, orEqual bool) Predicate {
	return func(i int) bool {
		if orEqual {
			return s[i] >= value
		}
		return s[i] > value
	}
}

//Lesser check if s[i] is lesser or lesser or equal to value
func (s Int32Slice) Lesser(value int32, orEqual bool) Predicate {
	return func(i int) bool {
		if orEqual {
			return s[i] <= value
		}
		return s[i] < value
	}
}

//Equal check if s[i] is equal to value
func (s Int32Slice) Equal(value int32) Predicate {
	return func(i int) bool {
		return s[i] == value
	}
}

//Int64Slice collection of int32
type Int64Slice []int64

//Len len
func (s Int64Slice) Len() int {
	return len(s)
}

//Greater check if s[i] is greater or greater or equal to value
func (s Int64Slice) Greater(value int64, orEqual bool) Predicate {
	return func(i int) bool {
		if orEqual {
			return s[i] >= value
		}
		return s[i] > value
	}
}

//Lesser check if s[i] is lesser or lesser or equal to value
func (s Int64Slice) Lesser(value int64, orEqual bool) Predicate {
	return func(i int) bool {
		if orEqual {
			return s[i] <= value
		}
		return s[i] < value
	}
}

//Equal check if s[i] is equal to value
func (s Int64Slice) Equal(value int64) Predicate {
	return func(i int) bool {
		return s[i] == value
	}
}

//UIntSlice collection of uint
type UIntSlice []uint

//Len len
func (s UIntSlice) Len() int {
	return len(s)
}

//Greater check if s[i] is greater or greater or equal to value
func (s UIntSlice) Greater(value uint, orEqual bool) Predicate {
	return func(i int) bool {
		if orEqual {
			return s[i] >= value
		}
		return s[i] > value
	}
}

//Lesser check if s[i] is lesser or lesser or equal to value
func (s UIntSlice) Lesser(value uint, orEqual bool) Predicate {
	return func(i int) bool {
		if orEqual {
			return s[i] <= value
		}
		return s[i] < value
	}
}

//Equal check if s[i] is equal to value
func (s UIntSlice) Equal(value uint) Predicate {
	return func(i int) bool {
		return s[i] == value
	}
}

//UInt8Slice collection of uint
type UInt8Slice []uint8

//Len len
func (s UInt8Slice) Len() int {
	return len(s)
}

//Greater check if s[i] is greater or greater or equal to value
func (s UInt8Slice) Greater(value uint8, orEqual bool) Predicate {
	return func(i int) bool {
		if orEqual {
			return s[i] >= value
		}
		return s[i] > value
	}
}

//Lesser check if s[i] is lesser or lesser or equal to value
func (s UInt8Slice) Lesser(value uint8, orEqual bool) Predicate {
	return func(i int) bool {
		if orEqual {
			return s[i] <= value
		}
		return s[i] < value
	}
}

//Equal check if s[i] is equal to value
func (s UInt8Slice) Equal(value uint8) Predicate {
	return func(i int) bool {
		return s[i] == value
	}
}

//UInt16Slice collection of uint
type UInt16Slice []uint16

//Len len
func (s UInt16Slice) Len() int {
	return len(s)
}

//Greater check if s[i] is greater or greater or equal to value
func (s UInt16Slice) Greater(value uint16, orEqual bool) Predicate {
	return func(i int) bool {
		if orEqual {
			return s[i] >= value
		}
		return s[i] > value
	}
}

//Lesser check if s[i] is lesser or lesser or equal to value
func (s UInt16Slice) Lesser(value uint16, orEqual bool) Predicate {
	return func(i int) bool {
		if orEqual {
			return s[i] <= value
		}
		return s[i] < value
	}
}

//Equal check if s[i] is equal to value
func (s UInt16Slice) Equal(value uint16) Predicate {
	return func(i int) bool {
		return s[i] == value
	}
}

//UInt32Slice collection of uint
type UInt32Slice []uint32

//Len len
func (s UInt32Slice) Len() int {
	return len(s)
}

//Greater check if s[i] is greater or greater or equal to value
func (s UInt32Slice) Greater(value uint32, orEqual bool) Predicate {
	return func(i int) bool {
		if orEqual {
			return s[i] >= value
		}
		return s[i] > value
	}
}

//Lesser check if s[i] is lesser or lesser or equal to value
func (s UInt32Slice) Lesser(value uint32, orEqual bool) Predicate {
	return func(i int) bool {
		if orEqual {
			return s[i] <= value
		}
		return s[i] < value
	}
}

//Equal check if s[i] is equal to value
func (s UInt32Slice) Equal(value uint32) Predicate {
	return func(i int) bool {
		return s[i] == value
	}
}

//UInt64Slice collection of uint
type UInt64Slice []uint64

//Len len
func (s UInt64Slice) Len() int {
	return len(s)
}

//Greater check if s[i] is greater or greater or equal to value
func (s UInt64Slice) Greater(value uint64, orEqual bool) Predicate {
	return func(i int) bool {
		if orEqual {
			return s[i] >= value
		}
		return s[i] > value
	}
}

//Lesser check if s[i] is lesser or lesser or equal to value
func (s UInt64Slice) Lesser(value uint64, orEqual bool) Predicate {
	return func(i int) bool {
		if orEqual {
			return s[i] <= value
		}
		return s[i] < value
	}
}

//Equal check if s[i] is equal to value
func (s UInt64Slice) Equal(value uint64) Predicate {
	return func(i int) bool {
		return s[i] == value
	}
}

//StringSlice collection of int
type StringSlice []string

//Len len
func (s StringSlice) Len() int {
	return len(s)
}

//Greater check if s[i] is greater or greater or equal to value
func (s StringSlice) Greater(value string, orEqual bool) Predicate {
	return func(i int) bool {
		if orEqual {
			return strings.Compare(s[i], value) >= 0
		}
		return strings.Compare(s[i], value) > 0
	}
}

//Lesser check if s[i] is lesser or lesser or equal to value
func (s StringSlice) Lesser(value string, orEqual bool) Predicate {
	return func(i int) bool {
		if orEqual {
			return strings.Compare(s[i], value) <= 0
		}
		return strings.Compare(s[i], value) < 0
	}
}

//Equal check if s[i] is equal to value
func (s StringSlice) Equal(value string) Predicate {
	return func(i int) bool {
		return strings.Compare(s[i], value) == 0
	}
}
