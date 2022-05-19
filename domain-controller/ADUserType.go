package main

import (
	"errors"
	"strconv"
)

type ADUSER map[string]interface{}

type ADGROUP map[string]interface{}

type ADOU map[string]interface{}

type ABLEGROUP map[string]interface{}

type ADCOMPUTER map[string]interface{}
type ADComputer struct {
	distinguishedName string `uri:"distinguishedName" form:"distinguishedName"`
	cn                string `uri:"cn" from:"cn"`

	name           string `uri:"name" from:"name"`
	sAMAccountName string `uri:"sAMAccountName" from:"sAMAccountName"`
}

type ADUser struct {
	Sn                         string   `json:"sn" uri:"sn" form:"sn" example:"user"`                                                                         //성
	GivenName                  string   `json:"givenName" uri:"givenName" form:"givenName" example:"user"`                                                    //이름
	Initials                   string   `json:"initials" uri:"initials" form:"initials" example:"user"`                                                       //이니셜
	Accountname                string   `uri:"accountname" json:"accountname" form:"accountname" example:"user"`                                              //어카운트명
	UserPrincipalName          string   `uri:"userPrincipalName" json:"userPrincipalName" form:"userPrincipalName" example:"user"`                            //로그온이름(accountname@ADdomain형식)
	Username                   string   `uri:"username" json:	"username" form:"username" example:"user"`                                                      //works로그인명
	SAMAccountName             string   `uri:"sAMAccountName" json:"sAMAccountName" form:"sAMAccountName" example:"user"`                                     //windows2000이전사용자로그온이름(ADdomain\sAMAccountName형식)
	Description                string   `uri:"description" json:"description" form:"description" example:"user"`                                              //설명
	Info                       string   `uri:"info" json:"info" form:"info" example:"user"`                                                                   //참고내용
	Title                      string   `uri:"title" json:"title" form:"title" example:"user"`                                                                //직함
	O                          string   `uri:"o" json:"o" form:"o" example:"user"`                                                                            //ldap회사명
	Company                    string   `uri:"company" json:"company" form:"company" example:"user"`                                                          //AD회사명
	PostOfficeBox              string   `uri:"postOfficeBox" json:"postOfficeBox" form:"postOfficeBox" example:"user"`                                        //사서함주소
	PhysicalDeliveryOfficeName string   `uri:"physicalDeliveryOfficeName" json:"physicalDeliveryOfficeName" form:"physicalDeliveryOfficeName" example:"user"` //사무실주소
	StreetAddress              string   `uri:"streetAddress" json:"streetAddress" form:"streetAddress" example:"user"`                                        //주소
	I                          string   `uri:"I" json:"I" form:"I" example:"user"`                                                                            //구/군/시
	St                         string   `uri:"st" json:"st" form:"st" example:"user"`                                                                         //시/도
	Department                 string   `uri:"department" json:"department" form:"department" example:"user"`                                                 //부서
	Mail                       string   `uri:"mail" json:"mail" form:"mail" example:"user"`                                                                   //메일주소
	TelephoneNumber            string   `uri:"telephoneNumber" json:"telephoneNumber" form:"telephoneNumber" example:"user"`                                  //일반->전화
	Pager                      string   `uri:"pager" json:"pager" form:"pager" example:"user"`                                                                //전화->호출기
	Mobile                     string   `uri:"mobile" json:"mobile" form:"mobile" example:"user"`                                                             //전화->휴대폰
	FacsimileTelephoneNumber   string   `uri:"facsimileTelephoneNumber" json:"facsimileTelephoneNumber" form:"facsimileTelephoneNumber" example:"user"`       //전화->팩스
	HomePhone                  string   `uri:"homePhone" json:"homePhone" form:"homePhone" example:"user"`                                                    //전화->집
	IpPhone                    string   `uri:"ipPhone" json:"ipPhone" form:"ipPhone" example:"user"`                                                          //전화->ip전화
	PostalCode                 string   `uri:"postalCode" json:"postalCode" form:"postalCode" example:"user"`                                                 //주소->우편번호
	Manager                    string   `uri:"manager" json:"manager" form:"manager" example:"user"`                                                          //상사dn("CN=User3,CN=Users,DC=dc1,DC=local")
	WWWHomePage                string   `uri:"wWWHomePage" json:"wWWHomePage" form:"wWWHomePage" example:"user"`                                              //홈페이지주소
	MemberOf                   []string `uri:"memberOf" json:"memberOf" form:"memberOf" example:"user"`                                                       //그룹dn목록
	CountryCode                int      `uri:"countryCode" json:"countryCode" form:"countryCode" example:"1"`                                                 //주소->국가숫자코드(한국:410,일본:392.미국:840)
	C                          string   `uri:"c" json:"c" form:"c" example:"user"`                                                                            //주소->국가영문코드명(한국:KR,일본:JP,미국:US)
	Co                         string   `uri:"co" json:"co" form:"co" example:"user"`                                                                         //:=c
	DistinguishedName          string   `uri:"distinguishedName" json:"distinguishedName" form:"distinguishedName" example:"user"`
}
type COUNTRY map[string]string
type COUNTRYMAP map[string]COUNTRY

