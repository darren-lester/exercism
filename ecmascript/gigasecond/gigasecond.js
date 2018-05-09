const gigasecondInMilliseconds = 1000000000000;

export default class Gigasecond {
  constructor(inputDate) {
    this.inputDate = inputDate;
  }

  date() {
    const gigasecondAnniversary = this.inputDate.getTime() + gigasecondInMilliseconds;
    return new Date(gigasecondAnniversary);
  }
}
