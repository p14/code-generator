// WORD BANKS FOR CODE COMBINATIONS
const firstWordBank = [
  'PEANUT',
  'EXIT',
  'TUNE',
  'BONFIRE',
  'JAW',
  'MOTHER',
  'FIELD', // 2.0
];

const secondWordBank = [
  'ASTRONAUT',
  'MONSTERS',
  'SOUR',
  'BLEND',
  'PLANE',
  'PLANTAIN', // 2.0
  'NUKE', // 2.0
];

const thirdWordBank = [
  'COCONUT',
  'SUIT',
  'BLUE',
  'ROCKY',
  'MOON',
  'TREETOP',
  'RASPBERRY', // 2.0
];

const fourthWordBank = [
  'FIELD',
  'PLANE',
  'CLOUD',
  'NUKE',
  'SHOWER',
  'BUGS',
  'BEAN', // 2.0
  'PEEK', // 2.0
];

// FUNCTION TO PERMEATE EVERY POSSIBLE COMBINATION
function getCodeOptions() {
  const options = [];

  for (const firstWord of firstWordBank) {
    for (const secondWord of secondWordBank) {
      for (const thirdWord of thirdWordBank) {
        for (const fourthWord of fourthWordBank) {
          options.push(`${firstWord} ${secondWord} ${thirdWord} ${fourthWord}`);
        }
      }
    }
  }

  return options;
};

// GET OPTIONS AND SHUFFLE LIST
const codeOptions = getCodeOptions();
const shuffledOptions = codeOptions.sort(() => Math.random() - 0.5);

// LOG EACH CODE OPTION LINE BY LINE
shuffledOptions.forEach((code) => {
  console.log(code);
});
