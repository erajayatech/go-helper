package helper

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/dustin/go-humanize"
	"github.com/erajayatech/go-helper/constants"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
	"github.com/shyandsy/ShyGinErrors"
)

func ValidateStruct(payload interface{}, payloadMessageError map[string]string) (errMessage map[string]string) {
	ge := ShyGinErrors.NewShyGinErrors(payloadMessageError)
	err := CheckStruct(payload)
	if err != nil {
		errMessage = ge.ListAllErrors(payload, err)

		return
	}

	return
}

func ValidateStructWithError(payload interface{}, payloadMessageError map[string]string) (errMessage map[string]string, err error) {
	ge := ShyGinErrors.NewShyGinErrors(payloadMessageError)
	err = CheckStruct(payload)
	if err != nil {
		errMessage = ge.ListAllErrors(payload, err)

		return
	}

	return
}

func CheckStruct(payload interface{}) (err error) {
	v := validator.New()

	return v.Struct(payload)
}

func FormatRupiah(amount int) string {
	humanizeValue := humanize.Comma(int64(amount))
	stringValue := strings.Replace(humanizeValue, ",", ".", -1)
	return "Rp " + stringValue
}

func FormatGender(gender int) string {
	var Gender string
	if gender == 1 {
		Gender = "M"
	} else if gender == 2 {
		Gender = "F"
	} else {
		Gender = ""
	}
	return Gender
}

func MustGetEnv(key string) string {
	err := godotenv.Load()
	if err != nil {
		log.Warn("Cannot load file .env: ", err)
	}

	value := os.Getenv(key)
	if len(value) == 0 {
		return ""
	}
	return value
}

func FormatInfoText(actionName, orderNumber, status, updatedBy string) string {
	return fmt.Sprintf("%s #%s %s - updated_by: %s", actionName, orderNumber, status, updatedBy)
}

func ExpectedInt(v interface{}) int {
	var result int
	switch v.(type) {
	case int:
		result = v.(int)
	case int64:
		result = int(v.(int64))
	case float64:
		result = int(v.(float64))
	case string:
		result, _ = strconv.Atoi(v.(string))
	}
	return result
}

func ExpectedInt64(v interface{}) int64 {
	var result int64
	switch v.(type) {
	case int:
		result = int64(v.(int))
	case float64:
		result = int64(v.(float64))
	case string:
		resultInt, _ := strconv.Atoi(v.(string))
		result = int64(resultInt)
	}
	return result
}

func ExpectedString(v interface{}) string {
	var result string
	switch v.(type) {
	case int:
		result = strconv.Itoa(v.(int))
	case int64:
		result = strconv.Itoa(int(v.(int64)))
	case float64:
		result = strconv.Itoa(int(v.(float64)))
	case string:
		result, _ = v.(string)
	}
	return result
}

func FloatToString(f float64) string {
	s := fmt.Sprintf("%f", f)
	return s
}

func ValidateDateFormat(p string) (result string, err error) {
	// date format harus dd-mm-yyyy atau dd-mm-yyyy
	result = strings.ReplaceAll(p, "/", "-")
	d := strings.Split(result, "-")
	if len(d) != 3 {
		err = errors.New("use format dd-mm-yyyy or dd-mm-yyyy")
	}

	for i, k := range d {
		ki := ExpectedInt(k)
		if ki <= 0 {
			err = errors.New("date cant be zero")
		} else {
			if i == 0 && ki > 31 {
				err = errors.New("use format dd-mm-yyyy or dd-mm-yyyy")
			}

			if i == 1 && ki > 12 {
				err = errors.New("use format dd-mm-yyyy or dd-mm-yyyy")
			}

			if i == 2 && ki < 1000 {
				err = errors.New("use format dd-mm-yyyy or dd-mm-yyyy")
			}
		}
	}
	return
}

