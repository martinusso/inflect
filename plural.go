// References:
// http://www.enchantedlearning.com/grammar/partsofspeech/nouns/plurals/

package inflect

import "strings"

var (
	irregular = map[string]string{
		// A few nouns have the same singular and plural forms
		"athletics":       "athletics",
		"barracks":        "barracks",
		"bellows":         "bellows",
		"billiards":       "billiards",
		"cattle":          "cattle",
		"customs":         "customs",
		"congratulations": "congratulations",
		"darts":           "darts",
		"dregs":           "dregs",
		"economics":       "economics",
		"linguistics":     "linguistics",
		"eyeglasses":      "eyeglasses",
		"gallows":         "gallows",
		"glasses":         "glasses",
		"goods":           "goods",
		"gymnastics":      "gymnastics",
		"headquarters":    "headquarters",
		"jeans":           "jeans",
		"mathematics":     "mathematics",
		"means":           "means",
		"measles":         "measles",
		"mumps":           "mumps",
		"news":            "news",
		"oats":            "oats",
		"outskirts":       "outskirts",
		"pants":           "pants",
		"pliers":          "pliers",
		"pajamas":         "pajamas",
		"savings":         "savings",
		"scissors":        "scissors",
		"series":          "series",
		"shears":          "shears",
		"shorts":          "shorts",
		"species":         "species",
		"spectacles":      "spectacles",
		"stairs":          "stairs",
		"statistics":      "statistics",
		"steps":           "steps",
		"thanks":          "thanks",
		"tongs":           "tongs",
		"tropics":         "tropics",
		"trousers":        "trousers",
		"tweezers":        "tweezers",
		"vespers":         "vespers",
		"wages":           "wages",
		"wits":            "wits",
		"words":           "words",

		// Singular and plural are the same
		"advice":      "advice",
		"aircraft":    "aircraft",
		"bison":       "bison",
		"corn":        "corn",
		"deer":        "deer",
		"equipment":   "equipment",
		"evidence":    "evidence",
		"fish":        "fish",
		"gold":        "gold",
		"information": "information",
		"jewelry":     "jewelry",
		"kin":         "kin",
		"legislation": "legislation",
		"luck":        "luck",
		"luggage":     "luggage",
		"moose":       "moose",
		"music":       "music",
		"offspring":   "offspring",
		"sheep":       "sheep",
		"silver":      "silver",
		"swine":       "swine",
		"trout":       "trout",
		"wheat":       "wheat",

		// Irregular Plurals of Nouns
		"child":  "children",
		"die":    "dice",
		"foot":   "feet",
		"goose":  "geese",
		"louse":  "lice",
		"man":    "men",
		"mouse":  "mice",
		"person": "people",
		"that":   "those",
		"this":   "these",
		"tooth":  "teeth",
		"woman":  "women",

		"ox": "oxen",

		// ending in f, fe
		"belief":       "beliefs",
		"chef":         "chefs",
		"cliff":        "cliffs",
		"dwarf":        "dwarfs",
		"grief":        "griefs",
		"gulf":         "gulfs",
		"handkerchief": "handkerchiefs",
		"kerchief":     "kerchiefs",
		"mischief":     "mischiefs",
		"muff":         "muffs",
		"oaf":          "oafs",
		"proof":        "proofs",
		"ref":          "refs",
		"roof":         "roofs",
		"safe":         "safes",
		"turf":         "turfs", // turves is acceptable

		// ending in o
		"albino":    "albinos",
		"armadillo": "armadillos",
		"auto":      "autos",
		"cameo":     "cameos",
		"cello":     "cellos",
		"combo":     "combos",
		"duo":       "duos",
		"ego":       "egos",
		"folio":     "folios",
		"halo":      "halos",
		"inferno":   "infernos",
		"kangaroo":  "kangaroos",
		"lasso":     "lassos",
		"memento":   "mementos",
		"memo":      "memos",
		"piano":     "pianos",
		"photo":     "photos",
		"portfolio": "portfolios",
		"pro":       "pros",
		"silo":      "silos",
		"solo":      "solos",
		"stereo":    "stereos",
		"studio":    "studios",
		"taco":      "tacos",
		"tattoo":    "tattoos",
		"tuxedo":    "tuxedos",
		"typo":      "typos",
		"veto":      "vetoes",
		"video":     "videos",
		"yo":        "yos",
		"zoo":       "zoos",

		"quiz": "quizzes",

		"alga":     "algae",
		"alumna":   "alumnae",
		"antenna":  "antennae",
		"larva":    "larvae",
		"nebula":   "nebulae",
		"pupa":     "pupae", //pupas is acceptable
		"vertebra": "vertebrae",
		"vita":     "vitae",

		"automaton":  "automata",
		"criterion":  "criteria",
		"phenomenon": "phenomena",
		"polyhedron": "polyhedra",

		"album":   "albums",
		"stadium": "stadiums",

		"appendix": "appendices", // appendixes is acceptable
		"vertex":   "vertices",
		"vortex":   "vortices",

		"soliloquy": "soliloquies",
		"money":     "monies",
		"graffito":  "graffiti",
		"trilby":    "trilbys",
		"numen":     "numina",
		"atman":     "atmas",
		"lowlife":   "lowlifes",
		"rom":       "roma",
		"carmen":    "carmina",

		"cactus":   "cacti", // cactuses is acceptable
		"focus":    "foci",
		"fungus":   "fungi",
		"nucleus":  "nuclei",
		"syllabus": "syllabi",

		"analysis":    "analyses",
		"axis":        "axes",
		"basis":       "bases",
		"crisis":      "crises",
		"diagnosis":   "diagnoses",
		"ellipsis":    "ellipses",
		"emphasis":    "emphases",
		"hypothesis":  "hypotheses",
		"neurosis":    "neuroses",
		"oasis":       "oases",
		"parenthesis": "parentheses",
		"paralysis":   "paralyses",
		"synopsis":    "synopses",
		"synthesis":   "syntheses",
		"thesis":      "theses",

		"datum": "data",

		"genus":    "genera",
		"mythos":   "mythoi",
		"testis":   "testes",
		"yes":      "yeses",
		"octopus":  "octopuses", // octopi is acceptable
		"genie":    "genies",
		"ganglion": "ganglia",
		"occiput":  "occipita",
		"corpus":   "corpora",
		"opus":     "opera",
		"penis":    "penes",
		"atlas":    "atlases",

		"CamelOctopus": "CamelOctopi",
		"ley":          "leyes",
	}
)

