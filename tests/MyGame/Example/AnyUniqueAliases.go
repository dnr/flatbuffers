// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package Example

import "strconv"

type AnyUniqueAliases byte

const (
	AnyUniqueAliasesNONE AnyUniqueAliases = 0
	AnyUniqueAliasesM    AnyUniqueAliases = 1
	AnyUniqueAliasesT    AnyUniqueAliases = 2
	AnyUniqueAliasesM2   AnyUniqueAliases = 3
)

var EnumNamesAnyUniqueAliases = map[AnyUniqueAliases]string{
	AnyUniqueAliasesNONE: "NONE",
	AnyUniqueAliasesM:    "M",
	AnyUniqueAliasesT:    "T",
	AnyUniqueAliasesM2:   "M2",
}

var EnumValuesAnyUniqueAliases = map[string]AnyUniqueAliases{
	"NONE": AnyUniqueAliasesNONE,
	"M":    AnyUniqueAliasesM,
	"T":    AnyUniqueAliasesT,
	"M2":   AnyUniqueAliasesM2,
}

func (v AnyUniqueAliases) String() string {
	if s, ok := EnumNamesAnyUniqueAliases[v]; ok {
		return s
	}
	return "AnyUniqueAliases(" + strconv.FormatInt(int64(v), 10) + ")"
}
