package zodiac

import "pkg/errors"

type Zodiac string

// enum:"aries,taurus,gemini,cancer,leo,virgo,libra,scorpio,sagittarius,capricorn,aquarius,pisces"
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

var Array = []Zodiac{
	Aries,
	Taurus,
	Gemini,
	Cancer,
	Leo,
	Virgo,
	Libra,
	Scorpio,
	Sagittarius,
	Capricorn,
	Aquarius,
	Pisces,
}

func (z Zodiac) ToRussian() string {
	switch z {
	case Aries:
		return "Овен"
	case Taurus:
		return "Телец"
	case Gemini:
		return "Близнецы"
	case Cancer:
		return "Рак"
	case Leo:
		return "Лев"
	case Virgo:
		return "Дева"
	case Libra:
		return "Весы"
	case Scorpio:
		return "Скорпион"
	case Sagittarius:
		return "Стрелец"
	case Capricorn:
		return "Козерог"
	case Aquarius:
		return "Водолей"
	case Pisces:
		return "Рыбы"
	default:
		return ""
	}
}

func (z Zodiac) Validate() error {
	switch z {
	case Aries, Taurus, Gemini, Cancer, Leo, Virgo, Libra, Scorpio, Sagittarius, Capricorn, Aquarius, Pisces:
		return nil
	default:
		return errors.BadRequest.New("Unknown zodiac", errors.ParamsOption("zodiac", z))
	}
}
