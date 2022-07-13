func rotateSlice(slice []float32, n int) (rotateSlice []float32) {
    var start []float32
    var end []float32
    size := len(slice)
    rotatedSlice = make([]float32, size)

    nAbs := math.Abs(float64(n))

    if int(nAbs) > size {
        remainder, _ := qandremainderfloat32(n), float32(size))
        n = int(remainder)
    }

    if n != 0 {
        if n > 0 {
            index := size - n
            begin = slice[index:]
            end = slice[0:index]
            copy(rotatedSlice, begin)
            copy(rotatedSlice[n:], end)
        } else {
            n = int(nAbs)
            index := size - n
            begin = slice[n:]
            end = slice[0:n]
            copy(rotatedSlice, begin)
            copy(rotatedSlice[index:], end)
        }
    } else {
        copy(rotatedSlice, slice)
    }

    return rotatedSlice
