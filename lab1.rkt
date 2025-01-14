;task 1

(- (+ 7 (* 13 22)) (* (/ 51 64) (- 19 (/ 45 (+ 32 11)))))

(define PI 3.1415)

(define x (/ PI 2))  ; x = π/4
(+ (expt (sin x) 2) (expt (cos x) 2))
(= (sin (* 2 x)) (* 2 (sin x) (cos x)))  
(= (cos (* 2 x)) (- (expt (cos x) 2) (expt (sin x) 2)))

;task 2

(define (fact n)
  (if (= n 0)
      1
      (* n (fact (- n 1)))
      )
  )
; Test cases for fact
(display (fact 5))
(newline)
(display (fact 7))
(newline)

(define (power x y)
  (if (= y 0)
      1
      (* x (power x (- y 1)))))

(display (power 4 5)) 
(newline)
(display (power 3 6))
(newline)

;task 3

(define factSum
  (lambda (x y)
    (+ (fact x) (fact y))))
(display (factSum 2 3))
(newline)
(display (factSum 5 6))