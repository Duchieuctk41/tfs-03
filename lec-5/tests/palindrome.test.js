const { test, expect } = require('@jest/globals')
const isPalindrome = require('../utils/palindrome')

test('test palindrome: case 1', () => {
  expect(isPalindrome('eye')).toBe(true)
})

test('test palindrome: case 2', () => {
  expect(isPalindrome('_eye')).toBe(true)
})

test('test palindrome: case 3', () => {
  expect(isPalindrome('race car')).toBe(true)
})

test('test palindrome: case 4', () => {
  expect(isPalindrome('not a palindrome')).toBe(false)
})

test('test palindrome: case 5', () => {
  expect(isPalindrome('A man, aplan, a canal. Panama')).toBe(true)
})

test('test palindrome: case 6', () => {
  expect(isPalindrome('My age is 0, 0 si ega ym.')).toBe(true)
})

test('test palindrome: case 7', () => {
  expect(isPalindrome('0_0 (: /-\ :) 0-0')).toBe(true)
})

test('test palindrome: case 8', () => {
  expect(isPalindrome('five|\_/|four')).toBe(false)
})

// npm test ispalindrome.test.js // test file
