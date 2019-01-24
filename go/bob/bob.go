package bob

import "regexp"

var shouting = regexp.MustCompile(`^[^a-z]+\!?$`)
var shoutingGibberish = regexp.MustCompile(`^[A-Z]+$`)
var question = regexp.MustCompile(`\?$`)
var forcefulQuestion = regexp.MustCompile(`^[A-Z\s]+\?$`)
var noLetters = regexp.MustCompile(`^[^a-zA-Z]+$`)
var silence = regexp.MustCompile(`^\s*$`)
var endWithWhitespace = regexp.MustCompile(`\s$`)
var questionEndWithWhitespace = regexp.MustCompile(`\?\s+$`)
var nonQuestionEndWithWhitespace = regexp.MustCompile(`[^\?]*\s$`)

const chillOut = `Whoa, chill out!`
const sure = `Sure.`
const whatever = `Whatever.`
const calmDown = `Calm down, I know what I'm doing!`
const beThatWay = `Fine. Be that way!`

func Hey(remark string) string {
	switch {
	case silence.MatchString(remark):
		return beThatWay
	case questionEndWithWhitespace.MatchString(remark):
		return sure
	case nonQuestionEndWithWhitespace.MatchString(remark):
		return whatever
	case endWithWhitespace.MatchString(remark):
		return sure
	case forcefulQuestion.MatchString(remark):
		return calmDown
	case question.MatchString(remark):
		return sure
	case noLetters.MatchString(remark):
		return whatever
	case shouting.MatchString(remark):
		return chillOut
	case shoutingGibberish.MatchString(remark):
		return chillOut
	default:
		return whatever
	}
}
