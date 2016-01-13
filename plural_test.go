package inflect

import "testing"

func TestPluralAddingSToTheEnd(t *testing.T) {
	nouns := map[string]string{
		"book":       "books",
		"bonus":      "bonuses",
		"brother":    "brothers",
		"cat":        "cats",
		"cow":        "cows",
		"cup":        "cups",
		"journey":    "journeys",
		"sprout":     "sprouts",
		"bottle":     "bottles",
		"pencil":     "pencils",
		"desk":       "desks",
		"sticker":    "stickers",
		"window":     "windows",
		"kangaroo":   "kangaroos",
		"piano":      "pianos",
		"video":      "videos",
		"sabretooth": "sabretooths",
		"sabertooth": "sabertooths",
		"flatfoot":   "flatfoots",
		"tenderfoot": "tenderfoots",
		"talouse":    "talouses",
		"blouse":     "blouses",
		"mongoose":   "mongooses",
		"fuse":       "fuses",
	}

	for word, expected := range nouns {
		got := Pluralize(word)
		if got != expected {
			t.Errorf("Expected '%s' got '%s'", expected, got)
		}
	}
}

func TestPluralMostNounsThatEndIn_ch_sh_s_x_z(t *testing.T) {
	nouns := map[string]string{
		"annex":   "annexes",
		"complex": "complexes",
		"duplex":  "duplexes",
		"hex":     "hexes",
		"index":   "indexes", // indices is acceptable
		"prize":   "prizes",
		"watch":   "watches",
		"buss":    "busses",
		"arch":    "arches",
		"atlas":   "atlases",
		"ax":      "axes",
		"bash":    "bashes",
		"bench":   "benches",
		"bias":    "biases",
		"botch":   "botches",
		"box":     "boxes",
		"brush":   "brushes",
		"bunch":   "bunches",
		"bus":     "buses",
		"bush":    "bushes",
		"canvas":  "canvases",
		"catch":   "catches",
		"church":  "churches",
		"class":   "classes",
		"compass": "compasses",
		"crash":   "crashes",
		"cross":   "crosses",
		"dais":    "daises",
		"dish":    "dishes",
		"dress":   "dresses",
		"equinox": "equinoxes",
		"etch":    "etches",
		"fetch":   "fetches",
		"fix":     "fixes",
		"fox":     "foxes",
		"gas":     "gases",
		"grass":   "grasses",
		"itch":    "itches",
		"kiss":    "kisses",
		"larch":   "larches",
		"lash":    "lashes",
		"latch":   "latches",
		"mantis":  "mantises",
		"march":   "marches",
		"marsh":   "marshes",
		"mash":    "mashes",
		"mass":    "masses",
		"match":   "matches",
		"moss":    "mosses",
		"mix":     "mixes",
		"pass":    "passes",
		"patch":   "patches",
		"pox":     "poxes",
		"radish":  "radishes",
		"sash":    "sashes",
		"sketch":  "sketches",
		"starch":  "starches",
		"stitch":  "stitches",
		"tax":     "taxes",
		"touch":   "touches",
		"trash":   "trashes",
		"twitch":  "twitches",
		"vehicle": "vehicles",
		"wish":    "wishes",
		"witch":   "witches",
		"wrench":  "wrenches",
		"buzz":    "buzzes",
		"fizz":    "fizzes",
		"klutz":   "klutzes",
		"topaz":   "topazes",
		"waltz":   "waltzes",
	}
	for s, p := range nouns {
		got := Pluralize(s)
		if got != p {
			t.Errorf("Expected '%s' got '%s'", p, got)
		}
	}
}

func TestPluralEndingWithConsonantPlusY(t *testing.T) {
	nouns := map[string]string{
		"activity":   "activities",
		"ally":       "allies",
		"army":       "armies",
		"baby":       "babies",
		"beauty":     "beauties",
		"berry":      "berries",
		"cherry":     "cherries",
		"city":       "cities",
		"colony":     "colonies",
		"country":    "countries",
		"daisy":      "daisies",
		"dictionary": "dictionaries",
		"duty":       "duties",
		"enemy":      "enemies",
		"fairy":      "fairies",
		"family":     "families",
		"ferry":      "ferries",
		"fly":        "flies",
		"gallery":    "galleries",
		"history":    "histories",
		"injury":     "injuries",
		"jelly":      "jellies",
		"jerry":      "jerries",
		"kitty":      "kitties",
		"lady":       "ladies",
		"lily":       "lilies",
		"navy":       "navies",
		"mary":       "maries",
		"party":      "parties",
		"pony":       "ponies",
		"reply":      "replies",
		"romany":     "romanies",
		"secretary":  "secretaries",
		"sky":        "skies",
		"spy":        "spies",
		"story":      "stories",
		"study":      "studies",
		"symphony":   "symphonies",
		"theory":     "theories",
		"trophy":     "trophies",
		"try":        "tries",
		"university": "universities",
		"variety":    "varieties",
		"victory":    "victories",
	}

	for s, p := range nouns {
		got := Pluralize(s)
		if got != p {
			t.Errorf("Expected '%s' got '%s'", p, got)
		}
	}
}

