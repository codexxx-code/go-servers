package zodiac

type Zodiac string

const (
	Aries       Zodiac = "Aries"
	Taurus      Zodiac = "Taurus"
	Gemini      Zodiac = "Gemini"
	Cancer      Zodiac = "Cancer"
	Leo         Zodiac = "Leo"
	Virgo       Zodiac = "Virgo"
	Libra       Zodiac = "Libra"
	Scorpio     Zodiac = "Scorpio"
	Sagittarius Zodiac = "Sagittarius"
	Capricorn   Zodiac = "Capricorn"
	Aquarius    Zodiac = "Aquarius"
	Pisces      Zodiac = "Pisces"
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
