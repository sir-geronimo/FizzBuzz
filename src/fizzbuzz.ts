function fizzbuzz(n: number): void {
  for (let i = 0; i <= n; i++) {
    let output = ''

    const isDivBy3 = i % 3 === 0
    const isDivBy5 = i % 5 === 0

    if (isDivBy3) output += 'Fizz'
    if (isDivBy5) output += 'Buzz'

    if (output === '') output = `${i}`

    console.log(output)
  }
}


const [num] = process.argv.slice(2)
fizzbuzz(Number(num ?? 100))