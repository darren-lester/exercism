class Year {
  constructor(year) {
    this.year = year; 
  }

  isLeap() {
    const divBy4 = this.year % 4 === 0;
    const divBy100 = this.year % 100 === 0;
    const divBy400 = this.year % 400 === 0;
    return divBy4 && (!divBy100 || divBy400);
  }
}

export default Year;
