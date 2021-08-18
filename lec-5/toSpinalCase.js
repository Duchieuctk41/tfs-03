const toSpinalCase = (str) => {

  let regex1 = /(?!^)([A-Z])/g // (?!^) = không tính ký tự đầu dòng
  let regex2 = /([^a-zA-Z])+/g

  let spinalStr = str.replace(regex1, ' $1')
  .replace(regex2, '-').toLowerCase()

  console.log(spinalStr)
  return spinalStr
}

toSpinalCase('My   ^Name_IsHieu')