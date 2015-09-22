# subtleist
> Anonymously remind one another of subtle-isms and other [Recurse Center social rules] in Slack.

## Usage
You can use `/socialrules` to either remind your current Slack channel about a
rule, or a specific team member privately.

    /socialrules [surprise|wellactually|backseat|subtle] [<@user>]

<img src="http://f.cl.ly/items/272F0h1h1l13241f0t12/subtleist.png" width="780">

## Motivation
I believe the [Recurse Center social rules] are fantastic guidelines, and
something most teams in the tech industry should adopt.

The social rules are intended to be _lightweight_, in that it's not a big deal
when people occasionally accidentally break them, and everyone just reminds
each other and moves on.

As noted in Allison Kaptur's [blog post], this can be notably more difficult
with the fourth social rule:

> Breaking the fourth social rule, like breaking any other social rule, is an
accident and a small thing. In theory, someone should be able to say “Hey, that
was subtly sexist,” get the response “Oops, sorry!” and move on just as easily
as if they’d well-actually'ed. In practice, people are less likely to point out
when this rule is broken, and more likely to be defensive if they were the
rule-breaker. We’d like to change this.

These social rules are deeply ingrained in the norms of the Recurse Center.
However, for other engineering teams trying to adopt them, it may be
intimidating for individuals to point out violations, especially at the
beginning.

The hope is that by creating a anonymous, no-fault way to remind one another of
the social rules, members of a team are more likely to say something, increasing
their adoption into the team culture.

[Recurse Center social rules]: https://www.recurse.com/manual#sub-sec-social-rules
[blog post]: https://www.recurse.com/blog/38-subtle-isms-at-hacker-school


## Setup
### Create an incoming webhook in Slack
https://my.slack.com/services/new/incoming-webhook/

You'll need the generated URL in the next step.

### Setup the bot

The bot will refuse to run unless a `SLACK_WEBHOOK_URL` environment variable
exists, which you should have generated in the previous step.

To make this easy, here is a one-step "Click to install" Heroku button that
handles all the installation and setup for you via the web, with free hosting!

[![Deploy](https://www.herokucdn.com/deploy/button.svg)](https://heroku.com/deploy)

### Create the slash command in Slack
[Create a new slash command](https://my.slack.com/services/new/slash-commands/),
and configure it with the following values:

- Command: `/socialrules`
- URL: `https://foo.bar/slack_hook` (replace foo.bar with your deploy)
- Method: `POST`
- Autocomplete help text
  - Description: `Remind someone about the social rules`
  - Usage hint: `[surprise|wellactually|backseat|subtle] [<destination>]`
