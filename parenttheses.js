function reverseInParentheses(inputString) {
    let arr = inputString
    let i = 0, start = 0, end = -1
    while (end < start && i < arr.length) {
        if (arr[i] === '(') {
            start = i
        }
        if (arr[i] === ')') {
            end = i
        }
        i++
    }
    let temp = arr.substring(start + 1, end)
    console.log(`temp:${temp} str:${inputString} start:${start} end:${end}`)
    console.log("-----------------------")
    if ( end !== -1) {
        let lStr = arr.substring(0, start);
        let rStr = arr.substring(end + 1);
        console.log(`arr.substring(0, start):${lStr}   arr.substring(end + 1):${rStr}`)
        console.log("--------------------")
        return reverseInParentheses(lStr + [...temp].reverse().join('') + rStr)
    }
    return arr
}

// console.log(reverseInParentheses('(bar)'))
// console.log(reverseInParentheses('foo(bar)baz'))
// console.log(reverseInParentheses('foo(bar(baz))blim'))
// console.log(reverseInParentheses('((see)(you))'))
console.log(reverseInParentheses('(abc)d(efg)'))

