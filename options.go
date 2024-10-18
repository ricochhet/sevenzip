package sevenzip

type Options struct {
	SzCompressionFormat         string
	SzCompressionLevel          string
	SzCompressionMethod         string
	SzCompressionDictionarySize string
	SzCompressionFastBytes      string
	SzCompressionSolidBlockSize string
	SzCompressionMultithreading string
	SzCompressionMemory         string
}

func getDefaultOptions() Options {
	return Options{
		SzCompressionFormat:         "7z",
		SzCompressionLevel:          "-mx9",
		SzCompressionMethod:         "-m0=lzma2",
		SzCompressionDictionarySize: "-md=64m",
		SzCompressionFastBytes:      "-mfb=64",
		SzCompressionSolidBlockSize: "-ms=4g",
		SzCompressionMultithreading: "-mmt=2",
		SzCompressionMemory:         "-mmemuse=26g",
	}
}

func assureOptions(opts ...Options) Options {
	defopt := getDefaultOptions()

	if len(opts) == 0 {
		return defopt
	}

	return opts[0]
}