var countries = COUNTRYMAP{
	"AF": {"name": "Afghanistan", "string": "AF", "code": "4"},
	"AL": {"name": "Albania", "string": "AL", "code": "8"},
	"DZ": {"name": "Algeria", "string": "DZ", "code": "12"},
	"AS": {"name": "American Samoa", "string": "AS", "code": "16"},
	"AD": {"name": "Andorra", "string": "AD", "code": "20"},
	"AO": {"name": "Angola", "string": "AO", "code": "24"},
	"AI": {"name": "Anguilla", "string": "AI", "code": "660"},
	"AQ": {"name": "Antarctica", "string": "AQ", "code": "10"},
	"AG": {"name": "Antigua and Barbuda", "string": "AG", "code": "28"},
	"AR": {"name": "Argentina", "string": "AR", "code": "32"},
	"AM": {"name": "Armenia", "string": "AM", "code": "51"},
	"AW": {"name": "Aruba", "string": "AW", "code": "533"},
	"AU": {"name": "Australia", "string": "AU", "code": "36"},
	"AT": {"name": "Austria", "string": "AT", "code": "40"},
	"AZ": {"name": "Azerbaijan", "string": "AZ", "code": "31"},
	"BS": {"name": "Bahamas", "string": "BS", "code": "44"},
	"BH": {"name": "Bahrain", "string": "BH", "code": "48"},
	"BD": {"name": "Bangladesh", "string": "BD", "code": "50"},
	"BB": {"name": "Barbados", "string": "BB", "code": "52"},
	"BY": {"name": "Belarus", "string": "BY", "code": "112"},
	"BE": {"name": "Belgium", "string": "BE", "code": "56"},
	"BZ": {"name": "Belize", "string": "BZ", "code": "84"},
	"BJ": {"name": "Benin", "string": "BJ", "code": "204"},
	"BM": {"name": "Bermuda", "string": "BM", "code": "60"},
	"BT": {"name": "Bhutan", "string": "BT", "code": "64"},
	"BO": {"name": "Bolivia", "string": "BO", "code": "68"},
	"BQ": {"name": "Bonaire", "string": "BQ", "code": "535"},
	"BA": {"name": "Bosnia and Herzegovina", "string": "BA", "code": "70"},
	"BW": {"name": "Botswana", "string": "BW", "code": "72"},
	"BV": {"name": "Bouvet Island", "string": "BV", "code": "74"},
	"BR": {"name": "Brazil", "string": "BR", "code": "76"},
	"IO": {"name": "British Indian Ocean Territory", "string": "IO", "code": "86"},
	"BN": {"name": "Brunei Darussalam", "string": "BN", "code": "96"},
	"BG": {"name": "Bulgaria", "string": "BG", "code": "100"},
	"BF": {"name": "Burkina Faso", "string": "BF", "code": "854"},
	"BI": {"name": "Burundi", "string": "BI", "code": "108"},
	"KH": {"name": "Cambodia", "string": "KH", "code": "116"},
	"CM": {"name": "Cameroon", "string": "CM", "code": "120"},
	"CA": {"name": "Canada", "string": "CA", "code": "124"},
	"CV": {"name": "Cape Verde", "string": "CV", "code": "132"},
	"KY": {"name": "Cayman Islands", "string": "KY", "code": "136"},
	"CF": {"name": "Central African Republic", "string": "CF", "code": "140"},
	"TD": {"name": "Chad", "string": "TD", "code": "148"},
	"CL": {"name": "Chile", "string": "CL", "code": "152"},
	"CN": {"name": "China", "string": "CN", "code": "156"},
	"CX": {"name": "Christmas Island", "string": "CX", "code": "162"},
	"CC": {"name": "Cocos (Keeling) Islands", "string": "CC", "code": "166"},
	"CO": {"name": "Colombia", "string": "CO", "code": "170"},
	"KM": {"name": "Comoros", "string": "KM", "code": "174"},
	"CG": {"name": "Congo", "string": "CG", "code": "178"},
	"CD": {"name": "Democratic Republic of the Congo", "string": "CD", "code": "180"},
	"CK": {"name": "Cook Islands", "string": "CK", "code": "184"},
	"CR": {"name": "Costa Rica", "string": "CR", "code": "188"},
	"HR": {"name": "Croatia", "string": "HR", "code": "191"},
	"CU": {"name": "Cuba", "string": "CU", "code": "192"},
	"CW": {"name": "Curacao", "string": "CW", "code": "531"},
	"CY": {"name": "Cyprus", "string": "CY", "code": "196"},
	"CZ": {"name": "Czech Republic", "string": "CZ", "code": "203"},
	"CI": {"name": "Cote d'Ivoire", "string": "CI", "code": "384"},
	"DK": {"name": "Denmark", "string": "DK", "code": "208"},
	"DJ": {"name": "Djibouti", "string": "DJ", "code": "262"},
	"DM": {"name": "Dominica", "string": "DM", "code": "212"},
	"DO": {"name": "Dominican Republic", "string": "DO", "code": "214"},
	"EC": {"name": "Ecuador", "string": "EC", "code": "218"},
	"EG": {"name": "Egypt", "string": "EG", "code": "818"},
	"SV": {"name": "El Salvador", "string": "SV", "code": "222"},
	"GQ": {"name": "Equatorial Guinea", "string": "GQ", "code": "226"},
	"ER": {"name": "Eritrea", "string": "ER", "code": "232"},
	"EE": {"name": "Estonia", "string": "EE", "code": "233"},
	"ET": {"name": "Ethiopia", "string": "ET", "code": "231"},
	"FK": {"name": "Falkland Islands (Malvinas)", "string": "FK", "code": "238"},
	"FO": {"name": "Faroe Islands", "string": "FO", "code": "234"},
	"FJ": {"name": "Fiji", "string": "FJ", "code": "242"},
	"FI": {"name": "Finland", "string": "FI", "code": "246"},
	"FR": {"name": "France", "string": "FR", "code": "250"},
	"GF": {"name": "French Guiana", "string": "GF", "code": "254"},
	"PF": {"name": "French Polynesia", "string": "PF", "code": "258"},
	"TF": {"name": "French Southern Territories", "string": "TF", "code": "260"},
	"GA": {"name": "Gabon", "string": "GA", "code": "266"},
	"GM": {"name": "Gambia", "string": "GM", "code": "270"},
	"GE": {"name": "Georgia", "string": "GE", "code": "268"},
	"DE": {"name": "Germany", "string": "DE", "code": "276"},
	"GH": {"name": "Ghana", "string": "GH", "code": "288"},
	"GI": {"name": "Gibraltar", "string": "GI", "code": "292"},
	"GR": {"name": "Greece", "string": "GR", "code": "300"},
	"GL": {"name": "Greenland", "string": "GL", "code": "304"},
	"GD": {"name": "Grenada", "string": "GD", "code": "308"},
	"GP": {"name": "Guadeloupe", "string": "GP", "code": "312"},
	"GU": {"name": "Guam", "string": "GU", "code": "316"},
	"GT": {"name": "Guatemala", "string": "GT", "code": "320"},
	"GG": {"name": "Guernsey", "string": "GG", "code": "831"},
	"GN": {"name": "Guinea", "string": "GN", "code": "324"},
	"GW": {"name": "Guinea-Bissau", "string": "GW", "code": "624"},
	"GY": {"name": "Guyana", "string": "GY", "code": "328"},
	"HT": {"name": "Haiti", "string": "HT", "code": "332"},
	"HM": {"name": "Heard Island and McDonald Mcdonald Islands", "string": "HM", "code": "334"},
	"VA": {"name": "Holy See (Vatican City State)", "string": "VA", "code": "336"},
	"HN": {"name": "Honduras", "string": "HN", "code": "340"},
	"HK": {"name": "Hong Kong", "string": "HK", "code": "344"},
	"HU": {"name": "Hungary", "string": "HU", "code": "348"},
	"IS": {"name": "Iceland", "string": "IS", "code": "352"},
	"IN": {"name": "India", "string": "IN", "code": "356"},
	"ID": {"name": "Indonesia", "string": "ID", "code": "360"},
	"IR": {"name": "Iran, Islamic Republic of", "string": "IR", "code": "364"},
	"IQ": {"name": "Iraq", "string": "IQ", "code": "368"},
	"IE": {"name": "Ireland", "string": "IE", "code": "372"},
	"IM": {"name": "Isle of Man", "string": "IM", "code": "833"},
	"IL": {"name": "Israel", "string": "IL", "code": "376"},
	"IT": {"name": "Italy", "string": "IT", "code": "380"},
	"JM": {"name": "Jamaica", "string": "JM", "code": "388"},
	"JP": {"name": "Japan", "string": "JP", "code": "392"},
	"JE": {"name": "Jersey", "string": "JE", "code": "832"},
	"JO": {"name": "Jordan", "string": "JO", "code": "400"},
	"KZ": {"name": "Kazakhstan", "string": "KZ", "code": "398"},
	"KE": {"name": "Kenya", "string": "KE", "code": "404"},
	"KI": {"name": "Kiribati", "string": "KI", "code": "296"},
	"KR": {"name": "Korea, Republic of", "string": "KR", "code": "410"},
	"KW": {"name": "Kuwait", "string": "KW", "code": "414"},
	"KG": {"name": "Kyrgyzstan", "string": "KG", "code": "417"},
	"LA": {"name": "Lao People's Democratic Republic", "string": "LA", "code": "418"},
	"LV": {"name": "Latvia", "string": "LV", "code": "428"},
	"LB": {"name": "Lebanon", "string": "LB", "code": "422"},
	"LS": {"name": "Lesotho", "string": "LS", "code": "426"},
	"LR": {"name": "Liberia", "string": "LR", "code": "430"},
	"LY": {"name": "Libya", "string": "LY", "code": "434"},
	"LI": {"name": "Liechtenstein", "string": "LI", "code": "438"},
	"LT": {"name": "Lithuania", "string": "LT", "code": "440"},
	"LU": {"name": "Luxembourg", "string": "LU", "code": "442"},
	"MO": {"name": "Macao", "string": "MO", "code": "446"},
	"MK": {"name": "Macedonia, the Former Yugoslav Republic of", "string": "MK", "code": "807"},
	"MG": {"name": "Madagascar", "string": "MG", "code": "450"},
	"MW": {"name": "Malawi", "string": "MW", "code": "454"},
	"MY": {"name": "Malaysia", "string": "MY", "code": "458"},
	"MV": {"name": "Maldives", "string": "MV", "code": "462"},
	"ML": {"name": "Mali", "string": "ML", "code": "466"},
	"MT": {"name": "Malta", "string": "MT", "code": "470"},
	"MH": {"name": "Marshall Islands", "string": "MH", "code": "584"},
	"MQ": {"name": "Martinique", "string": "MQ", "code": "474"},
	"MR": {"name": "Mauritania", "string": "MR", "code": "478"},
	"MU": {"name": "Mauritius", "string": "MU", "code": "480"},
	"YT": {"name": "Mayotte", "string": "YT", "code": "175"},
	"MX": {"name": "Mexico", "string": "MX", "code": "484"},
	"FM": {"name": "Micronesia, Federated States of", "string": "FM", "code": "583"},
	"MD": {"name": "Moldova, Republic of", "string": "MD", "code": "498"},
	"MC": {"name": "Monaco", "string": "MC", "code": "492"},
	"MN": {"name": "Mongolia", "string": "MN", "code": "496"},
	"ME": {"name": "Montenegro", "string": "ME", "code": "499"},
	"MS": {"name": "Montserrat", "string": "MS", "code": "500"},
	"MA": {"name": "Morocco", "string": "MA", "code": "504"},
	"MZ": {"name": "Mozambique", "string": "MZ", "code": "508"},
	"MM": {"name": "Myanmar", "string": "MM", "code": "104"},
	"NA": {"name": "Namibia", "string": "NA", "code": "516"},
	"NR": {"name": "Nauru", "string": "NR", "code": "520"},
	"NP": {"name": "Nepal", "string": "NP", "code": "524"},
	"NL": {"name": "Netherlands", "string": "NL", "code": "528"},
	"NC": {"name": "New Caledonia", "string": "NC", "code": "687"},
	"NZ": {"name": "New Zealand", "string": "NZ", "code": "554"},
	"NI": {"name": "Nicaragua", "string": "NI", "code": "558"},
	"NE": {"name": "Niger", "string": "NE", "code": "562"},
	"NG": {"name": "Nigeria", "string": "NG", "code": "566"},
	"NU": {"name": "Niue", "string": "NU", "code": "570"},
	"NF": {"name": "Norfolk Island", "string": "NF", "code": "574"},
	"MP": {"name": "Northern Mariana Islands", "string": "MP", "code": "580"},
	"NO": {"name": "Norway", "string": "NO", "code": "578"},
	"OM": {"name": "Oman", "string": "OM", "code": "512"},
	"PK": {"name": "Pakistan", "string": "PK", "code": "586"},
	"PW": {"name": "Palau", "string": "PW", "code": "585"},
	"PS": {"name": "Palestine, State of", "string": "PS", "code": "275"},
	"PA": {"name": "Panama", "string": "PA", "code": "591"},
	"PG": {"name": "Papua New Guinea", "string": "PG", "code": "598"},
	"PY": {"name": "Paraguay", "string": "PY", "code": "600"},
	"PE": {"name": "Peru", "string": "PE", "code": "604"},
	"PH": {"name": "Philippines", "string": "PH", "code": "608"},
	"PN": {"name": "Pitcairn", "string": "PN", "code": "612"},
	"PL": {"name": "Poland", "string": "PL", "code": "616"},
	"PT": {"name": "Portugal", "string": "PT", "code": "620"},
	"PR": {"name": "Puerto Rico", "string": "PR", "code": "630"},
	"QA": {"name": "Qatar", "string": "QA", "code": "634"},
	"RO": {"name": "Romania", "string": "RO", "code": "642"},
	"RU": {"name": "Russian Federation", "string": "RU", "code": "643"},
	"RW": {"name": "Rwanda", "string": "RW", "code": "646"},
	"RE": {"name": "Reunion", "string": "RE", "code": "638"},
	"BL": {"name": "Saint Barthelemy", "string": "BL", "code": "652"},
	"SH": {"name": "Saint Helena", "string": "SH", "code": "654"},
	"KN": {"name": "Saint Kitts and Nevis", "string": "KN", "code": "659"},
	"LC": {"name": "Saint Lucia", "string": "LC", "code": "662"},
	"MF": {"name": "Saint Martin (French part)", "string": "MF", "code": "663"},
	"PM": {"name": "Saint Pierre and Miquelon", "string": "PM", "code": "666"},
	"VC": {"name": "Saint Vincent and the Grenadines", "string": "VC", "code": "670"},
	"WS": {"name": "Samoa", "string": "WS", "code": "882"},
	"SM": {"name": "San Marino", "string": "SM", "code": "674"},
	"ST": {"name": "Sao Tome and Principe", "string": "ST", "code": "678"},
	"SA": {"name": "Saudi Arabia", "string": "SA", "code": "682"},
	"SN": {"name": "Senegal", "string": "SN", "code": "686"},
	"RS": {"name": "Serbia", "string": "RS", "code": "688"},
	"SC": {"name": "Seychelles", "string": "SC", "code": "690"},
	"SL": {"name": "Sierra Leone", "string": "SL", "code": "694"},
	"SG": {"name": "Singapore", "string": "SG", "code": "702"},
	"SX": {"name": "Sint Maarten (Dutch part)", "string": "SX", "code": "534"},
	"SK": {"name": "Slovakia", "string": "SK", "code": "703"},
	"SI": {"name": "Slovenia", "string": "SI", "code": "386"},
	"SB": {"name": "Solomon Islands", "string": "SB", "code": "90"},
	"SO": {"name": "Somalia", "string": "SO", "code": "706"},
	"ZA": {"name": "South Africa", "string": "ZA", "code": "710"},
	"GS": {"name": "South Georgia and the South Sandwich Islands", "string": "GS", "code": "239"},
	"SS": {"name": "South Sudan", "string": "SS", "code": "728"},
	"ES": {"name": "Spain", "string": "ES", "code": "724"},
	"LK": {"name": "Sri Lanka", "string": "LK", "code": "144"},
	"SD": {"name": "Sudan", "string": "SD", "code": "729"},
	"SR": {"name": "Suriname", "string": "SR", "code": "740"},
	"SJ": {"name": "Svalbard and Jan Mayen", "string": "SJ", "code": "744"},
	"SZ": {"name": "Swaziland", "string": "SZ", "code": "748"},
	"SE": {"name": "Sweden", "string": "SE", "code": "752"},
	"CH": {"name": "Switzerland", "string": "CH", "code": "756"},
	"SY": {"name": "Syrian Arab Republic", "string": "SY", "code": "760"},
	"TW": {"name": "Taiwan, Province of China", "string": "TW", "code": "158"},
	"TJ": {"name": "Tajikistan", "string": "TJ", "code": "762"},
	"TZ": {"name": "United Republic of Tanzania", "string": "TZ", "code": "834"},
	"TH": {"name": "Thailand", "string": "TH", "code": "764"},
	"TL": {"name": "Timor-Leste", "string": "TL", "code": "626"},
	"TG": {"name": "Togo", "string": "TG", "code": "768"},
	"TK": {"name": "Tokelau", "string": "TK", "code": "772"},
	"TO": {"name": "Tonga", "string": "TO", "code": "776"},
	"TT": {"name": "Trinidad and Tobago", "string": "TT", "code": "780"},
	"TN": {"name": "Tunisia", "string": "TN", "code": "788"},
	"TR": {"name": "Turkey", "string": "TR", "code": "792"},
	"TM": {"name": "Turkmenistan", "string": "TM", "code": "795"},
	"TC": {"name": "Turks and Caicos Islands", "string": "TC", "code": "796"},
	"TV": {"name": "Tuvalu", "string": "TV", "code": "798"},
	"UG": {"name": "Uganda", "string": "UG", "code": "800"},
	"UA": {"name": "Ukraine", "string": "UA", "code": "804"},
	"AE": {"name": "United Arab Emirates", "string": "AE", "code": "784"},
	"GB": {"name": "United Kingdom", "string": "GB", "code": "826"},
	"US": {"name": "United States", "string": "US", "code": "840"},
	"UM": {"name": "United States Minor Outlying Islands", "string": "UM", "code": "581"},
	"UY": {"name": "Uruguay", "string": "UY", "code": "858"},
	"UZ": {"name": "Uzbekistan", "string": "UZ", "code": "860"},
	"VU": {"name": "Vanuatu", "string": "VU", "code": "548"},
	"VE": {"name": "Venezuela", "string": "VE", "code": "862"},
	"VN": {"name": "Vietnam", "string": "VN", "code": "704"},
	"VG": {"name": "British Virgin Islands", "string": "VG", "code": "92"},
	"VI": {"name": "US Virgin Islands", "string": "VI", "code": "850"},
	"WF": {"name": "Wallis and Futuna", "string": "WF", "code": "876"},
	"EH": {"name": "Western Sahara", "string": "EH", "code": "732"},
	"YE": {"name": "Yemen", "string": "YE", "code": "887"},
	"ZM": {"name": "Zambia", "string": "ZM", "code": "894"},
	"ZW": {"name": "Zimbabwe", "string": "ZW", "code": "716"},
}

