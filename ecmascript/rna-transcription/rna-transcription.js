class Transcriptor {
  toRna(dnaStrand) {
    const dnaArray = Array.from(dnaStrand);
    const rnaArray = dnaArray.map(this.toComplement);
    return rnaArray.join('');  
  }

  toComplement(nucleotide) {
    const complement = complements[nucleotide];
    if (complement === undefined) {
      throw new Error('Invalid input DNA.');
    }
    return complement;
  }
}

const complements = {
  C: 'G',
  G: 'C',
  A: 'U',
  T: 'A'
};

export default Transcriptor;
