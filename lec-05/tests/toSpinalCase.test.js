const { test, expect } = require('@jest/globals')
const toSpinalCase = require('../utils/toSpinalCase')

test('test toSpinalCase: case 1', () => {
  expect(toSpinalCase('MyNameIsQuan')).toEqual('my-name-is-quan')
})

test('test toSpinalCase: case 2', () => {
  expect(toSpinalCase('my_name_is_Quan')).toEqual('my-name-is-quan')
})

test('test toSpinalCase: case 3', () => {
  expect(toSpinalCase('My Name Is Quan')).toEqual('my-name-is-quan')
})

test('test toSpinalCase: case 4', () => {
  expect(toSpinalCase('My Name Is-Quan')).toEqual('my-name-is-quan')
})
