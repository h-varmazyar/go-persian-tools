package phone

import (
	"errors"
	"fmt"
	"regexp"
)

const (
	IranCell = "ایرانسل"
	MCI      = "همراه اول"
	TCI      = "مخابرات ایران"
	RighTel  = "رایتل"
	Talya    = "تالیا"
	MTCE     = "اسپادان"
	TeleKish = "تله‌کیش"
	ApTel    = "آپ‌تل"
	Azartel  = "آذرتل"
	SamanTel = "سامانتل"
	LotusTel = "لوتوس‌تل"
	Shatel   = "شاتل موبایل"
)

const (
	PrePaid  = "اعتباری"
	PostPaid = "دائمی"
	Child    = "سیمکارت کودک"
	TDLte    = "TD-Lte"
	Unknown  = "نامشخص"
	Landline = "ثابت"
)

const (
	AllProvinces          = "همه استان‌ها"
	EastAzerbaijan        = "آذربایجان شرقی"
	WestAzerbaijan        = "آذربایجان غربی"
	Ardabil               = "اردبیل"
	Isfahan               = "اصفهان"
	Alborz                = "البرز"
	Ilam                  = "ایلام"
	Boushehr              = "بوشهر"
	Tehran                = "تهران"
	ChaharmahalBakhtiyari = "چهارمحال و بختیاری"
	NorthKhorasan         = "خراسان شمالی"
	RazaviKhorasan        = "خراسان رضوی"
	SouthKhorasan         = "خراسان جنوبی"
	Khouzestan            = "خوزستان"
	Zanjan                = "زنجان"
	Semnan                = "سمنان"
	SistanBalouchestan    = "سیستان و بلوچستان"
	Fars                  = "فارس"
	Qazvin                = "قزوین"
	Qom                   = "قم"
	Kordestan             = "کردستان"
	Kerman                = "کرمان"
	Kermanshah            = "کرمانشاه"
	KohlilouyeBoyerahmad  = "کهکیلویه و بویراحمد"
	Golestan              = "گلستان"
	Gilan                 = "گیلان"
	Lorestan              = "لرستان"
	Mazandaran            = "مازندران"
	Markazi               = "مرکزی"
	Hormozgan             = "هرمزگان"
	Hamadan               = "همدان"
	Yazd                  = "یزد"
)

var (
	ErrPhoneNotValid   = errors.New(`phone number not valid`)
	ErrMalformedPhone  = errors.New(`malformed phone number`)
	ErrInvalidCityCode = errors.New(`invalid city code`)
)

type Phone struct {
	Code          string
	Base          string
	FullNumber    string
	TrimmedNumber string
	Operator      string
	Provinces     []string
	Type          string
}

func GetPhoneNumberDetails(phoneNumber string) (*Phone, error) {
	if r, b := IsMobile(phoneNumber); b {
		return parseMobile(r, phoneNumber)
	} else if r, b := IsLandline(phoneNumber); b {
		return parseLandline(r, phoneNumber)
	} else {
		return nil, ErrPhoneNotValid
	}
}

func IsMobile(number string) (*regexp.Regexp, bool) {
	r := regexp.MustCompile(`^(\+98|98|0098)?0?(9\d{9}$)`)
	return r, r.MatchString(number)
}

func IsLandline(number string) (*regexp.Regexp, bool) {
	r := regexp.MustCompile(`^(\+98|98|0098)?0?(\d{10}$)`)
	return r, r.MatchString(number)
}

