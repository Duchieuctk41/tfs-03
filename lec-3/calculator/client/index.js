let dataPush = {
  operators: [],
  numbers: []
}


function addOperators(operator) {
  dataPush.operators.push(operator)
  console.log(dataPush.operators)
}

function addNumbers(number) {
  dataPush.numbers.push(Number(number))
  console.log(dataPush)
}

function handleSubmit() {
  console.log(dataPush)
  postData('http://localhost:3000/calc', dataPush)
  .then((response) => response.json())
  .then((data) => {console.log(data)});
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
