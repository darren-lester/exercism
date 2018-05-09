class Pangram {
  constructor(testString) {
    this.testString = testString;
  }

  isPangram() {
    if (!this.testString) {
      return false;
    }

    const preparedTestString = this.prepareTestString(this.testString);
    
    return Pangram.alphabet.every((character) => {
      return preparedTestString.includes(character);
    });
  }

  prepareTestString(testString) {
    return testString.toLowerCase();
  }
}

Pangram.alphabet = [
  'a', 'b', 'c', 'd',
  'e', 'f', 'g', 'h',
  'i', 'j', 'k', 'l',
  'm', 'n', 'o', 'p',
  'q', 'r', 's', 't',
  'u', 'v', 'w', 'x',
  'y', 'z'
];

export default Pangram;
