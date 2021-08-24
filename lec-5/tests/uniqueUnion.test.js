const { test, expect } = require('@jest/globals')
const uniqueUnion = require('../utils/uniqueUnion')

test('test uniqueUnion: case 1', () => {
  expect(uniqueUnion ([1, 3, 2], [5, 2, 1, 4], [2, 1])).toEqual([1, 3, 2, 5, 4])
})

test('test uniqueUnion: case 1', () => {
  expect(uniqueUnion ([1, 2, 3], [5, 2, 1])).toEqual([1, 2, 3, 5])
})

test('test uniqueUnion: case 1', () => {
  expect(uniqueUnion ([1, 2, 3], [5, 2, 1, 4], [2, 1], [6, 7, 8])).toEqual([1, 2, 3, 5, 4, 6, 7, 8])
})
