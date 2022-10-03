const runningSum = function (nums) {

  let sum = 0;

  for (let i = 0; i < nums.length; i++) {
    sum = sum + nums[i]
    nums[i] = sum
  }

  return nums
};

console.log(runningSum([1, 2, 3, 4]))
console.log(runningSum([1, 1, 1, 1, 1]))
console.log(runningSum([3, 1, 2, 10, 1]))