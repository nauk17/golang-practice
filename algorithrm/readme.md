# Sum of all primes less than or equal N
### Mindset
1. Pseudo code
 ```
 loop i <= Sqrt N
    loop j = i*i, i*i + 1*i, i*i + 2*i +... util j < n
      sumNotPrimes += j
 resl = N * (N + 1) / 2 - sumNotPrimes
 ```
 
 2. Run
 ```go
 go run main.go -n=2000000
 ```
