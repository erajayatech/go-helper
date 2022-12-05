# go-helper
helper for go app

## Installation

```bash
go get github.com/erajayatech/go-helper
```

## Usage

### ValidateStruct
```go
import "github.com/erajayatech/go-helper"

// struct
type Request struct {
	Email  string `json:"email" example:"andrietrilaksono@gmail.com" validate:"required" msg:"error_invalid_email"`
	Name string `json:"name" example:"andrie" validate:"required" msg:"error_invalid_name"`
}

var RequestErrorMessage = map[string]string{
	"error_invalid_email":  "email is required",
	"error_invalid_name":  "name is required",
}

func main() {
    var (
        request Request
    )

    errorMessage := helper.ValidateStruct(request, RequestErrorMessage)

    a, _ := json.MarshalIndent(errorMessage, "", "\t")
	fmt.Print(string(a))
}
```

```bash
{
	"email": "email is required",
	"name": "name is required"
}
```

### FormatRupiah
```go
import "github.com/erajayatech/go-helper"

func main() {
	amountRupiah := helper.FormatRupiah(3000)

	fmt.Println(amountRupiah)
}
```

```bash
"Rp 3.000"
```

### FormatGender
```go
import "github.com/erajayatech/go-helper"

func main() {
	gender := helper.FormatGender(1)

	fmt.Println(gender)
}
```

```bash
"M"
```

### MustGetEnv
```go
import "github.com/erajayatech/go-helper"

func main() {
	env := helper.MustGetEnv("MODE")

	fmt.Println(env)
}
```

```bash
"local"
```

### FormatInfoText
```go
import "github.com/erajayatech/go-helper"

func main() {
	env := helper.FormatInfoText("update data","11000034","processing","system")

	fmt.Println(env)
}
```

```bash
"update data #11000034 processing - updated_by: system"
```

### ExpectedInt
```go
import "github.com/erajayatech/go-helper"

func main() {
	value := helper.ExpectedInt(6.5)

	fmt.Println(value)
}
```

```bash
6
```

### ExpectedInt64
```go
import "github.com/erajayatech/go-helper"

func main() {
	var amount int64 = 67

	value := helper.ExpectedInt64(amount)

	fmt.Println(value)
}
```

```bash
67
```

### ExpectedString
```go
import "github.com/erajayatech/go-helper"

func main() {
	value := helper.ExpectedString(67)

	fmt.Println(value)
}
```

```bash
"67"
```

### FloatToString
```go
import "github.com/erajayatech/go-helper"

func main() {
	value := helper.FloatToString(0.06)

	fmt.Println(value)
}
```

```bash
"0.060000"
```

### ValidateDateFormat
```go
import "github.com/erajayatech/go-helper"

func main() {
	date, err := helper.ValidateDateFormat("01-09-2010")

	fmt.Println(date)
}
```

```bash
"01-09-2010"
```

### ConvertIsoDateFormat
```go
import "github.com/erajayatech/go-helper"

func main() {
	date, err := helper.ConvertIsoDateFormat("2022/07/10")

	fmt.Println(date)
}
```

```bash
"10-07-2022"
```

### SanitizeSpecialChar
```go
import "github.com/erajayatech/go-helper"

func main() {
	data, err := helper.SanitizeSpecialChar("jalan\nnamajalan\n")

	fmt.Println(data)
}
```

```bash
"jalan namajalan"
```

### ContainsSliceString
```go
import "github.com/erajayatech/go-helper"

func main() {
	array := []string{"abc", "def"}
	data, err := helper.ContainsSliceString(array,"def")

	fmt.Println(data)
}
```

```bash
true
```

### CreateKeyValuePairs
```go
import "github.com/erajayatech/go-helper"

func main() {
	array := map[string]string{
				"name":  "name is required",
				"email": "email is required",
			}
	data, err := helper.CreateKeyValuePairs(array)

	fmt.Println(data)
}
```

```bash
name="name is required"
email="email is required"
```

## License
[MIT](https://choosealicense.com/licenses/mit/)