package app

import (
	"regexp"
)

// NormalMatch matchs just a normal match
func NormalMatch(regx string, arg string) bool {
	matched, err := regexp.MatchString(regx, arg)
	if err != nil {
		return false
	}
	return matched
}

// Match matches a regex from begin of a string to the end
func Match(regx string, arg string) bool {
	matched, err := regexp.MatchString("^("+regx+")$", arg)
	if err != nil {
		return false
	}
	return matched
}

// MatchFullFlag matches full flags like --force and -f with as input "force" and "f"
func MatchFullFlag(full string, letter string, arg string) bool {
	return Match("--"+full+"|-\\w{0,}"+letter+"\\w{0,}", arg)
}

// FindMatch finds a specific match and returns a specific part of the match
func FindMatch(toMatch string, regx string, selecter int) string {
	re := regexp.MustCompile(regx)
	out := re.FindStringSubmatch(string(toMatch))
	if len(out) > selecter {
		return out[selecter]
	}
	return ""
}

// FindAllMatch find all matches
func FindAllMatch(toMatch string, regx string, selecter int) []string {
	re := regexp.MustCompile(regx)
	out := re.FindAllStringSubmatch(string(toMatch), -1)
	toReturn := []string{}
	for _, value := range out {
		toReturn = append(toReturn, value[selecter])
	}
	return toReturn
}

// MatchFlag just matches a flag like --yes
func MatchFlag(full string, arg string) bool {
	return Match("--"+full, arg)
}

// CleanupCli clansup the cli ouput,
// it removes unnecessary spaces and like breaks
func CleanupCli(input string) string {
	return Replace(input, " ", `(\s{2,}|\t{1,}|\n{1,})+`)
}

// Replace with a regex
func Replace(toReplace string, Replaceval string, regx string) string {
	re := regexp.MustCompile(regx)
	return re.ReplaceAllString(toReplace, Replaceval)
}

// CleanFindMatch cleans directly the output of FindMatch
func CleanFindMatch(toMatch string, regx string, selecter int) string {
	return CleanupCli(FindMatch(toMatch, regx, selecter))
}
