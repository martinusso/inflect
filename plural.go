package inflect

import "strings"

var (
	// Irregular Plurals of Nouns
	irregulars = map[string]string{
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
	}

	// Singular and plural are the same
	unchanging = map[string]string{
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
	}

	// Only the plural exists
	onlyPluralForms = map[string]string{
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
	}

	exceptions = map[string]string{
		// ending in f, fe
		"belief":       "beliefs",
		"chef":         "chefs",
		"cliff":        "cliffs",
		"dwarf":        "dwarfs",
		"grief":        "griefs",
		"gulf":         "gulfs",
		"handkerchief": "handkerchiefs",
		"kerchief":     "kerchiefs",
		"lowlife":      "lowlifes",
		"mischief":     "mischiefs",
		"muff":         "muffs",
		"oaf":          "oafs",
		"proof":        "proofs",
		"ref":          "refs",
		"roof":         "roofs",
		"safe":         "safes",
		"turf":         "turfs", // turves is acceptable

		// ending in o +es
		"buffalo":  "buffaloes",
		"cargo":    "cargoes",
		"echo":     "echoes",
		"embargo":  "embargoes",
		"grotto":   "grottoes",
		"hero":     "heroes",
		"mosquito": "mosquitoes",
		"motto":    "mottoes",
		"potato":   "potatoes",
		"tomato":   "tomatoes",
		"torpedo":  "torpedoes",
		"veto":     "vetoes",
		"volcano":  "volcanoes",
		"zero":     "zeroes",

		// ending with Y
		"soliloquy": "soliloquies",
		"trilby":    "trilbys",
		"quiz":      "quizzes",
		"ley":       "leyes",

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

		// ending with 'um'
		"stadium": "stadiums",
		"album":   "albums",

		"ox":       "oxen",
		"appendix": "appendices", // appendixes is acceptable
		"vertex":   "vertices",
		"vortex":   "vortices",

		"graffito": "graffiti",
		"numen":    "numina",
		"atman":    "atmas",
		"rom":      "roma",
		"carmen":   "carmina",

		// Ends with 'us'
		"alumnus":      "alumni",
		"bacillus":     "bacilli",
		"cactus":       "cacti", // cactuses is acceptable
		"CamelOctopus": "CamelOctopi",
		"focus":        "foci",
		"fungus":       "fungi",
		"locus":        "loci",
		"nucleus":      "nuclei",
		"radius":       "radii",
		"stimulus":     "stimuli",
		"syllabus":     "syllabi",
		"terminus":     "termini",
		"torus":        "tori",

		"corpus": "corpora",
		"genus":  "genera",
		"opus":   "opera",

		"testis": "testes",
		"penis":  "penes",

		"mythos": "mythoi",

		"ganglion": "ganglia",
		"occiput":  "occipita",
	}
)

// Pluralize generates the plurals of nouns
func Pluralize(word string) string {
	if val, ok := irregulars[word]; ok {
		return val
	}
	if val, ok := onlyPluralForms[word]; ok {
		return val
	}
	if val, ok := unchanging[word]; ok {
		return val
	}
	if val, ok := exceptions[word]; ok {
		return val
	}

	switch {
	case endWithConsonantPlusY(word):
		return word[:len(word)-1] + "ies"
	case endIn(word, "f"):
		return word[:len(word)-1] + "ves"
	case endIn(word, "fe"):
		return word[:len(word)-2] + "ves"
	case endIn(word, "um"):
		return word[:len(word)-2] + "a"
	case endIn(word, "psis"):
		return word[:len(word)-4] + "pses"
	case endIn(word, "sis"):
		return word[:len(word)-3] + "ses"
	case endIn(word, "xis"):
		return word[:len(word)-3] + "xes"
	case endIn(word, "ch", "sh", "s", "x", "z"):
		return word + "es"
	default:
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