func ConvertIsoDateFormat(p string) (result string, err error) {
	result = strings.ReplaceAll(p, "/", "-")
	d := strings.Split(result, "-")
	if len(d) != 3 {
		err = errors.New("use format dd-mm-yyyy or dd-mm-yyyy")
	}

	result = d[2] + "-" + d[1] + "-" + d[0]
	return
}

func SanitizeSpecialChar(word string) string {
	space := regexp.MustCompile(`\s+`)
	re := strings.NewReplacer("/n", " ", "\n", " ")

	return space.ReplaceAllString(strings.TrimSpace(re.Replace(word)), " ")
}

func ContainsSliceString(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func CreateKeyValuePairs(m map[string]string) string {
	b := new(bytes.Buffer)
	for key, value := range m {
		fmt.Fprintf(b, "%s=\"%s\"\n", key, value)
	}
	return b.String()
}

func IsSourceJdsportKiosk(source string) bool {
	switch source {
	case
		constants.XSource_JDSport_Kiosk,
		constants.XSource_JDSport_Mkg_k1,
		constants.XSource_JDSport_Lmp_k1,
		constants.XSource_JDSport_Pim_k1,
		constants.XSource_JDSport_Snc_k1,
		constants.XSource_JDSport_Ctp_k1,
		constants.XSource_JDSport_Sms_k1,
		constants.XSource_JDSport_Lmk_k1,
		constants.XSource_JDSport_Mkg_k2,
		constants.XSource_JDSport_Lmp_k2,
		constants.XSource_JDSport_Pim_k2,
		constants.XSource_JDSport_Snc_k2,
		constants.XSource_JDSport_Ctp_k2,
		constants.XSource_JDSport_Sms_k2,
		constants.XSource_JDSport_Lmk_k2,
		constants.XSource_JDSport_Smb_k1,
		constants.XSource_JDSport_Smb_k2,
		constants.XSource_JDSport_Spm_k1,
		constants.XSource_JDSport_Spm_k2,
		constants.XSource_JDSport_Pms_k1,
		constants.XSource_JDSport_Pms_k2,
		constants.XSource_JDSport_Tps_k1,
		constants.XSource_JDSport_Tps_k2,
		constants.XSource_JDSport_Smd_k1,
		constants.XSource_JDSport_Smd_k2,
		constants.XSource_JDSport_Tmm_k1,
		constants.XSource_JDSport_Tmm_k2,
		constants.XSource_JDSport_Hos_k1,
		constants.XSource_JDSport_Hos_k2:
		return true
	}

	return false
}

func IsSourceJdsport(source string) bool {
	if source == constants.XSource_JDSport || IsSourceJdsportKiosk(source) {
		return true
	}

	return false
}

func IsSourcePaymentLink(source string) bool {
	if IsSourceJdsportKiosk(source) {
		return true
	}

	switch source {
	case
		constants.XSource_Telemarketing,
		constants.XSource_ReverseO2O:
		return true
	}

	return false
}

func IsValidSource(source string) bool {
	if IsSourceJdsport(source) {
		return true
	}

	switch source {
	case
		constants.XSourceRepair,
		constants.XSource_Eraspace,
		constants.XSource_IBox,
		constants.XSource_Repair,
		constants.XSource_TheFaceShop,
		constants.XSource_Telemarketing,
		constants.XSource_ImmersiveTech,
		constants.XSource_ReverseO2O,
		constants.XSource_Erafonenmore,
		constants.XSource_Enterprise,
		constants.XSource_IboxEdu,
		constants.XSource_IboxEnt,
		constants.XSource_UrbanRepublic:
		return true
	}

	return false
}

func SourceToStoreCodeCCM(source string) (storeCode string) {
	switch source {
	case constants.XSource_Enterprise:
		return "en"
	case
		"",
		constants.XSource_Telemarketing,
		constants.XSource_ReverseO2O,
		constants.XSource_Erafonenmore:
		return constants.XSource_Eraspace
	}

	return source
}
