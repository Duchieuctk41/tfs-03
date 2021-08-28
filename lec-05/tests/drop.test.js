const { test, expect } = require('@jest/globals')
const drop = require('../utils/drop')

test('test drop: case 1', () => {
  expect(drop([1, 2, 3, 4], n => n >= 3)).toEqual([3, 4])
})

test('test drop: case 2', () => {
  expect(drop([0, 1, 0, 1], n => n === 1)).toEqual([1,0,1])
})

test('test drop: case 3', () => {
  expect(drop([1, 2, 3, 4], n => n >5)).toEqual([])
})

// npm test drop.test.js