func codeToString(code int) (country COUNTRY, err error) {
	for _, val := range countries {
		valcode, _ := strconv.Atoi(val["code"])
		if valcode == code {
			return val, nil
		}
	}
	return nil, errors.New("Can not find country")
}

type ADGroup struct {
	groupname         string   `uri:"groupname" form:"groupname"`           //works 로그인명
	sAMAccountName    string   `uri:"sAMAccountName" form:"sAMAccountName"` //windows 2000 이전 사용자 로그온 이름(ADdomain\sAMAccountName 형식)
	description       string   `uri:"description" form:"description"`       //설명
	memberOf          []string `uri:"memberOf" form:"memberOf"`             //그룹 dn 목록
	member            []string `uri:"member" form:"member"`                 //그룹 dn 목록
	distinguishedName string   `uri:"distinguishedName" form:"distinguishedName"`
}

func NewADGroup(adgroup ADGROUP) (adgroupstruct *ADGroup) {
	adgroupstruct = &ADGroup{}

	if val, ok := adgroup["groupname"]; ok {
		adgroupstruct.groupname = val.(string)
	}
	if val, ok := adgroup["sAMAccountName"]; ok {
		adgroupstruct.sAMAccountName = val.([]string)[0]
	}
	if val, ok := adgroup["description"]; ok {
		adgroupstruct.description = val.([]string)[0]
	}
	if val, ok := adgroup["memberOf"]; ok {
		adgroupstruct.memberOf = val.([]string)
	}
	if val, ok := adgroup["member"]; ok {
		adgroupstruct.member = val.([]string)
	}
	if val, ok := adgroup["distinguishedName"]; ok {
		adgroupstruct.distinguishedName = val.([]string)[0]
	}
	return adgroupstruct
}

