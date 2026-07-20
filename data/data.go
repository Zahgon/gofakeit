package data

var Data = map[string]map[string][]string{
	"person":    Person,
	"auth":      Auth,
	"address":   Address,
	"airline":   Airline,
	"company":   Company,
	"job":       Job,
	"lorem":     Lorem,
	"language":  Languages,
	"internet":  Internet,
	"file":      Files,
	"color":     Colors,
	"computer":  Computer,
	"hipster":   Hipster,
	"beer":      Beer,
	"hacker":    Hacker,
	"animal":    Animal,
	"currency":  Currency,
	"log_level": LogLevels,
	"timezone":  TimeZone,
	"car":       Car,
	"emoji":     Emoji,
	"word":      Word,
	"text":      Text,
	"food":      Food,
	"minecraft": Minecraft,
	"celebrity": Celebrity,
	"error":     Error,
	"html":      Html,
	"book":      Books,
	"movie":     Movies,
	"school":    School,
	"song":      Songs,
	"product":   Product,
	"bank":      Bank,
}

func List() map[string][]string { _ = "STUB: not implemented"; return nil }

func Get(key string) map[string][]string { _ = "STUB: not implemented"; return nil }

func Set(key string, data map[string][]string) { _ = "STUB: not implemented"; return }

func Remove(key string) { _ = "STUB: not implemented"; return }

func GetSubData(key, subkey string) []string { _ = "STUB: not implemented"; return nil }

func SetSub(key, subkey string, data []string) { _ = "STUB: not implemented"; return }

func RemoveSub(key, subkey string) { _ = "STUB: not implemented"; return }