func TestPluralEndingWithVowelPlusY(t *testing.T) {
	nouns := map[string]string{
		"alley":    "alleys",
		"attorney": "attorneys",
		"essay":    "essays",
		"boy":      "boys",
		"day":      "days",
		"delay":    "delays",
		"guy":      "guys",
		"jay":      "jays",
		"key":      "keys",
		"money":    "moneys",
		"osprey":   "ospreys",
		"play":     "plays",
		"ray":      "rays",
		"stray":    "strays",
		"toy":      "toys",
		"tray":     "trays",
		"turkey":   "turkeys",
		"valley":   "valleys",
		"way":      "ways",
	}
	for s, p := range nouns {
		got := Pluralize(s)
		if got != p {
			t.Errorf("Expected '%s' got '%s'", p, got)
		}
	}
}

func TestPluralEndingWithFeOrF(t *testing.T) {
	nouns := map[string]string{
		"beef":  "beeves", // beefs is acceptable
		"calf":  "calves",
		"elf":   "elves",
		"half":  "halves",
		"hoof":  "hooves", // hoofs is acceptable
		"knife": "knives",
		"leaf":  "leaves",
		"life":  "lives",
		"loaf":  "loaves",
		"scarf": "scarves",
		"self":  "selves",
		"shelf": "shelves",
		"thief": "thieves", // thiefs is acceptable
		"wife":  "wives",
		"wolf":  "wolves",
	}
	for s, p := range nouns {
		got := Pluralize(s)
		if got != p {
			t.Errorf("Expected '%s' got '%s'", p, got)
		}
	}
}

func TestPluralEndingWithO(t *testing.T) {
	nouns := map[string]string{
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
		"lasso":     "lassos",
		"memento":   "mementos",
		"memo":      "memos",
		"piano":     "pianos",
		"photo":     "photos",
		"portfolio": "portfolios",
		"pro":       "pros",
		"silo":      "silos",
		"solo":      "solos",
		"taco":      "tacos",
		"tuxedo":    "tuxedos",
		"typo":      "typos",
		"veto":      "vetoes",
		"yo":        "yos",
		"kangaroo":  "kangaroos",
		"stereo":    "stereos",
		"studio":    "studios",
		"tattoo":    "tattoos",
		"video":     "videos",
		"zoo":       "zoos",
	}
	for s, p := range nouns {
		got := Pluralize(s)
		if got != p {
			t.Errorf("Expected '%s' got '%s'", p, got)
		}
	}
}

func TestPluralEndingWithOn(t *testing.T) {
	nouns := map[string]string{
		"criterion":  "criteria",
		"phenomenon": "phenomena",
		"automaton":  "automata",
	}
	for s, p := range nouns {
		got := Pluralize(s)
		if got != p {
			t.Errorf("Expected '%s' got '%s'", p, got)
		}
	}
}

func TestPluralEndingWithUm(t *testing.T) {
	nouns := map[string]string{
		"bacterium":  "bacteria",
		"curriculum": "curricula",
		"datum":      "data",
		"erratum":    "errata",
		"gymnasium":  "gymnasia",
		"medium":     "media",
		"memorandum": "memoranda",
		"ovum":       "ova",
		"stratum":    "strata",
		"album":      "albums",
		"stadium":    "stadiums",
	}
	for s, p := range nouns {
		got := Pluralize(s)
		if got != p {
			t.Errorf("Expected '%s' got '%s'", p, got)
		}
	}
}

func TestPluralIrregularNouns(t *testing.T) {
	for s, p := range irregular {
		got := Pluralize(s)
		if got != p {
			t.Errorf("Expected '%s' got '%s'", p, got)
		}
	}
}
