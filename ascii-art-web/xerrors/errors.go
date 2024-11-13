package xerrors

type AsciiError interface {
	Error() string
	Type() string
}

type asciiError struct {
	mError string
	mType  string
}

func (a asciiError) Error() string {
	return a.mError
}

func (a asciiError) Type() string {
	return a.mType
}

func New(err, _type string) AsciiError {
	return asciiError{
		mError: err,
		mType:  _type,
	}
}

const TypeInvalidAscii = "invalid ascii"

func InvalidAscii(err string) AsciiError {
	return New(err, TypeInvalidAscii)
}

const TypeInvalidGraphics = "invalid graphics"

func InvalidGraphics(err string) AsciiError {
	return New(err, TypeInvalidGraphics)
}

const TypeInvalidBanner = "invalid banner"

func InvalidBanner(err string) AsciiError {
	return New(err, TypeInvalidBanner)
}
