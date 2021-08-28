let dataPush = {
  operators: [],
  numbers: []
}

let screening = {
  input: "0",
  output: "Result"
}
let screenInput, screenOutput
let isTrue = false, isClear = false, isComma = true

window.onload = function () {
  screenInput = document.getElementById("inputscr")
  screenOutput = document.getElementById("ouputscr")
};


function addOperators(operator) {
  if (isClear) { clearAll() }
  if (isTrue) {
    screening.input += " " + operator + " "
    screenInput.innerText = screening.input
  }
  isTrue = false
  isComma = true
}

function addNumbers(number) {
  if (isClear) { clearAll() }
  if (screening.input == "0") { screening.input = "" }
  screening.input += number
  screenInput.innerText = screening.input
  isTrue = true
}

function addComma() {
  if (isClear) { clearAll() }
  if (isComma) {
    screening.input += "."
    console.log(screening.input)
    screenInput.innerText = screening.input
    isComma = false
  }
}

function clearAll() {
  dataPush = {
    operators: [],
    numbers: []
  }
  screening = {
    input: "0",
    output: "Result"
  }
  isTrue = false
  isClear = false
  isComma = true
  screenInput.innerText = screening.input
  screenOutput.innerText = screening.output
}


function deleteInput() {
  console.log(9 * 0)
  if (screening.input.length === 1) { clearAll() }
  if (screening.input !== "0") {
    let lastElement = screening.input[screening.input.length - 1]
    if(lastElement === " ") {
      screening.input = screening.input.substr(0, screening.input.length - 2) 
      isTrue = true
    }
    screening.input = screening.input.substr(0, screening.input.length - 1)
    screenInput.innerText = screening.input
  }
}

function handleSubmit() {
  isClear = true
  isTrue = false
  let substrings = screening.input.split(" ");
  for (let i = 0; i < substrings.length; i++) {
    i % 2 === 1
      ? dataPush.operators.push(substrings[i])
      : dataPush.numbers.push(Number(substrings[i]))
  }
  console.log(dataPush)
  postData('http://localhost:3000/calc', dataPush)
    .then((response) => response.json())
    .then((data) => {
      screening.output = data.msg
      screenOutput.innerText = screening.output
      dataPush = {
        operators: [],
        numbers: []
      }
    })
}
async function postData(url = '', data = {}) {
  const response = await fetch(url, {
    method: 'POST',
    mode: 'cors',
    headers: {
      'Content-Type': 'application/json',
      'Accept': 'application/json'
    },
    body: JSON.stringify(data)
  });
  return response;
}
