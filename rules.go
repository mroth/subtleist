package main // import "github.com/mroth/subtleist"

const manualURL = "https://www.recurse.com/manual"

// Rule represents a Recurse Center social rule
type Rule struct {
	title       string // title of the social rule
	description string // the full descriptive text
	id          string // section id in recurse center manual, e.g. #no-feigned-surprise
	trigger     string // trigger the rule if command contains this string
}

func (r Rule) uri() string {
	return manualURL + "/#" + r.id
}

var surpriseRule = Rule{
	title:       "No feigning surprise",
	description: `The first rule means you shouldn't act surprised when people say they don't know something. This applies to both technical things ("What?! I can't believe you don't know what the stack is!") and non-technical things ("You don't know who RMS is?!"). Feigning surprise has absolutely no social or educational benefit: When people feign surprise, it's usually to make them feel better about themselves and others feel worse. And even when that's not the intention, it's almost always the effect. As you've probably already guessed, this rule is tightly coupled to our belief in the importance of people feeling comfortable saying "I don't know" and "I don't understand."`,
	id:          "no-feigned-surprise",
	trigger:     "surprise",
}

var wellActuallyRule = Rule{
	title:       "No well-actually's",
	description: `A well-actually happens when someone says something that's almost—but not entirely—correct, and you say, "well, actually…" and then give a minor correction. This is especially annoying when the correction has no bearing on the actual conversation. This doesn't mean that we aren't about truth-seeking or that we don't care about being precise. Almost all well-actually's in our experience are about grandstanding, not truth-seeking.`,
	id:          "no-well-actuallys",
	trigger:     "actually",
}

var backseatRule = Rule{
	title:       "No back-seat driving",
	description: `If you overhear people working through a problem, you shouldn't intermittently lob advice across the room. This can lead to the "too many cooks" problem, but more important, it can be rude and disruptive to half-participate in a conversation. This isn't to say you shouldn't help, offer advice, or join conversations. On the contrary, we encourage all those things. Rather, it just means that when you want to help out or work with others, you should fully engage and not just butt in sporadically.`,
	id:          "no-backseat-driving",
	trigger:     "backseat",
}

var subtleRule = Rule{
	title:       "No subtle -isms",
	description: `Our last social rule bans subtle racism, sexism, homophobia, transphobia, and other kinds of bias. This one is different from the rest, because it covers a class of behaviors instead of one very specific pattern. Subtle -isms are small things that make others feel uncomfortable, things that we all sometimes do by mistake. For example, saying "It's so easy my grandmother could do it" is a subtle -ism. Like the other three social rules, this one is often accidentally broken. Like the other three, it's not a big deal to mess up – you just apologize and move on.`,
	id:          "no-subtle-isms",
	trigger:     "subtle",
}

var rules = []Rule{surpriseRule, wellActuallyRule, backseatRule, subtleRule}
