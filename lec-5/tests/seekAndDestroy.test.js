const { expect, test } = require('@jest/globals')
const seekAndDestroy = require('../utils/seekAndDestroy')

test('test seekAndDestroy: case 1', () => {
  expect(seekAndDestroy([1, 2, 3, 1, 2, 3], 2, 3)).toEqual([1, 1])
})

test('test seekAndDestroy: case 2', () => {
  expect(seekAndDestroy([1, 2, 3, 5, 1, 2, 3], 2, 3)).toEqual([1, 5, 1])
})

test('test seekAndDestroy: case 3', () => {
  expect(seekAndDestroy(["foo", "bar", 1], "foo", 1)).toEqual(["bar"])
})