func parseMobile(r *regexp.Regexp, number string) (*Phone, error) {
	matched := r.FindStringSubmatch(number)
	if len(matched) != 3 {
		return nil, ErrMalformedPhone
	}
	resp := new(Phone)
	resp.FullNumber = fmt.Sprintf("0%s", matched[2])
	resp.TrimmedNumber = matched[2]
	resp.Code = resp.FullNumber[:4]
	resp.Base = resp.FullNumber[4:]

	mvno := false
	switch resp.Code {
	case "0930", "0933", "0935", "0936", "0937", "0938", "0939", "0900", "0903", "0905":
		resp.Operator = IranCell
		resp.Type = Unknown
		resp.Provinces = []string{AllProvinces}
	case "0901":
		resp.Operator = IranCell
		resp.Type = PrePaid
		resp.Provinces = []string{AllProvinces}
	case "0902":
		resp.Operator = IranCell
		resp.Type = PostPaid
		resp.Provinces = []string{AllProvinces}
	case "0904":
		resp.Operator = IranCell
		resp.Type = Child
		resp.Provinces = []string{AllProvinces}
	case "0941":
		resp.Operator = IranCell
		resp.Type = TDLte
		resp.Provinces = []string{AllProvinces}

	case "0920":
		resp.Operator = RighTel
		resp.Type = PostPaid
		resp.Provinces = []string{AllProvinces}
	case "0921":
		resp.Operator = RighTel
		resp.Type = PrePaid
		resp.Provinces = []string{AllProvinces}
	case "0922":
		resp.Operator = RighTel
		resp.Type = PrePaid
		resp.Provinces = []string{AllProvinces}

	case "0910":
		resp.Operator = MCI
		resp.Type = Unknown
		resp.Provinces = []string{AllProvinces}
	case "0911":
		resp.Operator = MCI
		resp.Type = Unknown
		resp.Provinces = []string{Mazandaran, Golestan, Gilan}
	case "0912":
		resp.Operator = MCI
		resp.Type = PostPaid
		resp.Provinces = []string{Tehran, Alborz, Zanjan, Semnan, Qazvin, Qom, Markazi}
	case "0913":
		resp.Operator = MCI
		resp.Type = Unknown
		resp.Provinces = []string{Isfahan, Yazd, ChaharmahalBakhtiyari, Kerman}
	case "0914":
		resp.Operator = MCI
		resp.Type = Unknown
		resp.Provinces = []string{EastAzerbaijan, WestAzerbaijan, Ardabil, Isfahan}
	case "0915":
		resp.Operator = MCI
		resp.Type = Unknown
		resp.Provinces = []string{RazaviKhorasan, SouthKhorasan, NorthKhorasan, SistanBalouchestan}
	case "0916":
		resp.Operator = MCI
		resp.Type = Unknown
		resp.Provinces = []string{Khouzestan, Lorestan, Fars, Isfahan}
	case "0917":
		resp.Operator = MCI
		resp.Type = PrePaid
		resp.Provinces = []string{Fars, Boushehr, KohlilouyeBoyerahmad, Hormozgan}
	case "0918":
		resp.Operator = MCI
		resp.Type = PrePaid
		resp.Provinces = []string{Kermanshah, Kordestan, Ilam, Hamadan}
	case "0919":
		resp.Operator = MCI
		resp.Type = PrePaid
		resp.Provinces = []string{AllProvinces}
	case "0990":
		resp.Operator = MCI
		resp.Type = PrePaid
		resp.Provinces = []string{AllProvinces}
	case "0991":
		resp.Operator = MCI
		resp.Type = Unknown
		resp.Provinces = []string{AllProvinces}
	case "0992":
		resp.Operator = MCI
		resp.Type = PrePaid
		resp.Provinces = []string{AllProvinces}
	case "0993":
		resp.Operator = MCI
		resp.Type = PrePaid
		resp.Provinces = []string{AllProvinces}
	case "0994":
		resp.Operator = MCI
		resp.Type = PrePaid
		resp.Provinces = []string{AllProvinces}

	case "0932":
		resp.Operator = Talya
		resp.Type = PrePaid
		resp.Provinces = []string{AllProvinces}

	case "0931":
		resp.Operator = MTCE
		resp.Type = PrePaid
		resp.Provinces = []string{AllProvinces}

	case "0934":
		resp.Operator = TeleKish
		resp.Type = PostPaid
		resp.Provinces = []string{Hormozgan}
	default:
		resp.Code = resp.FullNumber[:6]
		resp.Base = resp.FullNumber[6:]
		mvno = true

	}
	if mvno {
		switch resp.Code {
		case "099910":
			resp.Operator = ApTel
			resp.Type = PostPaid
			resp.Provinces = []string{AllProvinces}
		case "099911":
			resp.Operator = ApTel
			resp.Type = PostPaid
			resp.Provinces = []string{AllProvinces}
		case "099913":
			resp.Operator = ApTel
			resp.Type = PostPaid
			resp.Provinces = []string{AllProvinces}

		case "099914":
			resp.Operator = Azartel
			resp.Type = PrePaid
			resp.Provinces = []string{AllProvinces}

		case "099999":
			resp.Operator = SamanTel
			resp.Type = PostPaid
			resp.Provinces = []string{AllProvinces}
		case "099998":
			resp.Operator = SamanTel
			resp.Type = PrePaid
			resp.Provinces = []string{AllProvinces}
		case "099997":
			resp.Operator = SamanTel
			resp.Type = PrePaid
			resp.Provinces = []string{AllProvinces}
		case "099996":
			resp.Operator = SamanTel
			resp.Type = PostPaid
			resp.Provinces = []string{AllProvinces}

		case "099810":
			resp.Operator = Shatel
			resp.Type = PrePaid
			resp.Provinces = []string{AllProvinces}
		case "099811":
			resp.Operator = Shatel
			resp.Type = PrePaid
			resp.Provinces = []string{AllProvinces}
		case "099812":
			resp.Operator = Shatel
			resp.Type = PrePaid
			resp.Provinces = []string{AllProvinces}
		case "099814":
			resp.Operator = Shatel
			resp.Type = PrePaid
			resp.Provinces = []string{AllProvinces}
		case "099815":
			resp.Operator = Shatel
			resp.Type = PrePaid
			resp.Provinces = []string{AllProvinces}
		}
	}
	return resp, nil
}

