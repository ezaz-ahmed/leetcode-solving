

const pivotIndex = function (nums) {

  const sum = nums.reduce((a, b) => a + b)
  let leftSum = 0

  for (let i = 0; i < nums.length; i++) {


    if (leftSum == (sum - leftSum - nums[i])) {
      return i
    }

    leftSum += nums[i]
  }

  return -1
};

console.log(pivotIndex([1, 7, 3, 6, 5, 6]))
console.log(pivotIndex([1, 2, 3]))
console.log(pivotIndex([2, 1, -1]))