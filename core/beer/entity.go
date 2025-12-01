package beer

//Tags json = `` 
//
type Beer struct {
	ID    int64     `json:"id"`
	Name  string    `json:"name"`
	Type  BeerType  `json:"type"`
	Style BeerStyle `json:"style"`
}


type BeerType int 


const (	
	TypeAle = 1 
	TypeLager = 2
	TypeStout = 3
	TypePilsner = 4
	TypeWheat = 5
)

func (t BeerType) String() string {
	switch t {
	case TypeAle:
		return "Ale"
	case TypeLager:
		return "Lager"
	case TypeStout:
		return "Stout"
	case TypePilsner:
		return "Pilsner"
	case TypeWheat:
		return "Wheat"
	default:
		return "Unknown"
	}


}

type BeerStyle int


const (
	StyleAmber = iota + 1 
	StyleBlonde
	StyleBrown
	StyleIPA
	StylePorter
	StyleRed
	StyleSour
	StyleStrong
	StyleLight
	StyleDark
)

func (t BeerStyle) String() string {
	switch t {
	case StyleAmber:
		return "Amber"
	case StyleBlonde:
		return "Blonde"
	case StyleBrown:
		return "Brown"
	case StyleIPA:
		return "IPA"
	case StylePorter:
		return "Porter"
	case StyleRed:
		return "Red"
	case StyleSour:
		return "Sour"
	case StyleStrong:
		return "Strong"
	case StyleLight:
		return "Light"
	case StyleDark:
		return "Dark"
	default:
		return "Unknown"
	}
}