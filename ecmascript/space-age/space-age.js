const SECONDS_IN_EARTH_YEAR = 31557600;

const ORBITAL_PERIODS = {
  earth: 1,
  mercury:  0.2408467,
  venus: 0.61519726,
  mars: 1.8808158,
  jupiter: 11.862615,
  saturn: 29.447498,
  uranus: 84.016846,
  neptune: 164.79132
};

export default class SpaceAge {
  constructor(seconds) {
    this.seconds = seconds;
    
    const createAgeCalculator = orbitalPeriod => () => {
      const age = calculateEarthYears() / orbitalPeriod;
      return formatResponse(age);
    };

    const calculateEarthYears = () => this.seconds / SECONDS_IN_EARTH_YEAR;
    const formatResponse = rawResponse => parseFloat(rawResponse.toFixed(2));

    this.onEarth = createAgeCalculator(ORBITAL_PERIODS.earth);
    this.onMercury = createAgeCalculator(ORBITAL_PERIODS.mercury);
    this.onVenus = createAgeCalculator(ORBITAL_PERIODS.venus);
    this.onMars = createAgeCalculator(ORBITAL_PERIODS.mars);
    this.onJupiter = createAgeCalculator(ORBITAL_PERIODS.jupiter);
    this.onSaturn = createAgeCalculator(ORBITAL_PERIODS.saturn);
    this.onUranus = createAgeCalculator(ORBITAL_PERIODS.uranus);
    this.onNeptune = createAgeCalculator(ORBITAL_PERIODS.neptune);
  }
}
