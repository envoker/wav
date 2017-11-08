package samples

const (
	minInt8 = -1 << (8 - 1)
	maxInt8 = (1 << (8 - 1)) - 1

	minInt16 = -1 << (16 - 1)
	maxInt16 = (1 << (16 - 1)) - 1

	minInt24 = -1 << (24 - 1)
	maxInt24 = (1 << (24 - 1)) - 1

	minInt32 = -1 << (32 - 1)
	maxInt32 = (1 << (32 - 1)) - 1
)

func MaxSample(bitsPerSample int) int32 {
	switch bitsPerSample {
	case 8:
		return maxInt8
	case 16:
		return maxInt16
	case 24:
		return maxInt24
	case 32:
		return maxInt32
	default:
		return 0
	}
}
