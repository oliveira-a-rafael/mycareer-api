package domains

var Positions = map[string]string{
	"GK":  "GOALKEEPER",
	"D":   "DEFENSE_MEN",
	"SW":  "SWEEPER",
	"RWB": "WINGBACK_RWB",
	"LWB": "WINGBACK_LWB",
	"RB":  "RIGHT_BACK_RB",
	"LB":  "RIGHT_BACK_LB",
	"CB":  "CENTRE_BACK",
	"DM":  "DEFENSIVE_MIDFIELDER",
	"RW":  "WING_RW",
	"LW":  "WING_LW",
	"LM":  "MIDFIELDERS_LM",
	"RM":  "MIDFIELDERS_RM",
	"CM":  "MIDFIELDERS_CM",
	"AM":  "ATTACKING_MIDFIELDERS",
	"CF":  "FORWARD_CENTRE",
	"RF":  "FORWARD_RIGHT",
	"LF":  "FORWARD_LEFT",
	"ST":  "STRIKERS",
}

type Position struct {
	Value       string `json:"value"`
	Description string `json:"description"`
}
