package zodiac

import "pkg/errors"

type Zodiac string

// enums:"aries,taurus,gemini,cancer,leo,virgo,libra,scorpio,sagittarius,capricorn,aquarius,pisces"
const (
	Aries       Zodiac = "aries"
	Taurus      Zodiac = "taurus"
	Gemini      Zodiac = "gemini"
	Cancer      Zodiac = "cancer"
	Leo         Zodiac = "leo"
	Virgo       Zodiac = "virgo"
	Libra       Zodiac = "libra"
	Scorpio     Zodiac = "scorpio"
	Sagittarius Zodiac = "sagittarius"
	Capricorn   Zodiac = "capricorn"
	Aquarius    Zodiac = "aquarius"
	Pisces      Zodiac = "pisces"
)

func (z Zodiac) Validate() error {
	switch z {
	case Aries, Taurus, Gemini, Cancer, Leo, Virgo, Libra, Scorpio, Sagittarius, Capricorn, Aquarius, Pisces:
		return nil
	default:
		return errors.BadRequest.New("Unknown zodiac", errors.ParamsOption("zodiac", z))
	}
}
