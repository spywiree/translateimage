package main

import (
	"log"

	languagecodes "github.com/spywiree/langcodes"
)

func StringToLanguageCode(s string) (languagecodes.LanguageCode, bool) {
	// 假设有一个映射表将字符串映射到 LanguageCode
	languageCodeMap := map[string]languagecodes.LanguageCode{
		"DETECT_LANGUAGE":     languagecodes.DETECT_LANGUAGE,
		"AFRIKAANS":           languagecodes.AFRIKAANS,
		"ALBANIAN":            languagecodes.ALBANIAN,
		"AMHARIC":             languagecodes.AMHARIC,
		"ARABIC":              languagecodes.ARABIC,
		"ARMENIAN":            languagecodes.ARMENIAN,
		"ASSAMESE":            languagecodes.ASSAMESE,
		"AYMARA":              languagecodes.AYMARA,
		"AZERBAIJANI":         languagecodes.AZERBAIJANI,
		"BAMBARA":             languagecodes.BAMBARA,
		"BASQUE":              languagecodes.BASQUE,
		"BELARUSIAN":          languagecodes.BELARUSIAN,
		"BENGALI":             languagecodes.BENGALI,
		"BHOJPURI":            languagecodes.BHOJPURI,
		"BOSNIAN":             languagecodes.BOSNIAN,
		"BULGARIAN":           languagecodes.BULGARIAN,
		"CATALAN":             languagecodes.CATALAN,
		"CEBUANO":             languagecodes.CEBUANO,
		"CHICHEWA":            languagecodes.CHICHEWA,
		"CHINESE_SIMPLIFIED":  languagecodes.CHINESE_SIMPLIFIED,
		"CHINESE_TRADITIONAL": languagecodes.CHINESE_TRADITIONAL,
		"CORSICAN":            languagecodes.CORSICAN,
		"CROATIAN":            languagecodes.CROATIAN,
		"CZECH":               languagecodes.CZECH,
		"DANISH":              languagecodes.DANISH,
		"DHIVEHI":             languagecodes.DHIVEHI,
		"DOGRI":               languagecodes.DOGRI,
		"DUTCH":               languagecodes.DUTCH,
		"ENGLISH":             languagecodes.ENGLISH,
		"ESPERANTO":           languagecodes.ESPERANTO,
		"ESTONIAN":            languagecodes.ESTONIAN,
		"EWE":                 languagecodes.EWE,
		"FILIPINO":            languagecodes.FILIPINO,
		"FINNISH":             languagecodes.FINNISH,
		"FRENCH":              languagecodes.FRENCH,
		"FRISIAN":             languagecodes.FRISIAN,
		"GALICIAN":            languagecodes.GALICIAN,
		"GEORGIAN":            languagecodes.GEORGIAN,
		"GERMAN":              languagecodes.GERMAN,
		"GREEK":               languagecodes.GREEK,
		"GUARANI":             languagecodes.GUARANI,
		"GUJARATI":            languagecodes.GUJARATI,
		"HAITIAN_CREOLE":      languagecodes.HAITIAN_CREOLE,
		"HAUSA":               languagecodes.HAUSA,
		"HAWAIIAN":            languagecodes.HAWAIIAN,
		"HEBREW":              languagecodes.HEBREW,
		"HINDI":               languagecodes.HINDI,
		"HMONG":               languagecodes.HMONG,
		"HUNGARIAN":           languagecodes.HUNGARIAN,
		"ICELANDIC":           languagecodes.ICELANDIC,
		"IGBO":                languagecodes.IGBO,
		"ILOCANO":             languagecodes.ILOCANO,
		"INDONESIAN":          languagecodes.INDONESIAN,
		"IRISH":               languagecodes.IRISH,
		"ITALIAN":             languagecodes.ITALIAN,
		"JAPANESE":            languagecodes.JAPANESE,
		"JAVANESE":            languagecodes.JAVANESE,
		"KANNADA":             languagecodes.KANNADA,
		"KAZAKH":              languagecodes.KAZAKH,
		"KHMER":               languagecodes.KHMER,
		"KINYARWANDA":         languagecodes.KINYARWANDA,
		"KONKANI":             languagecodes.KONKANI,
		"KOREAN":              languagecodes.KOREAN,
		"KRIO":                languagecodes.KRIO,
		"KURDISH_KURMANJI":    languagecodes.KURDISH_KURMANJI,
		"KURDISH_SORANI":      languagecodes.KURDISH_SORANI,
		"KYRGYZ":              languagecodes.KYRGYZ,
		"LAO":                 languagecodes.LAO,
		"LATIN":               languagecodes.LATIN,
		"LATVIAN":             languagecodes.LATVIAN,
		"LINGALA":             languagecodes.LINGALA,

		// ... 其他语言代码映射
	}

	lc, ok := languageCodeMap[s]
	if !ok {
		// 如果没有找到映射，可以返回默认值或者错误
		log.Printf("Language code not found: %s", s)
		return lc, false
	}
	return lc, true
}
