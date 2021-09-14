let a = [3, 2, 1];

function getMuls(a) {
  const reducer = (pre, cur) => pre + cur;
  let total = a.reduce(reducer);

  a.forEach((v, i) => {
    a[i] = total / v;
  });
  return a;
}

console.log(getMuls(a));
