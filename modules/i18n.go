package modules

import (
	E "github.com/yuw-mvc/yuw/exceptions"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"io/ioutil"
	"strings"
)

const (
	dirLanguage string = "./resources/lang/"
)

var (
	Translated *i18n
	translations map[language.Tag]*message.Printer
)

type i18n struct {
	util *Utils
}

func NewI18N() *i18n {
	return &i18n{
		util: NewUtils(),
	}
}

func (translate *i18n) Loading() (err error) {
	dir := I.Get("Languages.Dir", dirLanguage).(string)

	f, err := ioutil.ReadDir(dir)
	if err != nil {
		return
	}

	translations = make(map[language.Tag]*message.Printer, 0)
	for _, val := range f {
		if val.IsDir() {
			continue
		}

		fn := strings.Split(val.Name(), ".")
		ln := fn[0]

		if fn[1] != "json" {
			continue
		}

		byteJson, err := ioutil.ReadFile(dir + val.Name())
		if err != nil {
			return err
		}

		for key, translated := range translate.util.JsonToMap(byteJson) {
			message.SetString(language.MustParse(ln), key, translated.(string))
		}

		Tag, err := translate.LanguageFormat(ln)
		if err != nil {
			return err
		}

		translations[Tag] = message.NewPrinter(Tag)
	}

	return
}

func (translate *i18n) LanguageFormat(ln string) (tag language.Tag, err error) {
	if ok, _ := translate.util.StrContains(ln, "cn"); ok {
		tag = language.Chinese
		return
	}

	languages := I.Get("Languages.Lns", []interface{}{}).([]interface{})
	if ok, _ := translate.util.StrContains(ln, languages ...); ok == false {
		return language.English, E.Err("yuw^m_in_a", E.ErrPosition())
	}

	tag = language.MustParse(ln)
	return
}


