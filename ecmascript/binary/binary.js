export default class Binary {
  constructor(binary) {
    this.binary = binary;
  }

  toDecimal() {
    if (!this.isValid()) {
      return 0;
    }
    
    const digits = Array.from(this.binary).map((val) => parseInt(val));
    const reversedDigits = digits.reverse();
    
    return reversedDigits.reduce((sum, digit, index) => {
      return digit ? sum + Math.pow(2, index) : sum;
    }, 0);
  }

  isValid() {
    return /^[01]+$/.test(this.binary);
  }
}