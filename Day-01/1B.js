const fs = require('fs')
const arr = fs.readFileSync('./input.txt').toString().split('\n')

const addIfBigger = (arr, newValue) => {
    if (arr.length < 3) {
        return [...arr, newValue]
    }

    const newArr = [...arr].sort((a, b) => b - a)

    if (newArr[2] < newValue) {
        newArr.splice(2, 1, newValue)
    }
    return newArr
}

const getCalsForTop3 = (array) => {
    let top3 = []
    let temp = 0

    array.forEach((value) => {
        if (value === '') {
            top3 = addIfBigger(top3, temp)
            temp = 0
        } else {
            temp += +value
        }
    })

    return top3
}

const answer = getCalsForTop3(arr).reduce((acc, val) => (acc += val), 0)

console.log(answer)
console.assert(answer === 209603, 'equals correct answer')
