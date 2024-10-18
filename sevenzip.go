package sevenzip

import (
	"github.com/ricochhet/simplefs"
	"github.com/ricochhet/simpleprocess"
)

func SzExtract(src, dest string, silent bool) (ErrorCode, error) {
	if simplefs.Exists("redist/win64/7z.exe") {
		if err := simpleprocess.RunFile("redist/win64/7z.exe", true, true, false, "x", src, "-o"+dest+"/*"); err != nil {
			return CouldNotCompress, err
		}

		return NoError, nil
	}

	if !simpleprocess.DoesFileExist("7z") {
		return ProcessNotFound, ErrSevenZipNotFound
	}

	if err := simpleprocess.RunFile("7z", true, false, silent, "x", src, "-o"+dest+"/*"); err != nil {
		return CouldNotExtract, err
	}

	return NoError, nil
}

func SzBinExtract(src, dest, bin string, silent bool) (ErrorCode, error) {
	if !simplefs.Exists(bin) {
		return ProcessNotFound, ErrSevenZipNotFound
	}

	if err := simpleprocess.RunFile(bin, true, true, silent, "x", src, "-o"+dest+"/*"); err != nil {
		return CouldNotCompress, err
	}

	return NoError, nil
}

func SzCompress(src, dest string, silent bool, opts ...Options) (ErrorCode, error) {
	opt := assureOptions(opts...)

	if !simpleprocess.DoesFileExist("7z") {
		return ProcessNotFound, ErrSevenZipNotFound
	}

	//nolint:lll // wontfix
	if err := simpleprocess.RunFile("7z", true, false, silent, "a", "-t"+opt.SzCompressionFormat, dest, src+"/*", opt.SzCompressionLevel, opt.SzCompressionMethod, opt.SzCompressionDictionarySize, opt.SzCompressionFastBytes, opt.SzCompressionSolidBlockSize, opt.SzCompressionMultithreading, opt.SzCompressionMemory); err != nil {
		return CouldNotCompress, err
	}

	return NoError, nil
}

func SzBinCompress(src, dest, bin string, silent bool, opts ...Options) (ErrorCode, error) {
	opt := assureOptions(opts...)

	if !simplefs.Exists(bin) {
		return ProcessNotFound, ErrSevenZipNotFound
	}

	//nolint:lll // wontfix
	if err := simpleprocess.RunFile(bin, true, true, silent, "a", "-t"+opt.SzCompressionFormat, dest, src+"/*", opt.SzCompressionLevel, opt.SzCompressionMethod, opt.SzCompressionDictionarySize, opt.SzCompressionFastBytes, opt.SzCompressionSolidBlockSize, opt.SzCompressionMultithreading, opt.SzCompressionMemory); err != nil {
		return CouldNotCompress, err
	}

	return NoError, nil
}
