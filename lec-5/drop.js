const drop = (arr, func) => {
  for(i in arr) 
    if (func(arr[i]) === true )
      return arr.slice(i)
  return []
}

let result = drop([1,2,3,4,5], n => n >= 2)

console.log(result)
