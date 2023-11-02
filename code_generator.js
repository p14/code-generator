// Word banks for code combinations
const firstWordBank = [
    'PEANUT',
    'EXIT',
    'TUNE',
    'BONFIRE',
    'JAW',
];

const secondWordBank = [
    'ASTRONAUT',
    'MONSTERS',
    'SOUR',
    'BLEND',
    'PLANE',
    'PEANUT',
];

const thirdWordBank = [
    'COCONUT',
    'SUIT',
    'BLUE',
    'ROCKY',
    'MOON',
    'TREETOP',
];

const fourthWordBank = [
    'FIELD',
    'PLANE',
    'CLOUD',
    'NUKE',
    'SHOWER',
    'BUGS',
    'PEAK',
];

// Function to permeate every possible combination
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

// Get options and shuffle the list
const codeOptions = getCodeOptions();
const shuffledOptions = codeOptions.sort(() => Math.random() - 0.5);

// Log each code option line-by-line
shuffledOptions.forEach((code) => {
    console.log(code);
});
