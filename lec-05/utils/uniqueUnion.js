const uniqueUnion = (...arrs) => {
  let newArr = [].concat(...arrs)
  uniArr = [...new Set(newArr)]

  return uniArr
}

// uniqueUnion([1,2,3], [2,3,4])

module.exports = uniqueUnion
