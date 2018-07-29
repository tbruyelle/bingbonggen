package item

import (
	"math/rand"
	"time"
)

type Item struct {
	Body    string
	Eyes    string
	Mouth   string
	Glasses string
	Hat     string
}

func Random() Item {
	rand.Seed(time.Now().UnixNano())
	return Item{
		Body:    Bodies[rand.Intn(len(Bodies))],
		Mouth:   Mouthes[rand.Intn(len(Mouthes))],
		Eyes:    Eyes[rand.Intn(len(Eyes))],
		Glasses: Glasses[rand.Intn(len(Glasses))],
		Hat:     Hats[rand.Intn(len(Hats))],
	}
}

var Bodies = []string{
	"body_base",
	"body_sick",
	"body_angry",
	"body_orange",
	"body_indian",
	"body_brown",
	"body_ertha",
}

var Eyes = []string{
	"eyes_base",
	"eyes_girl",
	"eyes_boy_tired",
	"eyes_girl_tired",
	"eyes_boy_angry",
	"eyes_girl_angry",
	"eyes_closed",
	"eyes_closed2",
	"eyes_happy",
	"eyes_cross",
	"eyes_diff",
	"eyes_diff_girl",
	"eyes_fronce",
	"eyes_tornado",
}

var Mouthes = []string{
	"mouth_base",
	"mouth_happy",
	"mouth_o",
	"mouth_smile",
	"mouth_pascontent",
	"mouth_miam",
	"mouth_slobber",
}

var Glasses = []string{
	"glasses_empty",
	"glasses_pink",
	"glasses_blue",
	"glasses_brown",
	"glasses_star",
	"glasses_heart",
}

var Hats = []string{
	"hat_empty",
	"hat_sheriff",
	"hat_chine",
	"hat_king",
}
