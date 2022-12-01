const fs = require('fs')
const arr = fs.readFileSync('./input.txt').toString().split('\n')

const getTopCals = (array) => {
    let result = 0
    let temp = 0
    array.forEach((value) => {
        if (value === '') {
            result = Math.max(result, temp)
            temp = 0
        } else {
            temp += +value
        }
    })
    return result
}

const answer = getTopCals(arr)

console.log(answer)
console.assert(answer === 71506, 'equals correct answer')
