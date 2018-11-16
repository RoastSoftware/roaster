# Pull request formatting
+ The PR title is formatted as follows: `controller/static: frob the quux before blarfing`
  + The package name goes before the colon
  + The part after the colon uses the verb tense + phrase that completes the blank in,
    "This change modifies Roaster to ___________"
  + Lowercase verb after the colon
  + No trailing period
  + Keep the title as short as possible. ideally under 76 characters or shorter
+ No Markdown
+ The first PR comment (this one) is wrapped at 76 characters, unless it's
  really needed (ASCII art, table, or long link)
+ If there is a corresponding issue, add either `(fixes #1234)` or `(updates #1234)`
(the latter if this is not a complete fix) to this comment
