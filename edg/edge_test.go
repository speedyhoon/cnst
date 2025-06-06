package edg_test

import (
	"github.com/speedyhoon/cnst/edg"
	"strconv"
	"testing"
)

func TestUnique(t *testing.T) {
	list := map[int][]int{
		1:       {edg.Monogon},
		2:       {edg.Digon},
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
		62:      {edg.Hexacontadigon, edg.Triacontahenagram},
		63:      {edg.Hexacontatrigon},
		64:      {edg.Hexacontatetragon, edg.Triacontadigram},
		65:      {edg.Hexacontapentagon},
		66:      {edg.Hexacontahexagon, edg.Triacontatrigram},
		67:      {edg.Hexacontaheptagon},
		68:      {edg.Hexacontaoctagon, edg.Triacontatetragram},
		69:      {edg.Hexacontaenneagon},
		70:      {edg.Heptacontagon, edg.Triacontapentagram},
		71:      {edg.Heptacontahenagon},
		72:      {edg.Heptacontadigon, edg.Triacontahexagram},
		73:      {edg.Heptacontatrigon},
		74:      {edg.Heptacontatetragon, edg.Triacontaheptagram},
		75:      {edg.Heptacontapentagon},
		76:      {edg.Heptacontahexagon, edg.Triacontaoctagram},
		77:      {edg.Heptacontaheptagon},
		78:      {edg.Heptacontaoctagon, edg.Triacontaenneagram},
		79:      {edg.Heptacontaenneagon},
		80:      {edg.Octacontagon, edg.Tetracontagram},
		81:      {edg.Octacontahenagon},
		82:      {edg.Octacontadigon, edg.Tetracontahenagram},
		83:      {edg.Octacontatrigon},
		84:      {edg.Octacontatetragon, edg.Tetracontadigram},
		85:      {edg.Octacontapentagon},
		86:      {edg.Octacontahexagon, edg.Tetracontatrigram},
		87:      {edg.Octacontaheptagon},
		88:      {edg.Octacontaoctagon, edg.Tetracontatetragram},
		89:      {edg.Octacontaenneagon},
		90:      {edg.Enneacontagon, edg.Tetracontapentagram},
		91:      {edg.Enneacontahenagon},
		92:      {edg.Enneacontadigon, edg.Tetracontahexagram},
		93:      {edg.Enneacontatrigon},
		94:      {edg.Enneacontatetragon, edg.Tetracontaheptagram},
		95:      {edg.Enneacontapentagon},
		96:      {edg.Enneacontahexagon, edg.Tetracontaoctagram},
		97:      {edg.Enneacontaheptagon},
		98:      {edg.Enneacontaoctagon, edg.Tetracontaenneagram},
		99:      {edg.Enneacontaenneagon},
		100:     {edg.Hectogon, edg.Pentacontagram},
		101:     {edg.Hectahenagon},
		102:     {edg.Hectadigon, edg.Pentacontahenagram},
		103:     {edg.Hectatrigon},
		104:     {edg.Hectatetragon, edg.Pentacontadigram},
		105:     {edg.Hectapentagon},
		106:     {edg.Hectahexagon, edg.Pentacontatrigram},
		107:     {edg.Hectaheptagon},
		108:     {edg.Hectaoctagon, edg.Pentacontatetragram},
		109:     {edg.Hectaenneagon},
		110:     {edg.Pentacontapentagram},
		112:     {edg.Pentacontahexagram},
		114:     {edg.Pentacontaheptagram},
		116:     {edg.Pentacontaoctagram},
		118:     {edg.Pentacontaenneagram},
		120:     {edg.Hexacontagram},
		122:     {edg.Hexacontahenagram},
		124:     {edg.Hexacontadigram},
		126:     {edg.Hexacontatrigram},
		128:     {edg.Hexacontatetragram},
		130:     {edg.Hexacontapentagram},
		132:     {edg.Hexacontahexagram},
		134:     {edg.Hexacontaheptagram},
		136:     {edg.Hexacontaoctagram},
		138:     {edg.Hexacontaenneagram},
		140:     {edg.Heptacontagram},
		142:     {edg.Heptacontahenagram},
		144:     {edg.Heptacontadigram},
		146:     {edg.Heptacontatrigram},
		148:     {edg.Heptacontatetragram},
		150:     {edg.Heptacontapentagram},
		152:     {edg.Heptacontahexagram},
		154:     {edg.Heptacontaheptagram},
		156:     {edg.Heptacontaoctagram},
		158:     {edg.Heptacontaenneagram},
		160:     {edg.Octacontagram},
		162:     {edg.Octacontahenagram},
		164:     {edg.Octacontadigram},
		166:     {edg.Octacontatrigram},
		168:     {edg.Octacontatetragram},
		170:     {edg.Octacontapentagram},
		172:     {edg.Octacontahexagram},
		174:     {edg.Octacontaheptagram},
		176:     {edg.Octacontaoctagram},
		178:     {edg.Octacontaenneagram},
		180:     {edg.Enneacontagram},
		182:     {edg.Enneacontahenagram},
		184:     {edg.Enneacontadigram},
		186:     {edg.Enneacontatrigram},
		188:     {edg.Enneacontatetragram},
		190:     {edg.Enneacontapentagram},
		192:     {edg.Enneacontahexagram},
		194:     {edg.Enneacontaheptagram},
		196:     {edg.Enneacontaoctagram},
		198:     {edg.Enneacontaenneagram},
		200:     {edg.Dihectogon, edg.Hectogram},
		202:     {edg.Hectahenagram},
		204:     {edg.Hectadigram},
		206:     {edg.Hectatrigram},
		208:     {edg.Hectatetragram},
		210:     {edg.Hectapentagram},
		212:     {edg.Hectahexagram},
		214:     {edg.Hectaheptagram},
		216:     {edg.Hectaoctagram},
		218:     {edg.Hectaenneagram},
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
		const expectedQty, expectedValues = 242, 181
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
