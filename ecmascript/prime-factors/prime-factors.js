export default class PrimeFactors {
  for(number) {
    const primeFactors = [];
    let dividend = number;
    let divisor = 2;

    while (dividend > 1) {
      if (dividend % divisor === 0) {
        primeFactors.push(divisor);
        dividend /= divisor;  
      } else {
        ++divisor;
      }
   }

    return primeFactors;
  }
}