function isValid(str: string): boolean {
  const parens: string[] = []
  const brackets: Record<string, string> = {
    ')': '(',
    '}': '{',
    ']': '[',
  }
  const opened = Object.values(brackets)
  const closed = Object.keys(brackets)

  for (const char of str) {
    const isOpening = opened.some(c => c === char)
    const isClosing = closed.some(c => c === char)

    if (isOpening) {
      parens.push(char)
    }

    if (isClosing) {
      return brackets[char] === parens.pop() // Matches Previous
    }
  }

  return parens.length === 0
}


[
  '({})',
  '(){}[]',
  '({[]})',
  '({)}',
].forEach(t => {
  console.log(isValid(t))
})