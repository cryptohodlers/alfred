package assets

import (
	"fmt"
	"strings"

	"github.com/stellar/go/build"
	"github.com/stellar/go/clients/horizon"
)

type Asset struct {
	BuilderAsset build.Asset
	Domain       string
	Instructions string
	Type         string
}

var (
	Assets = []Asset{
		{
			BuilderAsset: build.NativeAsset(),
			Domain:       "stellar.org",
		},
		{
			BuilderAsset: build.CreditAsset(
				"MOBI",
				"GA6HCMBLTZS5VYYBCATRBRZ3BZJMAFUDKYYF6AH6MVCMGWMRDNSWJPIH",
			),
			Domain: "mobius.network",
			Type:   "token",
		},
		{
			BuilderAsset: build.CreditAsset(
				"SLT",
				"GCKA6K5PCQ6PNF5RQBF7PQDJWRHO6UOGFMRLK3DYHDOI244V47XKQ4GP",
			),
			Domain:       "smartlands.io",
			Instructions: "https://smartlands.io",
			Type:         "token",
		},
		{
			BuilderAsset: build.CreditAsset(
				"CNY",
				"GAREELUB43IRHWEASCFBLKHURCGMHE5IF6XSE7EXDLACYHGRHM43RFOX",
			),
			Domain:       "ripplefox.com",
			Instructions: "Leave your address in the message to seller when you order theIte:https://shop109149722.taobao.com",
		},
		{
			BuilderAsset: build.CreditAsset(
				"RMT",
				"GCVWTTPADC5YB5AYDKJCTUYSCJ7RKPGE4HT75NIZOUM4L7VRTS5EKLFN",
			),
			Domain:       "sureremit.co",
			Instructions: "https://sureremit.co",
			Type:         "token",
		},
		{
			BuilderAsset: build.CreditAsset(
				"KIN",
				"GBDEVU63Y6NTHJQQZIKVTC23NWLQVP3WJ2RI2OTSJTNYOIGICST6DUXR",
			),
			Domain:       "apay.io",
			Instructions: "https://apay.io/",
		},
		{
			BuilderAsset: build.CreditAsset(
				"XCN",
				"GCNY5OXYSY4FKHOPT2SPOQZAOEIGXB5LBYW3HVU3OWSTQITS65M5RCNY",
			),
			Domain:       "fchain.io",
			Instructions: "fchain.io",
		},
		{
			BuilderAsset: build.CreditAsset(
				"USD",
				"GBSTRUSD7IRX73RQZBL3RQUH6KS3O4NYFY3QCALDLZD77XMZOPWAVTUK",
			),
			Domain:       "stronghold.co",
			Instructions: "USD",
		},
		{
			BuilderAsset: build.CreditAsset(
				"USD",
				"GDSRCV5VTM3U7Y3L6DFRP3PEGBNQMGOWSRTGSBWX6Z3H6C7JHRI4XFJP",
			),
			Domain:       "x.token.io",
			Instructions: "",
		},
		{
			BuilderAsset: build.CreditAsset(
				"WSD",
				"GDSVWEA7XV6M5XNLODVTPCGMAJTNBLZBXOFNQD3BNPNYALEYBNT6CE2V",
			),
			Domain:       "thewwallet.com",
			Instructions: "thewwallet.com",
		},
		{
			BuilderAsset: build.CreditAsset(
				"HKDT",
				"GABSZVZBYEO5F4V5LZKV7GR4SAJ5IKJGGOF43BIN42FNDUG7QPH6IMRQ",
			),
			Domain:       "cryptomover.com",
			Instructions: "cryptomover.com",
		},
		{
			BuilderAsset: build.CreditAsset(
				"EURT",
				"GAP5LETOV6YIE62YAM56STDANPRDO7ZFDBGSNHJQIYGGKSMOZAHOOS2S",
			),
			Domain:       "tempo.eu.com",
			Instructions: "tempo.eu.com",
		},
		{
			BuilderAsset: build.CreditAsset(
				"BTC",
				"GBSTRH4QOTWNSVA6E4HFERETX4ZLSR3CIUBLK7AXYII277PFJC4BBYOG",
			),
			Domain:       "stronghold.co",
			Instructions: "stronghold.co",
		},
		{
			BuilderAsset: build.CreditAsset(
				"BTC",
				"GATEMHCCKCY67ZUCKTROYN24ZYT5GK4EQZ65JJLDHKHRUZI3EUEKMTCH",
			),
			Domain:       "naobtc.com",
			Instructions: "naobtc.com",
		},
		{
			BuilderAsset: build.CreditAsset(
				"BTC",
				"GAUTUYY2THLF7SGITDFMXJVYH3LHDSMGEAKSBU267M2K7A3W543CKUEF",
			),
			Domain:       "apay.io",
			Instructions: "https://apay.io/",
		},
		{
			BuilderAsset: build.CreditAsset(
				"ETH",
				"GBSTRH4QOTWNSVA6E4HFERETX4ZLSR3CIUBLK7AXYII277PFJC4BBYOG",
			),
			Domain:       "stronghold.co",
			Instructions: "stronghold.co",
		},
		{
			BuilderAsset: build.CreditAsset(
				"ETH",
				"GBETHKBL5TCUTQ3JPDIYOZ5RDARTMHMEKIO2QZQ7IOZ4YC5XV3C2IKYU",
			),
			Domain:       "fchain.io",
			Instructions: "fchain.io",
		},
		{
			BuilderAsset: build.CreditAsset(
				"XRP",
				"GBVOL67TMUQBGL4TZYNMY3ZQ5WGQYFPFD5VJRWXR72VA33VFNL225PL5",
			),
			Domain:       "stellarport.io",
			Instructions: "stellarport.io",
		},
		{
			BuilderAsset: build.CreditAsset(
				"SPQR",
				"GDGBLAE6EIE45E7WK2H7L2YBFPJBGOMYAM3IVJ4OKFQTZ643NJTHSWJS",
			),
			Domain:       "spqr.co",
			Instructions: "https://spqr.co",
			Type:         "token",
		},
	}
)

// indexes
var (
	CodeToAsset = map[string][]Asset{}
)

func init() {
	for _, a := range Assets {
		code := a.BuilderAsset.Code
		if a.BuilderAsset.Native {
			code = "XLM"
		}
		CodeToAsset[code] = append(CodeToAsset[code], a)
	}
}

func GetAssets(code string) []Asset {
	return CodeToAsset[strings.ToUpper(code)]
}

func GetByCodeIssuer(code, issuer string) *Asset {
	for _, a := range CodeToAsset[code] {
		if a.BuilderAsset.Issuer == issuer {
			return &a
		}
	}

	return nil
}

func (a Asset) ToHorizonAsset() horizon.Asset {
	if a.BuilderAsset.Native {
		return horizon.Asset{
			Type: "native",
		}
	}

	typ := "credit_alphanum4"
	if len(a.BuilderAsset.Code) > 4 {
		typ = "credit_alphanum12"
	}

	return horizon.Asset{
		Type:   typ,
		Code:   a.BuilderAsset.Code,
		Issuer: a.BuilderAsset.Issuer,
	}
}

func (a Asset) String() string {
	if a.BuilderAsset.Native {
		return "XLM stellar.org"
	}

	return fmt.Sprintf("%s %s (%s)",
		a.BuilderAsset.Code, a.Domain,
		a.BuilderAsset.Issuer)
}

func (a Asset) CodeString() string {
	if a.BuilderAsset.Native {
		return "XLM"
	}

	return a.BuilderAsset.Code
}
