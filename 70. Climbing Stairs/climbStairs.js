const climbStairs = function (n, memo = {}) {
  if (n in memo) return memo[n];
  if (n <= 3) return n;
  memo[n] = climbStairs(n - 1, memo) + climbStairs(n - 2, memo);
  return memo[n];
};

console.log(climbStairs(4));
console.log(climbStairs(50));
