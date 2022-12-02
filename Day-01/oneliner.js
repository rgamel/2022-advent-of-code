const fs = require('fs')

const sum = (acc, val) => (acc += +val)
const desc = (a, b) => b - a

const topElf = (filename, n = 1) =>
    fs
        .readFileSync(filename)
        .toString()
        .split('\n\n')
        .map((elf) => elf.split('\n').reduce(sum, 0))
        .sort(desc)
        .slice(0, n)
        .reduce(sum, 0)

console.log('1a:', topElf('./input.txt', 1))
console.log('1b:', topElf('./input.txt', 3))
