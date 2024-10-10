package sevenzip

import (
	"errors"

	"github.com/ricochhet/simplefs"
	"github.com/ricochhet/simpleprocess"
)

type ErrorCode int

const (
	NoError ErrorCode = iota
	ProcessNotFound
	CouldNotExtract
	CouldNotCompress
)

const (
	SzCompressionFormat         string = "7z"
	SzCompressionLevel          string = "-mx9"
	SzCompressionMethod         string = "-m0=lzma2"
	SzCompressionDictionarySize string = "-md=64m"
	SzCompressionFastBytes      string = "-mfb=64"
	SzCompressionSolidBlockSize string = "-ms=4g"
	SzCompressionMultithreading string = "-mmt=2"
	SzCompressionMemory         string = "-mmemuse=26g"
)

var errSevenZipNotFound = errors.New("7zip was not found")

func SzExtract(src, dest string, silent bool) (ErrorCode, error) {
	if simplefs.Exists("redist/win64/7z.exe") {
		if err := simpleprocess.RunFile("redist/win64/7z.exe", true, true, false, "x", src, "-o"+dest+"/*"); err != nil {
			return CouldNotCompress, err
		}

		return NoError, nil
	}

	if !simpleprocess.DoesFileExist("7z") {
		return ProcessNotFound, errSevenZipNotFound
	}

	if err := simpleprocess.RunFile("7z", true, false, silent, "x", src, "-o"+dest+"/*"); err != nil {
		return CouldNotExtract, err
	}

	return NoError, nil
}

func SzBinExtract(src, dest, bin string, silent bool) (ErrorCode, error) {
	if !simplefs.Exists(bin) {
		return ProcessNotFound, errSevenZipNotFound
	}

	if err := simpleprocess.RunFile(bin, true, true, silent, "x", src, "-o"+dest+"/*"); err != nil {
		return CouldNotCompress, err
	}

	return NoError, nil
}

func SzCompress(src, dest string, silent bool) (ErrorCode, error) {
	if !simpleprocess.DoesFileExist("7z") {
		return ProcessNotFound, errSevenZipNotFound
	}

	//nolint:lll // wontfix
	if err := simpleprocess.RunFile("7z", true, false, silent, "a", "-t"+SzCompressionFormat, dest, src+"/*", SzCompressionLevel, SzCompressionMethod, SzCompressionDictionarySize, SzCompressionFastBytes, SzCompressionSolidBlockSize, SzCompressionMultithreading, SzCompressionMemory); err != nil {
		return CouldNotCompress, err
	}

	return NoError, nil
}

func SzBinCompress(src, dest, bin string, silent bool) (ErrorCode, error) {
	if !simplefs.Exists(bin) {
		return ProcessNotFound, errSevenZipNotFound
	}

	//nolint:lll // wontfix
	if err := simpleprocess.RunFile(bin, true, true, silent, "a", "-t"+SzCompressionFormat, dest, src+"/*", SzCompressionLevel, SzCompressionMethod, SzCompressionDictionarySize, SzCompressionFastBytes, SzCompressionSolidBlockSize, SzCompressionMultithreading, SzCompressionMemory); err != nil {
		return CouldNotCompress, err
	}

	return NoError, nil
}
