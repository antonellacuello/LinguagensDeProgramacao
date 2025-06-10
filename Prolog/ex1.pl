progenitor(x, y) :- mae(x, y).
progenitor(x, y) :- pai(x, y).

filho(Y, X) :- pai(X, Y).
filho(Y, X) :- mae(X, Y).

irmao(X, Y) :-
    pai(P, X), pai(P, Y),
    mae(M, X), mae(M, Y),
    X \= Y.

mae(mona, herb).
mae(mona, homer).
pai(abraham, herb).
pai(abraham, homer).

mae(jackie, marge).
mae(jackie, patty).
mae(jackie, selma).
pai(clancy, marge).
pai(clancy, patty).
pai(clancy, selma).

mae(marge, bart).
mae(marge, lisa).
mae(marge, maggie).
pai(homer, bart).
pai(homer, lisa).
pai(homer, maggie).

mae(selma, ling).
pai(ausente, ling).