var divideArray = function (nums, k) {
  nums.sort((a, b) => a - b);

  const results = [];

  let counter = 0;
  let row = [];

  for (let i = 0; i <= nums.length; i++) {
    console.log(i, results, row);
    if (counter === 0) {
      console.log(row);
      results.push(row);
      counter = 0;
      row = [];
    }

    if (counter < 3) {
      row[counter] = nums[i];
    }

    counter++;
  }

  return results;
};

console.log(divideArray([1, 3, 4, 8, 7, 9, 3, 5, 1]), 2);
