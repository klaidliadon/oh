// Released under an MIT-style license. See LICENSE.

// Generated by generate.oh

package boot

var Script string = `
define $connect: syntax (type out) as {
    set type: eval type
    syntax e (left right) as {
        define p: type
        spawn {
            eval: quasiquote: dynamic (unquote out) p
            e::eval left
            p::writer-close
        }
	block {
            dynamic $stdin = p
            e::eval right
            p::reader-close
	}
    }
}
define $redirect: syntax (chan mode mthd) as {
    syntax e (c cmd) as: make-env {
        define c: e::eval c
        define f = ()
        if (not: or (is-channel c) (is-pipe c)) {
            set f: open mode c
            set c = f
        }
        eval: quasiquote: dynamic (unquote chan) c
        e::eval cmd
        if (not: is-null f): eval: quasiquote: f::(unquote mthd)
    }
}
define ...: method (: args) as {
    cd $origin
    define path: car args
    if (eq 2: length args) {
        cd: car args
        set path: cadr args
    }
    while true {
        define abs: symbol: "/"::join $cwd path
        if (exists abs): return abs
	if (eq $cwd /): return path
        cd ..
    }
}
define and: syntax e (: lst) as {
    define r = false
    while (not: is-null: car lst) {
        set r: e::eval: car lst
        if (not r): return r
        set lst: cdr lst
    }
    return r
}
define append-stderr: $redirect $stderr "a" writer-close
define append-stdout: $redirect $stdout "a" writer-close
define apply: method (f: args) as: f @args
define backtick: syntax e (cmd) as {
    define p: pipe
    spawn {
        dynamic $stdout = p
        e::eval cmd
        p::writer-close
    }
    define r: cons () ()
    define c = r
    while (define l: p::readline) {
	set-cdr c: cons l ()
	set c: cdr c
    }
    p::reader-close
    return: cdr r
}
define caar: method (l) as: car: car l
define cadr: method (l) as: car: cdr l
define cdar: method (l) as: cdr: car l
define cddr: method (l) as: cdr: cdr l
define caaar: method (l) as: car: caar l
define caadr: method (l) as: car: cadr l
define cadar: method (l) as: car: cdar l
define caddr: method (l) as: car: cddr l
define cdaar: method (l) as: cdr: caar l
define cdadr: method (l) as: cdr: cadr l
define cddar: method (l) as: cdr: cdar l
define cdddr: method (l) as: cdr: cddr l
define caaaar: method (l) as: caar: caar l
define caaadr: method (l) as: caar: cadr l
define caadar: method (l) as: caar: cdar l
define caaddr: method (l) as: caar: cddr l
define cadaar: method (l) as: cadr: caar l
define cadadr: method (l) as: cadr: cadr l
define caddar: method (l) as: cadr: cdar l
define cadddr: method (l) as: cadr: cddr l
define cdaaar: method (l) as: cdar: caar l
define cdaadr: method (l) as: cdar: cadr l
define cdadar: method (l) as: cdar: cdar l
define cdaddr: method (l) as: cdar: cddr l
define cddaar: method (l) as: cddr: caar l
define cddadr: method (l) as: cddr: cadr l
define cdddar: method (l) as: cddr: cdar l
define cddddr: method (l) as: cddr: cddr l
define catch: syntax e (pattern handler) as {
	define entry: cons (symbol: e::eval pattern) (e::eval handler)
	dynamic $handlers: cons entry $handlers
}
define channel-stderr: $connect channel $stderr
define channel-stdout: $connect channel $stdout
define echo: builtin (: args) as {
    if (is-null args) {
        $stdout::write: symbol ""
    } else {
        $stdout::write @(for args symbol)
    }
}
define error: builtin (: args) as: $stderr::write @args
define for: method (l m) as {
    define r: cons () ()
    define c = r
    while (not: is-null l) {
        set-cdr c: cons (m: car l) ()
        set c: cdr c
        set l: cdr l
    }
    return: cdr r
}
define glob: builtin (: args) as: return args
define import: syntax e (name) as {
    set name: e::eval name
    define m: module name
    if (or (is-null m) (is-object m)) {
        return m
    }

    define l: list (quote source) name
    set l: cons (quote object): cons l ()
    set l: list (quote $root::define) m l
    e::eval l
}
define is-list: method (l) as {
    if (is-null l): return false
    if (not: is-cons l): return false
    if (is-null: cdr l): return true
    is-list: cdr l
}
define is-text: method (t) as: or (is-string t) (is-symbol t)
define list-ref: method (k x) as: car: list-tail k x
define list-tail: method (k x) as {
    if k {
        list-tail (sub k 1): cdr x
    } else {
        return x
    }
}
define object: syntax e (: body) as {
    e::eval: cons (quote block): append body (quote: context)
}
define or: syntax e (: lst) as {
    define r = false
    while (not: is-null: car lst) {
	set r: e::eval: car lst
        if r: return r
        set lst: cdr lst 
    }
    return r
}
define pipe-stderr: $connect pipe $stderr
define pipe-stdout: $connect pipe $stdout
define printf: method (f: args) as: echo: f::sprintf @args
define quasiquote: syntax e (cell) as {
    if (not: is-cons cell): return cell
    if (is-null cell): return cell
    if (eq (quote unquote): car cell): return: e::eval: cadr cell
    cons {
        e::eval: list (quote quasiquote): car cell
        e::eval: list (quote quasiquote): cdr cell
    }
}
define quote: syntax (cell) as: return cell
define read: builtin () as: $stdin::read
define readline: builtin () as: $stdin::readline
define redirect-stderr: $redirect $stderr "w" writer-close
define redirect-stdin: $redirect $stdin "r" reader-close
define redirect-stdout: $redirect $stdout "w" writer-close
define source: syntax e (name) as {
	define basename: e::eval name
	define paths = ()
	define name = basename

	if (has $OHPATH): set paths: (string $OHPATH)::split ":"
	while (and (not: is-null paths) (not: exists name)) {
		set name: "/"::join (car paths) basename
		set paths: cdr paths
	}

	if (not: exists name): set name = basename

        define r: cons () ()
        define c = r
	define f: open r- name
	while (define l: f::read) {
		set-cdr c: cons l ()
		set c: cdr c
	}
	set c: cdr r
	f::close
	define eval-list: method (rval first rest) as {
		if (is-null first): return rval
		eval-list (e::eval first) (car rest) (cdr rest)
	}
	eval-list (status 0) (car c) (cdr c)
}
define substitute-stdin: syntax () as {
    echo "process substitution not yet implemented."
}
define substitute-stdout: syntax () as {
    echo "process substitution not yet implemented."
}
define throw: method (ex: args) as {
	define handlers = $handlers
	while (not: is-null handlers) {
		define entry: car handlers
		define pattern: car entry
		define handler: cdr entry
		if (or (eq * pattern) (match pattern ex)) {
			if (handler ex @args) {
				return true
			}
		}
		set handlers: cdr handlers
	}
	return false
}
define write: method (: args) as: $stdout::write @args
dynamic $handlers = ()

catch *: method (ex: msg) as {
	error: interpolate "unhandled ${ex}: ${msg}"
	return true
}

exists ("/"::join $HOME .ohrc) && source ("/"::join $HOME .ohrc)

`

//go:generate ./generate.oh
//go:generate go fmt generated.go
