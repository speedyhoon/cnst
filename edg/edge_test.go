package edg_test

import (
	"github.com/speedyhoon/cnst/edg"
	"strconv"
	"testing"
)

func TestUnique(t *testing.T) {
	list := map[int][]int{
		3:       {edg.Triangle},
		4:       {edg.Square, edg.Rectangle, edg.Quadrilateral, edg.Diamond},
		5:       {edg.Pentagon},
		6:       {edg.Hexagon, edg.Tristar, edg.Triquetra},
		7:       {edg.Heptagon},
		8:       {edg.Octagon, edg.Tetragram},
		9:       {edg.Nonagon},
		10:      {edg.Decagon, edg.Pentagram},
		11:      {edg.Hendecagon},
		12:      {edg.Dodecagon, edg.Hexagram},
		13:      {edg.Tridecagon},
		14:      {edg.Tetradecagon, edg.Heptagram},
		15:      {edg.Pentadecagon},
		16:      {edg.Hexadecagon, edg.Octagram},
		17:      {edg.Heptadecagon},
		18:      {edg.Octadecagon, edg.Enneagram},
		19:      {edg.Enneadecagon},
		20:      {edg.Icosagon, edg.Decagram},
		21:      {edg.Icosikaihenagon},
		22:      {edg.Icosikaidigon, edg.Hendecagram},
		23:      {edg.Icositrigon},
		24:      {edg.Icositetragon, edg.Dodecagram},
		25:      {edg.Icosikaipentagon},
		26:      {edg.Icosikaihexagon, edg.Tridecagram},
		27:      {edg.Icosikaiheptagon},
		28:      {edg.Icosikaioctagon, edg.Tetradecagram},
		29:      {edg.Icosikaienneagon},
		30:      {edg.Triacontagon, edg.Pentadecagram},
		31:      {edg.Triacontahenagon},
		32:      {edg.Triacontadigon, edg.Hexadecagram},
		33:      {edg.Triacontatrigon},
		34:      {edg.Triacontatetragon, edg.Heptadecagram},
		35:      {edg.Triacontapentagon},
		36:      {edg.Triacontahexagon, edg.Octadecagram},
		37:      {edg.Triacontaheptagon},
		38:      {edg.Triacontaoctagon, edg.Enneadecagram},
		39:      {edg.Triacontaenneagon},
		40:      {edg.Tetracontagon, edg.Icosagram},
		41:      {edg.Tetracontahenagon},
		42:      {edg.Tetracontadigon, edg.Icosikaihenagram},
		43:      {edg.Tetracontatrigon},
		44:      {edg.Tetracontatetragon, edg.Icosikaidigram},
		45:      {edg.Tetracontapentagon},
		46:      {edg.Tetracontahexagon, edg.Icositrigram},
		47:      {edg.Tetracontaheptagon},
		48:      {edg.Tetracontaoctagon, edg.Icositetragram},
		49:      {edg.Tetracontaenneagon},
		50:      {edg.Pentacontagon, edg.Icosikaipentagram},
		51:      {edg.Pentacontahenagon},
		52:      {edg.Pentacontadigon, edg.Icosikaihexagram},
		53:      {edg.Pentacontatrigon},
		54:      {edg.Pentacontatetragon, edg.Icosikaiheptagram},
		55:      {edg.Pentacontapentagon},
		56:      {edg.Pentacontahexagon, edg.Icosikaioctagram},
		57:      {edg.Pentacontaheptagon},
		58:      {edg.Pentacontaoctagon, edg.Icosikaienneagram},
		59:      {edg.Pentacontaenneagon},
		60:      {edg.Hexacontagon, edg.Triacontagram},
		61:      {edg.Hexacontahenagon},
		62:      {edg.Hexacontadigon},
		63:      {edg.Hexacontatrigon},
		64:      {edg.Hexacontatetragon},
		65:      {edg.Hexacontapentagon},
		66:      {edg.Hexacontahexagon},
		67:      {edg.Hexacontaheptagon},
		68:      {edg.Hexacontaoctagon},
		69:      {edg.Hexacontaenneagon},
		70:      {edg.Heptacontagon},
		71:      {edg.Heptacontahenagon},
		72:      {edg.Heptacontadigon},
		73:      {edg.Heptacontatrigon},
		74:      {edg.Heptacontatetragon},
		75:      {edg.Heptacontapentagon},
		76:      {edg.Heptacontahexagon},
		77:      {edg.Heptacontaheptagon},
		78:      {edg.Heptacontaoctagon},
		79:      {edg.Heptacontaenneagon},
		80:      {edg.Octacontagon, edg.Tetracontagram},
		81:      {edg.Octacontahenagon},
		82:      {edg.Octacontadigon},
		83:      {edg.Octacontatrigon},
		84:      {edg.Octacontatetragon},
		85:      {edg.Octacontapentagon},
		86:      {edg.Octacontahexagon},
		87:      {edg.Octacontaheptagon},
		88:      {edg.Octacontaoctagon},
		89:      {edg.Octacontaenneagon},
		90:      {edg.Enneacontagon},
		91:      {edg.Enneacontahenagon},
		92:      {edg.Enneacontadigon},
		93:      {edg.Enneacontatrigon},
		94:      {edg.Enneacontatetragon},
		95:      {edg.Enneacontapentagon},
		96:      {edg.Enneacontahexagon},
		97:      {edg.Enneacontaheptagon},
		98:      {edg.Enneacontaoctagon},
		99:      {edg.Enneacontaenneagon},
		100:     {edg.Hectogon, edg.Pentacontagram},
		101:     {edg.Hectahenagon},
		102:     {edg.Hectadigon},
		103:     {edg.Hectatrigon},
		104:     {edg.Hectatetragon},
		105:     {edg.Hectapentagon},
		106:     {edg.Hectahexagon},
		107:     {edg.Hectaheptagon},
		108:     {edg.Hectaoctagon},
		109:     {edg.Hectaenneagon},
		120:     {edg.Hexacontagram},
		140:     {edg.Heptacontagram},
		160:     {edg.Octacontagram},
		180:     {edg.Enneacontagram},
		200:     {edg.Dihectogon, edg.Hectogram},
		300:     {edg.Trihectogon},
		400:     {edg.Tetrahectogon, edg.Dihectogram},
		500:     {edg.Pentahectogon},
		600:     {edg.Hexahectogon, edg.Trihectogram},
		700:     {edg.Heptahectogon},
		800:     {edg.Octahectogon, edg.Tetrahectogram},
		900:     {edg.Enneahectogon},
		1000:    {edg.Chiliagon, edg.Pentahectogram},
		1200:    {edg.Hexahectogram},
		1400:    {edg.Heptahectogram},
		1600:    {edg.Octahectogram},
		1800:    {edg.Enneahectogram},
		2000:    {edg.Chiliagram},
		10000:   {edg.Myriagon},
		20000:   {edg.Myriagram},
		1000000: {edg.Megagon},
		2000000: {edg.Megagram},
	}

	t.Run("expected quantity", func(t *testing.T) {
		const expectedQty, expectedValues = 168, 129
		if qty := len(list); qty != expectedValues {
			t.Errorf("have %d map items, expected %d\n", qty, expectedValues)
		}

		var qty int
		for _, names := range list {
			qty += len(names)
		}
		if qty != expectedQty {
			t.Errorf("have %d total edges, expected %d\n", qty, expectedQty)
		}
	})

	for expected, names := range list {
		t.Run(strconv.Itoa(expected), func(t *testing.T) {
			for _, name := range names {
				if expected != name {
					t.Errorf("expected edg type `%d`, got: `%d`", name, expected)
				}
			}
		})
	}
}
