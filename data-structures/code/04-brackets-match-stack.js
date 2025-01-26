function main(str){
  const stack = new Array()

  let isValid = true

  for (const char of str) {
    if (char === "(" || char === "[" || char === "{") {
      stack.push(char)
      // console.log('push', char)
      continue
    }

    if (stack.length === 0) {
      isValid = false
      break
    }

    const lastOpen = stack.pop()

    // console.log({ lastOpen, char })

    if (char === ")" && lastOpen !== "(") {
      isValid = false
      break
    }
    if (char === "]" && lastOpen !== "[") {
      isValid = false
      break
    }
    if (char === "}" && lastOpen !== "{") {
      isValid = false
      break
    }
  }

  return isValid;
}

const cases = [
  "[{}]",
  "(()())",
  "[}",
  "[()]))()",
  "[]{}[{}]",
]

cases.forEach((c) => {
  console.log(`${c} is valid? ${main(c)}`)
})
