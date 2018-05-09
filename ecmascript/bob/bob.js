const isShout = message => message.match(/[A-Z]/) && message.toUpperCase() === message;
const isQuestion = message => message.endsWith('?');
const isForceful = message => message.endsWith('!');
const isForcefulQuestion = message => isQuestion(message) && isShout(message);
const isSilence = message => message.match(/^\s*$/) !== null;

class Bob {
  hey(message) {
    const testResult = tests.find(test => test.predicate(message));
    return testResult ? testResult.response : Bob.responses.whatever;
  }
}

Bob.responses = {
  whatever: 'Whatever.',
  whoa: 'Whoa, chill out!',
  sure: 'Sure.',
  calmDown: 'Calm down, I know what I\'m doing!',
  fine: 'Fine. Be that way!'
};

const tests = [
  {predicate: isSilence, response: Bob.responses.fine},
  {predicate: isForcefulQuestion, response: Bob.responses.calmDown},
  {predicate: isQuestion, response: Bob.responses.sure},
  {predicate: isShout, response: Bob.responses.whoa},
  {predicate: isForceful, response: Bob.responses.whatever},
];

export default Bob;
