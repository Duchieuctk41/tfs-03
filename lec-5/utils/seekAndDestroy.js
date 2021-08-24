const seekAndDestroy = (arr, ...args) => {
  let newArr = arr.filter(e => !args.includes(e))
  
  return newArr
}

// seekAndDestroy([1, 2, 3, 5, 1, 2, 3], 2, 3) // [1, 5, 1]

module.exports = seekAndDestroy