// Plural generates the plurals of nouns
func Plural(word string) string {
	if val, ok := irregular[word]; ok {
		return val
	}

	switch {
	case endWithConsonantPlusY(word):
		// Most nouns that end in a consonant and y, y becomes ies
		return word[:len(word)-1] + "ies"
	case endIn(word, "f"):
		// Most nouns that end in f ,	f becomes ves
		return word[:len(word)-1] + "ves"
	case endIn(word, "fe"):
		// Most nouns that end in fe,	fe becomes ves
		return word[:len(word)-2] + "ves"
	case endIn(word, "um"):
		// Most nouns that end in um. Change final 'um' to 'a'
		return word[:len(word)-2] + "a"
	case endIn(word, "on"):
		// Nouns ending in -on becoming -a
		return word[:len(word)-2] + "a"
	case endIn(word, "o"):
		// Ends with 'o'	Add 'es'
		return word + "es"
	case endIn(word, "ch", "sh", "s", "x", "z"):
		// Most nouns that end in ch, sh, s, x, or z	add es
		return word + "es"
	default:
		// Most nouns make their plurals by simply adding –s to the end
		// Most nouns that end in a vowel and y, add s
		return word + "s"
	}
}

func endIn(word string, s ...string) bool {
	for _, v := range s {
		if strings.HasSuffix(strings.ToUpper(word), strings.ToUpper(v)) {
			return true
		}
	}
	return false
}

func endWithConsonantPlusY(word string) bool {
	if !strings.HasSuffix(strings.ToUpper(word), "Y") {
		return false
	}
	return IsConsonant(word[len(word)-2 : len(word)-1])
}

func endWithVowelPlusY(word string) bool {
	if !strings.HasSuffix(strings.ToUpper(word), "Y") {
		return false
	}
	return IsVowel(word[len(word)-2 : len(word)-1])
}
