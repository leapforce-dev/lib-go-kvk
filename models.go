package kvk

type Resultaat struct {
	Pagina     int32           `json:"pagina"`
	Aantal     int32           `json:"aantal"`
	Totaal     int32           `json:"totaal"`
	Vorige     string          `json:"vorige"`
	Volgende   string          `json:"volgende"`
	Resultaten []ResultaatItem `json:"resultaten"`
	Links      []Link          `json:"links"`
}

type ResultaatItem struct {
	KvkNummer            string `json:"kvkNummer"`
	Rsin                 string `json:"rsin"`
	Vestigingsnummer     string `json:"vestigingsnummer"`
	Handelsnaam          string `json:"handelsnaam"`
	Straatnaam           string `json:"straatnaam"`
	Huisnummer           int64  `json:"huisnummer"`
	HuisnummerToevoeging string `json:"huisnummerToevoeging"`
	Postcode             string `json:"postcode"`
	Plaats               string `json:"plaats"`
	Type                 string `json:"type"`
}

type Adres struct {
	Type                 string   `json:"type"`
	IndAfgeschermd       string   `json:"indAfgeschermd"`
	VolledigAdres        string   `json:"volledigAdres"`
	Straatnaam           string   `json:"straatnaam"`
	Huisnummer           int32    `json:"huisnummer"`
	HuisnummerToevoeging string   `json:"huisnummerToevoeging"`
	Huisletter           string   `json:"huisletter"`
	ToevoegingAdres      string   `json:"toevoegingAdres"`
	Postcode             string   `json:"postcode"`
	Postbusnummer        int32    `json:"postbusnummer"`
	Plaats               string   `json:"plaats"`
	StraatHuisnummer     string   `json:"straatHuisnummer"`
	PostcodeWoonplaats   string   `json:"postcodeWoonplaats"`
	Regio                string   `json:"regio"`
	Land                 string   `json:"land"`
	GeoData              *GeoData `json:"geoData"`
}

type Basisprofiel struct {
	KvkNummer               string               `json:"kvkNummer"`
	IndNonMailing           string               `json:"indNonMailing"`
	Naam                    string               `json:"naam"`
	FormeleRegistratiedatum string               `json:"formeleRegistratiedatum"`
	MaterieleRegistratie    MaterieleRegistratie `json:"materieleRegistratie"`
	TotaalWerkzamePersonen  int                  `json:"totaalWerkzamePersonen"`
	StatutaireNaam          string               `json:"statutaireNaam"`
	Handelsnamen            []Handelsnaam        `json:"handelsnamen"`
	SbiActiviteiten         []SbiActiviteit      `json:"sbiActiviteiten"`
	Links                   []Link               `json:"links"`
	Embedded                EmbeddedContainer    `json:"_embedded"`
}

type Eigenaar struct {
	Rsin                  string   `json:"rsin"`
	Rechtsvorm            string   `json:"rechtsvorm"`
	UitgebreideRechtsvorm string   `json:"uitgebreideRechtsvorm"`
	Adressen              []Adres  `json:"adressen"`
	Websites              []string `json:"websites"`
	Links                 []Link   `json:"links"`
}

type EmbeddedContainer struct {
	Hoofdvestiging Vestiging `json:"hoofdvestiging"`
	Eigenaar       Eigenaar  `json:"eigenaar"`
}

type GeoData struct {
	AddresseerbaarObjectId string  `json:"addresseerbaarObjectId"`
	NummerAanduidingId     string  `json:"nummerAanduidingId"`
	GpsLatitude            float64 `json:"gpsLatitude"`
	GpsLongitude           float64 `json:"gpsLongitude"`
	RijksdriehoekX         float64 `json:"rijksdriehoekX"`
	RijksdriehoekY         float64 `json:"rijksdriehoekY"`
	RijksdriehoekZ         float64 `json:"rijksdriehoekZ"`
}

type Handelsnaam struct {
	Naam     string `json:"naam"`
	Volgorde int    `json:"volgorde"`
}

type Link struct {
	Rel         string `json:"rel"`
	Href        string `json:"href"`
	HrefLang    string `json:"hreflang"`
	Media       string `json:"media"`
	Title       string `json:"title"`
	Type        string `json:"type"`
	Deprecation string `json:"deprecation"`
	Profile     string `json:"profile"`
	Name        string `json:"name"`
}

type MaterieleRegistratie struct {
	DatumAanvang string `json:"datumAanvang"`
	DatumEinde   string `json:"datumEinde"`
}

type SbiActiviteit struct {
	SbiCode            string `json:"sbiCode"`
	SbiOmschrijving    string `json:"sbiOmschrijving"`
	IndHoofdactiviteit string `json:"indHoofdactiviteit"`
}

type Vestiging struct {
	Vestigingsnummer         string               `json:"vestigingsnummer"`
	KvkNummer                string               `json:"kvkNummer"`
	Rsin                     string               `json:"rsin"`
	IndNonMailing            string               `json:"indNonMailing"`
	MaterieleRegistratie     MaterieleRegistratie `json:"materieleRegistratie"`
	EersteHandelsnaam        string               `json:"eersteHandelsnaam"`
	IndHoofdvestiging        string               `json:"indHoofdvestiging"`
	IndCommercieleVestiging  string               `json:"indCommercieleVestiging"`
	VoltijdWerkzamePersonen  int                  `json:"voltijdWerkzamePersonen"`
	TotaalWerkzamePersonen   int                  `json:"totaalWerkzamePersonen"`
	DeeltijdWerkzamePersonen int                  `json:"deeltijdWerkzamePersonen"`
	Handelsnamen             []Handelsnaam        `json:"handelsnamen"`
	Adressen                 []Adres              `json:"adressen"`
	Websites                 []string             `json:"websites"`
	SbiActiviteiten          []SbiActiviteit      `json:"sbiActiviteiten"`
	Links                    []Link               `json:"links"`
}

type VestigingBasis struct {
	Vestigingsnummer        string `json:"vestigingsnummer"`
	KvkNummer               string `json:"kvkNummer"`
	EersteHandelsnaam       string `json:"eersteHandelsnaam"`
	IndHoofdvestiging       string `json:"indHoofdvestiging"`
	IndAdresAfgeschermd     string `json:"indAdresAfgeschermd"`
	IndCommercieleVestiging string `json:"indCommercieleVestiging"`
	VolledigAdres           string `json:"volledigAdres"`
	Links                   []Link `json:"links"`
}

type VestigingList struct {
	KvkNummer                        string           `json:"kvkNummer"`
	AantalCommercieleVestigingen     int64            `json:"aantalCommercieleVestigingen"`
	AantalNietCommercieleVestigingen int64            `json:"aantalNietCommercieleVestigingen"`
	TotaalAantalVestigingen          int64            `json:"totaalAantalVestigingen"`
	Vestigingen                      []VestigingBasis `json:"vestigingen"`
	Links                            []Link           `json:"links"`
}
