package gofakeit

const lowerStr = "abcdefghijklmnopqrstuvwxyz"
const upperStr = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const numericStr = "0123456789"
const specialStr = "@#$%&?|!(){}<>=*+-_:;,."
const specialSafeStr = "!@.-_*"
const spaceStr = " "
const allStr = lowerStr + upperStr + numericStr + specialStr + spaceStr
const vowels = "aeiou"
const hashtag = '#'
const questionmark = '?'
const dash = '-'
const base58 = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"
const minUint = 0
const maxUint = ^uint(0)
const minInt = -maxInt - 1
const maxInt = int(^uint(0) >> 1)
const is32bit = ^uint(0)>>32 == 0

func dataCheck(dataVal []string) bool { _ = "STUB: not implemented"; return false }

func getRandValue(f *Faker, dataVal []string) string { _ = "STUB: not implemented"; return "" }

func replaceWithNumbers(f *Faker, str string) string { _ = "STUB: not implemented"; return "" }

func replaceWithLetters(f *Faker, str string) string { _ = "STUB: not implemented"; return "" }

func replaceWithHexLetters(f *Faker, str string) string { _ = "STUB: not implemented"; return "" }

func randLetter(f *Faker) rune { _ = "STUB: not implemented"; return 0 }

func randCharacter(f *Faker, s string) string { _ = "STUB: not implemented"; return "" }

func randHexLetter(f *Faker) rune { _ = "STUB: not implemented"; return 0 }

func randDigit(f *Faker) rune { _ = "STUB: not implemented"; return 0 }

func randIntRange(f *Faker, min, max int) int { _ = "STUB: not implemented"; return 0 }

func randUintRange(f *Faker, min, max uint) uint { _ = "STUB: not implemented"; return 0 }

func toFixed(num float64, precision int) float64 { _ = "STUB: not implemented"; return 0 }

func equalSliceString(a, b []string) bool { _ = "STUB: not implemented"; return false }

func equalSliceInt(a, b []int) bool { _ = "STUB: not implemented"; return false }

func equalSliceInterface(a, b []any) bool { _ = "STUB: not implemented"; return false }

func stringInSlice(a string, list []string) bool { _ = "STUB: not implemented"; return false }

func anyToString(a any) string { _ = "STUB: not implemented"; return "" }

func title(s string) string { _ = "STUB: not implemented"; return "" }

func funcLookupSplit(str string) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func parseNameAndParamsFromTag(tag string) (string, string) {
	_ = "STUB: not implemented"
	return "", ""
}

func parseMapParams(info *Info, fParams string) (*MapParams, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func addSplitValsToMapParams(splitVals []string, info *Info, mapParams *MapParams) (*MapParams, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
