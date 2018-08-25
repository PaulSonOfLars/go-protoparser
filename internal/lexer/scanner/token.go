package scanner

// Token represents a lexical token.
type Token int

// The result of Scan is one of these tokens.
const (
	// Special tokens
	TILLEGAL Token = iota
	TEOF

	// Identifiers
	TIDENT

	// Literals
	TINTLIT
	TFLOATLIT
	TBOOLLIT
	TSTRLIT

	// Comment
	TCOMMENT

	// Misc characters
	TSEMICOLON   // ;
	TCOLON       // :
	TEQUALS      // =
	TQUOTE       // " or '
	TLEFTPAREN   // (
	TRIGHTPAREN  // )
	TLEFTCURLY   // {
	TRIGHTCURLY  // }
	TLEFTSQUARE  // [
	TRIGHTSQUARE // ]
	TLESS        // <
	TGREATER     // >
	TCOMMA       // ,
	TDOT         // .

	// Keywords
	TSYNTAX
	TSERVICE
	TRPC
	TRETURNS
	TMESSAGE
	TIMPORT
	TPACKAGE
	TOPTION
	TREPEATED
	TWEAK
	TPUBLIC
	TONEOF
	TMAP
	TRESERVED
	TENUM
	TSTREAM
)

func asMiscToken(ch rune) Token {
	m := map[rune]Token{
		',': TCOMMA,
		'.': TDOT,
	}
	if t, ok := m[ch]; ok {
		return t
	}
	return TILLEGAL
}

func asKeywordToken(st string) Token {
	m := map[string]Token{
		"service": TSERVICE,
		"rpc":     TRPC,
	}

	if t, ok := m[st]; ok {
		return t
	}
	return TILLEGAL
}
