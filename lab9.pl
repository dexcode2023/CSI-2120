% Dexter Day, Lab 9, 300192263
% Task 1

duplicateKth(List, K, Result) :-
    duplicateKth(List, K, K, Result).

duplicateKth([], _, _, []).

duplicateKth([H|T], K, 1, [H, H|RT]) :-
    duplicateKth(T, K, K, RT).

duplicateKth([H|T], K, C, [H|RT]) :-
    C > 1,
    C1 is C - 1,
    duplicateKth(T, K, C1, RT).


% Task 2

count([], _, 0).
count([H|T], N, C) :-
    H =:= N,
    count(T, N, C1),
    C is C1 + 1.

count([H|T], N, C) :-
    H =\= N,
    count(T, N, C).


% Task 3

removeRepetition([], []).
removeRepetition([H|T], [H|RT]) :-
    \+ member(H, T),
    removeRepetition(T, RT).
removeRepetition([H|T], RT) :-
    member(H, T),
    removeRepetition(T, RT).

run_tests :-
    % Task 1 Tests
    duplicateKth([a, b, c, d, e, f, g], 3, L1),
    writeln('Task 1 Test 1:'),
    writeln(L1),

    duplicateKth([], 2, L2),
    writeln('Task 1 Test 2:'),
    writeln(L2),

    duplicateKth([2, 4], 4, L3),
    writeln('Task 1 Test 3:'),
    writeln(L3),

    % Task 2 Test
    count([2, 5, 1, 2, 3, 4, 1, 6, 1], 1, C),
    writeln('Task 2 Test:'),
    writeln(C),

    % Task 3 Test
    removeRepetition([10, 20, 3, 21, 20, 10], L4),
    writeln('Task 3 Test:'),
    writeln(L4).


:- initialization(run_tests).
