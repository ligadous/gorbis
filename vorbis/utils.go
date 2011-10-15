package vorbis

import(
  "math"
)
func iPow(b,e int) int {
  if e == 0 { return 1 }
  if e % 2 == 1 {
    return b * iPow(b, e-1)
  }
  return iPow(b*b, e/2)
}

func ilog(n uint32) int {
  e := 31
  bit := uint32(1) << 31
  for e >= 0 {
    if (n & bit) != 0 { return e + 1 }
    bit = bit >> 1
    e--
  }
  return 0
}

// From the spec: The return value for this function is defined to be ’the greatest integer
// value for which [return_value] to the power of [codebook_dimensions] is less than or equal to
// [codebook_ entries]’.
func Lookup1Values(entries,dimensions int) int {
  if dimensions <= 0 {
    panic("Can't do a lookup1Values with dimensions <= 0")
  }
  low := 0
  high := 1
  for iPow(high, dimensions) <= entries {
    low = high
    high *= 2
  }
  var mid int
  for high - low > 1 {
    mid = (high + low) / 2
    if iPow(mid, dimensions) <= entries {
      low = mid
    } else {
      high = mid
    }
  }
  if iPow(mid+1, dimensions) <= entries { return mid+1 }
  if iPow(mid, dimensions) <= entries { return mid }
  return mid-1
}

// Translated from some java source somewhere, uses an iterative approach instead of a
// binary search
func Lookup1ValuesJava(entries,dimensions int) int {
  vals := int(math.Floor(math.Pow(float64(entries), 1.0 / float64(dimensions))))
  for {
    acc := 1
    acc1 := 1
    for i := 0; i < dimensions; i++ {
      acc *= vals
      acc1 *= vals+1
    }
    if acc <= entries && acc1 > entries {
      break
    } else if acc > entries {
      vals--
    } else {
      vals++
    }
  }
  return vals
}