func (aduserstruct *ADGroup) ToMap() (adgroup ADGROUP) {
	adgroup = ADGROUP{
		"groupname":         aduserstruct.groupname,
		"sAMAccountName":    aduserstruct.sAMAccountName,
		"description":       aduserstruct.description,
		"memberOf":          aduserstruct.memberOf,
		"member":            aduserstruct.member,
		"distinguishedName": aduserstruct.distinguishedName,
	}
	return adgroup
}

func NewADUser(aduser ADUSER) (aduserstruct *ADUser) {
	aduserstruct = &ADUser{}

	if val, ok := aduser["sn"]; ok {
		aduserstruct.Sn = val.([]string)[0]
	}
	if val, ok := aduser["givenName"]; ok {
		aduserstruct.GivenName = val.([]string)[0]
	}
	if val, ok := aduser["initials"]; ok {
		aduserstruct.Initials = val.([]string)[0]
	}
	if val, ok := aduser["accountname"]; ok {
		aduserstruct.Accountname = val.([]string)[0]
	}
	if val, ok := aduser["userPrincipalName"]; ok {
		aduserstruct.UserPrincipalName = val.([]string)[0]
	}
	if val, ok := aduser["sAMAccountName"]; ok {
		aduserstruct.SAMAccountName = val.([]string)[0]
	}
	if val, ok := aduser["description"]; ok {
		aduserstruct.Description = val.([]string)[0]
	}
	if val, ok := aduser["info"]; ok {
		aduserstruct.Info = val.([]string)[0]
	}
	if val, ok := aduser["title"]; ok {
		aduserstruct.Title = val.([]string)[0]
	}
	if val, ok := aduser["o"]; ok {
		aduserstruct.O = val.([]string)[0]
	}
	if val, ok := aduser["company"]; ok {
		aduserstruct.Company = val.([]string)[0]
	}
	if val, ok := aduser["postOfficeBox"]; ok {
		aduserstruct.PostOfficeBox = val.([]string)[0]
	}
	if val, ok := aduser["physicalDeliveryOfficeName"]; ok {
		aduserstruct.PhysicalDeliveryOfficeName = val.([]string)[0]
	}
	if val, ok := aduser["streetAddress"]; ok {
		aduserstruct.StreetAddress = val.([]string)[0]
	}
	if val, ok := aduser["I"]; ok {
		aduserstruct.I = val.([]string)[0]
	}
	if val, ok := aduser["st"]; ok {
		aduserstruct.St = val.([]string)[0]
	}
	if val, ok := aduser["department"]; ok {
		aduserstruct.Department = val.([]string)[0]
	}
	if val, ok := aduser["mail"]; ok {
		aduserstruct.Mail = val.([]string)[0]
	}
	if val, ok := aduser["telephoneNumber"]; ok {
		aduserstruct.TelephoneNumber = val.([]string)[0]
	}
	if val, ok := aduser["pager"]; ok {
		aduserstruct.Pager = val.([]string)[0]
	}
	if val, ok := aduser["mobile"]; ok {
		aduserstruct.Mobile = val.([]string)[0]
	}
	if val, ok := aduser["facsimileTelephoneNumber"]; ok {
		aduserstruct.FacsimileTelephoneNumber = val.([]string)[0]
	}
	if val, ok := aduser["homePhone"]; ok {
		aduserstruct.HomePhone = val.([]string)[0]
	}
	if val, ok := aduser["ipPhone"]; ok {
		aduserstruct.IpPhone = val.([]string)[0]
	}
	if val, ok := aduser["postalCode"]; ok {
		aduserstruct.PostalCode = val.([]string)[0]
	}
	if val, ok := aduser["manager"]; ok {
		aduserstruct.Manager = val.([]string)[0]
	}
	if val, ok := aduser["wWWHomePage"]; ok {
		aduserstruct.WWWHomePage = val.([]string)[0]
	}
	if val, ok := aduser["username"]; ok {
		aduserstruct.Username = val.(string)
	}
	if val, ok := aduser["memberOf"]; ok {
		aduserstruct.MemberOf = val.([]string)
	}
	if val, ok := aduser["distinguishedName"]; ok {
		aduserstruct.DistinguishedName = val.(string)
	}
	if val, ok := aduser["countryCode"]; ok {

		code, _ := strconv.Atoi(val.([]string)[0])
		c, _ := codeToString(code)
		aduserstruct.C = c["string"]
		aduserstruct.Co = c["name"]
		aduserstruct.CountryCode = code

	} else if val, ok := aduser["c"]; ok {
		c := val.([]string)[0]
		aduserstruct.C = countries[c]["string"]
		aduserstruct.Co = countries[c]["name"]
		aduserstruct.CountryCode, _ = strconv.Atoi(countries[c]["code"])
	}
	return aduserstruct
}