func parseLandline(r *regexp.Regexp, number string) (*Phone, error) {
	matched := r.FindStringSubmatch(number)
	if len(matched) != 3 {
		return nil, ErrMalformedPhone
	}
	resp := new(Phone)
	resp.FullNumber = fmt.Sprintf("0%s", matched[2])
	resp.TrimmedNumber = matched[2]
	resp.Code = resp.FullNumber[:3]
	resp.Base = resp.FullNumber[3:]

	switch resp.Code {
	case "041":
		resp.Provinces = []string{EastAzerbaijan}
	case "044":
		resp.Provinces = []string{WestAzerbaijan}
	case "045":
		resp.Provinces = []string{Ardabil}
	case "031":
		resp.Provinces = []string{Isfahan}
	case "026":
		resp.Provinces = []string{Alborz}
	case "084":
		resp.Provinces = []string{Ilam}
	case "077":
		resp.Provinces = []string{Boushehr}
	case "021":
		resp.Provinces = []string{Tehran}
	case "038":
		resp.Provinces = []string{ChaharmahalBakhtiyari}
	case "056":
		resp.Provinces = []string{SouthKhorasan}
	case "051":
		resp.Provinces = []string{RazaviKhorasan}
	case "058":
		resp.Provinces = []string{NorthKhorasan}
	case "061":
		resp.Provinces = []string{Khouzestan}
	case "024":
		resp.Provinces = []string{Zanjan}
	case "023":
		resp.Provinces = []string{Semnan}
	case "054":
		resp.Provinces = []string{SistanBalouchestan}
	case "071":
		resp.Provinces = []string{Fars}
	case "028":
		resp.Provinces = []string{Qazvin}
	case "025":
		resp.Provinces = []string{Qom}
	case "087":
		resp.Provinces = []string{Kordestan}
	case "034":
		resp.Provinces = []string{Kerman}
	case "083":
		resp.Provinces = []string{Kermanshah}
	case "074":
		resp.Provinces = []string{KohlilouyeBoyerahmad}
	case "017":
		resp.Provinces = []string{Golestan}
	case "013":
		resp.Provinces = []string{Gilan}
	case "066":
		resp.Provinces = []string{Lorestan}
	case "011":
		resp.Provinces = []string{Mazandaran}
	case "086":
		resp.Provinces = []string{Markazi}
	case "076":
		resp.Provinces = []string{Hormozgan}
	case "081":
		resp.Provinces = []string{Hamadan}
	case "035":
		resp.Provinces = []string{Yazd}
	default:
		return nil, ErrInvalidCityCode
	}
	resp.Operator = TCI
	resp.Type = Landline
	return resp, nil
}
