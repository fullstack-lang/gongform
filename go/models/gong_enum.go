// generated code - do not edit
package models

// insertion point of enum utility functions
// Utility function for FormGroupName
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (formgroupname FormGroupName) ToString() (res string) {

	// migration of former implementation of enum
	switch formgroupname {
	// insertion code per enum code
	case FormGroupDefaultName:
		res = "Form"
	}
	return
}

func (formgroupname *FormGroupName) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Form":
		*formgroupname = FormGroupDefaultName
	default:
		return errUnkownEnum
	}
	return
}

func (formgroupname *FormGroupName) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "FormGroupDefaultName":
		*formgroupname = FormGroupDefaultName
	default:
		return errUnkownEnum
	}
	return
}

func (formgroupname *FormGroupName) ToCodeString() (res string) {

	switch *formgroupname {
	// insertion code per enum code
	case FormGroupDefaultName:
		res = "FormGroupDefaultName"
	}
	return
}

func (formgroupname FormGroupName) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "FormGroupDefaultName")

	return
}

func (formgroupname FormGroupName) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Form")

	return
}

// Utility function for InputTypeEnum
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (inputtypeenum InputTypeEnum) ToString() (res string) {

	// migration of former implementation of enum
	switch inputtypeenum {
	// insertion code per enum code
	case Text:
		res = "text"
	case Password:
		res = "password"
	case Number:
		res = "number"
	case Email:
		res = "email"
	case Tel:
		res = "tel"
	case Date:
		res = "date"
	case Datetime:
		res = "datetime-local"
	case Time:
		res = "time"
	case URL:
		res = "url"
	case Search:
		res = "search"
	case Range:
		res = "range"
	case Color:
		res = "color"
	case File:
		res = "file"
	case Hidden:
		res = "hidden"
	case Month:
		res = "month"
	case Week:
		res = "week"
	}
	return
}

func (inputtypeenum *InputTypeEnum) FromString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "text":
		*inputtypeenum = Text
	case "password":
		*inputtypeenum = Password
	case "number":
		*inputtypeenum = Number
	case "email":
		*inputtypeenum = Email
	case "tel":
		*inputtypeenum = Tel
	case "date":
		*inputtypeenum = Date
	case "datetime-local":
		*inputtypeenum = Datetime
	case "time":
		*inputtypeenum = Time
	case "url":
		*inputtypeenum = URL
	case "search":
		*inputtypeenum = Search
	case "range":
		*inputtypeenum = Range
	case "color":
		*inputtypeenum = Color
	case "file":
		*inputtypeenum = File
	case "hidden":
		*inputtypeenum = Hidden
	case "month":
		*inputtypeenum = Month
	case "week":
		*inputtypeenum = Week
	default:
		return errUnkownEnum
	}
	return
}

func (inputtypeenum *InputTypeEnum) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "Text":
		*inputtypeenum = Text
	case "Password":
		*inputtypeenum = Password
	case "Number":
		*inputtypeenum = Number
	case "Email":
		*inputtypeenum = Email
	case "Tel":
		*inputtypeenum = Tel
	case "Date":
		*inputtypeenum = Date
	case "Datetime":
		*inputtypeenum = Datetime
	case "Time":
		*inputtypeenum = Time
	case "URL":
		*inputtypeenum = URL
	case "Search":
		*inputtypeenum = Search
	case "Range":
		*inputtypeenum = Range
	case "Color":
		*inputtypeenum = Color
	case "File":
		*inputtypeenum = File
	case "Hidden":
		*inputtypeenum = Hidden
	case "Month":
		*inputtypeenum = Month
	case "Week":
		*inputtypeenum = Week
	default:
		return errUnkownEnum
	}
	return
}

func (inputtypeenum *InputTypeEnum) ToCodeString() (res string) {

	switch *inputtypeenum {
	// insertion code per enum code
	case Text:
		res = "Text"
	case Password:
		res = "Password"
	case Number:
		res = "Number"
	case Email:
		res = "Email"
	case Tel:
		res = "Tel"
	case Date:
		res = "Date"
	case Datetime:
		res = "Datetime"
	case Time:
		res = "Time"
	case URL:
		res = "URL"
	case Search:
		res = "Search"
	case Range:
		res = "Range"
	case Color:
		res = "Color"
	case File:
		res = "File"
	case Hidden:
		res = "Hidden"
	case Month:
		res = "Month"
	case Week:
		res = "Week"
	}
	return
}

func (inputtypeenum InputTypeEnum) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "Text")
	res = append(res, "Password")
	res = append(res, "Number")
	res = append(res, "Email")
	res = append(res, "Tel")
	res = append(res, "Date")
	res = append(res, "Datetime")
	res = append(res, "Time")
	res = append(res, "URL")
	res = append(res, "Search")
	res = append(res, "Range")
	res = append(res, "Color")
	res = append(res, "File")
	res = append(res, "Hidden")
	res = append(res, "Month")
	res = append(res, "Week")

	return
}

func (inputtypeenum InputTypeEnum) CodeValues() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "text")
	res = append(res, "password")
	res = append(res, "number")
	res = append(res, "email")
	res = append(res, "tel")
	res = append(res, "date")
	res = append(res, "datetime-local")
	res = append(res, "time")
	res = append(res, "url")
	res = append(res, "search")
	res = append(res, "range")
	res = append(res, "color")
	res = append(res, "file")
	res = append(res, "hidden")
	res = append(res, "month")
	res = append(res, "week")

	return
}

// end of insertion point for enum utility functions

type GongstructEnumStringField interface {
	string | FormGroupName | InputTypeEnum
	Codes() []string
	CodeValues() []string
}

type PointerToGongstructEnumStringField interface {
	*FormGroupName | *InputTypeEnum
	FromCodeString(input string) (err error)
}

type GongstructEnumIntField interface {
	int
	Codes() []string
	CodeValues() []int
}

type PointerToGongstructEnumIntField interface {
	
	FromCodeString(input string) (err error)
}

// Last line of the template