func (aduserstruct *ADUser) ToMap() (aduser ADUSER) {
	aduser = ADUSER{
		"sn":                         aduserstruct.Sn,
		"givenName":                  aduserstruct.GivenName,
		"initials":                   aduserstruct.Initials,
		"accountname":                aduserstruct.Accountname,
		"userPrincipalName":          aduserstruct.UserPrincipalName,
		"sAMAccountName":             aduserstruct.SAMAccountName,
		"description":                aduserstruct.Description,
		"info":                       aduserstruct.Info,
		"title":                      aduserstruct.Title,
		"o":                          aduserstruct.O,
		"company":                    aduserstruct.Company,
		"postOfficeBox":              aduserstruct.PostOfficeBox,
		"physicalDeliveryOfficeName": aduserstruct.PhysicalDeliveryOfficeName,
		"streetAddress":              aduserstruct.StreetAddress,
		"I":                          aduserstruct.I,
		"st":                         aduserstruct.St,
		"department":                 aduserstruct.Department,
		"mail":                       aduserstruct.Mail,
		"telephoneNumber":            aduserstruct.TelephoneNumber,
		"pager":                      aduserstruct.Pager,
		"mobile":                     aduserstruct.Mobile,
		"facsimileTelephoneNumber":   aduserstruct.FacsimileTelephoneNumber,
		"homePhone":                  aduserstruct.HomePhone,
		"ipPhone":                    aduserstruct.IpPhone,
		"postalCode":                 aduserstruct.PostalCode,
		"manager":                    aduserstruct.Manager,
		"wWWHomePage":                aduserstruct.WWWHomePage,
		"memberOf":                   aduserstruct.MemberOf,
		"countryCode":                aduserstruct.CountryCode,
		"c":                          aduserstruct.C,
		"co":                         aduserstruct.Co,
		"username":                   aduserstruct.Username,
	}
	return aduser
}
