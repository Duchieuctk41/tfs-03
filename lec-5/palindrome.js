const isPalindrome = (strInput) => {
  const regex = /[\W_]/g

  strInput = strInput.toLowerCase().replace(regex, '')
  strReverse = strInput.split('').reverse('').join('')
  
  return strInput === strReverse
}

isPalindrome('1 ^^^0.1')