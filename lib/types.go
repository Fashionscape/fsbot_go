package lib

// Outfit hold information related to the outfits
type Outfit struct {
	Id        string
	Link      string
	Submitter string
	Tag       string
	Featured  bool
	Deleted   bool
	Date      string
}

type GuildConf struct {
	Id               string
	Prefix           string
	LogChannel       string
	DisabledModules  map[string]bool
	DisabledCommands map[string]bool
}

type UserBlacklist struct {
	Id          string
	Blacklisted bool
	Reason      string
}

type GuildBlacklist struct {
	Id          string
	Blacklisted bool
	Reason      string
}

type ColorRole struct {
}
