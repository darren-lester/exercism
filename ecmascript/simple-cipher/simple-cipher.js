class Cipher {
  constructor(key) {
    if (key !== undefined) {
      this.validateKey(key);
    }
    this.key = key || this.generateKey();
  }

  validateKey(key) {
    if (!key.match(/^[a-z]+$/)) {
      throw new Error('Bad key');
    }
  }

  generateKey() {
    let key = '';
    for (let i = 0; i < 100; ++i) {
      key += this.generateKeyElement();
    }
    return key;
  }

  generateKeyElement() {
    const charCode = Math.floor(Cipher.minChar + Math.random() * this.calculateRange());
    return String.fromCharCode(charCode);
  }

  calculateRange() {
    return Cipher.maxChar - Cipher.minChar + 1;
  }

  encode(plainText) {
    return Array.prototype.map.call(plainText,
      this.rightShift.bind(this)).join('');
  }

  decode(encodedText) {
    return Array.prototype.map.call(encodedText,
      this.leftShift.bind(this)).join('');
  }

  rightShift(character, index) {
    return this.shift(character, index, 1);
  }

  leftShift(character, index) {
    return this.shift(character, index, -1);
  }

  shift(character, index, shiftDirection) {
    const shift = this.calculateShiftSize(index) * shiftDirection;
    const charCode = character.charCodeAt(0);
    const shiftedCharCode = charCode + shift;
    const normalisedShiftedCharCode = this.normaliseCharCode(shiftedCharCode);
    return String.fromCharCode(normalisedShiftedCharCode);
  }

  calculateShiftSize(index) {
    return this.calculateKeyChar(index).charCodeAt(0) - Cipher.minChar;
  }

  calculateKeyChar(index) {
    return this.key[index % this.key.length];
  }

  normaliseCharCode(charCode) {
    const range = this.calculateRange();
    if (charCode < Cipher.minChar) {
      charCode += range;
    } else if (charCode > Cipher.maxChar) {
      charCode -= range;
    }
    return charCode;
  }
}

Cipher.minChar = 'a'.charCodeAt(0);
Cipher.maxChar = 'z'.charCodeAt(0);

export default Cipher;
